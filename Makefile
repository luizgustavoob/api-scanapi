REGISTRY      	= registry.neoway.com.br
REGISTRY_GROUP 	= solutiondelivery
GITLAB_GROUP   	= sd-projects
BUILD 			= latest
NAME           	= $(shell basename $(CURDIR))
IMAGE          	= $(REGISTRY)/$(REGISTRY_GROUP)/$(NAME):$(BUILD)
POSTGRES_NAME 	= postgres_$(NAME)_$(BUILD)
NETWORK_NAME  	= network_$(NAME)_$(BUILD)
IMAGESCANAPI	= scanapi_$(NAME)_$(BUILD)
APIDOCKERNAME	= api-$(NAME)


clean: ##@dev Remove folder vendor, public and coverage.
	rm -rf vendor public coverage


install: clean ##@dev Download dependencies via go mod.
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor


env: ##@environment Create network and run postgres container.
	POSTGRES_NAME=${POSTGRES_NAME} \
	NETWORK_NAME=${NETWORK_NAME} \
	docker-compose -f ./test/docker-compose.yml up -d


env-stop: ##@environment Remove postgres container and remove network.
	POSTGRES_NAME=${POSTGRES_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose -f ./test/docker-compose.yml kill
	POSTGRES_NAME=${POSTGRES_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose -f ./test/docker-compose.yml rm -vf
	docker network rm $(NETWORK_NAME)


test: clean ##@check Run tests and coverage.
	docker build \
		--progress=plain \
		--network $(NETWORK_NAME) \
		--tag $(IMAGE) \
		--build-arg POSTGRES_URL=postgres://pg:pg@${POSTGRES_NAME}:5432/db?sslmode=disable \
		--target=test \
		--file=./build/Dockerfile .

	-mkdir -p coverage
	docker create --name $(NAME)-$(BUILD) $(IMAGE)
	docker cp $(NAME)-$(BUILD):/index.html ./coverage/.
	docker rm -vf $(NAME)-$(BUILD)


build: clean ##@build Build image.
	DOCKER_BUILDKIT=1 \
	docker build \
		--progress=plain \
		--tag $(IMAGE) \
		--target=build \
		--file=./build/Dockerfile .


image: clean ##@build Create release docker image.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--tag $(IMAGE) \
		--target=image \
		--file=./build/Dockerfile .


run-local: ##@dev Run locally.
	POSTGRES_URL=postgres://pg:pg@$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(POSTGRES_NAME)):5432/db?sslmode=disable \
	go run cmd/server/main.go


run-docker: remove-docker ##@docker Run docker container. BUILD and IMAGE before.
	docker run \
		--name $(APIDOCKERNAME) \
		--network $(NETWORK_NAME) \
		-e POSTGRES_URL=postgres://pg:pg@$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(POSTGRES_NAME)):5432/db?sslmode=disable \
		-p 9998:9998 \
		-d \
		$(IMAGE)


remove-docker: 
	docker stop $(APIDOCKERNAME)
	docker rm $(APIDOCKERNAME)


scan-internal: ##@check Run integration tests with ScanAPI. ENV, BUILD, IMAGE and RUN-DOCKER before.
	cat ./api/scanapi/scan_internal.yml | sed "s/{{HOST}}/$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(APIDOCKERNAME))/g" > ./api/scanapi/scan_myapi.yml
	cat ./api/scanapi/Dockerfile.tmpl | sed "s/{{FULLFILE}}/\/api\/scanapi\/scan_myapi.yml/g" | sed "s/{{FILE}}/scan_myapi.yml/g" > ./api/scanapi/Dockerfile

	docker build \
		--progress=plain \
		--network $(NETWORK_NAME) \
		--tag $(IMAGESCANAPI) \
		--file=./api/scanapi/Dockerfile .
	
	-mkdir -p coverage
	docker create --name $(NAME)-scan-internal $(IMAGESCANAPI)
	docker cp $(NAME)-scan-internal:/report.html ./coverage/scan-internal-report.html
	docker rm -vf $(NAME)-scan-internal
	docker rmi $(IMAGESCANAPI)
	rm -rf ./api/scanapi/Dockerfile
	rm -rf ./api/scanapi/scan_myapi.yml


scan-external: ##@check Run integration tests of external APIs with ScanAPI
	cat ./api/scanapi/Dockerfile.tmpl | sed "s/{{FULLFILE}}/\/api\/scanapi\/scan_external.yml/g" | sed "s/{{FILE}}/scan_external.yml/g" > ./api/scanapi/Dockerfile

	docker build \
		--progress=plain \
		--tag $(IMAGESCANAPI) \
		--file=./api/scanapi/Dockerfile .
	
	-mkdir -p coverage
	docker create --name $(NAME)-scan-external $(IMAGESCANAPI)
	docker cp $(NAME)-scan-external:/report.html ./coverage/scan-external-report.html
	docker rm -vf $(NAME)-scan-external
	docker rmi $(IMAGESCANAPI)
	rm -rf ./api/scanapi/Dockerfile
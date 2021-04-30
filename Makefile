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


run-docker: ##@docker Run docker container. BUILD and IMAGE before.
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
	-mkdir -p coverage
	cat ./api/scanapi/spec_internal.yml | sed "s/{{HOST}}/$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(APIDOCKERNAME))/g" > ./api/scanapi/final_spec_internal.yml
	scanapi run ./api/scanapi/final_spec_internal.yml -o ./coverage/spec-internal-report.html
	rm -rf ./api/scanapi/final_spec_internal.yml


scan-external: ##@check Run integration tests of external APIs with ScanAPI
	-mkdir -p coverage
	scanapi run ./api/scanapi/spec_external.yml -o ./coverage/spec-external-report.html
	
	
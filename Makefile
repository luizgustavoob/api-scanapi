REGISTRY      	= registry.neoway.com.br
REGISTRY_GROUP 	= solutiondelivery
GITLAB_GROUP   	= sd-projects
BUILD 			= latest
NAME           	= $(shell basename $(CURDIR))
IMAGE          	= $(REGISTRY)/$(REGISTRY_GROUP)/$(NAME):$(BUILD)
SCANAPI_IMAGE	= $(REGISTRY)/$(REGISTRY_GROUP)/scanapi:$(BUILD)
POSTGRES_NAME 	= postgres_$(NAME)_$(BUILD)
NETWORK_NAME  	= network_$(NAME)_$(BUILD)
APIDOCKERNAME	= api-$(NAME)


.PHONY: clean
clean: ##@dev Remove folder vendor, public and coverage.
	rm -rf vendor public coverage


.PHONY: install
install: clean ##@dev Download dependencies via go mod.
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor


.PHONY: env
env: ##@environment Create network and run postgres container.
	POSTGRES_NAME=${POSTGRES_NAME} \
	NETWORK_NAME=${NETWORK_NAME} \
	docker-compose -f ./test/docker-compose.yml up -d


.PHONY: env-stop
env-stop: ##@environment Remove postgres container and remove network.
	POSTGRES_NAME=${POSTGRES_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose -f ./test/docker-compose.yml kill
	POSTGRES_NAME=${POSTGRES_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose -f ./test/docker-compose.yml rm -vf
	docker network rm $(NETWORK_NAME)


.PHONY: test
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


.PHONY: build
build: clean ##@build Build image.
	DOCKER_BUILDKIT=1 \
	docker build \
		--progress=plain \
		--tag $(IMAGE) \
		--target=build \
		--file=./build/Dockerfile .


.PHONY: image
image: clean ##@build Create release docker image.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--tag $(IMAGE) \
		--target=image \
		--file=./build/Dockerfile .


.PHONY: run-local
run-local: ##@dev Run locally.
	POSTGRES_URL=postgres://pg:pg@$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(POSTGRES_NAME)):5432/db?sslmode=disable \
	go run cmd/server/main.go


.PHONY: run-docker
run-docker: ##@docker Run docker container. BUILD and IMAGE before.
	docker run \
		--name $(APIDOCKERNAME) \
		--network $(NETWORK_NAME) \
		-e POSTGRES_URL=postgres://pg:pg@$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(POSTGRES_NAME)):5432/db?sslmode=disable \
		-p 9998:9998 \
		-d \
		$(IMAGE)


.PHONY: remove-docker
remove-docker: 
	docker stop $(APIDOCKERNAME)
	docker rm $(APIDOCKERNAME)
	docker rmi $(IMAGE)
	docker rmi -f $(SCANAPI_IMAGE)


.PHONY: image-scanapi
image-scanapi:
	docker image build \
		--progress=plain \
		--tag $(SCANAPI_IMAGE) \
		--file api/scanapi/Dockerfile \
		api/scanapi/


.PHONY: scan-external
scan-external: image-scanapi
	-mkdir -p coverage
	docker container run --rm \
		-v $(CURDIR)/coverage/:/app/coverage/ \
		-v $(CURDIR)/api/scanapi/spec_external.yml:/app/spec/scanapi.yaml \
		$(SCANAPI_IMAGE) \
		"spec/scanapi.yaml" "-o" "coverage/scanapi_external_report.html"


.PHONY: scan-internal
scan-internal: image-scanapi
	-mkdir -p coverage
	cat ./api/scanapi/spec_internal.yml | sed "s/{{HOST}}/$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(APIDOCKERNAME))/g" > ./api/scanapi/final_spec_internal.yml	
	docker container run --rm \
		-v $(CURDIR)/coverage/:/app/coverage/ \
		-v $(CURDIR)/api/scanapi/final_spec_internal.yml:/app/spec/scanapi.yaml \
		--network $(NETWORK_NAME) \
		$(SCANAPI_IMAGE) \
		"spec/scanapi.yaml" "-o" "coverage/scanapi_internal_report.html"	
	rm -rf ./api/scanapi/final_spec_internal.yml
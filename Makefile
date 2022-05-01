SERVER_TAG = vladazn/wow/server
CLIENT_TAG = vladazn/wow/client
SWAGGER_TAG = vladazn/wow/swagger
VERSION = test


all: generate swagger

generate:
	rm -rf proto/gen
	cd proto && buf generate

.PHONY: swagger
swagger:
	rm -f swagger/ui/swagger.json
	cp proto/gen/openapiv2/proto/wow/wow.swagger.json swagger/ui/swagger.json

docker_swagger:
	docker build -t $(SWAGGER_TAG):$(VERSION) -f Docker/Dockerfile.swagger .

docker_server:
	docker build -t $(SERVER_TAG):$(VERSION) -f Docker/Dockerfile.server .

docker_client:
	docker build -t $(CLIENT_TAG):$(VERSION) -f Docker/Dockerfile.client .

docker: docker_server docker_client docker_swagger


mock:
	mockgen -source=internal/repository/repo.go \
	-destination=test/mocks/repository/repo_mock.go
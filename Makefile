# https://sohlich.github.io/post/go_makefile/

build:
	gofmt -l -s -w .
	go mod tidy
	go build -o bin/chefconnect cmd/chefconnect/main.go
	go build -o bin/manage cmd/manager/main.go
swagger:
	go get github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate spec -o ./swagger.json
run:
	go run main.go
swagger-ui:
	docker run --rm -it -p 8081:8080 -e SWAGGER_JSON=/tmp/swagger.json -v ${PWD}:/tmp swaggerapi/swagger-ui
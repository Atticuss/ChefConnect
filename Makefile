# https://sohlich.github.io/post/go_makefile/

build:
	go mod tidy
	go build main.go
swagger-spec:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate spec -o ./swagger.json
run:
	go run main.go
swagger-ui:
	docker run --rm -it -p 8081:8080 swaggerapi/swagger-ui
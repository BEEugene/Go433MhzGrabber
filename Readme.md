


# Swagger


`go get -u github.com/swaggo/swag/cmd/swag
`# go 1.16 or newer
`go install github.com/swaggo/swag/cmd/swag@latest
`go get -u github.com/swaggo/http-swagger
`go get -u github.com/alecthomas/template
`swag init -g ../../cmd/apiserver/main.go -d ./internal/app


Run and acceess http://localhost:8080/swagger/index.html
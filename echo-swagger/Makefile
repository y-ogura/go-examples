REPO=github.com/go-examples/echo-swagger

## generate swagger
gen-swagger:
	# swaggerファイル生成
	swagger -apiPackage="$(REPO)" -mainApiFile="$(REPO)/main.go" -output=./swagger-ui
	sed -i -e "s/package main/package swaggerui/g" ./swagger-ui/docs.go
	sed -i -e "s/\"basePath\": \"http:\/\/localhost:8080\/swagger-ui\"/\"basePath\": \"http:\/\/localhost:8080\"/g" swagger-ui/docs.go
	sed -i -e "1,/\"basePath\": \"http:\/\/localhost:8080\"/s/\"basePath\": \"http:\/\/localhost:8080\"/\"basePath\": \"http:\/\/localhost:8080\/swagger-ui\"/" swagger-ui/docs.go


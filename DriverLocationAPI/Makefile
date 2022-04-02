install-swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger: install-swagger
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
swagger-serve:
	GO111MODULE=off swagger serve -F=swagger --port=9091 --host=localhost swagger.yaml
include .env
export

compose-up:
	docker-compose up --build
.PHONY: compose-up

compose-down:
	docker-compose down
.PHONY: compose-down

linter-golangci:
	golangci-lint run
.PHONY: linter-golangci

test:
	go test -v ./internal/service/impl

mockgen:
	mockgen -source=C:\Users\Lenovo\Desktop\go-market\internal\service\service.go -destination=C:\Users\Lenovo\Desktop\go-market\internal\mocks\servicemocks\servicemocks.go -package=servicemocks
	mockgen -source=C:\Users\Lenovo\Desktop\go-market\internal\repository\repo.go -destination=C:\Users\Lenovo\Desktop\go-market\internal\mocks\repomocks\servicemocks.go -package=servicemocks
.PHONY: mockgen

swag:
	swag init -g .\cmd\app\main.go 
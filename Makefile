binary-name=raspi-htmx

build: templ-build
	@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-win.exe ./cmd/api/main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./cmd/api/main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./cmd/api/main.go

run: build
	@./bin/${binary-name}-linux

arm-build:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${binary-name}-arm64 ./cmd/api/main.go

arm-run: arm-build
	@./bin/${binary-name}-arm64

start-manager:
	@go run ./cmd/manager/main.go

protogen:
	@protoc \
		--proto_path=proto "proto/manager.proto" \
		--go_out=services/common/genproto/services --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/services \
		--go-grpc_opt=paths=source_relative

clean:
	@rm -rf ./bin/*
	@go clean

css-build:
	@tailwindcss -i ./services/web/public/static/css/input.css -o ./services/web/public/static/css/style.css

css-watch:
	@tailwindcss -i ./services/web/public/static/css/input.css -o ./services/web/public/static/css/style.css --watch

templ-build:
	@templ generate

templ-watch:
	@templ generate --watch

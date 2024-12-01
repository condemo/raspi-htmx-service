binary-name=raspi-htmx
service1-name=manager
service2-name=sysinfo
service3-name=weather
service4-name=logger

full-build: build-manager build-sysinfo build-weather build-logger build-htmx
kill-all: kill-services kill-htmx

amd64-build: templ-build
	# API
	@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-win.exe ./cmd/api/main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./cmd/api/main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./cmd/api/main.go

amd64-run: build
	@./bin/${binary-name}-linux

build-htmx:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${binary-name}-arm64 ./cmd/api/main.go

build-manager:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${service1-name}-arm64 ./cmd/manager/main.go

build-sysinfo:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${service2-name}-arm64 ./cmd/sys_info/main.go

build-weather:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${service3-name}-arm64 ./cmd/raspi_services/weather_service/main.go

build-logger:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${service4-name}-arm64 ./cmd/logger/main.go

run-htmx: build-htmx
	@./bin/${binary-name}-arm64

run-manager: build-manager
	@./bin/${service1-name}-arm64

run-sysinfo: build-sysinfo
	@./bin/${service2-name}-arm64

run-weather: build-weather
	@./bin/${service3-name}-arm64

run-logger: build-logger
	@./bin/${service4-name}-arm64

run-all:
	@make run-weather & make run-sysinfo & \
		sleep 1 && make run-manager & make run-htmx

run-services:
	@make run-weather & make run-sysinfo & \
			sleep 1 && make run-manager &

protogen:
	@protoc \
		--proto_path=proto "proto/manager.proto" \
		--go_out=services/common/genproto/services --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/services \
		--go-grpc_opt=paths=source_relative

	@protoc \
		--proto_path=proto "proto/sys_info.proto" \
		--go_out=services/common/genproto/services/sys_info --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/services/sys_info \
		--go-grpc_opt=paths=source_relative

	@protoc \
		--proto_path=proto "proto/raspi_services.proto" \
		--go_out=services/common/genproto/services/raspi_services --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/services/raspi_services \
		--go-grpc_opt=paths=source_relative

	@protoc \
		--proto_path=proto "proto/logger.proto" \
		--go_out=services/common/genproto/services/logger --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/services/logger \
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

kill-services:
	@lsof -t -i:7000 | xargs -r kill
	@lsof -t -i:8000 | xargs -r kill
	@lsof -t -i:8010 | xargs -r kill
	@lsof -t -i:9000 | xargs -r kill

kill-htmx:
	@lsof -t -i:4000 | xargs -r kill

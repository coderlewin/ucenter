
.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: run
run: tidy
	@go run cmd/server/main.go

.PHONY: gorm.gen
gorm.gen: tidy
	@go run cmd/gorm_generate/main.go

.PHONY: build.linux
build.linux: wire
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/server cmd/server/main.go cmd/server/wire_gen.go cmd/server/app.go

.PHONY: wire
wire: tidy
	@wire gen ./cmd/server

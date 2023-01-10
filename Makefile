DOCKER_REPO="hub.easystack.io/production/gosecure-container"
DOCKER_TAG="v1"
NAME = gosecure-container
PROJECT = gosecure-container
PROJECT_DIR = $(GOPATH)src

all: build
build:
	@go build ./cmd/server 

build-arm:
	@GOARM=7 GOARCH=arm64 GOOS=linux go build ./cmd/server 
	
# 安装 swagger 二进制文件
.PHONY: install-swagger
install-swagger:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.27.0

# 给文件添加 copyright
.PHONY: copyright
copyright:
	bash ./scripts/copyright.sh

# 检查swagger 文件是否有错误
.PHONY: valid-swagger
valid-swagger:
	swagger validate ./api/swagger.yaml

# 生成 swagger 
.PHONY: gen-swagger
gen-swagger:
	swagger validate ./api/swagger.yaml
	rm -rf ./api/restapi
	rm -rf ./api/models
	swagger generate server --target ./api --spec ./api/swagger.yaml --exclude-main  --name vanus
	bash ./scripts/copyright.sh

.PHONY: copy-gosecure-container
copy-gosecure-container:
	@mkdir -p $(PROJECT_DIR)
	@cp -rf $(shell command pwd;) $(PROJECT_DIR)
	@mkdir -p $(PROJECT_DIR)/$(NAME)/bin
	@cd $(PROJECT_DIR)/$(NAME)
	@echo $(PROJECT_DIR)/$(NAME)


.PHONY: test-style
test-style:
	@echo "PASS"
#	@golangci-lint run ./...
#	@scripts/teststyle.sh

.PHONY: test-unit
test-unit:
	@echo "PASS"
#	@go test ./...
#	@scripts/testunit.sh

.PHONY: coverage
coverage:
	@echo "PASS"
#	@scripts/coverage.sh

.PHONY: fmt
fmt:
	gofmt -s -w .


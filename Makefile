PROJECT_NAME=gorm_mysql_di
GO_PACKAGE_NAME=github.com/althenlimzixuan/gorm_mysql_di
DOCKER_TAG=latest
ENV=prod
PORT=8080:8080

init_go_project:
	go mod init ${GO_PACKAGE_NAME}	

init_essential_library:
	go install github.com/vektra/mockery/v2@latest && \
	go install github.com/google/wire/cmd/wire@latest && \
	go get -v github.com/google/wire/cmd/wire@latest && \
	go get -v "github.com/google/wire/cmd/wire@latest" && \
	go get -v "github.com/kelseyhightower/envconfig" && \
	go get -v "github.com/sirupsen/logrus"

init_library_project_dir:
	mkdir app
	mkdir app/apis
	mkdir app/constants
	mkdir app/domain
	mkdir app/domain/entities
	mkdir app/domain/repositories
	mkdir app/middlewares
	mkdir app/services
	echo "package constants" > app/constants/string.go
	echo "package constants" > app/constants/integer.go
	echo > .env
	echo > .env_test
	echo > .env_dev
	echo > sample.env
	echo "//go:build wireinject" > app/wire.go
	echo "//+build wireinject" >> app/wire.go
	echo >> app/wire.go
	echo "package app" >> app/wire.go
	echo "package app" > app/application.go
	echo "package app" > app/config.go
	echo "package main" > main.go
		
remove_go:
	rm -rf /usr/local/go && tar -C /usr/loc al -xzf go${GO_VERSION}.linux-amd64.tar.gz

install_go:
	sudo rm -rf /usr/local/go* && sudo rm -rf /usr/local/go
	wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz
	sudo tar -xvf go${GO_VERSION}.linux-amd64.tar.gz
	sudo mv go /usr/local
	sudo rm go${GO_VERSION}.linux-amd64.tar.gz
	echo "export GOROOT=/usr/local/go" >> ~/.bashrc
	echo "export GOPATH=$HOME/go" >> ~/.bashrc
	echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> ~/.bashrc
	nano ~/.bashrc

build:
	docker build --rm -t $(PROJECT_NAME):$(DOCKER_TAG) .

run:
	docker run -d --name=$(PROJECT_NAME) --network bridge --restart=always \
	--env-file ${ENV_FILE} \
	-p $(PORT) \
	$(PROJECT_NAME):$(DOCKER_TAG) 

wire:
	cd app && wire
	cd ..

clean:
	docker rm -f $(PROJECT_NAME):$(DOCKER_TAG)

destroy:
	docker stop $(PROJECT_NAME) && docker rm -f $(PROJECT_NAME)

restart:
	docker restart $(PROJECT_NAME)

test:
	/usr/bin/go test ./app/repository -v
	/usr/bin/go test ./app/usecase -v

test_coverage:
	touch .go-code-cover
	/usr/bin/go test -timeout 30s -coverprofile=./.go-code-cover $(GO_PACKAGE_NAME)/app/domain/repository
	/usr/bin/go test -timeout 30s -coverprofile=./.go-code-cover $(GO_PACKAGE_NAME)/app/domain/usecase

gen_test:
	go test ./... -coverprofile=test_coverage.out
	go tool cover -html=test_coverage.out -o test_coverage.html

install:
	go get \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		github.com/envoyproxy/protoc-gen-validate \
		github.com/bufbuild/buf/cmd/buf

mocks:
	mockery --dir $(dir) --all
	
mock:
	mockery --dir $(dir) --name $(name)

install_wire:
	go install github.com/vektra/mockery/v2@latest
	go install github.com/google/wire/cmd/wire@latest
	go get github.com/google/wire/cmd/wire@latest

	
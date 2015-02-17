APP_NAME = bintray
VERSION = latest

all: clean build

clean:
	@echo "--> Cleaning build"
	@go clean -i
	@rm -rf ./build

prepare:
	@mkdir -p build/bin/{linux-386,linux-amd64,darwin-amd64}
	@mkdir -p build/test
	@mkdir -p build/doc
	@mkdir -p build/tar

deps:
	@echo "--> Getting dependecies"
	@go get -v ./...
	@go get -v github.com/stretchr/testify/mock

format:
	@echo "--> Formatting source code"
	@go fmt ./...

test: prepare deps format
	@echo "--> Testing application"
	@go test -v -outputdir build/test ./...

build: test
	@echo "--> Building local application"
	@go build -o build/bin/linux-386/${VERSION}/${APP_NAME} -v .

build-all: test
	@echo "--> Building all application"
	@echo "... linux-386"
	@GOOS=linux GOARCH=386 go build -o build/bin/linux-386/${VERSION}/${APP_NAME} -v .
	@echo "... linux-amd64"
	@GOOS=linux GOARCH=amd64 go build -o build/bin/linux-amd64/${VERSION}/${APP_NAME} -v .
	@echo "... darwin-amd64"
	@GOOS=darwin GOARCH=amd64 go build -o build/bin/darwin-amd64/${VERSION}/${APP_NAME} -v .

package: build-all
	@echo "--> Packaging application"
	@tar cfz build/tar/${APP_NAME}-${VERSION}-linux-386.tar.gz -C build/bin/linux-386/${VERSION} ${APP_NAME}
	@tar cfz build/tar/${APP_NAME}-${VERSION}-linux-amd64.tar.gz -C build/bin/linux-amd64/${VERSION} ${APP_NAME}
	@tar cfz build/tar/${APP_NAME}-${VERSION}-darwin-amd64.tar.gz -C build/bin/darwin-amd64/${VERSION} ${APP_NAME}

release: package
ifeq ($(VERSION) , latest)
	@echo "--> Removing Latest Version"
	@curl -s -X DELETE -u ${ACCESS_KEY} https://api.bintray.com/packages/darkcrux/generic/${APP_NAME}/versions/${VERSION}
endif
	@echo "--> Releasing version: ${VERSION}"
	@curl -s -T "build/tar/${APP_NAME}-${VERSION}-linux-386.tar.gz" -u "${ACCESS_KEY}" "https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/${APP_NAME}-${VERSION}-linux-386.tar.gz"
	@echo "... linux-386"
	@curl -s -T "build/tar/${APP_NAME}-${VERSION}-linux-amd64.tar.gz" -u "${ACCESS_KEY}" "https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/${APP_NAME}-${VERSION}-linux-amd64.tar.gz"
	@echo "... linux-amd64"
	@curl -s -T "build/tar/${APP_NAME}-${VERSION}-darwin-amd64.tar.gz" -u "${ACCESS_KEY}" "https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/${APP_NAME}-${VERSION}-darwin-amd64.tar.gz"
	@echo "... darwin-amd64"
	@echo "--> Publishing version ${VERSION}"
	@curl -s -X POST -u ${ACCESS_KEY} https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/publish
	@echo 

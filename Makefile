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

format:
	@echo "--> Formatting source code"
	@go fmt ./...

test: prepare format
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
	@echo "--> Releasing version: ${VERSION}"
	@curl -T "build/tar/${APP_NAME}-${VERSION}-linux-386.tar.gz" -u "${ACCESS_KEY}" "https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/${APP_NAME}-${VERSION}-linux-386.tar.gz"
	@echo "... linux-386"
	@curl -T "build/tar/${APP_NAME}-${VERSION}-linux-amd64.tar.gz" -u "${ACCESS_KEY}" "https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/${APP_NAME}-${VERSION}-linux-amd64.tar.gz"
	@echo "... linux-amd64"
	@curl -T "build/tar/${APP_NAME}-${VERSION}-darwin-amd64.tar.gz" -u "${ACCESS_KEY}" "https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/${APP_NAME}-${VERSION}-darwin-386.tar.gz"
	@echo "... darwin-amd64"
	@echo "--> Publishing version ${VERSION}"
	@curl -X POST -u ${ACCESS_KEY} https://api.bintray.com/content/darkcrux/generic/${APP_NAME}/${VERSION}/publish

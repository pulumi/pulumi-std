VERSION := 1.0.0

build:
	mkdir -p bin
	cd terraformfns && go build \
		-o ../bin \
		-ldflags "-X github.com/pulumi/pulumi-terraformfns/terraformfns/version.Version=${VERSION}" ./...

tidy:
	cd terraformfns && go mod tidy

sdk_prep: build
	mkdir -p sdk

gen_sdks: gen_dotnet_sdk gen_nodejs_sdk gen_python_sdk gen_go_sdk gen_schema
	if [ -f sdk/go.mod ]; then rm sdk/go.mod; fi
	cd sdk && go mod init github.com/pulumi/pulumi-terraformfns/sdk

gen_schema: sdk_prep
	pulumi package get-schema bin/pulumi-resource-terraformfns > sdk/schema.json

gen_%_sdk: sdk_prep
	if [ -d sdk/$* ]; then rm -rf sdk/$*; fi
	pulumi package gen-sdk bin/pulumi-resource-terraformfns --language "$*" --out sdk

build_sdks: build_dotnet_sdk build_nodejs_sdk build_python_sdk build_go_sdk
	if ! [ -f sdk/go.mod ]; then \
		cd sdk && go mod init github.com/pulumi/pulumi-terraformfns/sdk; \
	fi
build_dotnet_sdk: gen_dotnet_sdk
	cd sdk/dotnet/ && \
		echo "${VERSION}" >version.txt && \
		dotnet build /p:Version=${VERSION}

build_nodejs_sdk: gen_nodejs_sdk
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc --version && \
		yarn run tsc && \
		cp -R scripts/ bin && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json && \
		rm ./bin/package.json.bak

build_python_sdk: gen_python_sdk
	cd sdk/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

build_go_sdk: gen_go_sdk
	# No-op

test: build
	cd tests && go run main.go

lint: lint-golang lint-copyright
lint-golang:
	cd terraformfns && golangci-lint run -c ../.golangci.yml --timeout 5m
lint-copyright:
	pulumictl copyright -x 'examples/**' -x 'sdk/**'

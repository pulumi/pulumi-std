VERSION := 1.7.2
JAVA_GEN := pulumi-java-gen
JAVA_GEN_VERSION := v0.7.1

build:
	mkdir -p bin
	cd std && go build \
		-o ../bin \
		-ldflags "-X github.com/pulumi/pulumi-std/std/version.Version=${VERSION}" ./...

tidy:
	cd std && go mod tidy
	cd sdk && go mod tidy

sdk_prep: build
	mkdir -p sdk

gen_sdks: gen_dotnet_sdk gen_nodejs_sdk gen_python_sdk gen_go_sdk gen_schema

gen_schema: sdk_prep
	pulumi package get-schema bin/pulumi-resource-std > sdk/schema.json

gen_%_sdk: sdk_prep
	if [ -d sdk/$* ]; then rm -rf sdk/$*; fi
	pulumi package gen-sdk sdk/schema.json --language "$*" --out sdk

build_sdks: build_dotnet_sdk build_nodejs_sdk build_python_sdk build_go_sdk

build_dotnet_sdk: gen_dotnet_sdk
	cd sdk/dotnet/ && \
		echo "${VERSION}" >version.txt && \
		dotnet build /p:Version=${VERSION}

build_nodejs_sdk: gen_nodejs_sdk
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc --version && \
		yarn run tsc
	cp README.md LICENSE sdk/nodejs/package.json sdk/nodejs/yarn.lock sdk/nodejs/bin/
	sed -i.bak 's/$${VERSION}/$(VERSION)/g' sdk/nodejs/bin/package.json
	rm sdk/nodejs/bin/package.json.bak

build_python_sdk: gen_python_sdk
	cd sdk/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

build_go_sdk: gen_go_sdk
	# No-op

gen_java_sdk: PACKAGE_VERSION := $(shell pulumictl get version --language generic)
gen_java_sdk: bin/pulumi-java-gen
	bin/$(JAVA_GEN) generate --schema sdk/schema.json --out sdk/java  --build gradle-nexus
	cd sdk/java/ && \
		echo "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17" > go.mod && \
		gradle --console=plain build

test: build
	cd tests && go run main.go

lint: lint-golang lint-copyright
lint-golang:
	cd std && golangci-lint run -c ../.golangci.yml --timeout 5m
lint-copyright:
	pulumictl copyright -x 'examples/**' -x 'sdk/**'

bin/pulumi-java-gen:
	$(shell pulumictl download-binary -n pulumi-language-java -v $(JAVA_GEN_VERSION) -r pulumi/pulumi-java)
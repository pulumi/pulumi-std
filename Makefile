VERSION := 2.2.0

build:
	mkdir -p bin
	cd provider && go build \
		-o ../bin \
		-ldflags "-X github.com/pulumi/pulumi-std/provider/pkg/version.Version=${VERSION}" ./...

tidy:
	cd provider && go mod tidy
	cd sdk && go mod tidy

sdk_prep: build
	mkdir -p sdk

gen_sdks: gen_dotnet_sdk gen_java_sdk gen_nodejs_sdk gen_python_sdk gen_go_sdk gen_schema

gen_schema: sdk_prep
	pulumi package get-schema ./bin/pulumi-resource-std > sdk/schema.json

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

build_java_sdk: gen_java_sdk
	cd sdk/java/ && \
		echo "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17" > go.mod && \
		gradle --console=plain build

test: build
	cd provider && go test ./...
	cd tests && go run main.go

lint: lint-golang lint-copyright
lint-golang:
	cd provider && golangci-lint run -c ../.golangci.yml --timeout 5m
lint-copyright:
	pulumictl copyright -x 'examples/**' -x 'sdk/**'

# To make an immediately observable change to .ci-mgmt.yaml:
#
# - Edit .ci-mgmt.yaml
# - Run make ci-mgmt to apply the change locally.
#
.PHONY: ci-mgmt
ci-mgmt: .ci-mgmt.yaml
	go run github.com/pulumi/ci-mgmt/provider-ci@master generate

.PHONY: codegen
codegen: # Required by CI
	mkdir -p bin
	touch bin/pulumi-gen-std

.PHONY: provider
provider: build # Required by CI

.PHONY: local_generate
local_generate: gen_sdks # Required by CI

.PHONY: test_provider
test_provider: test # Required by CI

.PHONY: generate_schema
generate_schema: gen_schema # Required by CI

.PHONY: generate_java_sdk build_go generate_go
generate_go: gen_go_sdk # Required by CI
build_go: build_go_sdk # Required by CI
install_go_sdk: # Required by CI

.PHONY: generate_java_sdk build_java install_java_sdk
generate_java: gen_java_sdk # Required by CI
build_java: build_java_sdk # Required by CI
install_java_sdk: # Required by CI

.PHONY: generate_python_sdk build_python install_python_sdk
generate_python: gen_python_sdk # Required by CI
build_python: build_python_sdk # Required by CI
install_python_sdk: # Required by CI

.PHONY: generate_nodejs_sdk build_nodejs install_nodejs_sdk
generate_nodejs: gen_nodejs_sdk # Required by CI
build_nodejs: build_nodejs_sdk # Required by CI
install_nodejs_sdk: # Required by CI

.PHONY: generate_dotnet_sdk build_dotnet_sdk install_dotnet_sdk
generate_dotnet: gen_dotnet_sdk # Required by CI
build_dotnet: build_dotnet_sdk # Required by CI
install_dotnet_sdk: # Required by CI

bin/pulumi-gen-${PACK}: # Required by CI
	touch bin/pulumi-gen-${PACK}

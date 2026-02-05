PROVIDER_VERSION ?= 2.3.0-alpha.0+dev
VERSION_GENERIC = $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")

build:
	mkdir -p bin
	cd provider && go build \
		-o ../bin \
		-ldflags "-X github.com/pulumi/pulumi-std/provider/v2/pkg/version.Version=${VERSION_GENERIC}" ./...

tidy:
	cd provider && go mod tidy
	cd sdk && go mod tidy

sdk_prep: build
	mkdir -p sdk

gen_sdks: gen_dotnet_sdk gen_java_sdk gen_nodejs_sdk gen_python_sdk gen_go_sdk gen_schema

gen_schema: sdk_prep
	pulumi package get-schema ./bin/pulumi-resource-std | jq 'del(.version)' > sdk/schema.json

gen_%_sdk: sdk_prep
	if [ -d sdk/$* ]; then rm -rf sdk/$*; fi
	pulumi package gen-sdk sdk/schema.json --language "$*" --version "${VERSION_GENERIC}"

build_sdks: build_dotnet_sdk build_nodejs_sdk build_python_sdk build_go_sdk

build_dotnet_sdk: gen_dotnet_sdk
	cd sdk/dotnet/ && \
		echo "${VERSION_GENERIC}" >version.txt && \
		dotnet build

build_nodejs_sdk: gen_nodejs_sdk
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc --version && \
		yarn run tsc
	cp README.md LICENSE sdk/nodejs/package.json sdk/nodejs/yarn.lock sdk/nodejs/bin/

build_python_sdk: gen_python_sdk
	cp README.md sdk/python/
	cd sdk/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && \
		../venv/bin/python -m build .

build_go_sdk: gen_go_sdk
	# No-op

build_java_sdk: PACKAGE_VERSION := ${VERSION_GENERIC}
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

# Set these variables to enable signing of the windows binary
AZURE_SIGNING_CLIENT_ID ?=
AZURE_SIGNING_CLIENT_SECRET ?=
AZURE_SIGNING_TENANT_ID ?=
AZURE_SIGNING_KEY_VAULT_URI ?=
SKIP_SIGNING ?=

bin/jsign-6.0.jar:
	mkdir -p bin
	wget https://github.com/ebourg/jsign/releases/download/6.0/jsign-6.0.jar --output-document=bin/jsign-6.0.jar

sign-goreleaser-exe-amd64: GORELEASER_ARCH := amd64_v1
sign-goreleaser-exe-arm64: GORELEASER_ARCH := arm64

# Set the shell to bash to allow for the use of bash syntax.
sign-goreleaser-exe-%: SHELL:=/bin/bash
sign-goreleaser-exe-%: bin/jsign-6.0.jar
	@# Only sign windows binary if fully configured.
	@# Test variables set by joining with | between and looking for || showing at least one variable is empty.
	@# Move the binary to a temporary location and sign it there to avoid the target being up-to-date if signing fails.
	@set -e; \
	if [[ "${SKIP_SIGNING}" != "true" ]]; then \
		if [[ "|${AZURE_SIGNING_CLIENT_ID}|${AZURE_SIGNING_CLIENT_SECRET}|${AZURE_SIGNING_TENANT_ID}|${AZURE_SIGNING_KEY_VAULT_URI}|" == *"||"* ]]; then \
			echo "Can't sign windows binaries as required configuration not set: AZURE_SIGNING_CLIENT_ID, AZURE_SIGNING_CLIENT_SECRET, AZURE_SIGNING_TENANT_ID, AZURE_SIGNING_KEY_VAULT_URI"; \
			echo "To rebuild with signing delete the unsigned windows exe file and rebuild with the fixed configuration"; \
			if [[ "${CI}" == "true" ]]; then exit 1; fi; \
		else \
			file=dist/build-provider-sign-windows_windows_${GORELEASER_ARCH}/pulumi-resource-std.exe; \
			mv $${file} $${file}.unsigned; \
			az login --service-principal \
				--username "${AZURE_SIGNING_CLIENT_ID}" \
				--password "${AZURE_SIGNING_CLIENT_SECRET}" \
				--tenant "${AZURE_SIGNING_TENANT_ID}" \
				--output none; \
			ACCESS_TOKEN=$$(az account get-access-token --resource "https://vault.azure.net" | jq -r .accessToken); \
			java -jar bin/jsign-6.0.jar \
				--storetype AZUREKEYVAULT \
				--keystore "PulumiCodeSigning" \
				--url "${AZURE_SIGNING_KEY_VAULT_URI}" \
				--storepass "$${ACCESS_TOKEN}" \
				$${file}.unsigned; \
			mv $${file}.unsigned $${file}; \
			az logout; \
		fi; \
	fi

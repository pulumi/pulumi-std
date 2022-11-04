# pulumi-terraformfns

Terraform built-in functions implemented as a native Pulumi provider. Mostly used as a compatibility layer when migrating terraform projects to Pulumi. 

### Build

```
make build
```

### Test

```
make test
```

### Generate inferred schema

```
make gen_schema
```

### Installation (SDKs to be published)

The Pulumi String provider is available as a package in all Pulumi languages:

- JavaScript/TypeScript: [`@pulumi/terraformfns`](https://www.npmjs.com/package/@pulumi/terraformfns)
- Python: [`pulumi-terraformfns`](https://pypi.org/project/pulumi-terraformfns/)
- Go: [`github.com/pulumi/pulumi-terraformfns/sdk/go`](https://pkg.go.dev/github.com/pulumi/pulumi-terraformfns/sdk/go)
- .NET: [`Pulumi.Terraformfns`](https://www.nuget.org/packages/Pulumi.Terraformfns)
- YAML: `pulumi plugin install resource terraformfns`
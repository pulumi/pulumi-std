# pulumi-std

Standard library functions implemented as a native Pulumi provider.

### [Function List](FUNCTION_LIST.md)


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

- JavaScript/TypeScript: [`@pulumi/std`](https://www.npmjs.com/package/@pulumi/std)
- Python: [`pulumi-std`](https://pypi.org/project/pulumi-std/)
- Go: [`github.com/pulumi/pulumi-std/sdk/go`](https://pkg.go.dev/github.com/pulumi/pulumi-std/sdk/go)
- .NET: [`Pulumi.Std`](https://www.nuget.org/packages/Pulumi.std)
- YAML: `pulumi plugin install resource std`

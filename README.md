# Terraform Provider Quortex

The Terraform Quortex provider is a plugin for Terraform that allows the full configuration of quortex solution.
This provider is maintained internally by the Quortex team.

## Documentation

Full, comprehensive documentation is available on the Terraform website:

https://registry.terraform.io/providers/quortex/quortex/latest/docs

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-quortex
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory.

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```

# Terraform VMware Workstation Pro Provider

This is the repository for the Terraform VMware Workstation Pro Provider, which one can use
with Terraform to work with [VMware Workstation][vmware-workstation] Products.

[vmware-workstation]: https://www.vmware.com/products/workstation-pro.html

This provider is inpirated on [vSphere Provider][vmware-vsphere]

[vmware-vsphere]: https://github.com/terraform-providers/terraform-provider-vsphere

For general information about Terraform, visit the [official
website][tf-website] and the [GitHub project page][tf-github].

[tf-website]: https://terraform.io/
[tf-github]: https://github.com/hashicorp/terraform

# Using the Provider

The current version of this provider requires Terraform v0.10.2 or higher to
run.

Note that you need to run `terraform init` to fetch the provider before
deploying. Read about the provider split and other changes to TF v0.10.0 in the
official release announcement found [here][tf-0.10-announce].

[tf-0.10-announce]: https://www.hashicorp.com/blog/hashicorp-terraform-0-10/

## Full Provider Documentation

The provider is documented in full on the [GitHub][github-docs] website.
Check the provider documentation for details on entering your connection information and
how to get started with writing configuration for vSphere resources.

[github-docs]: https://github.com/elsudano/terraform-provider-vmworkstation

### Controlling the provider version

Note that you can also control the provider version. This requires the use of a
`provider` block in your Terraform configuration if you have not added one
already.

The syntax is as follows:

```hcl
provider "vmworkstation" {
  version = "~> 1.5"
  ...
}
```

Version locking uses a pessimistic operator, so this version lock would mean
anything within the 1.x namespace, including or after 1.5.0. [Read
more][provider-vc] on provider version control.

[provider-vc]: https://www.terraform.io/docs/configuration/providers.html#provider-versions

# Building The Provider

**NOTE:** Unless you are [developing](#developing-the-provider) or require a
pre-release bugfix or feature, you will want to use the officially released
version of the provider (see [the section above](#using-the-provider)).

## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/elsudano/terraform-provider-vmworkstation`:

```sh
mkdir -p $GOPATH/src/github.com/elsudano
cd $GOPATH/src/github.com/elsudano
git clone git@github.com:elsudano/terraform-provider-vmworkstation
```

## Running the Build

After the clone has been completed, you can enter the provider directory and
build the provider.

```sh
cd $GOPATH/src/github.com/elsudano/terraform-provider-vmworkstation
make build
```

## Installing the Local Plugin

After the build is complete, copy the `terraform-provider-vmworkstation` binary into
the same path as your `terraform` binary, and re-run `terraform init`.

After this, your project-local `.terraform/plugins/ARCH/lock.json` (where `ARCH`
matches the architecture of your machine) file should contain a SHA256 sum that
matches the local plugin. Run `shasum -a 256` on the binary to verify the values
match.

# Developing the Provider

**NOTE:** Before you start work on a feature, please make sure to check the
[issue tracker][gh-issues] and existing [pull requests][gh-prs] to ensure that
work is not being duplicated. For further clarification, you can also ask in a
new issue.

[gh-issues]: https://github.com/elsudano/terraform-provider-vmworkstation/issues
[gh-prs]: https://github.com/elsudano/terraform-provider-vmworkstation/pulls

If you wish to work on the provider, you'll first need [Go][go-website]
installed on your machine (version 1.9+ is **required**). You'll also need to
correctly setup a [GOPATH][gopath], as well as adding `$GOPATH/bin` to your
`$PATH`.

[go-website]: https://golang.org/
[gopath]: http://golang.org/doc/code.html#GOPATH

See [Building the Provider](#building-the-provider) for details on building the provider.

# Testing the Provider

**NOTE:** Pending

## Configuring Environment Variables

Most of the tests in this provider require a comprehensive list of environment
variables to run. See the individual `*_test.go` files in the
[`vsphere/`](vsphere/) directory for more details. The next section also
describes how you can manage a configuration file of the test environment
variables.

### Using the `.tf-vmworkstation-devrc.mk` file

The [`tf-vmworkstation-devrc.mk.example`](tf-vmworkstation-devrc.mk.example) file contains
an up-to-date list of environment variables required to run the acceptance
tests. Copy this to `$HOME/.tf-vmworkstation-devrc.mk` and change the permissions to
something more secure (ie: `chmod 600 $HOME/.tf-vmworkstation-devrc.mk`), and
configure the variables accordingly.

## Running the Acceptance Tests

After this is done, you can run the acceptance tests by running:

```sh
$ make testacc
```

If you want to run against a specific set of tests, run `make testacc` with the
`TESTARGS` parameter containing the run mask as per below:

```sh
make testacc TESTARGS="-run=TestAccVMWVirtualMachine"
```

This following example would run all of the acceptance tests matching
`TestAccVMWVirtualMachine`. Change this for the specific tests you want to
run.


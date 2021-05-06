# nr-dashboard-hcl-gen

Generates HCL configuration for newrelic_one_dashboard resources from exported JSON documents.

## Installing

Install from the binaries on the release page, or via the install script:

```
curl https://raw.githubusercontent.com/ctrombley/nr-dashboard-hcl-gen/main/scripts/install.sh | bash
```

## Usage

Exported JSON can be fed to the tool via stdin or filename.

The `-l` argument can be used to provide a resource label to apply to the generated resource.

### Using stdin

```
cat dashboard.json | nr-dashboard-hcl-gen -l "my_dashboard" > main.tf
```

### Using a file

```
nr-dashboard-hcl-gen -l "my_dashboard" -i dashboard.json > main.tf
```
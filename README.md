# nr-dashboard-hcl-gen

Generates HCL configuration for newrelic_one_dashboard resources from exported JSON documents.

## Usage

Exported JSON can be fed to the tool via stdin or filename.

### Using stdin

```
cat dashboard.json | nr-dashboard-hcl-gen -l "my_dashboard" -o main.tf
```

### Using a file

```
nr-dashboard-hcl-gen -l "my_dashboard" -i dashboard.json -o main.tf
```
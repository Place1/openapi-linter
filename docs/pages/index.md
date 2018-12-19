# Getting Started

## Installation

**[Download - Github Releases Page](https://github.com/Place1/openapi-linter/releases)**

OpenAPI Linter is a single executable binary. You can download the latest release for your platform from the Github Releases page linked above.

Or you can build it from source:

```bash
git clone https://github.com/place1/openapi-linter
cd openapi-linter
go build
./openapi-linter --help
```

## Quickstart

Once you've got the binary (make sure it's on your `$PATH`) you can use it like this:

```bash
openapi-linter ./path/to/openapi.yaml
```

By default `openapi-linter` will look for a config file in the current directory called `openapi-linter.yaml`.
Here's an example config file with some comments:

```yaml
# openapi-linter.yaml
rules:
  noEmptyOperationIDs: true  # make sure all operations have an ID
  noEmptyDescriptions:
    operations: true         # operations must have descriptions
    parameters: true         # parameters must have descriptions
    properties: false        # properties may omit the description field
  naming:                    # naming conventions for different components of the spec
    paths: KebabCase
    tags: PascalCase
    operation: CamelCase
    parameters: SnakeCase
```

**[See Rules Docs](./rules)**

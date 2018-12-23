# OpenAPI Linter

A simple linter for openapi (swagger) specs.

**[Documentation](https://place1.github.io/openapi-linter/)**

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

You can also use the `--config <file>` flag to specify a different file.

**[See Rules Docs](https://place1.github.io/openapi-linter/rules/)**

## Contributing

If you find this project useful please consider contributing back!

If you have a new features (i.e. new rules) or you've encountered a bug
that you'd like to fix then please create a GitHub issue 😃.

### Building & Test
```bash
# Build
go build ./main.go

# Test
go test ./...

# Create release binaries
make

# Cleanup
make clean
```

### Build & Serve Docs
```bash
cd docs
pip install -r requirements.txt

# Serve on localhost:8000 (for development)
mkdocs serve

# Build static docs site
mkdocs build
```

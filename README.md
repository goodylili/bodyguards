# Bodyguards (WIP)

Bodyguards is a collection of Go tools for improving the quality and performance of your Go code. The bodyguards keep
your code safe from bugs and help you write more idiomatic, functional Go code. This project is inspired by Go Report
Card.

## Installation and Setup

To install Bodyguards, run the following command:

```shell
go install github.com/Goodnessuc/bodyguards@latest
```

Next, run the `bodyguards init` command to initialize `bodyguards` for your project.

```shell
bodyguards init
```

The `bodyguards init` command will automatically create a `boyguards.yaml` file in your project's root directory (the
directory where there's a `go.mod` file).

## Usage

To use any of the `bodyguards` tools, run the following command:

```shell
bodyguards [command] [path]
```

The `path` is the path to the Go source code that you want to analyze. If no path is specified, the entire package will
be analyzed or fixed.

## The `bodyguards.yaml` Configuration File

The configuration file is a YAML file where you can specify options for Bodyguards. Here are some available options:

- `formatter`: Format your Go code with `gofmt`.
- `linter`: Lint your Go code with `golangci-lint`.
- `architecture`: Set up the file architecture for your program following one of the popular architectural patterns.
- `report`: Generate a comprehensive report and rating of your project using `Go Report Card`.
- `documentation`: Generate the documentation for your project with `Godoc`.
- `vet`: Vet your Go project with `govet`.

Here's the contents of an example `bodyguards.yaml` file

```yaml
# Formatting options
formatter: gofmt

# Linter options
linter: golangci-lint

# File architecture options
architecture: clean

# Report options
report: Go Report Card

# Documentation options
documentation: Godoc

# Vet options
vet: govet
```

Once you've selected the configuration options for your app; You can use the `run` command to run `bodyguards` to
analyze/fix code errors with the specified options.

```shell
bodyguards run 
```

Additionally, You can generate a comprehensive report of the fixes/issues in your project with the `report` command


```shell
bodyguards report 
```


## Contributing

If you would like to contribute to this project, please feel free to fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
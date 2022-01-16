# gh-environments

[![ci](https://github.com/chelnak/gh-environments/actions/workflows/ci.yml/badge.svg)](https://github.com/chelnak/gh-environments/actions/workflows/ci.yml) [![Release](https://img.shields.io/github/release/chelnak/gh-environments.svg)](https://github.com/chelnak/gh-environments/releases/latest)

A gh-cli extension for managing environments.

## Installation and Upgrades

```bash
gh extension install chelnak/gh-environments
```

```bash
gh extension upgrade chelnak/gh-environments
```

## Usage

``` bash
Work with GitHub environments

Usage:
  environments [command] [flags]
  environments [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Delete an environment.
  help        Help about any command
  list        List environments for a repository
  view        View details about an environment.

Flags:
  -h, --help      help for environments
  -v, --version   version for environments

Use "environments [command] --help" for more information about a command.
```

<!-- ## TODO

[] Usage docs
[] Better command structure
[] Get version of extension (-v?)
[] Why is help in available commands for root with no args
[] Fix help message usage
[] Handle pagination
[] list should be filterable
[] tests -->

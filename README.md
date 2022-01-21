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
gh environments --help
```

## Setting aliases

It's possible to set command aliases with `gh alias set`. Here are some examples for the `gh environments` extension:

```bash
# Set a short name for the command
gh alias set env environments

# Or save a complex jq query
gh alias set myalias "environments list --json -q '.[] | select(.name | contains(""\"te""\"))'"
```

## Advanced usage

Remove multiple environments at once

```bash
#! /bin/bash

set -e

envs=$(gh environments list --json -q '.[] | select(.name | contains("temp-")) | .name')
for row in $(echo "${envs}" | jq -r '.[]'); do
    echo "Removing environment $row"
    gh environments delete $row --force
done
```

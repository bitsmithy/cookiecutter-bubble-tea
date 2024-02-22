#!/usr/bin/env bash

set -euo pipefail

mise install golang@latest
mise use golang@latest
mise exec golang@latest -- go mod init '{{ cookiecutter.app_path }}'
make bootstrap
mise exec just@latest go@latest -- just lint
mise exec just@latest go@latest -- just test
git init && git add . && git commit -m "Initial commit, generated with cookiecutter"
lefthook install

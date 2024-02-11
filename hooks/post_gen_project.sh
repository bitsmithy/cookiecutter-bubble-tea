#!/usr/bin/env bash

mise install golang@latest
mise use golang@latest
go mod init '{{ cookiecutter.app_path }}'
make bootstrap
just lint
just test
git init && git add . && git commit -m "Initial commit, generated with cookiecutter"
lefthook install

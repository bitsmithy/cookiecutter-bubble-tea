#!/usr/bin/env bash

asdf install golang latest
asdf local golang latest
go mod init '{{ cookiecutter.app_path }}'
make install
goimports-reviser -rm-unused -set-alias -format ./...
gofumpt -l -w .
git init && git add . && git commit -m "Initial commit, generated with cookiecutter"

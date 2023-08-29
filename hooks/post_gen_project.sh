#!/usr/bin/env bash

go get -u && go mod tidy
make install
git init && git add . && git commit -m "Initial commit, generated with cookiecutter"

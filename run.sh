#!/usr/bin/env bash
set -e

src="${1%.go}.go"
bin="${src%.go}"
go build -o "$bin" "$src"
"./$bin" < in.txt

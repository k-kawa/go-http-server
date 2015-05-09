#! /bin/bash

TARGET_ARCHS="386 amd64"
TARGET_OS="linux darwin windows"

gox -arch="${TARGET_ARCHS}" \
    -os="${TARGET_OS}" \
    -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

#!/usr/bin/env bash

mkdir -p builds

GOOS=linux   GOARCH=amd64  go build
tar czf builds/bbmri-fhir-gen-${VERSION}-linux-amd64.tar.gz bbmri-fhir-gen
rm bbmri-fhir-gen

GOOS=linux   GOARCH=arm64  go build
tar czf builds/bbmri-fhir-gen-${VERSION}-linux-arm64.tar.gz bbmri-fhir-gen
rm bbmri-fhir-gen

GOOS=darwin  GOARCH=amd64  go build
tar czf builds/bbmri-fhir-gen-${VERSION}-darwin-amd64.tar.gz bbmri-fhir-gen
rm bbmri-fhir-gen

GOOS=darwin  GOARCH=arm64  go build
tar czf builds/bbmri-fhir-gen-${VERSION}-darwin-arm64.tar.gz bbmri-fhir-gen
rm bbmri-fhir-gen

GOOS=windows GOARCH=amd64  go build
zip -q builds/bbmri-fhir-gen-${VERSION}-windows-amd64.zip bbmri-fhir-gen.exe
rm bbmri-fhir-gen.exe

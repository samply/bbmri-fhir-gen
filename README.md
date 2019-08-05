[![Build Status](https://travis-ci.org/samply/bbmri-fhir-gen.svg?branch=master)](https://travis-ci.org/samply/bbmri-fhir-gen)
[![Go Report Card](https://goreportcard.com/badge/github.com/samply/bbmri-fhir-gen)](https://goreportcard.com/report/github.com/samply/bbmri-fhir-gen)

# BBMRI FHIR Test Data Generator

bbmri-fhir-gen is a command line tool to generate FHIR test data following the [BBMRI Implementation Guide][1].

## Installation

bbmri-fhir-gen is written in Go. All you need is a single binary which is available for Linux, macOS and Windows.

### Linux

1. Download the latest release with the command:

   ```bash
   curl -LO https://github.com/samply/bbmri-fhir-gen/releases/download/v0.2.0/bbmri-fhir-gen-0.1.0-linux-amd64.tar.gz
   ```

1. Untar the binary:

   ```bash
   tar xzf bbmri-fhir-gen-0.1.0-linux-amd64.tar.gz
   ```
   
1. Move the binary in to your PATH.

   ```bash
   sudo mv ./bbmri-fhir-gen /usr/local/bin/bbmri-fhir-gen
   ```

1. Test to ensure the version you installed is up-to-date:

   ```bash
   bbmri-fhir-gen --version
   ```

### macOS

1. Download the latest release with the command:

   ```bash
   curl -LO https://github.com/samply/bbmri-fhir-gen/releases/download/v0.2.0/bbmri-fhir-gen-0.1.0-darwin-amd64.tar.gz
   ```

1. Untar the binary:

   ```bash
   tar xzf bbmri-fhir-gen-0.1.0-darwin-amd64.tar.gz
   ```
   
1. Move the binary in to your PATH.

   ```bash
   sudo mv ./bbmri-fhir-gen /usr/local/bin/bbmri-fhir-gen
   ```

1. Test to ensure the version you installed is up-to-date:

   ```bash
   bbmri-fhir-gen --version
   ```

### Windows

1. Download the latest release [here][2]

1. Unzip the binary.

1. Add the binary in to your PATH.

1. Test to ensure the version you downloaded is up-to-date:

   ```
   bbmri-fhir-gen --version
   ```
   
## Usage

```
$ bbmri-fhir-gen
Usage:
  bbmri-fhir-gen [directory] [flags]

Flags:
  -h, --help          help for bbmri-fhir-gen
  -n, --num int       number of patients to generate (default 100)
  -s, --start int     patient index to start with
      --tx-size int   number of patients per transaction (default 100)
      --version       version for bbmri-fhir-gen
```

The generator will generate FHIR transaction bundles in JSON format and put each bundle as file in the directory specified. The defaults will generate one transaction bundle called `transaction-0.json` with 100 patients. If you specify a total number (-n) higher than 100, more than one file will be generated. You can generate additional patients with new identifiers by specifying a custom start value (-s).

## License

Copyright © 2019 The Samply Development Community

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

[1]: <https://github.com/samply/bbmri-fhir-ig>
[2]: <https://github.com/samply/bbmri-fhir-gen/releases/download/v0.1.0/bbmri-fhir-gen-0.1.0-windows-amd64.zip>

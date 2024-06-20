# POC that parses yaml and json payloads to obtain the kubevirt container image

## Overview

A simple POC to read the release file 0000_50_installer_coreos-bootimages.yaml parse it, read the configmap json contents, parse that and get the kubevirt container by digest


## Usage

```bash

# clone and build

git clone https://github.com/lmzuccarelli/golang-oc-mirror-kubevirt

cd golang-oc-mirror-kubevirt

go build -o kubevirt-poc

```

Copy the file **0000_50_installer_coreos-bootimages.yaml** to the current directory

Execute to verify

```bash

./kubevirt-poc

```

# Guide

## Installation

1. Download the latest cli from our github release per your OS. Or you can build from source in `cli` folder.
2. Copy the binary to your system PATH, and make it executable. For example:
```
cp api-insights-cli /usr/local/bin
chmod +x /usr/local/bin/api-insights-cli
```

## folders

```
v0.0-rev1: raw spec
v0.0-rev2: perfect catalogue spec
v0.1-rev1: add some violations against perfect spec for catalogue
v0.1-rev2: fix all the issues in v0.1-rev1 for catalogue
```

## Get started

The `setup.sh` will create 5 services, and upload two versions of specs to it.
You can specify the version in argument, default to `v0.0-rev1`, `v0.0-rev2`, `v0.1-rev1`.
Change `host` in `.env` file can point to the api server, default to `http://localhost:8081`.
```
./setup.sh v0.0-rev2 v0.1-rev1 v0.1-rev2
```

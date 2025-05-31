# 📷 OpenMVGO - Go wrapper for OpenMVG and OpenMVS
[![Go](https://github.com/2024-dissertation/openmvgo/actions/workflows/ci.yml/badge.svg)](https://github.com/2024-dissertation/openmvgo/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/2024-dissertation/openmvgo.svg)](https://pkg.go.dev/github.com/2024-dissertation/openmvgo)

### Usage

Run demo:

```
NAME:
   OpenMVGO - A CLI tool for OpenMVG and OpenMVS operations

USAGE:
   go run cmd/cli/main.go [input dir] [output dir] [arguments...]

GLOBAL OPTIONS:
   --maxThreads int  (default: 1)
   --help, -h        show help
```

```sh
go run cmd/cli/main.go  ${PWD}/testdata/input/ ${PWD}/testdata/output --maxThreads=0
```

**Note: max threads 0 will use all available threads**

### Docker

A standalone Dockerfile is provided `Dockerfile.prod` [here](Dockerfile.prod). For development, use the `docker-compose-dev.yml` [here](docker-compose-dev.yml).

These will install [OpenMVS](https://github.com/cdcseacave/openMVS) and [OpenMVG](https://github.com/openMVG/openMVG)

Library coming soon!
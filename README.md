# About

[![Go Report Card](https://goreportcard.com/badge/github.com/tusharlock10/pego)](https://goreportcard.com/report/github.com/tusharlock10/pego)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/tusharlock10/pego)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/tusharlock10/pego)](https://pkg.go.dev/mod/github.com/tusharlock10/pego)

`pego` is a Golang API Library for the Hi-Rez developer API [PDF reference](https://docs.google.com/document/d/1OFS-3ocSx-1Rvg4afAnEHlT3917MAK_6eJTR6rzr-BM/edit)

Add it with `go get`
```bash
go get github.com/tusharlock10/pego
```

## Getting Started

```go
package main

import (
  "github.com/tusharlock10/pego/hirezapi"
  "github.com/tusharlock10/pego/models"
)

func main() {
  // Recommend storing these as environment variables or secrets
  devID := "yourDevID" // os.Getenv("MY_SECRET_DEV_ID")
  authKey := "yourAuthKey" // os.Getenv("MY_SECRET_AUTH_KEY")

  client, err := hirezapi.Init(devID, authKey)
}
```
# trusty-sdk-go - Enhanced
A Golang SDK for trusty

The SDK has been enhanced to fully decode the package reports.

```go
package main

import (
	"fmt"
	"context"

	"github.com/mintoolkit/mint/pkg/util/jsonutil"
	"github.com/autonomous-plane/trusty-sdk-go/pkg/client"
	"github.com/autonomous-plane/trusty-sdk-go/pkg/types"
)

func main() {
	trusty := client.New()

	if report, err := trusty.Report(context.TODO(), 
		&types.Dependency{
			Name: "express",
			Version: "4.17.1",
			Ecosystem: types.EcosystemNpm}); err == nil {
		fmt.Println(jsonutil.ToPretty(report))
	}
}
```

## Overview

Supported package types/ecosystems:

* `NPM` - `types.EcosystemNpm` - Node.js packages (purl type: `npm`)
* `Go` - `types.EcosystemGo` - Go packages (purl type: `golang`)
* `PyPI` - `types.EcosystemPypi` - Python packages (purl type: `pypi`)
* `Maven` - `types.EcosystemMaven` - Java packages (purl type: `maven`)
* `Crates` - `types.EcosystemCrates` - Rust/Cargo packages (purl type: `cargo`)

## Changes

* Ecosystem constants follow Go standards (e.g., `EcosystemNpm` instead of `ECOSYSTEM_NPM`)

## References

Online portal to lookup dependency information: https://www.trustypkg.dev/



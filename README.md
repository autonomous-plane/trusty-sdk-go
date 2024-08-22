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
			Ecosystem: types.ECOSYSTEM_NPM}); err == nil {
		fmt.Println(jsonutil.ToPretty(report))
	} else {
		panic(err)
	}
}
```


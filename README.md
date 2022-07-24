<div align="center">
  <h1><code>go-utils</code></h1>
  <p>
    <strong>A go package for various useful utilities.</strong>
  </p>
</div>

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/pkgsite.svg)](https://pkg.go.dev/github.com/mrinjamul/go-utils)
[![GoReportCard](https://goreportcard.com/badge/github.com/mrinjamul/go-utils)](https://goreportcard.com/report/github.com/mrinjamul/go-utils)
[![Code style: Standard](https://img.shields.io/badge/code%20style-Standard-blue.svg)]()
[![License: MIT ](https://img.shields.io/badge/License-MIT%20License-blue.svg)](https://github.com/mrinjamul/go-utils/blob/main/LICENSE)

To use this module,

```sh
go get github.com/mrinjamul/go-timeago
```

and then,

```go
import "github.com/mrinjamul/go-utils/[sub-package]"
```

### Examples

```go
package main

import (
	"fmt"

	"github.com/mrinjamul/go-utils/utils"
)

func main() {
    sha256sum := utils.Sha256sum("hello")
	fmt.Println(sha256sum)
}
```

## License

This go module is licensed under MIT License.
<br/>
Copyright © 2021 Injamul Mohammad Mollah <mrinjamul@gmail.com>

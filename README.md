# Ginka

A mocha like reporter for [Ginkgo](https://onsi.github.io/ginkgo/)

## Usage

Just import and use as custom reporters:

```go
package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/noodles723/ginka"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Book Suite", ginka.NewReporter())
}
```

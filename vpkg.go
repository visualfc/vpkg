package vpkg

import (
	"fmt"

	"github.com/visualfc/modtest"
	"github.com/visualfc/modtest/pkg"
)

func Test() {
	fmt.Println(modtest.Version())
	fmt.Println(pkg.Version())
}

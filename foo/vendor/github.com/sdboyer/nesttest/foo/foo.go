package main

import (
	"fmt"

	"github.com/sdboyer/nesttest"
	"github.com/sdboyer/nesttest/bar"
	"github.com/sdboyer/nesttest/foo/baz"
)

func main() {
	fmt.Println("Hello from foo/main in vendor area")
	bar.DoBar()
	baz.DoBaz()
	quux.DoQuux()
}

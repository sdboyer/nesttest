package main

import (
	"fmt"

	"github.com/sdboyer/nesttest"
	"github.com/sdboyer/nesttest/bar"
	"github.com/sdboyer/nesttest/foo/baz"
)

func main() {
	fmt.Println("Hello from foo/main in main area")
	ohai()
	DoFooOtherFile()
	bar.DoBar()
	baz.DoBaz()
	quux.DoQuux()
}

func ohai() {
	fmt.Println("ohai from main")
}

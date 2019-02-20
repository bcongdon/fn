package main

import (
	"fmt"

	"github.com/bcongdon/fn"
)

func main() {
	fNamer := fn.New()

	// Basic Name
	fmt.Println(fNamer.Name())
	// "200260220-072532-392d644-6193ecb1"

	// Name w/ Prefix
	fNamer.Prefix = "foo"
	fmt.Println(fNamer.Name())
	// "foo-200260220-072532-392d644-4dffcc5b"

	// Name w/ Postfix
	fNamer.Postfix = "bar"
	fNamer.Prefix = ""
	fmt.Println(fNamer.Name())
	// "200260220-072532-392d644-c25334f9-bar"

	// Name w/ file extension
	fNamer.Postfix = ""
	fmt.Println(fNamer.NameWithFileType("png"))
	// "200260220-072532-392d644-a165c554.png"
}

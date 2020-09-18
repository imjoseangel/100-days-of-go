///bin/true; exec /usr/bin/env go run "" "$@"
package main

import (
	"fmt"

	"github.com/imjoseangel/functions-go"
)

func main() {

	untaken := functions.RangeList{
		MinList: 1,
		MaxList: 10,
	}.MakeRange()

	fmt.Println(untaken)
}

package main

import (
	"fmt"

	"github.com/dostonlv/fpm/fpm"
)

func main() {
	opt := fpm.Some(42)

	doubled := fpm.Map(opt, func(x int) int {
		return x * 2
	})

	if doubled.IsSome() {
		fmt.Println("Result:", doubled.Unwrap())
	}

	chained := fpm.FlatMap(opt, func(x int) fpm.Option[string] {
		if x > 40 {
			return fpm.Some("Greater than 40")
		}
		return fpm.None[string]()
	})
	fmt.Println("Chained:", chained.String())

	none := fpm.None[int]()
	defaultVal := none.OrElse(999)
	fmt.Println("None default:", defaultVal)

	chained.Match(
		func(val string) { fmt.Println("Match:", val) },
		func() { fmt.Println("No value") },
	)
}

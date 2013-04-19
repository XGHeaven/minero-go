package config_test

import (
	"fmt"

	"github.com/toqueteos/minero/config"
)

func ExampleSortedMap() {
	c := config.New()
	c.Parse("a:\n b:2\n c:3\nd:4")

	m := c.Copy()
	fmt.Println(config.SortedMap(m))
	// Output:
	// "a.b": 2
	// "a.c": 3
	// "d": 4
}
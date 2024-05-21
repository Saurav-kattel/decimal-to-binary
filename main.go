package main

import (
	"fmt"
	"side-project/decimal-bin/converter"
)

func main() {
	c := converter.Converter{}
	binary := c.Convert()
	fmt.Println(binary)
}

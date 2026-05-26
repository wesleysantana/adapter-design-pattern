package main

import (
	"fmt"

	report "github.com/wesleysantana/adapter-design-pattern/internal/sale/report"
)

func main() {
	fileName, err := report.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileName)
}

package main

import (
	go_excel "apple-x-co/go-excel"
	"fmt"
	"os"
)

var version string
var revision string

func main() {
	goExcel := go_excel.NewGoExcel(version, revision)
	err := goExcel.Execute()
	if err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

package main

import (
	"fmt"

	"github.com/jsoriano/testbuilds"
)

func main() {
	fmt.Println("Plugins:", len(testbuilds.Plugins))
	for _, p := range testbuilds.Plugins {
		fmt.Println(p.GetName())
		for _, k := range p.Keys() {
			fmt.Printf(" - %v: %v\n", k, p.Setting(k))
		}
	}
}

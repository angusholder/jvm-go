package main

import (
	"fmt"
	"github.com/angusholder/jvm-go/classfile"
	"io/ioutil"
)

func main() {
	raw, err := ioutil.ReadFile("Utf8.class")
	if err != nil {
		panic(err)
	}

	cls, err := classfile.ParseFrom(raw)
	if err != nil {
		panic(err)
	}

	fmt.Printf("cls: %#v\n", cls)
}

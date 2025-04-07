package puppy

import (
	"fmt"

	"github.com/ojaspar/go-dog"
)

func Bark() string {
	return "woof!!"
}

func Barks() string {
	return "woof woof!!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func From11() {
	fmt.Println("I'am from version 1.1.0")
}

func From12() {
	fmt.Println("I'am from version 1.2.0")
}

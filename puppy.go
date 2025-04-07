package puppy

import (
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

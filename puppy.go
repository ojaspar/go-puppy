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

func Name() string {
	return dog.WhenGrowUp("Krane")
}

package pkg1

import (
	"fmt"

	"syreclabs.com/go/faker"
)

func HolaPkg() {
	fmt.Println("hola package 1 for test.")
}

func GiveMeFakeFullName() string {
	return faker.Name().Name()
}

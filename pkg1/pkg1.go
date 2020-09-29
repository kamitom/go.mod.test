package pkg1

import (
	"fmt"

	"syreclabs.com/go/faker"
)

func HolaPkg1() {
	fmt.Println("Hola pkg1 for test.")
}

func GiveMeFakeFullName() string {
	return faker.Name().Name()
}

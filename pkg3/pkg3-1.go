package pkg3

import (
	"fmt"
)

//TestArraySlice is ...
func TestArraySlice() {

	names := []string{"tom", "nsa", "csw"}

	fmt.Println(names[2])
	// fmt.Println("hello?! TestArraySlice?!")

}

// KeyValuePairTest is ...
func KeyValuePairTest(urname string) {
	kvs := map[string]string{
		"name":  urname,
		"email": urname + "@pm.me",
	}

	for key, value := range kvs {
		fmt.Println(key+":", value)
	}
}

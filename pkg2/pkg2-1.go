package pkg2

import (
	"fmt"

	"rsc.io/quote"
)

func QuoteToMeHello() string {
	return fmt.Sprintf("the quote HELLO is: %v", quote.Hello())
}
func QuoteToMeGo() string {
	return fmt.Sprintf("the quote Go is: %v", quote.Go())
}

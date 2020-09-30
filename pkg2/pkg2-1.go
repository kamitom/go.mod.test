package pkg2

import (
	"fmt"

	"rsc.io/quote"
)

func QuoteToMe() string {
	return fmt.Sprintf("the quote HELLO is: %v", quote.Hello())
}

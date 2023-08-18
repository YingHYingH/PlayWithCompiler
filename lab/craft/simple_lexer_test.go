package craft

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	script := "int age = 45;"
	fmt.Println("parse : ", script)
	simpleTokenReader := tokenize(script)
	dump(simpleTokenReader)
}

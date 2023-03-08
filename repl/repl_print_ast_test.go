package repl

import (
	"fmt"
	"strings"
	"testing"
)

func TestItShouldPrintAST(t *testing.T) {
	const program = "x * y / 2 + 3 * 8 - 123"
	input := strings.NewReader(fmt.Sprintf("%v\n%v\n", program, QUIT))
	var output Output
	Start(input, &output)

	t.Log(output)
}

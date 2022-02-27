package parser

import (
	"fmt"
)

func errWrongSyntax(i int) error {
	return fmt.Errorf("wrong syntax in line %v", i+1)
}

package parser

import (
	"fmt"
)

func errWrongSyntax(i int) error {
	return fmt.Errorf("wrong syntax in line %v", i+1)
}

func errUnexpectedType(types string) error {
	return fmt.Errorf("unexpected type of token: %v", types)
}
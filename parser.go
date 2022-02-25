package parser

import (
	"bufio"
	"os"

	"go.uber.org/zap"
)

type AST struct {
	Tokens []*Token
	logger *zap.Logger
}

type Token struct {
	Name  string
	Rules []*Rule
}

type Rule map[string]string

// ParseIntoAST css into AST
func ParseIntoAST(filename string) (*AST, error) {
	logger, _ := zap.NewDevelopment()
	AST := &AST{
		logger: logger}
	AST.scanFile(filename)
	return AST, nil
}

// ParseIntoCSS parse AST to css
func ParseIntoCSS(ast *AST) error {
	return nil
}

func (ast *AST) scanFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		ast.logger.Error("error", zap.Error(err))
	}
	sc := bufio.NewScanner(file)
	_, _ = createToken(sc)
	return nil
}

func createToken(sc *bufio.Scanner) ([]*Token, error) {
	return nil, nil
}

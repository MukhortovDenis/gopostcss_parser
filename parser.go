package parser

import (
	"bufio"
	"os"

	"go.uber.org/zap"
)

var (
	newline       = []byte("\n")
	selectorID    = []byte("#")
	selectorClass = []byte(".")
	selectorAll   = []byte("*")
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

	var i int
	cache := map[int][]byte{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		cache[i] = sc.Bytes()
		i++
	}

	_, _ = ast.createTokens(cache)
	return nil
}

func (ast *AST) createTokens(cache map[int][]byte) ([]*Token, error) {
	tokens := []*Token{}
	for i, key := range cache {
		switch key[0] {

		case newline[0]:
			continue

		case selectorAll[0]:
			tokenAll, err := tokenSelectorAll(i, cache)
			if err != nil {
				ast.logger.Error("error", zap.Error(err))
			}
			tokens = append(tokens, tokenAll)

		case selectorID[0]:
			tokenID, err := tokenSelectorID(i, cache)
			if err != nil {
				ast.logger.Error("error", zap.Error(err))
			}
			tokens = append(tokens, tokenID)

		case selectorClass[0]:
			tokenClass, err := tokenSelectorClass(i, cache)
			if err != nil {
				ast.logger.Error("error", zap.Error(err))
			}
			tokens = append(tokens, tokenClass)

		default:
			tokenTag, err := tokenSelectorTag(i, cache)
			if err != nil {
				ast.logger.Error("error", zap.Error(err))
			}
			tokens = append(tokens, tokenTag)
		}

	}
	return tokens, nil
}

func tokenSelectorID(i int, cache map[int][]byte) (*Token, error)

func tokenSelectorClass(i int, cache map[int][]byte) (*Token, error)

func tokenSelectorAll(i int, cache map[int][]byte) (*Token, error)

func tokenSelectorTag(i int, cache map[int][]byte) (*Token, error)
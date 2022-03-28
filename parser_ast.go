package parser

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"go.uber.org/zap"
)

const (
	newline       byte = 10
	selectorID    byte = 35
	selectorClass byte = 46
	selectorAll   byte = 42
	a_rule        byte = 64
)

const (
	aRuleType         = "@-Rules"
	selectorIDType    = "Selector ID"
	selectorClassType = "Selector Class"
	selectorAllType   = "Selector All"
	selectorTagType   = "Selector Tag"
)

var nullString []byte = []byte{}

type AST struct {
	Tokens []*Token
	logger *zap.Logger
}

type Token struct {
	Type  string
	Name  string
	Rules []*Rule
}

type Rule map[int]string

// ParseIntoAST css into AST
func ParseIntoAST(filename string) (*AST, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	AST := &AST{
		logger: logger}
	AST.scanFile(filename)
	return AST, nil
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
	err = ast.tokenizator(cache)
	if err != nil {
		ast.logger.Error("error", zap.Error(err))
	}
	return nil
}

func (ast *AST) tokenizator(cache map[int][]byte) error {
	var i int = 0
	for {
		if len(cache) == i {
			break
		}
		if len(cache[i]) == 0 {
			i++
			continue
		}

		switch cache[i][0] {

		case newline:
			i++
			continue

		case a_rule:
			tokenImport, newIndex, err := newTokenARule(i, cache)
			if err != nil {
				return err
			}
			ast.Tokens = append(ast.Tokens, tokenImport)
			i = newIndex

		case selectorAll:
			tokenAll, newIndex, err := newTokenSelectorAll(i, cache)
			if err != nil {
				return err
			}
			ast.Tokens = append(ast.Tokens, tokenAll)
			i = newIndex

		case selectorID:
			tokenID, newIndex, err := newTokenSelectorID(i, cache)
			if err != nil {
				return err
			}
			ast.Tokens = append(ast.Tokens, tokenID)
			i = newIndex

		default:
			if strings.Contains(string(cache[i][0]), ".") {
				tokenClass, newIndex, err := newTokenSelectorClass(i, cache)
				if err != nil {
					return err
				}
				ast.Tokens = append(ast.Tokens, tokenClass)
				i = newIndex
			} else {
				tokenTag, newIndex, err := newTokenSelectorTag(i, cache)
				if err != nil {
					return err
				}
				ast.Tokens = append(ast.Tokens, tokenTag)
				i = newIndex
			}
		}
	}
	return nil
}

func newTokenARule(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: aRuleType}
	rule := map[int]string{}
	var str []byte
	str = cache[i]
	if strings.Contains(string(cache[i]), "import") {
		token.Name = "@import"
		slice := strings.Fields(string(str))
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
		return token, i, nil
	}
	if strings.Contains(string(cache[i]), "charset") {
		token.Name = "@charset"
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
		return token, i, nil
	}
	if strings.Contains(string(cache[i]), "@font-face") || strings.Contains(string(cache[i]), "@page") {
		firstSlice := strings.Fields(string(cache[i]))
		token.Name = firstSlice[0]
		i++
		for {
			rule := map[int]string{}
			str = cache[i]
			if strings.Contains(string(str), "}") {
				break
			}
			slice := strings.Fields(string(str))
			for ind, word := range slice {
				word = strings.TrimSuffix(word, ",")
				word = strings.TrimSuffix(word, ";")
				rule[ind] = word
			}
			Rule := Rule(rule)
			token.Rules = append(token.Rules, &Rule)
			i++
		}
		i++
		return token, i, nil
	}
	return nil, 0, errors.New("invalid @-rule")
}

func newTokenSelectorID(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: selectorIDType}
	var str []byte
	firstSlice := strings.Fields(string(cache[i]))
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Fields(string(str))
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	i++
	return token, i, nil
}

func newTokenSelectorClass(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: selectorClassType}
	var str []byte
	firstSlice := strings.Fields(string(cache[i]))
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Fields(string(str))
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	i++
	return token, i, nil
}

func newTokenSelectorAll(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: selectorAllType}
	var str []byte
	firstSlice := strings.Fields(string(cache[i]))
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Fields(string(str))
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	i++
	return token, i, nil
}

func newTokenSelectorTag(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: selectorTagType}
	var str []byte
	firstSlice := strings.Fields(string(cache[i]))
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Fields(string(str))
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	i++
	return token, i, nil
}
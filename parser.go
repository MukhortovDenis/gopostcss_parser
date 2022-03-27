package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

var (
	newline       = []byte("\n")
	selectorID    = []byte("#")
	selectorClass = []byte(".")
	selectorAll   = []byte("*")
	a_rule        = []byte("@")
)

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

func (ast *AST) scanFile(filename string) ([]*Token, error) {
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
	tokens, err := ast.tokenizator(cache)
	if err != nil {
		ast.logger.Error("error", zap.Error(err))
	}
	return tokens, nil
}

func (ast *AST) tokenizator(cache map[int][]byte) ([]*Token, error) {
	tokens := []*Token{}
	for i, k := range cache {
		switch k[0] {

		case newline[0]:
			continue

		case a_rule[0]:
			tokenImport, newIndex, err := tokenARule(i, cache)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tokenImport)
			i = newIndex

		case selectorAll[0]:
			tokenAll, newIndex, err := tokenSelectorAll(i, cache)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tokenAll)
			i = newIndex

		case selectorID[0]:
			tokenID, newIndex, err := tokenSelectorID(i, cache)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tokenID)
			i = newIndex

		default:
			if strings.Contains(string(k[0]), ".") {
				tokenClass, newIndex, err := tokenSelectorClass(i, cache)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, tokenClass)
				i = newIndex
			}
			if !strings.Contains(string(k[0]), ".") {
				tokenTag, newIndex, err := tokenSelectorTag(i, cache)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, tokenTag)
				i = newIndex
			}
		}

	}
	return tokens, nil
}

func tokenARule(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: "@-Rules"}
	rule := map[int]string{}
	var str []byte
	str = cache[i]
	if strings.Contains(string(cache[i]), "import") {
		token.Name = "@import"
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			fmt.Println(word)
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
		firstSlice := strings.Split(string(cache[i]), " ")
		token.Name = firstSlice[0]
		i++
		for {
			rule := map[int]string{}
			str = cache[i]
			if strings.Contains(string(str), "}") {
				break
			}
			slice := strings.Split(string(str), " ")
			for ind, word := range slice {
				word = strings.TrimSuffix(word, ",")
				word = strings.TrimSuffix(word, ";")
				rule[ind] = word
			}
			Rule := Rule(rule)
			token.Rules = append(token.Rules, &Rule)
			i++
		}
		return token, i, nil
	}
	return nil, 0, errors.New("invalid @-rule")
}

func tokenSelectorID(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: "Selector ID"}
	var str []byte
	firstSlice := strings.Split(string(cache[i]), " ")
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	return token, i, nil
}

func tokenSelectorClass(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: "Selector Class"}
	var str []byte
	firstSlice := strings.Split(string(cache[i]), " ")
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	return token, i, nil
}

func tokenSelectorAll(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: "Selector All"}
	var str []byte
	firstSlice := strings.Split(string(cache[i]), " ")
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	return token, i, nil
}

func tokenSelectorTag(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Type: "Selector Tag"}
	var str []byte
	firstSlice := strings.Split(string(cache[i]), " ")
	token.Name = firstSlice[0]
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		if strings.Contains(string(str), "}") {
			break
		}
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
		}
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
		i++
	}
	return token, i, nil
}

// func isValidSyntax(slice []string) bool {
// 	if len(slice) != 2 || slice[1] != "{" {
// 		if len(slice) != 1 || !strings.Contains(slice[0], "{") {
// 			return false
// 		}
// 	}
// 	return true
// }

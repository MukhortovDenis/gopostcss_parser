package parser

import (
	"bufio"
	"os"
	"strings"

	"go.uber.org/zap"
)

var (
	newline       = []byte("\n")
	selectorID    = []byte("#")
	selectorClass = []byte(".")
	selectorAll   = []byte("*")
	importCss     = []byte("@")
)

type AST struct {
	Tokens []*Token
	logger *zap.Logger
}

type Token struct {
	Name  string
	Type  string
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
	for i, key := range cache {
		switch key[0] {

		case newline[0]:
			continue

		case importCss[0]:
			tokenImport, newIndex, err := tokenImport(i, cache)
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

		case selectorClass[0]:
			tokenClass, newIndex, err := tokenSelectorClass(i, cache)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tokenClass)
			i = newIndex

		default:
			tokenTag, newIndex, err := tokenSelectorTag(i, cache)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tokenTag)
			i = newIndex
		}

	}
	return tokens, nil
}

func tokenImport(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Name: "import"}
	rule := map[int]string{}
	var str []byte
	str = cache[i]
	slice := strings.Split(string(str), " ")
	for ind, word := range slice {
		word = strings.TrimSuffix(word, ",")
		word = strings.TrimSuffix(word, ";")
		rule[ind] = word
		Rule := Rule(rule)
		token.Rules = append(token.Rules, &Rule)
	}
	i++
	return token, i, nil
}

func tokenSelectorID(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Name: "Selector ID"}
	str := cache[i]
	token.Type = string(str[0])
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
			Rule := Rule(rule)
			token.Rules = append(token.Rules, &Rule)
		}
		i++
		if strings.Contains(string(str), "}") {
			break
		}
	}
	return token, i, nil
}

func tokenSelectorClass(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Name: "Selector Class"}
	str := cache[i]
	token.Type = string(str[0])
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
			Rule := Rule(rule)
			token.Rules = append(token.Rules, &Rule)
		}
		i++
		if strings.Contains(string(str), "}") {
			break
		}
	}
	return token, i, nil
}

func tokenSelectorAll(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Name: "Selector All"}
	str := cache[i]
	token.Type = string(str[0])
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
			Rule := Rule(rule)
			token.Rules = append(token.Rules, &Rule)
		}
		i++
		if strings.Contains(string(str), "}") {
			break
		}
	}
	return token, i, nil
}

func tokenSelectorTag(i int, cache map[int][]byte) (*Token, int, error) {
	token := &Token{Name: "Selector Tag"}
	str := cache[i]
	token.Type = string(str[0])
	i++
	for {
		rule := map[int]string{}
		str = cache[i]
		slice := strings.Split(string(str), " ")
		for ind, word := range slice {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ";")
			rule[ind] = word
			Rule := Rule(rule)
			token.Rules = append(token.Rules, &Rule)
		}
		i++
		if strings.Contains(string(str), "}") {
			break
		}
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

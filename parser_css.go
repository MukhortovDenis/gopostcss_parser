package parser

import (
	"os"
	"strings"
)

func ParseIntoCSS(ast *AST, filename string) error {
	if err := ast.createFile(filename); err != nil {
		return err
	}
	return nil
}

func (ast *AST) createFile(filename string) error {
	file, err := os.OpenFile("new_"+filename, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = ast.writeTokens(file); err != nil {
		return err
	}
	return nil
}

func (ast *AST) writeTokens(file *os.File) error {
	for token := range ast.Tokens {
		switch ast.Tokens[token].Type {

		case selectorClassType:
			if err := writeTokenSelectorClass(ast.Tokens[token], file); err != nil {
				return err
			}

		case selectorTagType:
			if err := writeTokenSelectorTag(ast.Tokens[token], file); err != nil {
				return err
			}

		case aRuleType:
			if err := writeTokenARule(ast.Tokens[token], file); err != nil {
				return err
			}

		case selectorIDType:
			if err := writeTokenSelectorID(ast.Tokens[token], file); err != nil {
				return err
			}

		case selectorAllType:
			if err := writeTokenSelectorAll(ast.Tokens[token], file); err != nil {
				return err
			}

		default:
			return errUnexpectedType(ast.Tokens[token].Type)
		}
		if len(ast.Tokens) == token+1 {
			return nil
		}
		_, err := file.WriteString("\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func writeTokenARule(token *Token, file *os.File) error {
	if token.Name == "@import" || token.Name == "@charset" {
		if len(token.Rules) != 1 {
			return errWrongNumberRules(token.Rules)
		}
		slice := *token.Rules[0]
		newSlice := make([]string, 0)
		for rule := range *token.Rules[0] {
			newSlice = append(newSlice, *slice[rule])
		}
		_, err := file.WriteString(strings.Join(newSlice, " ") + ";\n")
		if err != nil {
			return err
		}
		return nil
	}
	_, err := file.WriteString(token.Name + " {\n")
	if err != nil {
		return err
	}
	for i := range token.Rules {
		_, err := file.WriteString(" ")
		if err != nil {
			return err
		}
		for j := range *token.Rules[i] {
			str := *token.Rules[i]
			newstr := str[j]
			_, err := file.WriteString(" " + *newstr)
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString(";\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}
	return nil
}

func writeTokenSelectorID(token *Token, file *os.File) error {
	_, err := file.WriteString(token.Name + " {\n")
	if err != nil {
		return err
	}
	for i := range token.Rules {
		_, err := file.WriteString(" ")
		if err != nil {
			return err
		}
		for j := range *token.Rules[i] {
			str := *token.Rules[i]
			newstr := str[j]
			_, err := file.WriteString(" " + *newstr)
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString(";\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}
	return nil
}

func writeTokenSelectorClass(token *Token, file *os.File) error {
	_, err := file.WriteString(token.Name + " {\n")
	if err != nil {
		return err
	}
	for i := range token.Rules {
		_, err := file.WriteString(" ")
		if err != nil {
			return err
		}
		for j := range *token.Rules[i] {
			str := *token.Rules[i]
			newstr := str[j]
			_, err := file.WriteString(" " + *newstr)
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString(";\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}
	return nil
}

func writeTokenSelectorAll(token *Token, file *os.File) error {
	_, err := file.WriteString(token.Name + " {\n")
	if err != nil {
		return err
	}
	for i := range token.Rules {
		_, err := file.WriteString(" ")
		if err != nil {
			return err
		}
		for j := range *token.Rules[i] {
			str := *token.Rules[i]
			newstr := str[j]
			_, err := file.WriteString(" " + *newstr)
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString(";\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}
	return nil
}

func writeTokenSelectorTag(token *Token, file *os.File) error {
	_, err := file.WriteString(token.Name + " {\n")
	if err != nil {
		return err
	}
	for i := range token.Rules {
		_, err := file.WriteString(" ")
		if err != nil {
			return err
		}
		for j := range *token.Rules[i] {
			str := *token.Rules[i]
			newstr := str[j]
			_, err := file.WriteString(" " + *newstr)
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString(";\n")
		if err != nil {
			return err
		}
	}
	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}
	return nil
}

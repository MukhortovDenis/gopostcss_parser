package parser

import "testing"

// ================================================import===========================================
var testCaseImport map[int][]byte = map[int][]byte{0: []byte(`@import url("chrome://communicator/skin/");`)}
var testCaseCharset map[int][]byte = map[int][]byte{0: []byte(`@charset "iso-8859-15";`)}
var testCaseMedia []byte = []byte(`@charset "iso-8859-15";`) //Пока не реализовано
var testCaseFontFace map[int][]byte = map[int][]byte{0: []byte("@font-face {"), 1: []byte("font-family: Pompadur;"),
	2: []byte("src: url(fonts/pompadur.ttf);"), 3: []byte("}")}
var testCasePage map[int][]byte = map[int][]byte{0: []byte("@page :first {"), 1: []byte("margin: 1cm;"),
	2: []byte("}")}

// ==================================================================================================

// ================================================selector-id=======================================
var testCaseID map[int][]byte = map[int][]byte{
	0: []byte(`#help {`), 1: []byte(`position: absolute;`),
	2: []byte(`left: 160px;`), 3: []byte(`top: 50px;`),
	4: []byte(`width: 225px;`), 5 : []byte(`padding: 5px;`),
	6: []byte(`background: #f0f0f0;`), 7 : []byte(`}`)}
// ==================================================================================================

// ================================================selector-class====================================
var testCaseClass map[int][]byte = map[int][]byte{
	0: []byte(`p.cite {`), 1: []byte(`color: navy;`),
	2: []byte(`margin-left: 20px;`), 3: []byte(`border-left: 1px solid navy;`),
	4: []byte(`padding-left: 15px;`), 5 : []byte(`}`)}
// ==================================================================================================

// ================================================selector-all======================================
var testCaseAll map[int][]byte = map[int][]byte{
	0: []byte(`* {`), 1: []byte(`margin: 0;`),
	2: []byte(`padding: 0;`), 3: []byte(`}`)}
// ==================================================================================================

// ================================================selector-tag======================================
var testCaseTag map[int][]byte = map[int][]byte{
	7: []byte(`P {`), 8: []byte(`text-align: justify;`),
	9: []byte(`color: green;`), 10: []byte(`}`)}
// ==================================================================================================

func Test_tokenARule(t *testing.T) {
	testCases := []map[int][]byte{testCaseImport, testCaseCharset, testCaseFontFace, testCasePage}
	t.Run("Проверка создания токена с импортом", func(t *testing.T) {
		for ind := range testCases {
			token, index, err := tokenARule(0, testCases[ind])
			if err != nil {
				t.Error(err)
			}
			t.Log(token, "\n", index)
			for i := range token.Rules {
				t.Log(token.Rules[i], "\n")
			}
			t.Log("-----------------------------------------")
		}
	})
}

func Test_tokenSelectorID(t *testing.T) {
	t.Run("Проверка создания токена идентификатора", func(tt *testing.T) {
		token, index, err := tokenSelectorID(0, testCaseID)
		if err != nil {
			t.Error(err)
		}
		t.Log(token, "\n", index)
		for i := range token.Rules {
			t.Log(token.Rules[i], "\n")
		}
		t.Log("-----------------------------------------")
	})
}

func Test_tokenSelectorClass(t *testing.T) {
	t.Run("Проверка создания токена класса", func(tt *testing.T) {
		token, index, err := tokenSelectorClass(0, testCaseClass)
		if err != nil {
			t.Error(err)
		}
		t.Log(token, "\n", index)
		for i := range token.Rules {
			t.Log(token.Rules[i], "\n")
		}
		t.Log("-----------------------------------------")
	})
}

func Test_tokenSelectorAll(t *testing.T) {
	t.Run("Проверка создания токена универсального селектора", func(tt *testing.T) {
		token, index, err := tokenSelectorAll(0, testCaseAll)
		if err != nil {
			t.Error(err)
		}
		t.Log(token, "\n", index)
		for i := range token.Rules {
			t.Log(token.Rules[i], "\n")
		}
		t.Log("-----------------------------------------")
	})
}

func Test_tokenSelectorTag(t *testing.T) {
	t.Run("Проверка создания токена селектора тега", func(tt *testing.T) {
		token, index, err := tokenSelectorTag(7, testCaseTag)
		if err != nil {
			t.Error(err)
		}
		t.Log(token, "\n", index)
		for i := range token.Rules {
			t.Log(token.Rules[i], "\n")
		}
		t.Log("-----------------------------------------")
	})
}

func Test_ParseIntoAST(t *testing.T){
	t.Run("Проверка преобразования CSS в AST", func(tt *testing.T){
		ast, err := ParseIntoAST("stylelint.css")
		if err != nil {
			tt.Error(err)
		}
		t.Log(ast, "\n")
		for i := range ast.Tokens {
			t.Log(ast.Tokens[i], "\n")
			for k := range ast.Tokens[i].Rules{
				t.Log(ast.Tokens[i].Rules[k], "\n")
			}
		}
	})
}
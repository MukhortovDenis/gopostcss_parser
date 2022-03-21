package parser

import "testing"

var testCaseImport map[int][]byte = map[int][]byte{0: []byte(`@import url("chrome://communicator/skin/");`)}
var testCaseCharset map[int][]byte = map[int][]byte{0: []byte(`@charset "iso-8859-15";`)}
var testCaseMedia []byte = []byte(`@charset "iso-8859-15";`) //Пока не реализовано
var testCaseFontFace map[int][]byte = map[int][]byte{0: []byte("@font-face {"), 1: []byte("font-family: Pompadur;"),
	2: []byte("src: url(fonts/pompadur.ttf);"), 3: []byte("}")}
var testCasePage map[int][]byte = map[int][]byte{0: []byte("@page :first {"), 1: []byte("margin: 1cm;"),
	2: []byte("}")}

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

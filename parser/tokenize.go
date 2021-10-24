package main

import (
	"fmt"
	"unicode/utf8"
)

var singleQuote, singleQuoteSize = utf8.DecodeRuneInString("'")
var doubleQuote, doubleQuoteSize = utf8.DecodeRuneInString("\"")
var backslash, backslashSize = utf8.DecodeRuneInString("\\")
var slash, slashSize = utf8.DecodeRuneInString("/")
var newLine, newLineSize = utf8.DecodeRuneInString("\n")
var space, spaceSize = utf8.DecodeRuneInString(" ")
var feed, feedSize = utf8.DecodeRuneInString("\f")
var tab, tabSize = utf8.DecodeRuneInString("\t")
var cr, crSize = utf8.DecodeRuneInString("\r")
var openSquare, openSquareSize = utf8.DecodeRuneInString("[")
var closeSquare, closeSquareSize = utf8.DecodeRuneInString("]")
var openParentheses, openParenthesesSize = utf8.DecodeRuneInString("(")
var closeParentheses, closeParenthesesSize = utf8.DecodeRuneInString(")")
var openCurly, openCurlySize = utf8.DecodeRuneInString("{")
var closeCurly, closeCurlySize = utf8.DecodeRuneInString("}")
var semicolon, semicolonSize = utf8.DecodeRuneInString(";")
var asterisk, asteriskSize = utf8.DecodeRuneInString("*")
var colon, colonSize = utf8.DecodeRuneInString(":")
var at, atSize = utf8.DecodeRuneInString("@")

var position int = 0

// var reAtEnd = /[\t\n\f\r "#'()/;[\\\]{}]/g
// var reWordEnd = /[\t\n\f\r !"#'():;@[\\\]{}]|\/(?=\*)/g
// var reBadBracket = /.[\n"'(/\\]/
// var reHexEscape = /[\da-f]/i

func tokenize() {
	fmt.Println(space)
}

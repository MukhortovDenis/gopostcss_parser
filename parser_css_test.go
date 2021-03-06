package parser

import "testing"

func Test_ParseIntoCSS(t *testing.T) {
	t.Run("Проверка преобразования из AST в CSS", func(tt *testing.T) {
		ast, err := ParseIntoAST("parsetocss.css")
		if err != nil {
			tt.Error(err)
		}
		for i := range ast.Tokens {
			t.Log(ast.Tokens[i], "\n")
			for k := range ast.Tokens[i].Rules {
				t.Log(*ast.Tokens[i].Rules[k])
				for j := range *ast.Tokens[i].Rules[k] {
					slice := *ast.Tokens[i].Rules[k]
					t.Log(*slice[j])
				}
				t.Log("==============================================")
			}
		}
		if err = ParseIntoCSS(ast, "parsetocss.css"); err != nil {
			tt.Error(err)
		}
	})
}

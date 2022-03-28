# gopostcss_parser
Теперь является рабочим проектом для обрабатывания css файлом, предоставляет обЪект AST, с помощью которого можно изменить можно изменить содержание Ast и преобразовать в новый css файл
## Основные функции
### Получение AST
```
func ParseIntoAST(filename string) (*AST, error) ----> 
```
#### Структура AST
```
type AST struct {
	Tokens []*Token
	logger *zap.Logger
}

type Token struct {
	Type  string
	Name  string
	Rules []*Rule
}

type Rule []*string
```
### Получение нового CSS-файла
```
func ParseIntoCSS(ast *AST, filename string)
```
filename - название старого css-файла

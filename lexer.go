package mapq

const (
	TYPE_PLUS      = iota // "+"
	TYPE_SUB              // "-"
	TYPE_MUL              // "*"
	TYPE_DIV              // "/"
	TYPE_LP               // "("
	TYPE_RP               // ")"
	TYPE_VAR              // "([a-z]|[A-Z])([a-z]|[A-Z]|[0-9])*"
	TYPE_RES_TRUE         // "true"
	TYPE_RES_FALSE        // "false"
	TYPE_AND              // "&&"
	TYPE_OR               // "||"
	TYPE_EQ               // "=="
	TYPE_LG               // ">"
	TYPE_SM               // "<"
	TYPE_LEQ              // ">="
	TYPE_SEQ              // "<="
	TYPE_NEQ              // "!="
	TYPE_STR              // a quoted string(单引号)
	TYPE_INT              // an integer
	TYPE_FLOAT            // 小数，x.y这种
	TYPE_UNKNOWN          // 未知的类型
	TYPE_NOT              // "!"
	TYPE_DOT              // "."
	TYPE_RES_NULL         // "null"
)

// Lexer 词法分析器
type Lexer struct {
	input string //输入的文件
	pos   int    //当作一个指针
	runes []rune //不知道什么意思
}

// SetInput 设置输入
func (l *Lexer) SetInput(s string) {
	//panic("not implemented")
	l.input = s
}

// Peek 看下一个字符
func (l *Lexer) Peek() (ch rune, end bool) {
	//ch 返回的类型值，end 不知道什么意思

	//panic("not implemented")
	tem := rune(l.input[l.pos])
	var type0 rune
	//if end&&isLetterOrUnderscore(tem){
	//	type0 = TYPE_VAR
	//}
	if isLetter(tem) {
		type0 = TYPE_VAR
	} else if isNum(tem) {
		type0 = TYPE_INT
	} else if tem == '+' {
		type0 = TYPE_PLUS
	} else if tem == '-' {
		type0 = TYPE_SUB
	} else if tem == '*' {
		type0 = TYPE_MUL
	} else if tem == '/' {
		type0 = TYPE_DIV
	} else if tem == '(' {
		type0 = TYPE_LP
	} else if tem == ')' {
		type0 = TYPE_RP
	} else if tem == '>' {
		type0 = TYPE_LG
	} else if tem == '<' {
		type0 = TYPE_SM
	} else if tem == '!' {
		type0 = TYPE_NOT
	} else if tem == '.' {
		type0 = TYPE_DOT
	} else {
		type0 = TYPE_UNKNOWN
	}
	return type0, true
}

// some finction maybe useful for your implementation
func isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
func isLetterOrUnderscore(ch rune) bool {
	return isLetter(ch) || ch == '_'
}
func isNum(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// Checkpoint 检查点
type Checkpoint struct {
	pos int
}

// SetCheckpoint 设置检查点
func (l *Lexer) SetCheckpoint() Checkpoint {
	panic("not implemented")
}

// GobackTo 回到一个检查点
func (l *Lexer) GobackTo(c Checkpoint) {
	panic("not implemented")
}

// ScanType 扫描一个特定Token，下一个token不是这个类型则自动回退，返回err
func (l *Lexer) ScanType(code int) (token string, err error) {
	panic("not implemented")
}

// Scan scan a token
func (l *Lexer) Scan() (code int, token string, eos bool) {
	panic("not implemented")

}

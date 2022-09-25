package mapq

import (
	"errors"
)

const (
	TYPE_PLUS      = iota //0 "+"
	TYPE_SUB              //1 "-"
	TYPE_MUL              //2 "*"
	TYPE_DIV              //3 "/"
	TYPE_LP               //4 "("
	TYPE_RP               //5 ")"
	TYPE_VAR              //6 "([a-z]|[A-Z])([a-z]|[A-Z]|[0-9])*"
	TYPE_RES_TRUE         //7 "true"
	TYPE_RES_FALSE        //8 "false"
	TYPE_AND              //9 "&&"
	TYPE_OR               //10 "||"
	TYPE_EQ               //11 "=="
	TYPE_LG               //12 ">"
	TYPE_SM               //13 "<"
	TYPE_LEQ              //14 ">="
	TYPE_SEQ              //15 "<="
	TYPE_NEQ              //16 "!="
	TYPE_STR              //17 a quoted string(单引号)
	TYPE_INT              //18 an integer
	TYPE_FLOAT            //19 小数，x.y这种
	TYPE_UNKNOWN          //20 未知的类型
	TYPE_NOT              //21 "!"
	TYPE_DOT              //22 "."
	TYPE_RES_NULL         //23 "null"
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
	//ch val，end 越界判断
	//panic("not implemented")
	var type0 rune
	if l.pos >= len(l.input) {
		return 404, false
	}

	tem := rune(l.input[l.pos])

	//if end&&isLetterOrUnderscore(tem){
	//	type0 = TYPE_VAR
	//}

	if isLetterOrUnderscore(tem) {
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
	} else if tem == '\'' {
		type0 = TYPE_STR
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
	//panic("not implemented")
	ck := Checkpoint{pos: l.pos}
	return ck
}

// GobackTo 回到一个检查点
func (l *Lexer) GobackTo(c Checkpoint) {
	//panic("not implemented")
	l.pos = c.pos
}

// ScanType 扫描一个特定Token，下一个token不是这个类型则自动回退，返回err
func (l *Lexer) ScanType(code int) (token string, err error) {
	//panic("not implemented")
	ck := l.SetCheckpoint()
	type0, tem, epp := l.Scan()
	if !epp {
		return "", errors.New("not found wu")
	}
	if type0 == code {
		return tem, err
	} else {

		l.GobackTo(ck)
		return "", errors.New("not found you")
	}

}

// Scan scan a token
//code:TYPE token:val
func (l *Lexer) Scan() (code int, token string, eos bool) {
	//panic("not implemented")
	tem, end := l.Peek() //tem：类型码	end：是否越界
	l.runes = l.runes[:0]
	eos = true
	if !end {
		return 404, "", eos
	}

	switch tem {
	case TYPE_PLUS:
		l.pos++
		token = "+"
		return TYPE_PLUS, token, eos
	case TYPE_SUB:
		l.pos++
		token = "-"
		return TYPE_SUB, token, eos
	case TYPE_DIV:
		l.pos++
		token = "/"
		return TYPE_DIV, token, eos
	case TYPE_MUL:
		l.pos++
		token = "*"
		return TYPE_MUL, token, eos
	}
	//+ - * /

	if tem == TYPE_INT {
		l.runes = append(l.runes, rune(l.input[l.pos]))
		l.pos++
		dot := 0
		var end1 bool = true
		var tem1 rune
		tem1, end1 = l.Peek()
		for end1 && (tem1 == TYPE_INT || tem1 == TYPE_DOT) {
			tem1, end1 = l.Peek() //获取
			if end1 && (tem1 == TYPE_INT || tem1 == TYPE_DOT) {
				if tem1 == TYPE_DOT {
					dot = 1
				}
				l.runes = append(l.runes, rune(l.input[l.pos]))
			}
			l.pos++
		}
		token = string(l.runes)
		if dot == 1 {
			return TYPE_FLOAT, token, eos
		} else {
			return TYPE_INT, token, eos
		}
	}
	//数字：123456	浮点数：123.456

	if tem == TYPE_STR {
		//l.runes = append(l.runes, rune(l.input[l.pos]))
		l.pos++
		var tem1 rune
		var end1 bool = true
		tem1, end1 = l.Peek()
		for end1 && tem1 != TYPE_STR {
			tem1, end1 = l.Peek()
			if end1 && tem1 != TYPE_STR {
				l.runes = append(l.runes, rune(l.input[l.pos]))
			}
			//if tem1 == TYPE_STR {
			//l.runes = append(l.runes, rune(l.input[l.pos]))
			//}
			l.pos++
		}
		token = string(l.runes)
		return TYPE_STR, token, eos
	}
	//字符串：'46asd'

	if tem == TYPE_VAR {
		l.runes = append(l.runes, rune(l.input[l.pos]))
		l.pos++
		var tem1 rune
		var end1 bool = true
		tem1, end1 = l.Peek()
		for end1 && (tem1 == TYPE_VAR || tem1 == TYPE_INT) {
			tem1, end1 = l.Peek()
			if end1 && (tem1 == TYPE_VAR || tem1 == TYPE_INT) {
				l.runes = append(l.runes, rune(l.input[l.pos]))
			}
			l.pos++
		}
		token = string(l.runes)
		if token == "true" {
			return TYPE_RES_TRUE, token, eos
		} else if token == "false" {
			return TYPE_RES_FALSE, token, eos
		} else if token == "null" {
			return TYPE_RES_NULL, token, eos
		}
		return TYPE_VAR, token, eos
	}
	//标识符：a10	_456asd	true false null

	if tem == TYPE_LG {
		l.runes = append(l.runes, rune(l.input[l.pos]))
		l.pos++
		//var tem1 rune
		var end1 bool = true
		if end1 {
			_, end1 = l.Peek()
			if end1 && l.input[l.pos] == '=' {
				l.runes = append(l.runes, rune(l.input[l.pos]))
				l.pos++
				token = string(l.runes)
				return TYPE_LEQ, token, eos
			} else {
				token = string(l.runes)
				return TYPE_LG, token, eos
			}
		}
	}
	// > OR >=

	if tem == TYPE_SM {
		l.runes = append(l.runes, rune(l.input[l.pos]))
		l.pos++
		//var tem1 rune
		var end1 bool = true
		if end1 {
			_, end1 = l.Peek()
			if end1 && l.input[l.pos] == '=' {
				l.runes = append(l.runes, rune(l.input[l.pos]))
				l.pos++
				token = string(l.runes)
				return TYPE_SEQ, token, eos
			} else {
				token = string(l.runes)
				return TYPE_SM, token, eos
			}
		}
	}
	// < OR <=

	if tem == TYPE_NOT {
		l.runes = append(l.runes, rune(l.input[l.pos]))
		l.pos++
		//var tem1 rune
		var end1 bool = true
		if end1 {
			_, end1 = l.Peek()
			if end1 && l.input[l.pos] == '=' {
				l.runes = append(l.runes, rune(l.input[l.pos]))
				l.pos++
				token = string(l.runes)
				return TYPE_NEQ, token, eos
			} else {
				token = string(l.runes)
				return TYPE_NOT, token, eos
			}
		}
	}
	// ! OR !=

	if l.input[l.pos] == '=' {
		if l.pos+1 < len(l.input) && l.input[l.pos+1] == '=' {
			l.runes = append(l.runes, rune(l.input[l.pos]), rune(l.input[l.pos+1]))
			l.pos += 2
			token = string(l.runes)
			return TYPE_EQ, token, eos
		}
	}
	// ==

	if l.input[l.pos] == '&' {
		if l.pos+1 < len(l.input) && l.input[l.pos+1] == '&' {
			l.runes = append(l.runes, rune(l.input[l.pos]), rune(l.input[l.pos+1]))
			l.pos += 2
			token = string(l.runes)
			return TYPE_AND, token, eos
		}
	}
	// &&

	if l.input[l.pos] == '|' {
		if l.pos+1 < len(l.input) && l.input[l.pos+1] == '|' {
			l.runes = append(l.runes, rune(l.input[l.pos]), rune(l.input[l.pos+1]))
			l.pos += 2
			token = string(l.runes)
			return TYPE_AND, token, eos
		}
	}
	// ||
	eos = false
	return 404, "", eos
}

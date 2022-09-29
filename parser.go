package mapq

// Parser 语法分析器
type Parser struct {
	lexer *Lexer
}

func Getlevel(code int) (level int) {
	//0:	)
	//3:	+ -
	//4:	* /
	//1:	|| &&
	//2: 	== >= <= > < != !
	//5: 	(
	//code := Binnode.Op
	if code == TYPE_PLUS || code == TYPE_SUB {
		return 3
	} else if code == TYPE_MUL || code == TYPE_DIV {
		return 4
	} else if code == TYPE_OR || code == TYPE_AND || code == TYPE_NOT {
		return 1
	} else if code == TYPE_EQ || code == TYPE_LG || code == TYPE_SM || code == TYPE_LEQ || code == TYPE_SEQ || code == TYPE_NEQ {
		return 2
	} else if code == TYPE_LP {
		return 5
	} else if code == TYPE_RP {
		return 0
	}
	return 404

}

func IsOper(token string) bool {
	switch token {
	case "+", "-", "*", "/", "==", ">=", "<=", "!=", ">", "<", "!", "(", ")", "||", "&&":
		return true
	}
	return false
}

func PopNum(Numstack []BinNode) (x, y BinNode) {
	x = Numstack[len(Numstack)-1]
	y = Numstack[len(Numstack)-2]
	//Numstack = Numstack[0 : len(Numstack)-2]
	return x, y
}

func PopBin(Binstack []BinNode) (x BinNode) {
	x = Binstack[len(Binstack)-1]
	//Binstack = Binstack[0 : len(Binstack)-1]
	return x
}

// 你的递归下降分析代码（如果你使用递归下降的话
// func (p *Parser) boolexp() (node Node, err error) {
// 	panic("not implemented")
// }

// func (p *Parser) boolean() (node Node, err error) {
// 	panic("not implemented")
// }
// 别的分析函数
// 。。。。。

// Parse 生成ast
//str:"a+b==3"   类似表达式
//Parse 生成AST树
func (p *Parser) Parse(str string) (n BinNode, err error) {
	//panic("not implemented")
	var l Lexer
	p.lexer = &l
	p.lexer.input = str
	var Numstack []BinNode
	var Binstack []BinNode

	for code, token, eos := p.lexer.Scan(); eos; {
		if code == TYPE_INT || code == TYPE_FLOAT || code == TYPE_VAR || code == TYPE_RES_TRUE ||
			code == TYPE_RES_FALSE || code == TYPE_RES_NULL || code == TYPE_STR {
			Numnode := BinNode{Token: token, Diff: 0, Op: code}
			Numstack = append(Numstack, Numnode)
		}

		//code 为 +-*/ () && || == > < >= <= != !
		if IsOper(token) {
			if len(Binstack) == 0 || Getlevel(code) >= Getlevel(Binstack[len(Binstack)-1].Op) { //如果当前token的运算优先级大，则直接入栈
				var Binnode BinNode
				if token == "!" {
					Binnode = BinNode{Op: code, Token: token, Diff: 1}
				} else {
					Binnode = BinNode{Op: code, Token: token, Diff: 2}
				}
				Binstack = append(Binstack, Binnode)

			}
			if Getlevel(code) < Getlevel(Binstack[len(Binstack)-1].Op) {
				for len(Binstack) > 0 && Getlevel(code) < Getlevel(Binstack[len(Binstack)-1].Op) { //若小，则栈顶元素pop，左右节点为Num栈的顶二元素，同时Bin1再次入Num栈
					Bin1 := PopBin(Binstack)
					Binstack = Binstack[0 : len(Binstack)-1]
					if Bin1.Token == "(" || Bin1.Token == ")" {
						continue
					}
					if Bin1.Token == "!" {
						Num1 := Numstack[len(Numstack)-1]
						Numstack = Numstack[0 : len(Numstack)-1]
						Bin1.Left = &Num1
						Numstack = append(Numstack, Bin1)

					} else {
						Num1, Num2 := PopNum(Numstack) //直达Bin栈栈顶元素优先级小于token
						Numstack = Numstack[0 : len(Numstack)-2]
						Bin1.Left = &Num2
						Bin1.Right = &Num1
						Numstack = append(Numstack, Bin1)
					}

				}
				var Binnode BinNode
				if token == "!" {
					Binnode = BinNode{Op: code, Token: token, Diff: 1}
				} else {
					Binnode = BinNode{Op: code, Token: token, Diff: 2}
				}
				Binstack = append(Binstack, Binnode) ////已经小于，入栈
			}
		}
		code, token, eos = p.lexer.Scan()
	}

	for len(Binstack) > 0 {
		if Binstack[len(Binstack)-1].Diff == 1 {
			Bin := PopBin(Binstack)
			Binstack = Binstack[0 : len(Binstack)-1]

			Num1 := Numstack[len(Numstack)-1]
			Numstack = Numstack[0 : len(Numstack)-1]

			Bin.Left = &Num1
			Numstack = append(Numstack, Bin)
		} else {
			Bin1 := PopBin(Binstack)
			Binstack = Binstack[0 : len(Binstack)-1]
			if Bin1.Token == ")" {
				continue
			}

			Num1, Num2 := PopNum(Numstack)
			Numstack = Numstack[0 : len(Numstack)-2]

			Bin1.Left = &Num2
			Bin1.Right = &Num1
			Numstack = append(Numstack, Bin1)
		}
	}
	return Numstack[0], nil
}

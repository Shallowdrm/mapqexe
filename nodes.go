package mapq

import (
	"reflect"
	"strconv"
)

// Node 节点
type Node interface {
	Eval(data map[string]interface{}) interface{}
}

// 这些函数提供给你，也许可以帮上忙。。。
func toF64(i interface{}) float64 {
	switch v := i.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	}
	return 0
}
func trytoF64(i interface{}) interface{} {
	switch v := i.(type) { //类型
	case int, float64, int32, int64, uint32, uint64, float32:
		return toF64(v)

	}
	return i
}

func equal(left, right interface{}) bool {
	return reflect.DeepEqual(trytoF64(left), trytoF64(right))
}

// BinNode 双目表达式节点
type BinNode struct {
	Left, Right *BinNode //
	Op          int      //operation
	Token       string
	Diff        int //0 1 2表示有的孩子数量
}

// Eval 查询
//data:"a": 1, "b": 2
//true:1	false:0		null:-1

func (n *BinNode) Eval(data map[string]interface{}) interface{} {
	switch n.Diff {
	case 2: //operation
		{
			switch n.Token {
			case "==":
				if n.Left.Eval(data) == n.Right.Eval(data) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case "&&":
				if (n.Left.Eval(data).(float64) > 0) && (n.Right.Eval(data).(float64) > 0) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case "||":
				if (n.Left.Eval(data).(float64) > 0) || (n.Right.Eval(data).(float64) > 0) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case ">=":
				if n.Left.Eval(data).(float64) >= n.Right.Eval(data).(float64) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case "<=":
				if n.Left.Eval(data).(float64) <= n.Right.Eval(data).(float64) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case ">":
				if n.Left.Eval(data).(float64) > n.Right.Eval(data).(float64) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case "<":
				if n.Left.Eval(data).(float64) < n.Right.Eval(data).(float64) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case "!=":
				if n.Left.Eval(data) != n.Right.Eval(data) {
					return toF64(1)
				} else {
					return toF64(0)
				}
			case "+":
				return n.Left.Eval(data).(float64) + n.Right.Eval(data).(float64)
			case "-":
				return n.Left.Eval(data).(float64) - n.Right.Eval(data).(float64)
			case "/":
				return n.Left.Eval(data).(float64) / n.Right.Eval(data).(float64)
			case "*":
				return n.Left.Eval(data).(float64) * n.Right.Eval(data).(float64)
			}
		}
	case 1:
		{ //!
			if n.Left.Eval(data).(float64) > 0 {
				return toF64(0)
			} else {
				return toF64(1)
			}
		}
	case 0:
		{
			switch n.Op {
			case TYPE_INT, TYPE_FLOAT:
				num, _ := strconv.ParseFloat(n.Token, 64)
				return num
			case TYPE_RES_TRUE:
				return toF64(1)
			case TYPE_RES_FALSE:
				return toF64(0)
			case TYPE_RES_NULL:
				return toF64(-1)
			case TYPE_VAR:
				elem, ok := data[n.Token]
				if ok {
					return float64(elem.(int))
				} else {
					return toF64(-1)
				}
			}

		}
	}
	return 404
}

// 别的节点。。。。

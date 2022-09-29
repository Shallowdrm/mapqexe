package mapq

// QueryMap 查询map

//data:"a": 1, "b": 2	query:"a+b==3"
//QueryMap：bool返回对表达式的判断
func QueryMap(data map[string]interface{}, query string) (bool, error) {
	p := &Parser{}
	n, err := p.Parse(query)
	if err != nil {
		return false, err
	}
	if n.Eval(data).(float64) > 0 {
		return true, nil
	} else {
		return false, nil
	}
	//return n.Eval(data).(bool), nil
}

// RunQuery 查询
func RunQuery(root BinNode, data map[string]interface{}) (bool, error) {
	if root.Eval(data).(float64) > 0 {
		return true, nil
	} else {
		return false, nil
	}
	//return root.Eval(data).(bool), nil
}

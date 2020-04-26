package a_way_to_export_symbols

// 这个文件 inner.go 就是普通的 go 程序文件，这些变量和类型都没有导出
// 实际上，除了 exports.go ，其他文件中，都不应该有导出的变量和类型

type inner_typ_a int
type inner_typ_b int

func inner_fn_a(){}
func inner_fn_b(){}

var inner_var_a = 1
var inner_var_b = 1

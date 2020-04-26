package a_way_to_export_symbols

// 这个文件，exports.go，用于创建别名，这些别名会被导出，并且指向包内部的变量或类型
// 当我们需要接手新的项目时，阅读 exports.go 文件，就可以很清晰地明了这个包的工作，相当于 API 文档一般

type (
	TYP_a = inner_typ_a
	TYP_b = inner_typ_b
)

var (
	FN_a = inner_fn_a
	FN_b = inner_fn_b
)

var (
	VAR_a = inner_var_a
	VAR_b = inner_var_b
)

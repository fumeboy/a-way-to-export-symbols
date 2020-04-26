# 一个导出包内符号的合适的方法
用于 golang 工程

这个方法我以前也偶尔会用，但今天又想到 golang 奇妙的导出方式，所以决定为这个方法写一篇文章。

## 背景

在小型 golang 项目中，随意地使用大写字母开头的命名来导出符号是很轻松的，没有太大的阅读包袱。

但是当项目变得很大时，大写字母开头的命名则很难做好视觉上或其他方面上的提醒，在当项目转手时，它对于新程序员的上手是个不小的折磨。

这个时候我们需要一个好的方法来进行包内符号的导出

## 方法内容

很简单，没有什么魔法。

一句话概括： 把 包内符号的导出 放在特定的文件中进行。

假如 我有一个包，它是 pkg_example，这个包里我有这些符号需要导出：

```text
// inner.go
type inner_typ_a int
type inner_typ_b int

func inner_fn_a(){}
func inner_fn_b(){}

var inner_var_a = 1
var inner_var_b = 1
```

但是我并不直接修改它们的名字，而是给 pkg_example 创建一个特定的文件：exports.go

在这个文件里，创建这些符号的别名，这些别名是被导出的：

```text
// exports.go
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

```

就是这么简单

## 注意

仅当在给变量做别名时，需要额外的考虑：

`var VAR_a = inner_var_a` 这句赋值发生时，inner_var_a 是否已经初始化了？

若没有初始化，VAR_a 得到的是一个“空值”而不是初始化后的 inner_var_a
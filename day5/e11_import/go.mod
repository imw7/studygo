module import_demo

go 1.21

// require 格式必须是这样的，路径和需要导入的包的go.mod里面的一样
require "imw7.com/studygo/calc" v0.0.0
// replace 格式必须是这样的，最后的包名必须是存在的且不能修改成别名
replace "imw7.com/studygo/calc" => "../e10_calc"

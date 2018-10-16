// 每个源文件必须指定此文件所处的包（package）。
// 这里的包名为main（程序入口所处的包名必须为main）。
package main
// 引入一个标准库fmt（format的缩写）。
// 下面要用其中的Print和Println函数。
import "fmt"
// 定义一个自定义类型Person，它有3个成员变量。
// type和struct为关键字。
// 注意，在Golang变量声明中，变量名在前，类型在后。
type Person struct {
Name string // 字符串类型（内置类型）
Age int // 整数类型（内置类型）
IsGirl bool // 布尔类型（内置类型）
}
// 这是一个自定义函数PrintPerson，打印出一个Person变量的信息。
// func为关键字，是function的缩写。
// p为一个Person类型的输入变量。
// 输入变量p的声明也是变量名在前，类型在后。
func PrintPerson(p Person) {222
// 类型Person的各个成员变量对应着p的各个属性。
// 输入变量p的属性可以用p.Name、p.Age和p.IsGirl来表示。
// Print是fmt包中的一个函数，支持任意个数的参数。
// 所有的参数将被逐个用字符串的形式打印到标准输出。
// 下面的Println和Print的区别是输出各项将用空格分开
// 并在最后多输出一个换行。
fmt.Print(p.Name, "今年", p.Age, "岁了，")
// 常见的if-else流程控制语句。
// if后的表达式不需用()括起来。
if p.IsGirl {
fmt.Println("她是个女娃。")
} else {
fmt.Println("他是个男娃。")
}
}
// 这是程序入口，和其它很多语言非常类似，名称也是main。
func main() {
// 定义一个局部变量。
// var为关键字，是variable的缩写。
var huahua = Person {
Name: "花花",
Age: 17,
IsGirl: true, // 这里的逗号不可少
}
// 定义另外一个局部变量。
// 注意这里使用了和上面不一样的定义风格。
// 如果使用:=来声明变量，则不能前置var关键字。
xiaoming := Person {
Name: "小明",
Age: 18,
IsGirl: false,
}
// 调用自定义函数，输出信息。
PrintPerson(huahua)
PrintPerson(xiaoming)
}
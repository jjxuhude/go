package main

// 引入其它包
import (
	"demo/simple"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 定义一个用于打印程序使用指南的函数
var Usage = func() {
	fmt.Println("USAGE: demo command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}

// 程序入口函数
func main() {
	/*
	 * 用于获取命令行参数，注意程序名本身是第一个参数，
	 * 比如 calc add 1 2 这条指令，第一个参数是 calc
	 */
	args := os.Args
	// 除程序名本身外，至少需要传入两个其它参数，否则退出
	if args == nil || len(args) < 2 {
		error := errors.New("参数不能少于2个")
		fmt.Println(error)
		return
		//var myarr = [3]int{1, 2, 3}
		//fmt.Println(interface{}(myarr))
		//value, ok := interface{}(myarr).([3]int)
		//fmt.Println(value,ok)

		//=======================
		//num2是否属于接口Number1
		var num1 simple.Number = 1
		var num2 simple.Number2 = &num1
		if num3, ok := num2.(simple.Number1); ok {
			fmt.Println(num3.Equal(1))
		}
		return
	}

	// 第二个参数表示计算方法
	switch args[1] {
	// 如果是加法的话
	case "add":
		// 至少需要包含四个参数
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 获取待相加的数值，并将类型转化为整型
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		// 获取参数出错，则退出
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 从 simple 包引入 Add 方法进行加法计算
		ret := simple.Add(v1, v2)
		// 打印计算结果
		fmt.Println("Result: ", ret)
	// 如果是计算平方根的话
	case "sqrt":
		// 至少需要包含三个参数
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 获取待计算平方根的数值，并将类型转化为整型
		v, err := strconv.Atoi(args[2])
		// 获取参数出错，则退出
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 从 simple 包引入 Sqrt 方法进行平方根计算
		ret := simple.Sqrt(v)
		// 打印计算结果
		fmt.Println("Result: ", ret)
	// 如果计算方法不支持，打印程序使用指南
	case "string":
		name := os.Args[2]
		result := simple.Substr(name)
		fmt.Printf("%c,%c,%c", result[0], result[1], result[2])
	case "slice":
		slide := simple.Slice()
		fmt.Println(slide)

	case "map":
		simple.Map()

	case "ptr":
		a := 1
		b := 2
		simple.Ptr(&a, &b)
		fmt.Println(a, b)

	case "interface":
		simple.Interface(1, "aaa", true, [1]int{123}, []string{"aa", "bbb"}, map[string]string{"name": "jack"})

	case "struct":
		student := simple.NewStudent(1, "jack", true, 99)
		student.SetName("徐华德")
		fmt.Println(student)
	case "go":

	default:
		Usage()
	}

}

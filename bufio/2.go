package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
)

func file() {
	file, _ := os.OpenFile("errr.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	//buf := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(file)
	w.WriteString("你好呀中国\n")
	w.Flush() // 不要忘记刷新
	//fmt.Printf("%s",buf.string());

}

func exce() {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func csv1() {
	file, _ := os.OpenFile("product.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	file.WriteString("\xEF\xBB\xBF")
	defer file.Close()
	w := csv.NewWriter(file)
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	item := []string{"222", "3333", "4444"}
	records = append(records, item)

	fmt.Println(records)

	w.Write([]string{"ID", "名称", "大小"})
	w.WriteAll(records)
	w.Flush()

}

func main() {
	//file();
	//exce()
	csv1()
}

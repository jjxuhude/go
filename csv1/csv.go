package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "请输入路径")
		os.Exit(0)
	}

	imagePath := os.Args[1]

	imagePath, error := filepath.Abs(imagePath)
	if error == nil {
		imagePath += "\\*"
		skuArr, _ := filepath.Glob(imagePath)
		filename := "product.csv"
		os.Remove(filename)
		var fd_content [][]string
		fd_content = [][]string{{"SKU", "NAME", "DESC", "TYPE", "PRICE", "STOCK", "LEVEL", "IMAGES"}}
		for _, sku := range skuArr {
			id, images := readImages(sku)
			if images != "" {
				item := []string{id, "", "", "", "", "", "", images}
				fmt.Println(item)
				fd_content = append(fd_content, item)
			}
		}
		writeCsv(fd_content)

	} else {
		fmt.Println("路径错误")
	}

}

func writeCsv(records [][]string) {
	file, _ := os.OpenFile("product.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	//file.WriteString("\xEF\xBB\xBF")
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(records)
	w.Flush()
}

func readImages(imagepath string) (string, string) {
	imagepath += "\\*"
	images, _ := filepath.Glob(imagepath)

	old, _ := os.Getwd()
	var sku string
	var filesSlice = make([]string, 0)
	for i, image := range images {
		sku, _ = filepath.Split(image)
		sku = filepath.Base(sku)
		fi, _ := os.Stat(image)
		images[i] = strings.Replace(image, old, "", -1)
		images[i] = strings.Replace(images[i], "\\", "/", -1)

		if !fi.IsDir() {
			filesSlice = append(filesSlice, images[i])
		}
	}
	//fmt.Println(filesSlice)
	str := strings.Join(filesSlice, `;`)
	return sku, str
}

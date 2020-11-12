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
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("open file is failed, err: ", err)
		}
		defer file.Close()
		file.WriteString("\xEF\xBB\xBF")
		w := csv.NewWriter(file)
		var fd_content string
		fd_content = strings.Join([]string{"SKU", "IMAGES"}, ",") + "\n"
		for _, sku := range skuArr {
			id, images := readImages(sku)
			if images != "" {
				fd_content += strings.Join([]string{id, images}, ",") + "\n"
				fmt.Println(id, images)
			}
		}
		//fmt.Println(fd_content)
		buf := []byte(fd_content)
		file.Write(buf)
		w.Flush()

	} else {
		fmt.Println("路径错误")
	}

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
		if !fi.IsDir() {
			filesSlice = append(filesSlice, images[i])
		}
	}
	//fmt.Println(filesSlice)
	str := strings.Join(filesSlice, `;`)
	return sku, str
}

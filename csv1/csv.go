package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type fileSort struct {
	file string
	name int
}

func (f fileSort) String() string {
	return fmt.Sprintf("%s: %d", f.file, f.name)
}

type ByName []fileSort

func (f ByName) Len() int           { return len(f) }
func (f ByName) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByName) Less(i, j int) bool { return f[i].name < f[j].name }

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
	sort.Strings(images)

	var fileSortArr = make([]fileSort, 0)
	for _, fullfilename := range images {
		fielname := filepath.Base(fullfilename)
		filesuffix := path.Ext(fielname)
		fileprefix := fielname[0 : len(fielname)-len(filesuffix)]
		name, err := strconv.Atoi(fileprefix)
		if err != nil {
			fmt.Printf("文件名必须是数字：%s", fullfilename)
			os.Exit(0)
		}
		fileSortArr = append(fileSortArr, fileSort{fullfilename, name})
	}
	sort.Sort(ByName(fileSortArr))
	images = images[0:0]
	for _, fileSortIns := range fileSortArr {
		images = append(images, fileSortIns.file)
	}
	fmt.Println(images)

	old, _ := os.Getwd()
	var sku string
	var filesSlice = make([]string, 0)
	for i, image := range images {
		sku, _ = filepath.Split(image)
		sku = filepath.Base(sku)
		//fi, _ := os.Stat(image)
		images[i] = strings.Replace(image, old, "", -1)
		images[i] = strings.Replace(images[i], "\\", "/", -1)

		//if !fi.IsDir() {
		filesSlice = append(filesSlice, images[i])
		//}
	}
	//fmt.Println(filesSlice)
	str := strings.Join(filesSlice, `;`)
	return sku, str
}

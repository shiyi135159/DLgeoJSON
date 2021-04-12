package main

import (
	"DLgeoJSON/tool"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func DownloadGEOJSON(ver string) {
	f, err := excelize.OpenFile("cc_adcode.xlsx")
	if err != nil {
		fmt.Println("读取Excel发生异常：", err)
	}
	// Get all the rows in the Sheet "cc_adcode"
	rows, err := f.GetRows("cc_adcode")
	total := 0
	for i, row := range rows {
		code := ""
		name := ""
		level := 0
		if i > 0 {
			for idx, colCell := range row {
				if idx == 0 {
					code = colCell
				} else if idx == 1 {
					name = colCell
				} else if idx == 2 {
					level, _ = strconv.Atoi(colCell)
				}
			}
			url := ""
			fullFlag := ""
			if level < 3 {
				fullFlag = "_full"
			}
			url = fmt.Sprintf("https://geo.datav.aliyun.com/areas_v2/bound/%s%s.json", code, fullFlag)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("下载请求发生异常：", err)
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			// Save json to file
			os.MkdirAll("./geoJSON", os.ModePerm)
			f, err := os.Create(fmt.Sprintf("./geoJSON/%s.json", code))
			if err != nil {
				fmt.Println("geoJSON文件保存发生异常：", err.Error())
			}
			_, err = f.Write([]byte(bytes.NewBuffer(body).String()))
			defer f.Close()
			fmt.Println(fmt.Sprintf("√ %s %s", code, name))
			total++
		}
	}
	consoleTitle := fmt.Sprintf("geoJSON数据下载工具 v%s 数据下载完毕！", ver)
	tool.SetConsoleTitle(consoleTitle)
	fmt.Println()
	fmt.Println("> geoJSON已全部下载完成(", total, ") :)")
}

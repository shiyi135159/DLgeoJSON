package main

import (
	"fmt"
	"time"

	"DLgeoJSON/tool"
)

func main() {
	const VER = "1.0"
	consoleTitle := fmt.Sprintf("geoJSON数据下载工具 v%s 正在运行 - 请勿关闭！", VER)
	tool.SetConsoleTitle(consoleTitle)
	go DownloadGEOJSON(VER)
	for {
		time.Sleep(time.Minute)
	}
}

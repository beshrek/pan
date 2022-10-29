package main

import (
	"fmt"

	"github.com/beshrek/pan/file"
)

func main() {
	accessToken := "121.dc9cb720b7acca907f353df94d1c8a8d.YgobnAZMu1SAW5e0JMv4bAV6vERmGjONyI1dfAA.2mMfzA"
	path := "/apps/书梯/CHSS.mkv"
	localFilePath := `d:\tmp\20221026444-4套798店伊伊-0000.mp4`
	fileUploader := file.NewUploader(accessToken, path, localFilePath)
	res, err := fileUploader.Upload()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(res)
}

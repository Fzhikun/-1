package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件，准备写入
	file, err := os.Create("ninenine.txt")
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	// 生成99乘法表并写入文件
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			_, err := file.WriteString(fmt.Sprintf("%d*%d=%d\t", j, i, i*j))
			if err != nil {
				fmt.Println("写入文件时出错:", err)
				return
			}
		}
		_, err = file.WriteString("\n")
		if err != nil {
			fmt.Println("写入文件时出错:", err)
			return
		}
	}

	fmt.Println("99乘法表已成功写入ninenine.txt文件")
}

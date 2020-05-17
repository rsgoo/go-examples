package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	//打开源文件 && 目标文件
	//创建缓存读取器 && 写出器
	//分批次读取数据 并写入目标文件
	src, err1 := os.OpenFile("music.mp3", os.O_RDONLY, 0666)
	defer src.Close()

	if err1 != nil {
		fmt.Println("err1 is ", err1)
		return
	}

	dst, err2 := os.OpenFile("new.mp3", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer dst.Close()

	if err2 != nil {
		fmt.Println("err2 is ", err2)
		return
	}

	//源文件缓冲读取器
	srcReader := bufio.NewReader(src)

	//目标文件写出器
	dstWriter := bufio.NewWriter(dst)

	var buffer = make([]byte, 2048)
	for {
		//读取指定大小缓冲读取器数据
		n, err3 := srcReader.Read(buffer)
		if n == 0 {
			fmt.Println("行读取完成")
			break
		}
		if err3 != nil && err3 == io.EOF {
			fmt.Println("err is ", err3)
			break;
		}

		//将读取到的缓冲读取器数据写入到写出器中
		_, err4 := dstWriter.Write(buffer)
		if err4 != nil {
			fmt.Println("err4 is ", err4)
		}
	}
	fmt.Println("拷贝完成")

}

package main

import (

	"fmt"
	"os"
)

func main() {

	// f, err := os.Open("example.txt")
	// if err != nil {

	// 	//log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {

	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder",fileInfo.IsDir())
	// fmt.Println("file permisson",fileInfo.Mode())
	// fmt.Println("file size",fileInfo.Size())
	// fmt.Println("file modification time",fileInfo.ModTime())

	//read file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 10)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	//panic the error
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {

	// 	println("data:", d, string(buf[i]))

	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	//read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(dir)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(),fi.IsDir())
	// }

	//create files
	// f, err := os.Create("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()
	// f.WriteString("Hi, Golang")
	// f.WriteString("Nice Language")

	//read and write to another file(streaming fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// //destination file to copy the data
	// destFile, err := os.Create("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// //create reader and writer
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	f, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {

	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(f)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush()
	// fmt.Println("write data in new file successfully!")

	//delete file

	err:=os.Remove("test.txt")

	if err!= nil{
		panic(err)
	}

	fmt.Println("file deleted successfully!")

}

package main

import (
	"bufio"
	"fmt"
	"go_study/8.google_pipeline/pipeline"
	"os"
	"strconv"
)

//需要用pipelinedemon/main.go生成
const in = "/tmp/small.in"
const out = "/tmp/small.out"

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	pipeline.WriterSink(writer, p)
}

func createPipeline(filename string, filesize, chunkCount int) <-chan int {
	chunkSize := filesize / chunkCount
	pipeline.Init()
	//sortResults的类型是<-chan int类型的slice
	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemorySort(source))
	}

	return pipeline.MergeN(sortResults...)
}

func createNetworkPipeline(filename string, filesize, chunkCount int) <-chan int {
	chunkSize := filesize / chunkCount
	pipeline.Init()

	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		//tcp server
		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemorySort(source))
		sortAddr = append(sortAddr, addr)
	}

	//tcp client
	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResults...)
}

//测试单机版外部多路归并排序
func test1(in, out string, filesize, chunkcount int) {
	fileInfo, err := os.Stat(in)
	if err != nil {
		panic(err)
	}

	sz := fileInfo.Size()
	//fmt.Println(sz)//返回的是字节

	if int64(filesize) != sz {
		fmt.Println("file size not allow")
		return
	}

	p := createPipeline(in, filesize, chunkcount)
	writeToFile(p, out)
	printFile(out)
}

//测试网络版外部归并排序
func test2(in, out string, filesize, chunkcount int) {
	fileInfo, err := os.Stat(in)
	if err != nil {
		panic(err)
	}

	sz := fileInfo.Size()
	//fmt.Println(sz)//返回的是字节

	if int64(filesize) != sz {
		fmt.Println("file size not allow")
		return
	}

	p := createNetworkPipeline(in, filesize, chunkcount)
	writeToFile(p, out)
	printFile(out)
}

//实现外部排序
func main() {
	//测试单机版外部多路归并排序

	//512字节是生成64个整型时候/tmp/small.in的大小
	//512字节大小的文件分成了4块
	//test1("/tmp/small.in", "/tmp/small.out",512,4)

	//test1("/tmp/large.in", "/tmp/large.out", 800000000, 4)

	//==========================================================================
	//==========================================================================

	//测试网络版外部归并排序
	//test2("/tmp/small.in", "/tmp/small.out",512,4)
	test2("/tmp/large.in", "/tmp/large.out", 800000000, 4)
}

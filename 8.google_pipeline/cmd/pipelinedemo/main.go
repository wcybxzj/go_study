package main

import (
	"bufio"
	"fmt"
	"go_study/8.google_pipeline/pipeline"
	"os"
)

const n = 100000000 //生成的整数的数量,最终生成的文件大小,100000000数据量*8字节=800000000字节
//const n = 64 //64*8=512字节

//生成随机数大文件版本1:(太慢了)
//原因是都是系统调用，61%的时间都是在系统调用write
/*
strace -c -p 10336
% time     seconds  usecs/call     calls    errors syscall
------ ----------- ----------- --------- --------- ----------------
61.30    2.361628           9    264381           write
37.91    1.460611          10    139379      6743 futex
0.79    0.030322          12      2548           pselect6
0.00    0.000003           3         1           sched_yield
------ ----------- ----------- --------- --------- ----------------
100.00    3.852564                406309      6743 total
*/
func createbigfile_version1(filename string, n int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.RandmonSource(n)
	pipeline.WriterSink(file, p)
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReaderSource(file, -1)
	//只打印前100个数据
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

//生成随机数大文件版本2:(超级快)
/*
root@ybx-ThinkPad-T420:~/www/go_www/src/go_study# strace -c -p 11084
strace: Process 11084 attached
% time     seconds  usecs/call     calls    errors syscall
------ ----------- ----------- --------- --------- ----------------
96.89    0.530920          60      8873      3567 futex
2.54    0.013896           3      4546           pselect6
0.57    0.003150           5       665           write
0.00    0.000005           3         2           sched_yield
0.00    0.000001           1         2           epoll_pwait
------ ----------- ----------- --------- --------- ----------------
100.00    0.547972                 14088      3567 total
*/
func createbigfile_version2(filename string, n int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.RandmonSource(n)

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	//只打印前100个数据
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func test1() {
	p1 := pipeline.InMemorySort(pipeline.ArraySource(3, 2, 6, 7, 4))
	p2 := pipeline.InMemorySort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8))
	p := pipeline.Merge(p1, p2)
	for v := range p {
		fmt.Println(v)
	}
}

func main() {
	//test1()
	//fmt.Println("==============================")

	//createbigfile_version1("/tmp/small.in", 100000000)
	//createbigfile_version2("/tmp/small.in", 64)

	fmt.Println("==============================")
	createbigfile_version2("/tmp/large.in", 100000000)
}

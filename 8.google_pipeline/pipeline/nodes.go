package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

const useChannelCache = true

func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

var startTime time.Time

func Init() {
	startTime = time.Now()
}

//内排节点:
func InMemorySort(in <-chan int) <-chan int {
	var out chan int
	if useChannelCache {
		out = make(chan int, 1024)
	} else {
		out = make(chan int)
	}

	go func() {
		//read into memory:
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))

		//sort:
		sort.Ints(a)
		fmt.Println("InMemorySort done:", time.Now().Sub(startTime))

		//output:
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//归并节点:
//1.功能:
//从2个channel获取数据,可能一个channel有数据一个channel没数据

//2.必须close：
//否则main中读取这个channel会死锁

//3.当从in1和in2读取到数据的时候:
//merge协程阻塞等待到,另外2个内排协程完成工作
func Merge(in1, in2 <-chan int) <-chan int {
	var out chan int
	if useChannelCache {
		out = make(chan int, 1024)
	} else {
		out = make(chan int)
	}

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()
	return out
}

//reader节点:
//从文件读取数据,并且每个reader读取一部分
//chunkSize:
//一块文件的大小
//设置读取的个数, -1:读取个数不限制
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	var out chan int
	if useChannelCache {
		out = make(chan int, 1024)
	} else {
		out = make(chan int)
	}

	go func() {
		buffer := make([]byte, 8) //golang int 是64bit
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					panic(err)
				}
			}
			//fmt.Println("Read n:", n)//输出应该是8

			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			//结束条件
			//1.出错
			//2.读取到了制定的的块数
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

//sink节点:
func WriterSink(write io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)                     //golang int 是64bit
		binary.BigEndian.PutUint64(buffer, uint64(v)) //用大边,将v存入内存
		write.Write(buffer)
	}
}

func RandmonSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2

	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...),
	)
}

package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

//receiver是box指针
func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE) // int(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

//给Color加了一个String功能
//Color本来是int类型，让int可以打印值对应的颜色
func (c Color) String() string {
	strings := []string{"white", "black", "blue", "red", "yellow"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("box number is %d\n", len(boxes))
	fmt.Printf("第一个箱子的体积:%f\n", boxes[0].Volume())
	fmt.Printf("最后一个箱子的颜色是:%f\n", boxes[2].Volume())
	fmt.Println("最大尺寸箱子的颜色是:", boxes.BiggestColor().String())

	boxes.PaintItBlack()
	fmt.Println("最大尺寸箱子的颜色是:", boxes.BiggestColor().String())

	fmt.Printf("第一个箱子的颜色%s\n", boxes[0].color)

}

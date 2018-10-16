package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	var row, col int
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Fscanf(file, "%d %d", &row, &col)

	//[row][col]int
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int //row col
}

//{row, col}
var dirs = [4]point{
	{-1, 0}, //up
	{0, -1}, //left
	{1, 0},  //down
	{0, 1},  //right
}

//point + point, return new point
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//判断点在一个二维的slice中的位置是否越界
//越界:返回 0,false
//没越界:返回 所在位置的值,true
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

//返回的是走过的路径steps
func walk(maze [][]int, start, end point) [][]int {
	//steps是你的足迹, 初始化时候全是0
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//填充初始位置到队列{0, 0}
	Q := []point{start}
	//退出条件:
	//1:走到终点end
	//2.队列为空无路可走
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			//发现到终点了
			if cur == end {
				break
			}

			//获取next
			next := cur.add(dir)

			//在maze中next是会撞墙
			val, ok := next.at(maze)
			if !ok || val == 1 { //如果越界或者撞墙,跳过看下一个点
				continue
			}

			//在steps中next是否走过
			val, ok = next.at(steps)
			if !ok || val != 0 { //如果越界或者之前走过,跳过看下一个点
				continue
			}

			//next是否等于start,说明退回到起点了
			if next == start {
				continue
			}

			//获取当前步数
			curSteps, _ := cur.at(steps)
			//给steps中的next填充当前步数
			steps[next.i][next.j] = curSteps + 1
			//将next入队
			Q = append(Q, next)
		}
	}
	return steps
}

//根据正向走迷宫,反向走一遍迷宫 从end走向start
func backWalk(maze, steps [][]int, end, start point) {
	cur := end
	curVal := steps[cur.i][cur.j]
	nextVal := curVal - 1

	fmt.Printf("end point i:%d, j:%d, step val:%d\n",
		cur.i, cur.j, steps[cur.i][cur.j])

END:
	for {
		for _, dir := range dirs {
			if cur == start {
				fmt.Printf("start point i:%d, j:%d, step val:%d\n",
					cur.i, cur.j, steps[cur.i][cur.j])
				break END
			}

			next := cur.add(dir)

			//在maze中next是否会撞墙
			val, ok := next.at(maze)
			if !ok || val == 1 { //如果越界或者撞墙,跳过看下一个点
				continue
			}

			if steps[next.i][next.j] == nextVal {
				fmt.Printf("middle point i:%d j:%d step val:%d\n",
					next.i, next.j, steps[next.i][next.j])

				cur = next
				curVal = steps[cur.i][cur.j]
				nextVal = curVal - 1
				break
			}

		}
	}
}

func main() {
	//maze是迷宫
	//maze是5行,6列
	maze := readMaze("maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
	fmt.Println()

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}

	steps := walk(maze, start, end)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}

	backWalk(maze, steps, end, start)
}

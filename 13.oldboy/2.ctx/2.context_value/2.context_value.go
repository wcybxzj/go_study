package __context_value

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("abc").(int)
	if !ok {
		ret = 21342423
	}
	fmt.Printf("abc:%d\n", ret)

	ret, _ = ctx.Value("def").(int)
	fmt.Printf("def:%d\n", ret)
}

//输出:
//abc:21342423
//def:456
func test2_1() {
	ctx := context.WithValue(context.Background(), "abc", 13483434)
	ctx = context.WithValue(context.Background(), "def", 456)

	process(ctx)
}

//context上下文传值
func main() {
	test2_1()
}

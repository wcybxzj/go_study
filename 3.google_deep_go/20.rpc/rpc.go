package rpcdemon

import "github.com/pkg/errors"

//服务名称
type DemonService struct{}

type Args struct {
	A, B int
}

//服务方法:
//args:输入
//result:输出
func (DemonService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}

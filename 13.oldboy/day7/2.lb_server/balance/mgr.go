package balance

import "fmt"

//对所有LB算法进行管理
type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalanceMgr{
	allBalancer: make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

//对外提供注册
func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

//对外使用LB
func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}

	fmt.Printf("use %s balancer\n", name)
	inst, err = balancer.DoBalance(insts)
	return
}

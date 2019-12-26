package balance

import "fmt"

type BalancerMgr struct {
	balancerMap map[string]Balancer
}

var mgr = BalancerMgr{
	balancerMap: make(map[string]Balancer),
}

func (p *BalancerMgr) register(name string, balancer Balancer) {
	p.balancerMap[name] = balancer
}

func Register(name string, balancer Balancer) {
	mgr.register(name, balancer)
}

func (p *BalancerMgr) doBalance(balancerName string, servers []*Server) (server *Server, err error) {
	balancer, ok := p.balancerMap[balancerName]
	if !ok {
		err = fmt.Errorf("balancer %s is not support", balancerName)
		return
	}

	server, err = balancer.DoBalance(servers)
	return
}

func DoBalance(balancerName string, servers []*Server) (server *Server, err error) {
	return mgr.doBalance(balancerName, servers)
}

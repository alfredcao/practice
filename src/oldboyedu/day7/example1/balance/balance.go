package balance

import (
	"errors"
	"math/rand"
)

type Balancer interface {
	DoBalance(servers []*Server) (server *Server, err error)
}

type RandomBalancer struct {
}

func init() {
	Register("random", &RandomBalancer{})
	Register("roundrobin", &RoundRobinBalancer{})
}

func (p *RandomBalancer) DoBalance(servers []*Server) (server *Server, err error) {
	if servers == nil || len(servers) == 0 {
		err = errors.New("no server")
		return
	}

	index := rand.Intn(len(servers))
	server = servers[index]
	return
}

type RoundRobinBalancer struct {
	currIndex int
}

func (p *RoundRobinBalancer) DoBalance(servers []*Server) (server *Server, err error) {
	if servers == nil || len(servers) == 0 {
		err = errors.New("no server")
		return
	}
	length := len(servers)
	if p.currIndex >= length {
		p.currIndex = 0
	}
	server = servers[p.currIndex]
	p.currIndex = (p.currIndex + 1) % length
	return
}

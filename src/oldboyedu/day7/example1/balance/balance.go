package balance

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
)

type Balancer interface {
	DoBalance(servers []*Server, params ...interface{}) (server *Server, err error)
}

type RandomBalancer struct {
}

func init() {
	Register("random", &RandomBalancer{})
	Register("roundrobin", &RoundRobinBalancer{})
	Register("hash", &HashBalancer{})
}

func (p *RandomBalancer) DoBalance(servers []*Server, params ...interface{}) (server *Server, err error) {
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

func (p *RoundRobinBalancer) DoBalance(servers []*Server, params ...interface{}) (server *Server, err error) {
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

type HashBalancer struct {
}

func (p *HashBalancer) DoBalance(servers []*Server, params ...interface{}) (server *Server, err error) {
	if servers == nil || len(servers) == 0 {
		err = errors.New("no server")
		return
	}

	if params == nil || len(params) == 0 {
		err = errors.New("请传入hash key值")
		return
	}

	param := fmt.Sprintf("%v", params[0])
	table := crc32.MakeTable(crc32.IEEE)
	var hashCode = crc32.Checksum([]byte(param), table)

	length := uint32(len(servers))
	index := hashCode % length

	server = servers[index]
	return
}

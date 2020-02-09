package pool

import (
	"github.com/pkg/errors"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

const ACQUIRE_TIMEOUT_SECOND = 3

var (
	ErrPoolClosed = errors.New("pool already closed")
)

type Pool struct {
	m            sync.Mutex
	currentCount int64
	size         int64
	resources    chan io.Closer
	factory      func() (io.Closer, error)
	closed       bool
}

func NewPool(f func() (io.Closer, error), size int64) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size is too small")
	}

	return &Pool{
		currentCount: 0,
		size:         size,
		resources:    make(chan io.Closer, size),
		factory:      f,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	if p.closed {
		return nil, ErrPoolClosed
	}

	timeOutChan := time.After(time.Second * ACQUIRE_TIMEOUT_SECOND)

	select {
	case resource, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}

		return resource, nil
	case <-timeOutChan:
		if atomic.LoadInt64(&p.currentCount) >= p.size {
			return nil, errors.New("Acquire Resource Timeout")
		} else {
			r, err := p.factory()
			if err != nil {
				return nil, err
			} else {
				atomic.AddInt64(&p.currentCount, 1)
				return r, nil
			}

		}
	}
}

func (p *Pool) Ret(resource io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		resource.Close()
		return
	}

	select {
	case p.resources <- resource:
	default:
		resource.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}

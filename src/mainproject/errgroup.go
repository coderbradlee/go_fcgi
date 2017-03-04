package main

import (
	"sync"

	"golang.org/x/net/context"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type errgroup struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*errgroup, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &errgroup{cancel: cancel}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *errgroup) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *errgroup) Go(f func(t_purchase_order* purchase_order,t *DeliverGoodsForPO,sd *shared_data) error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}

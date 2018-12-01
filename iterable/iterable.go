// Package iterable provides an Iterable type that is capable of converting
// sequences of empty interface such as slice and channel to an Iterator.
package iterable

import "github.com/reactivex/rxgo/errors"

// Iterable converts channel and slice into an Iterator.
type Iterable <-chan interface{}

type RewindIterable struct {
	ch     chan interface{}
	index  int
	items  []interface{}
	rewind bool
}

func (r *RewindIterable) Next() (interface{}, error) {
	if r.ch != nil && !r.rewind {
		if next, ok := <-r.ch; ok {
			r.items = append(r.items, next)
			r.index = r.index + 1
			return next, nil
		}
		return nil, errors.New(errors.EndOfIteratorError)
	} else {
		if r.index < len(r.items) {
			next := r.index
			r.index = r.index + 1
			return next, nil
		} else {
			if r.ch != nil {
				r.rewind = false
				return r.Next()
			}

			return nil, errors.New(errors.EndOfIteratorError)
		}
	}
}

func (r *RewindIterable) Rewind() {
	r.index = 0
	r.items = make([]interface{}, 0)
}

// Next returns the next element in an Iterable sequence and an
// error when it reaches the end. Next registers Iterable to Iterator.
func (it Iterable) Next() (interface{}, error) {
	if next, ok := <-it; ok {
		return next, nil
	}
	return nil, errors.New(errors.EndOfIteratorError)
}

// New creates a new Iterable from a slice or a channel of empty interface.
func New(any interface{}) (Iterable, error) {
	switch any := any.(type) {
	case []interface{}:
		c := make(chan interface{}, len(any))
		go func() {
			for _, val := range any {
				c <- val
			}
			close(c)
		}()
		return Iterable(c), nil
	case chan interface{}:
		return Iterable(any), nil
	case <-chan interface{}:
		return Iterable(any), nil
	default:
		return nil, errors.New(errors.IterableError)
	}
}

// New creates a new Iterable from a slice or a channel of empty interface.
func NewRewindIterable(any interface{}) (RewindIterable, error) {
	r := RewindIterable{}
	switch any := any.(type) {
	case []interface{}:
		r.items = any
	case chan interface{}:
		r.ch = any
	default:
		return r, errors.New(errors.IterableError)
	}

	return r, nil
}

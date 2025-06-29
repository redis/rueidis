package rueidis

import (
	"iter"
)

type Scanner struct {
	next func(cursor uint64) (ScanEntry, error)
	err  error
}

func NewScanner(next func(cursor uint64) (ScanEntry, error)) *Scanner {
	return &Scanner{next: next}
}

func (s *Scanner) scan() iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		var e ScanEntry
		for e, s.err = s.next(0); s.err == nil && yield(e.Elements) && e.Cursor != 0; {
			e, s.err = s.next(e.Cursor)
		}
	}
}

func (s *Scanner) Iter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for vs := range s.scan() {
			for i := 0; i < len(vs) && yield(vs[i]); i++ {
			}
		}
	}
}

func (s *Scanner) Iter2() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for vs := range s.scan() {
			for i := 0; i+1 < len(vs) && yield(vs[i], vs[i+1]); i += 2 {
			}
		}
	}
}

func (s *Scanner) Err() error {
	return s.err
}

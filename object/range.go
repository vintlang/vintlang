package object

import "fmt"

type Range struct {
	Start   int64
	End     int64
	Current int64
}

func (r *Range) Type() ObjectType { return "RANGE" }

func (r *Range) Inspect() string {
	return fmt.Sprintf("%d..%d", r.Start, r.End)
}

func (r *Range) Next() (VintObject, VintObject) {
	if r.Current <= r.End {
		val := &Integer{Value: r.Current}
		r.Current++
		return val, val // For ranges, key and value are the same
	}
	return nil, nil
}

func (r *Range) Reset() {
	r.Current = r.Start
}

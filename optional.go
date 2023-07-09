package optional

import (
	"encoding/json"
)

type Optional[T any] struct {
	v *T
}

func (r Optional[T]) IsPresent() bool {
	return r.v != nil
}

func (r Optional[T]) Get() (T, bool) {
	if r.v == nil {
		var v T
		return v, false
	}
	return *r.v, true
}

func (r Optional[T]) OrElse(v T) T {
	if r.v == nil {
		return v
	}
	return *r.v
}

func (r Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.v)
}

func (r *Optional[T]) UnmarshalJSON(data []byte) error {
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	r.v = &v
	return nil
}

func New[T any](value T) Optional[T] {
	return Optional[T]{
		v: &value,
	}
}

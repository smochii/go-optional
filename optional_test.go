package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsPresent(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v := "hello"
		o := New(v)
		assert.True(t, o.IsPresent())
	})

	t.Run("int", func(t *testing.T) {
		v := 123
		o := New(v)
		assert.True(t, o.IsPresent())
	})

	t.Run("float", func(t *testing.T) {
		v := 123.456
		o := New(v)
		assert.True(t, o.IsPresent())
	})

	t.Run("bool", func(t *testing.T) {
		v := true
		o := New(v)
		assert.True(t, o.IsPresent())
	})

	t.Run("slice", func(t *testing.T) {
		v := []interface{}{
			"hello",
			123,
		}
		o := New(v)
		assert.True(t, o.IsPresent())
	})

	t.Run("struct", func(t *testing.T) {
		type S struct {
			A string
			B int
		}
		v := S{
			A: "hello",
			B: 123,
		}
		o := New(v)
		assert.True(t, o.IsPresent())
	})

	t.Run("map", func(t *testing.T) {
		v := map[string]interface{}{
			"a": "hello",
			"b": 123,
		}
		o := New(v)
		assert.True(t, o.IsPresent())
	})

	t.Run("nil", func(t *testing.T) {
		var o Optional[int]
		assert.False(t, o.IsPresent())
	})

	t.Run("struct", func(t *testing.T) {
		type S struct {
			A string
			B int
		}
		var o Optional[S]
		assert.False(t, o.IsPresent())
	})
}

func Test_Get(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v := "hello"
		o := New(v)
		actual, ok := o.Get()
		assert.Equal(t, v, actual)
		assert.True(t, ok)
	})

	t.Run("int", func(t *testing.T) {
		v := 123
		o := New(v)
		actual, ok := o.Get()
		assert.Equal(t, v, actual)
		assert.True(t, ok)
	})

	t.Run("float", func(t *testing.T) {
		v := 123.456
		o := New(v)
		actual, ok := o.Get()
		assert.Equal(t, v, actual)
		assert.True(t, ok)
	})

	t.Run("bool", func(t *testing.T) {
		v := true
		o := New(v)
		actual, ok := o.Get()
		assert.Equal(t, v, actual)
		assert.True(t, ok)
	})

	t.Run("slice", func(t *testing.T) {
		v := []interface{}{
			"hello",
			123,
		}
		o := New(v)
		actual, ok := o.Get()
		assert.Equal(t, v, actual)
		assert.True(t, ok)
	})

	t.Run("struct", func(t *testing.T) {
		v := struct {
			A string
			B int
		}{
			A: "hello",
			B: 123,
		}
		o := New(v)
		actual, ok := o.Get()
		assert.Equal(t, v, actual)
		assert.True(t, ok)
	})

	t.Run("map", func(t *testing.T) {
		v := map[string]interface{}{
			"a": "hello",
			"b": 123,
		}
		o := New(v)
		actual, ok := o.Get()
		assert.Equal(t, v, actual)
		assert.True(t, ok)
	})

	t.Run("nil", func(t *testing.T) {
		var o Optional[string]
		_, ok := o.Get()
		assert.False(t, ok)
	})
}

func Test_OrElse(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var o Optional[string]
		actual := o.OrElse("hello")
		assert.Equal(t, "hello", actual)
	})

	t.Run("not nil", func(t *testing.T) {
		o := New("world")
		actual := o.OrElse("hello")
		assert.Equal(t, "world", actual)
	})
}

func Test_MarshalJSON(t *testing.T) {
	type S struct {
		A Optional[string]
		B Optional[int]
		C Optional[float64]
		D Optional[bool]
		E Optional[[]interface{}]
		F Optional[struct {
			A string
			B int
		}]
		G Optional[map[string]interface{}]
	}

	t.Run("empty", func(t *testing.T) {
		s := S{}

		actual, err := json.Marshal(s)
		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
	})

	t.Run("not empty", func(t *testing.T) {
		s := S{
			A: New("hello"),
			B: New(123),
			C: New(123.456),
			D: New(true),
			E: New([]interface{}{
				"hello",
				123,
			}),
			F: New(struct {
				A string
				B int
			}{
				A: "hello",
				B: 123,
			}),
			G: New(map[string]interface{}{
				"a": "hello",
				"b": 123,
			}),
		}

		actual, err := json.Marshal(s)
		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
	})
}

func Test_UnmarshalJSON(t *testing.T) {
	type S struct {
		A Optional[string]
		B Optional[int]
		C Optional[float64]
		D Optional[bool]
		E Optional[[]interface{}]
		F Optional[struct {
			A string
			B int
		}]
		G Optional[map[string]interface{}]
	}

	t.Run("empty", func(t *testing.T) {
		s := S{}

		err := json.Unmarshal([]byte(`{}`), &s)
		assert.NoError(t, err)
		assert.False(t, s.A.IsPresent())
		assert.False(t, s.B.IsPresent())
		assert.False(t, s.C.IsPresent())
		assert.False(t, s.D.IsPresent())
		assert.False(t, s.E.IsPresent())
		assert.False(t, s.F.IsPresent())
		assert.False(t, s.G.IsPresent())
	})

	t.Run("not empty", func(t *testing.T) {
		s := S{}

		err := json.Unmarshal([]byte(`{
			"A": "hello",
			"B": 123,
			"C": 123.456,
			"D": true,
			"E": ["hello", 123],
			"F": {"A": "hello", "B": 123},
			"G": {"a": "hello", "b": 123}
		}`), &s)
		assert.NoError(t, err)
		assert.True(t, s.A.IsPresent())
		assert.True(t, s.B.IsPresent())
		assert.True(t, s.C.IsPresent())
		assert.True(t, s.D.IsPresent())
		assert.True(t, s.E.IsPresent())
		assert.True(t, s.F.IsPresent())
		assert.True(t, s.G.IsPresent())
	})
}

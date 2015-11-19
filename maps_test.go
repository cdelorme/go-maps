package maps

import (
	"testing"
)

/**
 * tables
 */

var gt = []struct {
	i struct {
		m map[string]interface{}
		k []string
	}
	o interface{}
	e error
}{
	{
		struct {
			m map[string]interface{}
			k []string
		}{
			m: map[string]interface{}{"key": "value"},
			k: []string{},
		},
		nil,
		errorNoKeys,
	},
	{
		struct {
			m map[string]interface{}
			k []string
		}{
			m: map[string]interface{}{"key": "value"},
			k: []string{"bad"},
		},
		nil,
		errorNotFound,
	},
	{
		struct {
			m map[string]interface{}
			k []string
		}{
			m: map[string]interface{}{"key": "value"},
			k: []string{"key"},
		},
		"value",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			k []string
		}{
			m: map[string]interface{}{"deep": map[string]interface{}{"key": "value"}},
			k: []string{"deep", "key"},
		},
		"value",
		nil,
	},
}

var bt = []struct {
	i struct {
		m map[string]interface{}
		d bool
		k []string
	}
	o bool
	e error
}{
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "true"},
			d: true,
			k: []string{},
		},
		true,
		errorNoKeys,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "true"},
			d: true,
			k: []string{"bad"},
		},
		true,
		errorNotFound,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int8(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int32(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int64(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint8(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint32(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint64(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": float32(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": float64(1)},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int8(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int32(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": int64(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint8(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint32(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": uint64(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": float32(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": float64(0)},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "true"},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": ""},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "false"},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "0"},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "null"},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "undefined"},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": "nil"},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": struct{}{}},
			d: false,
			k: []string{"key"},
		},
		true,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d bool
			k []string
		}{
			m: map[string]interface{}{"key": false},
			d: true,
			k: []string{"key"},
		},
		false,
		nil,
	},
}

var st = []struct {
	i struct {
		m map[string]interface{}
		d string
		k []string
	}
	o string
	e error
}{
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": "value"},
			d: "fallback",
			k: []string{},
		},
		"fallback",
		errorNoKeys,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": "value"},
			d: "fallback",
			k: []string{"bad"},
		},
		"fallback",
		errorNotFound,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"deep": map[string]interface{}{"key": "value"}},
			k: []string{"deep", "key"},
		},
		"value",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": struct {
				d string
			}{d: ""}},
			d: "fallback",
			k: []string{"key"},
		},
		"fallback",
		errorCastFailed,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": true},
			k: []string{"key"},
		},
		"true",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": int(6)},
			d: "7",
			k: []string{"key"},
		},
		"6",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": int8(6)},
			d: "7",
			k: []string{"key"},
		},
		"6",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": int32(6)},
			d: "7",
			k: []string{"key"},
		},
		"6",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": int64(6)},
			d: "7",
			k: []string{"key"},
		},
		"6",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": uint(9)},
			d: "7",
			k: []string{"key"},
		},
		"9",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": uint8(9)},
			d: "7",
			k: []string{"key"},
		},
		"9",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": uint32(9)},
			d: "7",
			k: []string{"key"},
		},
		"9",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": uint64(9)},
			d: "7",
			k: []string{"key"},
		},
		"9",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": float32(6.5)},
			d: "8.3",
			k: []string{"key"},
		},
		"6.5",
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d string
			k []string
		}{
			m: map[string]interface{}{"key": float64(6.5)},
			d: "8.3",
			k: []string{"key"},
		},
		"6.5",
		nil,
	},
}

var it = []struct {
	i struct {
		m map[string]interface{}
		d int64
		k []string
	}
	o int64
	e error
}{
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": 0},
			k: []string{},
		},
		0,
		errorNoKeys,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": 0},
			k: []string{"bad"},
		},
		0,
		errorNotFound,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": "not an int"},
			d: 6,
			k: []string{"key"},
		},
		6,
		errorCastFailed,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": struct{}{}},
			d: 6,
			k: []string{"key"},
		},
		6,
		errorCastFailed,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": "3"},
			k: []string{"key"},
		},
		3,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": "4.5"},
			k: []string{"key"},
		},
		4,
		nil,
	},

	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": int(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": int8(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": int32(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": int64(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": uint(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": uint8(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": uint32(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": uint64(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": float32(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": float64(7)},
			k: []string{"key"},
		},
		7,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": float32(8.3)},
			k: []string{"key"},
		},
		8,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d int64
			k []string
		}{
			m: map[string]interface{}{"key": float64(8.3)},
			k: []string{"key"},
		},
		8,
		nil,
	},
}

var ft = []struct {
	i struct {
		m map[string]interface{}
		d float64
		k []string
	}
	o float64
	e error
}{
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": "true"},
			k: []string{},
		},
		0,
		errorNoKeys,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": "true"},
			k: []string{"bad"},
		},
		0,
		errorNotFound,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": "true"},
			d: 3.3,
			k: []string{"key"},
		},
		3.3,
		errorCastFailed,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": struct{}{}},
			d: 3.3,
			k: []string{"key"},
		},
		3.3,
		errorCastFailed,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": "4.5"},
			d: 3.3,
			k: []string{"key"},
		},
		4.5,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": int(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": int8(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": int32(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": int64(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": uint(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": uint8(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": uint32(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": uint64(12)},
			d: 0,
			k: []string{"key"},
		},
		12,
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": float32(11.6)},
			d: 0,
			k: []string{"key"},
		},
		float64(float32(11.6)),
		nil,
	},
	{
		struct {
			m map[string]interface{}
			d float64
			k []string
		}{
			m: map[string]interface{}{"key": float64(11.6)},
			d: 0,
			k: []string{"key"},
		},
		11.6,
		nil,
	},
}

/**
 * tests
 */

func TestPlacebo(t *testing.T) {
	t.Parallel()

	if !true {
		t.FailNow()
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	for _, tt := range gt {
		o, e := Get(tt.i.m, tt.i.k...)
		if o != tt.o || e != tt.e {
			t.Logf("Expected (%+v, %+v), Got: (%+v, %+v) - %+v\n", tt.o, tt.e, o, e, tt)
			t.Fail()
		}
	}
}

func TestCastBool(t *testing.T) {
	t.Parallel()

	for _, tt := range bt {
		o, e := Bool(tt.i.m, tt.i.d, tt.i.k...)
		if o != tt.o || e != tt.e {
			t.Logf("Expected (%+v, %+v), Got: (%+v, %+v) - %+v\n", tt.o, tt.e, o, e, tt)
			t.Fail()
		}
	}
}

func TestCastString(t *testing.T) {
	t.Parallel()

	for _, tt := range st {
		o, e := String(tt.i.m, tt.i.d, tt.i.k...)
		if o != tt.o || e != tt.e {
			t.Logf("Expected (%+v, %+v), Got: (%+v, %+v) - %+v\n", tt.o, tt.e, o, e, tt)
			t.Fail()
		}
	}
}

func TestCastInt(t *testing.T) {
	t.Parallel()

	for _, tt := range it {
		o, e := Int(tt.i.m, tt.i.d, tt.i.k...)
		if o != tt.o || e != tt.e {
			t.Logf("Expected (%+v, %+v), Got: (%+v, %+v) - %+v\n", tt.o, tt.e, o, e, tt)
			t.Fail()
		}
	}
}

func TestCastFloat(t *testing.T) {
	t.Parallel()

	for _, tt := range ft {
		o, e := Float(tt.i.m, tt.i.d, tt.i.k...)
		if o != tt.o || e != tt.e {
			t.Logf("Expected (%+v, %+v), Got: (%+v, %+v) - %+v\n", tt.o, tt.e, o, e, tt)
			t.Fail()
		}
	}
}

func TestMerge(t *testing.T) {
	t.Parallel()

	// test case
	m1 := map[string]interface{}{"key": "value", "b": true, "deep": map[string]interface{}{"copy": "me"}, "fail": map[string]interface{}{"no": false}}
	m2 := map[string]interface{}{"key": "value2", "a": 1, "deep": map[string]interface{}{"next": "keypair"}, "fail": "test"}

	// result
	m3 := Merge(m1, m2)

	// extract & validate results
	if key, _ := String(m3, "", "key"); key != "value2" {
		t.Logf("Expected: value2, Got: %s\n", key)
		t.FailNow()
	}
	if a, _ := Int(m3, 0, "a"); a != 1 {
		t.Logf("Expected: 1, Got: %+v\n", a)
		t.FailNow()
	}
	if b, _ := Bool(m3, false, "b"); b != true {
		t.Logf("Expected: true, Got: %+v\n", b)
		t.FailNow()
	}
	if deepCopy, _ := String(m3, "", "deep", "copy"); deepCopy != "me" {
		t.Logf("Expected: me, Got: %+v\n", deepCopy)
		t.FailNow()
	}
	if deepNext, _ := String(m3, "", "deep", "next"); deepNext != "keypair" {
		t.Logf("Expected: keypair, Got: %+v\n", deepNext)
		t.FailNow()
	}
	if fail, _ := String(m3, "", "fail"); fail != "test" {
		t.Logf("Expected: test, Got: %+v\n", fail)
		t.FailNow()
	}

}

func TestTo(t *testing.T) {
	t.Parallel()

	// ref to locally defined struct
	r := &struct {
		Key   string `json:"key,omitempty"`
		Name  string `json:"name,omitempty"`
		Final string `json:"Final,omitempty"`
	}{}

	// create map[string]interface{} representing json /w one conflicting key
	d := map[string]interface{}{"key": "value", "Key": "value2", "name": "banana", "Final": "hammock"}

	// cast & validate
	To(r, d)
	if r.Key != "value" {
		t.Logf("Expected: value2, Got: %+v\n", r.Key)
		t.FailNow()
	}
	if r.Name != "banana" {
		t.Logf("Expected: banana, Got: %+v\n", r.Name)
		t.FailNow()
	}
	if r.Final != "hammock" {
		t.Logf("Expected: hammock, Got: %+v\n", r.Final)
		t.FailNow()
	}
}

/**
 * benchmarks
 */

func BenchmarkCastBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool(bt[i%len(bt)].i.m, bt[i%len(bt)].i.d, bt[i%len(bt)].i.k...)
	}
}

func BenchmarkCastString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(st[i%len(st)].i.m, st[i%len(st)].i.d, st[i%len(st)].i.k...)
	}
}

func BenchmarkCastInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(it[i%len(it)].i.m, it[i%len(it)].i.d, it[i%len(it)].i.k...)
	}
}

func BenchmarkCastFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float(ft[i%len(ft)].i.m, ft[i%len(ft)].i.d, ft[i%len(ft)].i.k...)
	}
}

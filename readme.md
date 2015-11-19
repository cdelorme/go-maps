
# go-maps

A support library for `map[string]interface{}` to handle deep merging and type casting without reflection library dependency.


## sales pitch

This library offers the following functionality:

- casting to `bool`
- casting to `string`
- casting to `int64`
- casting to `float64`
- merging of two or more `map[string]interface` with _map_ friendly deep copying
- casting to a defined struct leveraging the [`encoding/json`](https://golang.org/pkg/encoding/json/) package
- simply acquiring an `interface{}` at any depth in a map

The casting logic has custom behavior to deal with data types common to other, more dynamic, languages.  This gives it great utility with minimal modification external to go.  For example any non-zero number, unidentifiable type, or non-empty string without explicit "false", "nil", "null", or "0" value is inferred as true.  Similarly, converts numbers to strings sanely and vice-versa.

_There may be some strange behavior when casting from float32 to float64; not much I can do, review the IEEE-754, and avoid using float32 when using my library._

Since this package leverages the json library, behavior should be identical.

_Attempts were made to leverage the `reflect` package, but the tradeoff in code was not beneficial.  Near equal lines, and deferrederror-trapping were needed, and there was no support for safely translating between types such as displaying a number as a string or accepting the string "true" as a truthy boolean value.  I may try again and put together a benchmark comparison._

The decision to use 64-bit types was both performance and precision based; it also seems to be the default for most packages.

Latest iteration comes complete with table-drive tests to completely validate behavior, as well as benchmarks:

    BenchmarkCastBool-8     30000000            57.9 ns/op
    BenchmarkCastString-8   10000000           124 ns/op
    BenchmarkCastInt-8      20000000            67.1 ns/op
    BenchmarkCastFloat-8    20000000            66.4 ns/op


## intended purpose

This library was built to deal with `map[string]interace{}` data, specifically coming from three other libraries I wrote to gather application configuration:

- [go-config](https://github.com/cdelorme/go-config)
- [go-option](https://github.com/cdelorme/go-option)
- [go-env](https://github.com/cdelorme/go-env)


## usage

Merge example:

    m1 := make(map[string]interface{})
    m2 := make(map[string]interface{})
    m1["key"] = "value"
    m2["key"] = "value2"

    m1 = maps.Merge(m1, m2)

    if m1["key"] == "value2" {
		//this will be true
	}

Casting can also implicitly perform merges, it accepts a reference to a struct plus one or more maps:

    var data CustomStruct
    maps.To(&data, config, flags)

Casting accepts the map, a fallback or "default" value, and one or more keys (more keys is interpretted as nested map depth):

    m := make(map[string]interface{})
    m["boolExample"] = true
    b, err := maps.Bool(m, false, "boolExample")
    if err != nil {
        // handle or ignore error
    }

_You can choose to ignore the returned error when supplying a sane-default for the "fallback" value._

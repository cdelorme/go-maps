
# go-maps

A support library for `map[string]interface{}` to handle deep merging and type casting without reflection library dependency.


## sales pitch

This library offers basic casting to four common data types (`bool`, `string`, `int64`, and `float64`), as well as the option to merge one or more `map[string]interface{}` maps.

The `Merge()` functionality works with deep objects.

The choice of `int64` and `float64` data types is for performance reasons and ensuring same-behavior.

The translation of `bool` types is value when the data type is a boolean, but if converting from say a string or numeric value it implicitly assumes a truthy value (even if the value is `"false"` or `0` as a string or number).

The casting methods accept multiple string keys to traverse depth.

If data cannot be cast, or the value is not found an error is returned.

A fallback method can be supplied, so you can choose to deal with missing data or handle the error.

The `Merge()` method accepts N+1 maps sensibly combines data at different keys into a single `map[string]interface{}`.

It supports embedded `map[string]interface{}` data, such as objects within objects.

Conflicting results defer to the last map supplied.


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

Casting example:

    m := make(map[string]interface{})
    m["boolExample"] = true
    b, err := maps.Bool(m, false, "boolExample")
    if err != nil {
        // handle or ignore error
    }


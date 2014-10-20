
# go-maps

A golang map data type support library.  It focuses on `map[string]interface` providing merge and type casting with variable key depth.


## sales pitch

This library offers basic casting to four common data types (`bool`, `string`, `int64`, and `float64`), as well as the option to merge one or more `map[string]interface{}` maps.


## usage

Merge example:

    m1 := make(map[string]interface{})
    m2 := make(map[string]interface{})
    m1["key"] = "value"
    m2["key"] = "value2"

    m1 = maps.Merge(m1, m2)

    m1["key"] == "value2"

_Allows you to merge same results defering priority to the later maps.  Will also combine any data giving you all that exists in both._

Casting example:

    m := make(map[string]interface{})
    m["boolExample"] = true
    b, err := maps.Bool(&m, false, "boolExample")
    if err != nil {
        // handle error
    }

_All supplied casting methods expect the map by reference.  They accept multiple keys allowing to support key depth.  Additionally they expect a fallback value, but will also return errors to support both common scenarions; error handling and "default" values._

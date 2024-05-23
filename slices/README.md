* Need to understand slice length and capacity
  * Also, Slice operations e.g. append, slicing etc. and how they work on capacity
```
cap(s) -> gives the capacity of the slice s
len(s) -> gives the length of the slice s
```
* The `slice` of a `slice` shares the same memory as the original slice, so b[:] == b
  * Slicing does not copy the slice’s data. It creates a new slice value that points to the original array.
```
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:]
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
```
* Reference: https://go.dev/blog/slices-intro

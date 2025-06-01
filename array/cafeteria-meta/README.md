# Cafeteria

Meta Level 1

start: Wednesday, July 24, 2024

finish: Saturday, August 10, 2024

[Link](https://www.metacareers.com/profile/coding_puzzles?puzzle=203188678289677)

## Problem Statement

A cafeteria table consists of a row of ${N}$ seats, numbered from 1 to ${N}$ from left to right.

Social distancing guidelines require that every diner be seated such that ${K}$ seats to their left and ${K}$ seats to their right (or all the remaining seats if there are fewer than ${K}$) are empty.

There are currently ${M}$ diners seated at the table, where the ${i\text{th}}$ person is seated in seat, ${S_i}$.

Determine the `maximum number` of additional diners that can sit at the table without violating social distancing guidelines.

## Function Parameters

${N}$ = total number of seats at the table

${K}$ = number of seats to right and left of a diner that are closed

${M}$ = total number of seated diners at the table

${[S]}$ = slice representing the seat numbers for the seated diners

## Example

1. 
    N = 10
    
    K = 1
    
    M = 2
    
    S = [2, 6]

    ```
        _  S  _  _  _  S  _  _  _  _

        1  2  3  4  5  6  7  8  9  10

    ```

    with K = 1, the following seats are closed, ${x}$:

    ```
        x  S  x  _  x  S  x  _  _  _

        1  2  3  4  5  6  7  8  9  10

    ```
    Therefore, three additional diners, ${P}$, can sit without violating social distancing guidelines:

    ```
        x  S  x  P  x  S  x  P  x  P

        1  2  3  4  5  6  7  8  9  10

    ```

    Result = 3

## Orginial Algorithm, v0

This algorithm loops from 1 to ${M}$ to fill a hash table with a maximum size of ${O(N)}$.

1. We need to identify seats that are NOT:

    a) occuped by a person, OR

    b) restricted by social distancing guidelines

2. The occupied seats array is unsorted.  Loop through the occupied seats array, ${S}$, and construct a dictionary that identifies occupied seats and closed seats:

```text
    {
        6: "closed",
        7: "person",
        8: "closed
    }
```

3. Loop from seat 1 to ${N}$ and check the hash table to determine if the seat ${S_i}$ is occupied or closed:

    a. If seat is occupied, skip over ${K}$ seats.

    b. If seat is closed, move to the next seat in the table.

4. If the seat is empty, check the next ${K}$ seats to see if the seats are occupied or closed.  If all ${K}$ seats are closed, add a person in the current seat and increase the count.  Skip over ${K}$ seats and continue the loop.

5. When all ${N}$ seats have been examined, return the final count.

## Results

My v0 algorithm passes 29/32 tests.  Three tests fail, likely because of the 5-second time limit or space limitations.

I have tried numerous improvements to reduce compute time, but to no avail.

My dictionary adds $O{(N)}$ of memory, so that may be causing the test failures, especially for large ${N}$.

After researching [online](https://www.reddit.com/r/algorithms/comments/od01mp/stumped_on_an_algo_challenge_edge_case/), I learned that an algorithm exists to find the `maximum` number of seats for a given ${N}$ and ${K}$.  From there, you can modify the result to account for existing diners.

## Questions 

Would it be better to sort the array first?

---

## Modified Algorithm, v1

1. Sort ${[S]}$ from lowest to highest.

2. Break up the table into discrete segments based on the location of the seated diners:

    ```
        _  S  _  _  _  S  _  _  _  _

        1  2  3  4  5  6  7  8  9  10

    Segment 1:  start of table to first seated diner

        _  S
        1  2

    Segment 2:  segments between seated diners
              _  _  _  S
              3  4  5  6

    Segment 3:  last seated diner to end of table
                          _  _  _  _
                          7  8  9  10
    ```

3.  Loop over the elements in ${S[ ]}$ to find the break point of each segment.  Look for edge cases at the start and end of the table.

4. For each segment, determine the minimum number of seats required, accounting for the social distancing rules, to add one diner.

5.  Divide the empty seats by the minimum number to determine the maximum seats that can be added for that segment.

6.  Keep a tally of the count and return the count after all segments have been processed.

## Optimized Algorithm, v2

While v1 passed my unit tests, it did not perform significantly faster than my original v0.  In fact, some unit tests were periodically slower.

So I investigated other algorithms online and found a [Python solution](https://github.com/justin-qu/Meta_Coding_Challenges/blob/main/Level%201/Cafeteria.py) that limited the `for` loop to only those segments between seated diners.  The beginning table segment from the first chair to the first seated diner and the last table segment from the last seated diner to the end of the table were computed outside of a for loop.

This opitimization significantly improved performance across all unit tests.

## Sorting in Meta's Testing Platform

### Meta's Testing Platform

Unfortunately, Meta's testing platform is running an older version of Go and the new `slices` package is unavailable.  

I used the `runtime` package to check Meta's Go version:

```go
    import "runtime"

    fmt.Printf("The application was built with Go version: %s\n", runtime.Version())
```

As of August 2024, Meta is using `v1.16.6` released on July 12, 2021.

### Go Slices Package

In August 2023, Go v1.21 introduced a native [slices package](https://pkg.go.dev/slices) that makes quick work of sorting a slice of `int64` numbers like the input slice, $[S]$.

```go
    import "slices"

    S := []int64{11, 6, 14}
	slices.Sort(S)
```

The beauty of `slices.Sort()` is that it sorts a slice of _any_ order type in ascending order.  

The older `sort` package provides a native sort function for `float64` and `int` types through its `Float64Slice` and `IntSlice` data types, respectively.  But for other data types like `int64`, you have to implement the Sort interface.

### Sorting with the Sort Package

To sort a slice of `int64` integers, I had to implement the `sort.Sort` interface with three methods: `Len()`, `Swap()`, and `Less()`. 

```go
    type int64Slice []int64

    // define sort methods for sort.Sort interface
    func (a int64Slice) Len() int {
        return len(a)
    }

    func (a int64Slice) Swap(i int, j int) {
        a[i], a[j] = a[j], a[i]
    }

    func (a int64Slice) Less(i int, j int) bool {
        return a[i] < a[j]
    }
```

The `String()` method is optional:

```go
    func (a int64Slice) String() (s string) {
        sep := "" // for printing separating commas
        for _, el := range a {
            s += sep
            sep = ", "
            s += fmt.Sprintf("%d", el)
        }
        return
    }
```

With this interface defined, I could then sort the slice of `int64` integers, ${S}$:

```go
    import "sort"

	s := int64Slice(S)
	sort.Sort(s)
```


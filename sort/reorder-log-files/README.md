# Reorder Data in Log Files

[Leet #937](https://leetcode.com/problems/reorder-data-in-log-files/)

Easy

## Topics

• Sorting   
• Comparators   
• SliceStable   

## Code -> Packages -> Modules

Go code is grouped in packages and packages are grouped in modules.  You need to create a module to run the unit tests.

1. Create a Go module in the root directory
```go
    $ cd /reorder-log-files
    $ go mod init reorder-log-files
```
2. The above command creates a `go.mod` file in the root directory.


## Problem Statement

You are given an array of logs. A log comes in one of two types:

    1. Letter log - all words, except the identifier, contain letters

    2. Digit log - all words, except the identifier, contain digits

Each log is space-delimited, and the first word is the identifier

Sort the logs so that:

1. Letter logs come before Digit logs.

2. Letter logs are sorted alphabetically by content.  If content is equal, sort alphabetically by identifier.

3. Digit logs maintain their relative order.

## Examples

Input
```text
["a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"]
```

Output
```text
["g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"]
```
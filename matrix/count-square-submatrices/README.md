# Count Square Submatrices with All Ones

[Leet #1277](https://leetcode.com/problems/count-square-submatrices-with-all-ones/) - Medium

## Status

Started:    January 5, 2022    
Completed:  January 8, 2022

## Topics

• Matrices  
• Dynamic Programming   
• Google    

## Description

Given a `m * n` matrix of ones and zeros, return how many square submatrices have all ones.

## Examples

```text
  [
          m -->
          0  1  2  3
    n  0 [0, 1, 0, 1]       => 11 (1x1)
       1 [1, 1, 0, 1]       =>  1 (2x2)
       2 [0, 0, 1, 1]         _________
       3 [1, 1, 1, 1]          12 square submatrices
  ]


  [
          m -->
          0  1  2  3
    n  0 [0, 1, 0, 1]       => 12 (1 x 1)
       1 [1, 1, 1, 1]       =>  4 (2 x 2)
       2 [0, 1, 1, 1]       =>  1 (3 x 3)
       3 [0, 1, 1, 1]         ___________
                               17 square submatrices
  ]
```

## What is Dynamic Programming (DP)?

A dynamic programming (DP) problem has two key characteristics:

1.  __Overlapping Subproblems__ -- the problem can be broken down into smaller versions of the original problem that are reused multiple times.

2.  __Optimal Substructure__ -- an optimal solution for the original problem can be formed from optimal solutions to the overlapping subproblems.

To contrast, greedy problems have optimal substructure but not overlapping subproblems.  Divide-and-conquer algorithms break problems down into smaller subproblems, but they don't overlap.

### Advantages

•  DP aids in solving complex problems.

•  DP greatly improves the time complexity compared to brute force solutions.  For instance, the brute force solution for Fibonacci sequence requires exponential time complexity.  The DP solution is linear!

## Solution Approaches

1. Bottoms Up - 

2. Top Down - recursion with memoization

## How to Know When a DP Problem is Asked?

The first characteristic that is common in DP problems is that the problem will ask for:
a)  the optimum value (maximum or minimum) of something, or 
b)  the number of ways there are to do something. 

For example:

What is the minimum cost of doing...
What is the maximum profit from...
How many ways are there to do...
What is the longest possible...
Is it possible to reach a certain point...

The second characteristic common in DP problems is that future "decisions" depend on earlier decisions.

In DP problems, we need to factor in results from previous decisions.


## Algorithm

Each matrix element that equals `1` counts as a `1 x 1` square.

With the exception of the top row and far left column, each `1` could also be the lower right corner of larger squares, like `2 x 2`, `3 x 3`, etc.

1. Create a solution matrix `dp` of size (m - 1) * (n - 1).

2. Count the number of `1's` in row 0 from 0 to `m-1`.  This row is the upper bound of the matrix and these elements can't be the lower right corner of larger squares.  Add rowCount to sum.

3. Count the number of `1's` in column 0 from 1 to `n-1`.  This column is the left bound of the matrix and these elements can't be the lower right corner of larger squares.  Add columnCount to sum.

4. Starting at [1][1] of matrix, check if the value equals `1`.

    a. If YES
    Check left, diagonal, and up elements to see if they equal `1`.  Take the minimum of these three values.  Add the minimum plus 1 to represent the square itself and store this value in `dp` matrix.

    b. If NO
    Move to the next element in the row.

5. Once the `dp` matrix is filled, loop over that matrix to sum the values.  Add to sum and return.

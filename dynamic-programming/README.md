# Dynamic Programming (DP)

## What is Dynamic Programming?

DP is a programming technique that breaks down a complex problem into smaller subproblems.  DP can solve complex problems faster than brute force or other conventional algorithms.

A dynamic programming (DP) problem has two key characteristics:

1.  __Overlapping Subproblems__ -- the problem can be broken down into smaller versions of the original problem that are reused multiple times.

2.  __Optimal Substructure__ -- an optimal solution for the original problem can be formed from optimal solutions to the overlapping subproblems.

To contrast, greedy problems have optimal substructure but not overlapping subproblems.  Divide-and-conquer algorithms break problems down into smaller subproblems, but they don't overlap.

### Advantages

•  DP aids in solving complex problems.

•  DP greatly improves the time complexity compared to brute force solutions.  For instance, the brute force solution for Fibonacci sequence requires exponential time complexity.  The DP solution is linear!

## Solution Approaches

1. Top Down - recursion with memoization

2. Bottoms Up - tabulation (a `dp` table)

## Know When a DP Problem is Asked

DP problems have two trademark characteristics.  First, they ask to solve a certain type of problem.  Second, future alogrithmic decisions depend on earlier decisions.

### 1.  What is the problem asking to solve?

DP problems ask for either:

    a)  the optimum value (maximum or minimum) of something, or     
        
    b)  the number of ways there are to do something.   

Examples:

    •   What is the minimum cost of doing ...       
    •   What is the maximum profit from ... 
    •   How many ways can you do ...    
    •   What is the longest possible ...    
    •   Is it possible to reach a certain point ... 


### 2.  Future "decisions" depend on earlier decisions.

In DP problems, you need to factor in results from previous decisions.



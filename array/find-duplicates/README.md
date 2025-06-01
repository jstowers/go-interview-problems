# Find Duplicates

## Problem Statement

You have two sorted arrays, each with n-distinct numbers.

How do you find the common numbers between them?

## Brute Force - Nested For Loops

Use nested for loops.  Check each element of arr1 with each element of arr2.  If there are common elements between them, add them to a result array.

### Time Complexity

$O(M \cdot N)$

where M is the length of arr1 and N is the length of arr2

---

## Optimization 1 - Map

Use two separate for loops.  

In loop 1, iterate over arr2 and create a map of key:value pairs, where key is the element and the value is 1.

In loop 2, iterate over arr1 and check if each value is in the map.  If yes, add that element to a result array.

### Time Complexity

$O(M + N)$

where M is the length of arr1 and N is the length of arr2

---

## Optimization 2 - Binary Search

Loop through arr1 and for each element perform a binary search in arr2

### Time Complexity

$O(M + log(N))$

where M is the length of arr1 and N is the length of arr2.

This opitimization will be faster than Optimization #1.

# Find Difference of Two Arrays

[LeetCode #2215](https://leetcode.com/problems/find-the-difference-of-two-arrays/description/?envType=study-plan-v2&envId=leetcode-75)

Easy

## Problem Statement

Given two 0-indexed arrays, `nums1` and `nums2`, return _a list_ `answer` of size 2 where:

`answer[0]` = all _distinct_ integers in `nums1` that are not present in `nums2`

`answer[1]` = all _distinct_ integers in `nums2` that are not present in `nums1`

## Example

```go
nums1 := []int{1, 2, 3, 3, 4, 7, 8}

nums2 := []int{4, 5, 6, 8}
```

Result:

```go

answer = [[1 2 3 7] [5 6]]
```

## Solution

Runtime - 29 ms - better than 42.94%

Memory - 7.81 mb - better than 21.00%

### Time Complexity

$O(M \cdot N)$

where M is the length of nums1 and N is the length of nums2
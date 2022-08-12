# 1643. kth Smallest Instructions

## Solution 1 - Combinatorial

Time = O(h\*v)
Space = O(h+v)

Problems...
By factorial(21) we exceeds the largest integer possible for 64 bits, going unsigned will not help much.

## Solution 2 - Dynamic Programming

Time = O(h\*v)
Space = O(h\*v)

We use Dynamic Programming to calculate the number of paths to destination. This is actually also just memoizing the nCr used in solution 1.

as

```
nCr = (n-1)Cr + (n-1)C(r-1)
```

The number of paths from (x,y) = (h-x,v-y)C(v-y)

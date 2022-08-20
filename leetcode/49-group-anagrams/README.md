# 49. Group Anagrams

## Solution 1

1. Count unique characters
2. Construct a key from step 1
3. Append string into a map of array of strings (M), using the key from step 2
4. Iterate map and create array of array of string from the map (M)

Pseudo Code:

1. Create map M
2. For each word:
3. Create array A `[]int` of size 26
4. For each character in word:
5. `A[character]++`
6. Construct key from A
7. Append word into `M[key]`
8. Iterate `M` to construct final answer

Time = O(wc)
Space = O(c)

where w = number of words
where c = number of characters in a word

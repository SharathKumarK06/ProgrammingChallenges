# [Sparse Arrays](https://www.hackerrank.com/challenges/sparse-arrays/)

There is a collection of input strings and a collection of query strings. For
each query string, determine how many times it occurs in the list of input
strings. Return an array of the results.

## Example

stringList = ['ab', 'ab', 'abc']
queries = ['ab', 'abc', 'ab']

There are 2 instances of 'ab'. 1 of 'abc', and 0 of 'bc'. For each query, add an
element to the return array, `results = [2, 1, 0]`

## Function Description

Function must return an array of integers representing the frequency of
occurrence of each query string in `stringList`

`matchingStrings` has the following parameters:
- string stringList[n] - an array of string to search
- string queries[q] - an array of query strings

## Returns

- int[q] an array of results for each query

## Input Format

The first line contains and integer `n` the size of `stringList[]`
Each of the next `n` lines contains a string `stringList[i]`
The next line contains `q` the size of `queries[]`
Each of the next `q` lines contains a string `queries[i]`

## Constrains

1 <= n <= 1000
1 <= q <= 1000
1 <= |stringList[i],queries[i]| <= 20

## Sample Input 1

stringList = [ aba, baba, aba, xzxb ]
queries = [ aba, xzxb, ab ]

## Sample Output 1

2
1
0

## Explanation 1

Here, "aba" occurs twice, in the first and third string. The string "xzxb" occurs
once in the fourth string, and "ab" does not occur at all.

## Sample Input 2

stringList = [ def, de, fgh ]
queries = [ de, lmn, fgh ]


## Sample Output 2

1
0
1


## Sample Input 3

stringList = [ abcde, sdaklfj, asdjf, na, basdn, sdaklfj, asdjf, na, asdjf, na, basdn, sdaklfj, asdjf ]
queries = [ abcde, sdaklfj, asdjf, na, basdn ]

## Sample Output 3

1
3
4
3
2




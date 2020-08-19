
__Selection Sort__
---
in-place

worst-case performance: O(n^2)  
best-case performance: O(n^2)  
average-case performance: O(n^2)  

Inefficient for large lists, worse than insertion, 

---
####benefits:
- simple, useful where auxiliary memory is limited
- can be faster than others (e.g. mergesort) on small lists (<20 items)
    
---
####other:
worst case swaps is n-1 example list: []int{9, 7, 5, 3, 1, 2, 4, 6, 8}  
selection sort always has to make n^2 comparisons but the number of swaps 
varies. Typically performs 1 swap per scan in the best case, this can be reduced
to 0 with a check before swapping to ensure that you aren't 'swapping' a value 
with itself.

takes the same amount of time to run 
regardless of the order of the list
(could be helpful sometimes?)

---
####more info:
[wikipedia](https://en.wikipedia.org/wiki/Selection_sort)



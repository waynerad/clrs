# clrs

This repository has code that I wrote to master algorithm design techniques, currently "divide and conquer" (previously called recurrences) and "dynamic programming", which is a terrible name -- it implies there is code generation or somesuch going on, but there isn't. A better name would be "divide and conquer with memoization". The "memoization" is the practice of maintaining state that the recursive function calls can access to get results of previous computation. Dynamic programming algorithms can be translated into iterative (non-recursive) algorithms.

Here you will find quicksort as an exapmle of a "divide an conquer" algorithm. To test my mastery of "divide and conquer", I tested myself to see if I could implement quicksort knowing only the general idea -- that you take an unsorted array, pick a value as a "pivot" and move other values above and below the "pivot", then use recursion to sort the smaller arrays on each side of the pivot -- and timed myself. In other words, no Googling or referencing the book (CLRS) while taking the test. It took me 42 minutes to implement it. That's pretty good! Although not good enough to win a Google Code Jam competition. (If quicksort were a Code Jam problem, which of course it will never be because it is too well known.)

CLRS, by the way, refers to the book Introduction to Algorithms by Cormen, Leiserson, Rivest, & Stein. This is the book MIT uses for teaching algorithms.

I implemented my quicksort algorithm with a whacky 3-way swap, instead of only 2-way swaps. I know it can be done with only 2-way swaps, but didn't think of it in my 42-minute test period. I wrote the program in Golang, and afterward compared my time with the time of Go's built-in sort package. (qs.go is the original, done in 42 minutes, and qscompare.go is the version that does the comparison.) Mine was a lot faster. However, I thought that might be because I'd hard-coded mine to sort floats, rather than being able to sort "anything." So, I turned the compare and swap operations (keeping the 3-way-swap) into function calls through an interface. After making this change, my propgram's performance and Go's were about the same. Go's was actually more consistent, though, while mine had more unpredictable times. There's probably a reason for this -- looking at the Go code, which is possible because it's all open-source, I can see they did something much more complicated than what I did. Some of that extra code probably handles edge cases where it can speed things up, leading to more predictable times.

The version that can sort anything because it is abstracted through an interface is qsany.go. It follows the logic of the original code exactly, hence the interface has two compare and two swap functions (the two compares being greater than and less than, and the two swaps being 2-way and -3way). If we were going to make this part of a library for everyone to use, or part of a commercial system, it would probably we worth reducing the interface to one compare and one swap function. As it is, the built-in Go sort package interface works that way, so there is no need. However, this code is still useful as a succinct example of how to do polymorphism in Go, which might not be self-evident if you are used to object-oriented languages with a "class" keyword, or languages where it is usually done with closures (such as JavaScript).

The rod directory has the rod-cutting algorithm, again implemented without reference to the book, although this time I didn't time it. The rod-cutting problem is the following: Given a rod of length n inches and a table of prices p(i) for i = 1, 2, ..., n, determine the maximum revenue, r(n), obtainable by cutting up the rod and selling the pieces. Note that if the price p(n) for a rod of length n is large enough, an optimal solution may require no cutting at all.

This is solved by using recursion to work downward from the length asked for to the smallest length, then using "memoization" to remember the results and use them on the way back up so nothing needs to be recomputed. Without the memoization, once the sizes get reasonably large (say n = 40), the problem will take a long time (hours) even on a fast computer. To re-implement as an iterative solution, you start at the bottom and work you way up, remembering all your results along the way. Timing the program showed that it was faster in the iterative version. It is possible for the iterative version to be slower if the recursive solution would lead to fewer calculations -- the iterative solution calculates everything (for all n) but the recursive solution can, for certain inputs (n and price table) be more sparse. If it's not sufficiently sparse, then the iterative solution, without the function call overhead of recursion, will be faster.

Although the rod-cutting algorithm is a "toy" problem, the "memoization" technique is used for very advanced algorithms such as DNA comparison.


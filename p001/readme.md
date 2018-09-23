## Multiples of 3 and 5<br>Problem 1
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

## Solution
I found out that this problem can be solved on multiple ways.<br>
The **Slow** way is probably the most common way to solve this, you calculate with every number within the range.<br>
The **Medium** way calculates the valid numbers only.<br>
And the **Fast** way which calculates the result with a formula and caching some values.

## Benchmark
To test how fast they are I'm calculating the problem 10.000.000 times and see how long it takes on every way.<br>
The **Slow** way takes 28.3 seconds<br>
The **Medium** way takes 3.6 seconds<br>
The **Fast** way takes 680 milliseconds<br>
Quite a difference
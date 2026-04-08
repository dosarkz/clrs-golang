## Introduction

When you design and analyze algorithms, you need to be able to describe how they
operate and how to design them. You also need some mathematical tools to show
that your algorithms do the right thing and do it efficiently.

### 1.1 Algorithms

Informally, an algorithm is any well-defined computational procedure that takes
some value, or set of values, as input and produces some value, or set of values, as
output in a finite amount of time. An algorithm is thus a sequence of computational
steps that transform the input into the output.

As an example, suppose that you need to sort a sequence of numbers into mono-tonically increasing order.

Input: A sequence of n numbers {a1, a2, ..., an}

Output: A permutation of the numbers {a1, a2, ..., an} of the input sequence such that a1 ⇐ a2 ⇐ ... ⇐ an.

Thus, given the input sequence {3, 1, 2}, the output permutation would be {1, 2, 3}. Such an algorithm is called a sorting algorithm.

Which algorithm is best for a given application
depends on among other factors — the number of items to be sorted, the extent
to which the items are already somewhat sorted, possible restrictions on the item
values, the architecture of the computer, and the kind of storage devices to be used:
main memory, disks, or tapes.

An algorithm for a computational problem is correct if, for every problem instance provided as input, it halts - 
finished its computing in finite time and outputs the correct solution to the problem instance. 
A correct algorithm solves the given computational problem. An incorrect algorithm might not halt at all on some
input instances, or it might halt with an incorrect answer. Contrary to what you
might expect, incorrect algorithms can sometimes be useful, if you can control their error rate.


### Hard problems, NP-complete

Why are NP-complete problems interesting? First, although no efficient algorithm for an NP-complete problem has ever 
been found, nobody has ever proven that an efficient algorithm for one cannot exist. In other words, no one knows
whether efficient algorithms exist for NP-complete problems. 

As a concrete example, consider a delivery company with a central depot. Each day, it loads up delivery trucks at 
the depot and sends them around to deliver goods to several addresses. At the end of the day, each truck must end up 
back at the depot so that it is ready to be loaded for the next day. To reduce costs, the company wants to select an order
of delivery stops that yields the lowest overall distance traveled by each truck. This problem is the well-known 
'traveling-salesperson problem,' and it is NP-complete.


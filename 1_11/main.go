package main

import "fmt"

// A function f is defined by the rule that f(n) = n if n<3 and
// f(n) = f(n - 1) + 2f(n - 2) + 3f(n - 3) if n >= 3. Write a procedure that
// computes f by means of a recursive process. Write a procedure that computes f
// by means of an iterative process.

func FRecursive(n int) int {
	if n < 3 {
		return n
	} else {
		return FRecursive(n-1) + 2*FRecursive(n-2) + 3*FRecursive(n-3)
	}
}

// In general, an iterative process is one whose state can be summarized by a
// fixed number of state variables, together with a fixed rule that describes
// how the state variables should be updated as the process moves from state
// to state and an (optional) end test that specifies conditions under which
// the process should terminate. In computing n!, the number of steps required
// grows linearly with n. Such a process is called a linear iterative process.

func FIterative(n int) int {
	return fIterator(1, 0, n)
}

func fIterator(counter int, result int, n int) int {
	if n < 3 {
		return result
	} else {
		return fIterator(counter+1, result, n-1)
	}
}

func main() {
	fmt.Println(fib(5))
}

// (define (fib n)
//   (fib-iter 1 0 n))
//
// (define (fib-iter a b count)
//   (if (= count 0)
//     b
//     (fib-iter (+ a b) a (- count 1))))

func fib(n int) int {
	return fibIter(1, 0, n)
}

func fibIter(a, b, count int) int {
	if count == 0 {
		return b
	} else {
		return fibIter(a+b, a, count-1)
	}
}
// fib(6) = 8
// a  b count
// -----------
// 1  0     6
// 1  1     5
// 2  1     4
// 3  2     3
// 5  3     2
// 8  5     1
// 13 8     0

// f(5) = 25
// n1 n2 n3 count
// ---------------
//  2  1  0     5
//
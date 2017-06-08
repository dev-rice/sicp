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

// resTotal = res0 + 2*res1 + 3*res2

// f(3) = f(2) + 2*f(1) + 3*f(0) = 2 +   2*1 +  3*0 = 4
//                                 res0, res1, res2 = 2, 1, 0

// resTotal = 2 + 2*1 + 3*0 = 2 + 2 + 0 = 4
// f(4) = f(3) + 2*f(2) + 3*f(1) = 4 +   2*2 + 3*1  = 11
// 								   res0, res1, res2 = resTotal, res0, res1 = 4, 2, 1

// resTotal = 4 + 2*2 + 3*1 = 11
// f(5) = f(4) + 2*f(3) + 3*f(2) = 11 + 2*4 + 3*2 = 25 => res0, res1, res2



func FIterative(n int) int {
	if n < 3 {
		return n
	}
	return fIterator(2, 1, 0, n-3)
}

func fIterator(res0 int, res1 int, res2 int, counter int) int {
	resTotal := res0 + 2*res1 + 3*res2

	if counter == 0 {
		return resTotal
	} else {
		return fIterator(resTotal, res0, res1, counter-1)
	}
}

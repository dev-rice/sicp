package main

func FRecursive(n int) int {
	if n < 3 {
		return n
	} else {
		return FRecursive(n-1) + 2*FRecursive(n-2) + 3*FRecursive(n-3)
	}
}

func FIterative(n int) int {
	if n < 3 {
		return n
	}



	return 0
}

func main() {

}

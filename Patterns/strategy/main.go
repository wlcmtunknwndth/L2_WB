package main

import "fmt"

type partitionAlg interface {
	partition([]int) []int
}

type partitionTwoPointers struct {
}

func (p *partitionTwoPointers) partition(arr []int) []int {
	fmt.Println("sorting by two pointers")
	return arr
}

type partitionThreePointers struct {
}

func (p *partitionThreePointers) partition(arr []int) []int {
	fmt.Println("sorting by three pointers")
	return arr
}

type QuickSort struct {
	arr []int
	alg partitionAlg
}

func (s *QuickSort) sort() []int {
	s.arr = s.alg.partition(s.arr)
	fmt.Println("here what u got: ", s.arr)
	return s.arr
}

func (s *QuickSort) changePartAlg(alg partitionAlg) {
	s.alg = alg
}

func main() {
	arr := []int{12, 31, 341, 32, 1256, 54, 341}

	twoPs := partitionTwoPointers{}

	threePs := partitionThreePointers{}

	sort := QuickSort{
		arr: arr,
		alg: &twoPs,
	}
	sort.sort()

	sort.changePartAlg(&threePs)
	sort.sort()

}

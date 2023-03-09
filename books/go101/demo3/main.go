package main

func NoDiff[V comparable](vs ...V) bool {
	if len(vs) == 0 {
		return true
	}
	v := vs[0]
	for _, x := range vs[1:] {
		if v != x {
			return false
		}
	}
	return true
}

func main() {
	var NoDiffString = NoDiff[string]
	println(NoDiff("Go", "Go", "Go"))
	println(NoDiffString("Go", "go"))

	println(NoDiff(123, 123, 123, 123))
	println(NoDiff[int](123, 123, 789))

	type A = [2]int
	println(NoDiff(A{}, A{}, A{}))
	println(NoDiff(A{}, A{}, A{1, 2}))

	println(NoDiff(new(int)))
	println(NoDiff(new(int), new(int)))

	println(NoDiff[bool]())

	// _ = NoDiff() // error: 不能推断类型V
	// error: 类型没有实现comparable
	// _ = NoDiff([]int{}, []int{})
	// _ = NoDiff(map[string]int{})
	// _ = NoDiff(any(1), any(1))
}

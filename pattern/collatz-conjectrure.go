package pattern

func CollatzConjecture(num int, list []int) (int, []int) {
	if num == 1 {
		return 1, append(list, num)
	}
	if num%2 == 0 {
		return CollatzConjecture(num/2, append(list, num))
	}
	return CollatzConjecture((num*3)+1, append(list, num))
}

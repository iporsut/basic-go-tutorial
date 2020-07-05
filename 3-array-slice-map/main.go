package main

func main() {
	// Arrays
	// var nums [5]int
	// fmt.Println(nums)
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// nums[3] = 4
	// nums[4] = 5
	// // nums[5] = 6
	// fmt.Println(nums[0])
	// fmt.Println(nums[1])
	// fmt.Println(nums[2])
	// fmt.Println(nums[3])
	// fmt.Println(nums[4])
	// nums := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(nums)

	// Slice
	// var nums []int
	// fmt.Println(nums)
	// nums = append(nums, 1)
	// fmt.Println(nums)
	// nums[0] = 2
	// fmt.Println(nums[0])
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(nums)
	// nums := make([]int, 5)
	// fmt.Println(nums)
	// var nums []int
	// nums := make([]int, 0)
	// if nums == nil {
	// 	fmt.Println("nums is nil")
	// } else {
	// 	fmt.Println("nums is not nil")
	// }
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(len(nums))
	// nums = append(nums, 6)
	// fmt.Println(len(nums))

	// Slicing operator
	// nums := []int{1, 2, 3, 4, 5}
	// newSlice := nums[1:3]
	// fmt.Println(nums)
	// fmt.Println(newSlice)
	// newSlice = nums[0:3] // nums[:3]
	// fmt.Println(newSlice)
	// newSlice = nums[3:5] // nums[3:]
	// fmt.Println(newSlice)
	// nums = nums[:4]
	// fmt.Println(nums)
	// nums = nums[1:]
	// fmt.Println(nums)
	// nums = append(nums[:1], nums[2:]...)
	// fmt.Println(nums)

	// Map
	// var wordCount map[string]int
	// wordCount = make(map[string]int)
	// wordCount := make(map[string]int)
	// fmt.Println(wordCount)
	// wordCount["hello"] = 1
	// wordCount["world"] = 2
	// fmt.Println(wordCount)
	// fmt.Println(wordCount["hello"])
	// fmt.Println(wordCount["world"])
	// delete(wordCount, "world")
	// fmt.Println(wordCount)
	// fmt.Println(wordCount["world"])
	// wordCount := map[string]int{
	// 	"hello": 1,
	// 	"world": 2,
	// }
	// fmt.Println(wordCount)
	// fmt.Println(wordCount["hello"])
	// fmt.Println(wordCount["world"])

	// helloCount, exist := wordCount["hello"]
	// if exist {
	// 	fmt.Println("have word hello in map and has", helloCount, "words")
	// } else {
	// 	fmt.Println("word hello not exist")
	// }

	// if helloCount, ok := wordCount["hello"]; ok {
	// 	fmt.Println("have word hello in map and has", helloCount, "words")
	// } else {
	// 	fmt.Println("word hello not exist")
	// }

	// nums := []int{1, 2, 3, 4, 5}
	// for i := range nums {
	// 	fmt.Println(i)
	// }
	// for i, v := range nums {
	// 	fmt.Println(i, v)
	// }
	// for _, v := range nums {
	// 	fmt.Println(v)
	// }

	// wordCount := map[string]int{
	// 	"hello": 1,
	// 	"world": 2,
	// }
	// for k := range wordCount {
	// 	fmt.Println(k)
	// }
	// for k, v := range wordCount {
	// 	fmt.Println(k, v)
	// }
	// for _, v := range wordCount {
	// 	fmt.Println(v)
	// }
}

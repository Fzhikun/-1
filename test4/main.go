func twoSum(nums []int, target int) []int {
// 创建一个map来存储元素值及其索引
numMap := make(map[int]int)

// 遍历数组
for i, num := range nums {
// 计算与当前元素相加等于目标值的另一个元素
complement := target - num

// 检查map中是否存在这个另一个元素
if index, ok := numMap[complement]; ok {
// 如果存在，返回两个元素的索引
return []int{index, i}
}

// 如果不存在，将当前元素及其索引添加到map中
numMap[num] = i
}

// 如果没有找到满足条件的两个元素，返回一个空数组（或者根据实际情况处理）
return nil
}
/**
 *
 * File:    leetcode1.go
 *          两数之和
 * 
 * Author:  yiluqingtai(1572236483@qq.com)
 *          Created on 21/6/30
 *          
 **/



func twoSum(nums []int, target int) []int {
    hashTable := map[int]int{}
    for j, x := range nums {
        i, ok := hashTable[target - x]
        if ok {
            return []int{i, j}
        }
        hashTable[x] = j
    }
    return nil
}
package main

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if needle == "" {
		return 0
	} else if haystack == "" {
		return -1
	}
	hayLen := len(haystack)
	needLen := len(needle)
	for i := 0; i < hayLen && hayLen-i >= needLen; i++ {
		if haystack[i] == needle[0] {
			j := 1
			for ; j < needLen; j++ {
				if haystack[i+j] != needle[j] {
					break;
				}
			}
			if j == needLen {
				return i
			}
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	i := 0
	for ; i < numsLen; i++ {
		if target > nums[i] {
			continue
		} else if target <= nums[i] {
			return i
		}
	}
	return i
}

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sLen := len(s)
	var sum int = 0
	for i := 0; i < sLen; i++ {
		if i == sLen-1 || romanMap[string(s[i])] >= romanMap[string(s[i+1])] {
			sum += romanMap[string(s[i])]
		} else {
			sum -= romanMap[string(s[i])]
		}
	}
	return sum
}

func maxSubArray(nums []int) int {
	sumLen := len(nums)
	if sumLen == 0 {
		return 0
	}
	sum := nums[0]
	addSum := 0

	for i := 0; i < sumLen; i++ {
		addSum += nums[i]
		if addSum > sum {
			sum = addSum
		}
		if addSum < 0 {
			addSum = 0
		}
	}
	return sum
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			count++
		} else {
			break
		}
	}
	return count
}

func plusOne(digits []int) []int {
	digLen := len(digits)
	if digLen == 0 {
		return append(digits, 1)
	}
	sign := true
	for i := digLen - 1; i >= 0; i-- {
		if digits[i] != 9 {
			sign = false
			digits[i] ++
			break
		} else if i == 0 {
			digits[i] = 1
		} else {
			digits[i] = 0
		}
	}
	if sign {
		digits = append(digits, 0)
	}
	return digits
}

func isAnagram(s string, t string) bool {
	reserve := [26]int{}
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		reserve[int(s[i])-97] ++
		reserve[int(t[i])-97] --
	}

	for key := range reserve {
		if reserve[key] != 0 {
			return false
		}
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	numMap := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		numMap[nums2[i]] = 1
	}
	for i := 0; i < len(nums1); i++ {
		if val, err := numMap[nums1[i]]; err && val == 1 {
			result = append(result, nums1[i])
			numMap[nums1[i]] = 2
		}
	}
	return result
}

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	numMap := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		if _, err := numMap[nums2[i]]; err {
			numMap[nums2[i]] ++
		} else {
			numMap[nums2[i]] = 1
		}

	}
	for i := 0; i < len(nums1); i++ {
		if val, err := numMap[nums1[i]]; err && val > 0 {
			result = append(result, nums1[i])
			numMap[nums1[i]] --
		}
	}
	return result
}

type Interval struct {
	Start int
	End   int
}


func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	quickSort(intervals)
	result := [] Interval{}
	count := len(intervals)
	for i := 0; i < count-1; i++ {
		if intervals[i].End >= intervals[i+1].Start && intervals[i].End >= intervals[i+1].End {
			intervals[i].End = intervals[i+1].End
			result = append(result, intervals[i])
			i++
		} else {
			result = append(result, intervals[i])
			if i == count-2 {
				result = append(result, intervals[i+1])
			}
		}
	}
	return result
}

func quickSort(intervals []Interval)  {
	if len(intervals) <= 1 {
		return
	}
	index,mid := 1,intervals[0]
	left ,right := 0,len(intervals) - 1
	for left < right {
		if intervals[index].Start > mid.Start || (intervals[index].Start == mid.Start && intervals[index].End > mid.End) {
			intervals[index] , intervals[right] = intervals[right] , intervals[index]
			right--
		}  else {
			intervals[index] , intervals[left] = intervals[left] , intervals[index]
			left++
			index++
		}
	}
	intervals[left].Start = mid.Start
	intervals[left].End = mid.End
	quickSort(intervals[:left])
	quickSort(intervals[left+1:])
}

func sortColors(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	i,j,k := 0,0,0
	for _,val := range nums {
		if val == 0 {
			i++
		} else if val == 1 {
			j++
		} else {
			k++
		}
	}
	for p := 0; p < len(nums); p++ {
		if p < i {
			nums[p] = 0
		} else if p < j + i {
			nums[p] = 1
		} else {
			nums[p] = 2
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	quikSortList(head,nil)
	return head
}

func quikSortList(head *ListNode,tail *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// found the partion
	partionVal := head.Val
	p,q := head,head

	for q != tail {
		if q.Val < partionVal {
			p = p.Next
			p.Val ,q.Val = q.Val, p.Val
		}
		q = q.Next
	}
	head.Val ,p.Val = p.Val,head.Val

	quikSortList(head,p)
	quikSortList(p.Next,tail)
}


func main() {
	//fmt.Println(strStr("mississippi","issip"))
	//fmt.Println(searchInsert([]int{1,3,5},5))
	//fmt.Println(romanToInt("MDCIV"))
	//fmt.Println(maxSubArray([]int{-2,-1}))
	//fmt.Println(lengthOfLastWord(" a"))
	//fmt.Println(plusOne([]int{1}))
	//fmt.Println(isAnagram("anagram","nagaram"))
	//fmt.Println(intersection([]int{},[]int{}))
	//fmt.Println(intersect([]int{2,2},[]int{1,2,2,2,1}))
	//fmt.Println(merge([]Interval{
	//	{1,4},
	//	{1,5},
	//	//{8,10},
	//	//{15,18},
	//}))


}

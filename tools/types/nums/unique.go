package nums

import (
	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/tools/maths"
)

var PerfThr = 300

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop(s []int) []int {
	dup := make([]int, 0, len(s))
	for _, v := range s {
		if !IsInIntSlice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap(s []int) []int {
	dup := make([]int, 0, len(s))
	tmpMap := make(map[int]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop32(s []int32) []int32 {
	dup := make([]int32, 0, len(s))
	for _, v := range s {
		if !IsInInt32Slice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap32(s []int32) []int32 {
	dup := make([]int32, 0, len(s))
	tmpMap := make(map[int32]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop64(s []int64) []int64 {
	dup := make([]int64, 0, len(s))
	for _, v := range s {
		if !IsInInt64Slice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap64(s []int64) []int64 {
	dup := make([]int64, 0, len(s))
	tmpMap := make(map[int64]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// list自去重
func UniqueInt(s []int) []int {
	if len(s) < PerfThr {
		return uniqueIntByLoop(s)
	} else {
		return uniqueIntByMap(s)
	}
}

func UniqueInt32(s []int32) []int32 {
	if len(s) < PerfThr {
		return uniqueIntByLoop32(s)
	} else {
		return uniqueIntByMap32(s)
	}
}

func UniqueIntLen(s []int, l int) []int {
	ints := UniqueInt(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

func UniqueInt32Len(s []int32, l int) []int32 {
	ints := UniqueInt32(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

// s1对s2做差
func DifferInt(s1, s2 []int) []int {
	if len(s2) == 0 {
		return s1
	}
	ints := make([]int, 0, len(s1))
	for _, i := range s1 {
		if !IsInIntSlice(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

func DifferIntLen(s1, s2 []int, l int) []int {
	ints := DifferInt(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

func DifferInt32(s1, s2 []int32) []int32 {
	if len(s2) == 0 {
		return s1
	}
	ints := make([]int32, 0, len(s1))
	for _, i := range s1 {
		if !IsInInt32Slice(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

func DifferInt32Len(s1, s2 []int32, l int) []int32 {
	ints := DifferInt32(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}
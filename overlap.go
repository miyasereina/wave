package main

import (
	"fmt"
	"sort"
)

//提取所有波峰波谷

func extract() {
	wave := Restore()
	max := wave[0] //假设数组的首位为最大值
	//取出所有零轴上的波峰与零轴下的波谷
	nodes := []node{}
	for i := 1; i <= len(wave)-2; i++ {
		if max <= wave[i] && wave[i] > wave[i-1] && wave[i] > wave[i+1] {
			max = wave[i]
			nodes = append(nodes, node{max, i})
			continue
		}

	}
	min := wave[0]
	for i := 1; i <= len(wave)-2; i++ {
		if min >= wave[i] && wave[i] < wave[i-1] && wave[i] < wave[i+1] {
			min = wave[i]
			nodes = append(nodes, node{min, i})
			continue
		}
	}
	//重叠区由波峰波谷决定
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].index < nodes[j].index
	})
	if len(wave) <= 4 {
		return
	}
	overlaps := 1
	max, min = interval(wave[0:4])
	for i := 3; i <= len(wave)-1; i = i + 4 {
		//边界保护防止越界
		if len(wave)-i <= 5 {
			if outInterval(max, min, wave[i:]) {
				//最后一段超界则一定有新重叠区直接加一即可
				overlaps++
				break
			}
		} else {
			if outInterval(max, min, wave[i:i+4]) {
				max, min = interval(wave[i : i+4])
				overlaps++
			}
		}
	}
	fmt.Println(overlaps)

}
func outInterval(left, right float64, nodes []float64) bool {
	sort.Float64s(nodes)
	return nodes[0] >= right || nodes[len(nodes)-1] <= left
}
func interval(nodes []float64) (float64, float64) {
	sort.Float64s(nodes)
	return nodes[0], nodes[3]
}

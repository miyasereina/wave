package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

func Restore() []float64 {
	data, err := ioutil.ReadFile("wave.data")
	if err != nil {
		panic(err)
	}
	j := 0
	wave := []float64{}
	//双指针遍历还原为数组
	for i := 0; i <= len(data)-1; i++ {
		if data[i] == ' ' {
			str := string(data[j:i])
			fl64, err1 := strconv.ParseFloat(str, 64)
			if err1 != nil {
				panic(err1)
			}
			wave = append(wave, fl64)
			j = i + 1
			continue
		}
		if i == len(data)-1 {
			str := string(data[j : i+1])
			fl64, err1 := strconv.ParseFloat(str, 64)
			if err1 != nil {
				panic(err1)
			}
			wave = append(wave, fl64)
		}
	}
	return wave
}

var N int
var M int

type node struct {
	value float64
	index int
}

func count() {
	segment := Restore()
	fmt.Println(len(segment))
	max := segment[0] //假设数组的首位为最大值
	//取出所有零轴上的波峰与零轴下的波谷
	nodes := []node{}
	for i := 1; i <= len(segment)-2; i++ {
		if max <= segment[i] && segment[i] > 0 && segment[i] > segment[i-1] && segment[i] > segment[i+1] {
			max = segment[i]
			nodes = append(nodes, node{max, i})
			continue
		}

	}
	min := segment[0]
	for i := 1; i <= len(segment)-2; i++ {
		if min >= segment[i] && segment[i] < 0 && segment[i] < segment[i-1] && segment[i] < segment[i+1] {
			min = segment[i]
			nodes = append(nodes, node{min, i})
			continue
		}
	}
	//按横轴排序,二次模拟波
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].index < nodes[j].index
	})

	if len(nodes) <= 1 {

		fmt.Println(len(nodes) + 1)

	} else {
		for true {
			compress(nodes)
			if M == N {
				break
			}

			M = N
			N = 0
		}

	}
	fmt.Println(M + 1)

}

///递归模拟到分段完成

func compress(nodes []node) int {
	extremePs := []node{}
	for i := range nodes {
		if i == 0 {
			if (nodes[i].value > 0 && nodes[i].value > nodes[i+1].value) || (nodes[i].value < 0 && nodes[i].value < nodes[i+1].value) {
				extremePs = append(extremePs, nodes[i])
				N++
				continue
			}
		} else if i == len(nodes)-1 {
			if (nodes[i].value > 0 && nodes[i].value > nodes[i-1].value) || (nodes[i].value < 0 && nodes[i].value < nodes[i-1].value) {
				extremePs = append(extremePs, nodes[i])
				N++
				continue
			}
		} else {
			if (nodes[i-1].value < nodes[i].value && nodes[i].value < nodes[i+1].value) || (nodes[i-1].value > nodes[i].value && nodes[i].value > nodes[i+1].value) {
				extremePs = append(extremePs, nodes[i])
				N++
				continue
			}
		}
	}
	extremePs = extremePs[0:0]
	return len(extremePs)
}

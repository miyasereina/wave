package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func creatFile(n, k1, k2, j1, j2 int) {
	//打开文件不存在就创建
	file, err := os.OpenFile("wave.data", os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//随机数种子
	rand.Seed(time.Now().UnixNano())

	var data0 []string
	//todo:设计函数根据n,k1,k2,j1,j2的值给出合理的上升与下跌段的最大值
	//暂时没有思路所以写死
	segment := rand.Intn(11)
	//crest:=rand.Intn(5)
	//trough:=rand.Intn(5)
	//首尾看做波峰波谷
	//随机起始段处于上升段或者下降段
	start := rand.Intn(2)

	switch start {
	//初段为上升段
	case 0:
		{
			data := creatWave(segment, j1, k1, k2, start, n)
			for i, _ := range data {
				data0 = append(data0, fmt.Sprintf("%.2f", data[i]))
			}
		}
	//初段为下跌段
	case 1:
		{
			data := creatWave(segment, j2, k1, k2, start, n)
			for i, _ := range data {
				data0 = append(data0, fmt.Sprintf("%.2f", data[i]))
			}
		}

	}
	result := strings.Join(data0, " ")
	_, err = file.WriteString(result)
	if err != nil {
		panic(err)
	}
}

func creatWave(segment, j, k1, k2, start, n int) []float64 {
	fmt.Println(n)
	var tmp []float64
	var data []float64
	//设置初始范围
	max := 1001
	min := -1001
	//生成每一大段里的小波浪
	for i := 1; i <= segment; i++ {
		if n == 0 {
			return data
		}
		smallW := rand.Intn(j/2)*2 + 1
		//生成小波浪里的每个点
		for k := 1; k <= smallW; k++ {
			//判断上升或者下降
			if k%2 != start {
				num := (rand.Intn(k2-k1+1) + k2)
				if num >= n || i == segment {
					num = n
					n = 0
				} else {
					n = n - num
				}
				for l := 1; l <= num; l++ {
					tmp = append(tmp, rand.Float64()*float64(rand.Intn(max-min)+min-1))
					sort.Float64s(tmp)
					data = append(data, tmp...)
					max = int(tmp[len(tmp)-1])
					min = -1001
					//清空临时数组
					tmp = tmp[0:0]
				}
			} else {
				num := rand.Intn(k2-k1+1) + k2
				for l := 1; l <= num; l++ {
					tmp = append(tmp, rand.Float64()*float64(rand.Intn(max-min)+min-1))
					sort.Slice(tmp, func(i, j int) bool {
						return tmp[i] > tmp[j]
					})
					data = append(data, tmp...)
					max = 1001
					min = int(tmp[len(tmp)-1])
					tmp = tmp[0:0]
				}
			}
		}
	}
	return data
}

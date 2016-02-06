package main

import "fmt"

//go:generate yagi -tem=./temp/list.go -gen=int64;int32 -pkg=main

func main() {
	{
		m := ListInt64{}
		m.Add(1)
		m.Add(2)
		fmt.Println(m.Items())
	}
	{
		m := ListInt32{}
		m.Add(1)
		m.Add(2)
		fmt.Println(m.Items())
	}
}

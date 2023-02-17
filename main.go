package main

import (
	"sina/function"
	"time"
)

func main() {
	//先获取信息
	for {
		function.GetLuckSearchApi()
		function.FollowSet()
		function.HuaTiZhuanFa()
		function.LikeSet()
		time.Sleep(5*time.Minute)
	}

}

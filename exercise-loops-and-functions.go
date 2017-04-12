package main

import (
	"fmt"
	"math"
)

// 計算(なんとfloat64には^が使えない)
func getNext(input float64, target float64) float64 {
	return input - (input*input-target)/(input*2)
}

// 終了判定
func chkAcc(current float64, past1 float64, past2 float64) bool {
	// 前回と同じもしくは前々回と同じ場合trueを返す
	if current == past1 || current == past2 {
		return true
	}
	return false
}

// 指定のフォーマットでecho
func echoMsg(count int, value float64, real bool) int {
	imag := ""
	if !real {
		imag = "i"
	}
	fmt.Printf("%v回目の試行: \t%v%v\n", count, value, imag)
	// todo: Unit?の返し方
	return 0
}

func Sqrt(x float64, tryNum int) float64 {
	var (
		init = 1.0
		real = true
	)

	if x == 0 {
		// 0の場合は正常な値が出ない
		Println(0)
		return 0.0
	} else if x < 0 {
		// xが負の数のとき
		x = -x
		real = false
	}

	// 手動計算
	if tryNum > 0 {
		next := init
		for i := 1; i <= tryNum; i++ {
			next = getNext(next, x)
			echoMsg(i, next, real)
		}
		return next
	}

	// 自動計算
	var (
		next  = init
		past1 = -1.0
		past2 = -1.0
		max   = 100
	)

	// todo: なぜかこのprintlnが遅延する
	// println("---自動計算ログ---")

	// 念の為chkAccに漏れがあった場合に備えて試行回数上限を設ける
	for i := 1; !chkAcc(next, past1, past2) && i <= max; i++ {
		past2 = past1
		past1 = next
		next = getNext(next, x)
		// 計算ログ出力
		echoMsg(i, next, real)
		if i == max {
			println("試行上限に達した為停止します")
		}
	}
	return next
}

func main() {
	// Sqrt(2, 10)
	println("自作関数の答え: ", Sqrt(2, 0))
	println("組込関数の答え: ", math.Sqrt(2))
}

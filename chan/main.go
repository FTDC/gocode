package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

//generate int64 number
func generateRandomNumber(ch chan<- int64) {
	// var i int64
	for {
		ch <- int64(rand.Int63())
	}
}

//get the sum of all digit of input number
// 获取输入数的每一位的和
func GetSum(s int64) int64 {
	strNum := fmt.Sprintf("%d", s)
	lenNum := len(strNum)
	var sumOfAllDigit int64 = 0
	for i := lenNum - 1; i >= 0; i-- {
		temp := s / int64(math.Pow(10, float64(i)))
		fmt.Println(temp)
		s = s % int64(math.Pow(10, float64(i)))
		sumOfAllDigit += temp
	}
	fmt.Println("sum is ", sumOfAllDigit)
	return sumOfAllDigit
}

//24个线程消费生成的数据
func JobChan(wg *sync.WaitGroup, ch chan int64) (sumInfo chan int64) {
	// var chresult chan<- int64
	sumInfo = make(chan int64, 3000)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go func(ch <-chan int64) {

			for {
				temp := <-ch
				fmt.Println(temp)
				sum := GetSum(temp)
				fmt.Println("temp is ", temp)
				fmt.Println("sum is ", sum)
				sumInfo <- sum
			}
			wg.Done()
		}(ch)

	}
	return
}

func main() {
	wg := sync.WaitGroup{}
	chOrigin := make(chan int64, 30000)
	// ch1 := make(chan int64, 300)
	go generateRandomNumber(chOrigin)
	chSum := JobChan(&wg, chOrigin)
	for v := range chSum {
		fmt.Println("*&*&*&*&*&***&**", v)
	}
	// wg.Wait()
	// select {}
}

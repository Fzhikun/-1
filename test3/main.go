package main

import (
fmt
)

// generate函数生成一个自然数序列，从2开始，通过channel发送出去
func generate(ch chan int) {
  for i := 2; ; i++ {
    ch <- i
  }
}

// filter函数过滤掉能被prime整除的数，将不能被prime整除的数通过另一个channel发送出去
func filter(in chan int, out chan int, prime int) {
  for {
    num := <-in
    if num%prime != 0 {
      out <- num
    }
  }
}

func main() {
  ch := make(chan int) // 创建一个无缓冲channel
  go generate(ch)      // 启动生成数序列的goroutine
  for i := 0; i < 6; i++ { // 我们希望找到6个素数
    prime := <-ch // 从channel接收一个数，这个数一定是素数
    fmt.Printf("prime:%d\n", prime)
    out := make(chan int) // 创建一个新的channel用于下一轮过滤
    go filter(ch, out, prime) // 启动一个新的goroutine进行过滤
    ch = out // 将当前的channel更新为新的过滤后的channel
  }
}
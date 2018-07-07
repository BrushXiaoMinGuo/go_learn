#### 如何并发

在需要并发的函数前加上go， go funcname 启动一个goroutine来实现并发 

#### 并发之间的通信
go提供了sync包 和 channel机制来实现并发之间的通信

sync包提供来同步锁之类的原语，sync使用WaitGroup来实现同步，它的实现类似于一个队列结构，提供来add（任务+1），done（任务-1），wait（阻塞）

```
package test

import "fmt"
import "syc.WaitGroup"

var waitgroup sync.WaitGroup // 申明waitgroup变量

func add(num int)
{
  fmt.Println("finish %d", num)
  waitgroup.Done() //任务执行完，队列任务-1
}

func main()
{
   for i:=1; i<=10; i++{
       waitgroup.Add(go add(i)) // 队列任务+1
       go add(i) // 启动goroutine 执行任务
   }

   waitgroup.Wait() //队列阻塞
}
```

线程之间通信有共享内存和消息队列

go 使用channel来进行通信，

ch := make(chan typename) //无缓冲channel
ch := make(chan typename, 10)// 有缓冲的channel，缓冲大小是5
ch <- 4 向channel灌数据
res := <- ch channel 输出数据

对于不带缓冲的channel，写和读都会造成阻塞；
对于带缓冲的channel 当缓冲满了或者缓冲里面没有数据，就会造成阻塞；

```
package test

func add(ch chan int){
  <- ch
}

func main(){
  ch := make(chan int)
  go add(ch) // add 线程阻塞
  ch<-6 //阻塞解除
}
``` 

死锁版本
```
package test

func add(ch chan int){
  <- ch
}

func main(){
  ch := make(chan int)
  ch<-6 // 阻塞 
  go add(ch) // 造成死锁
}
```



package main 

import (
	"fmt"
	"sync"
	//"runtime"
	
)

var  counter  int  = 0 ;

func Count(lock * sync.Mutex){

	lock.Lock()
	counter++;
	fmt.Println("counter = ", counter)
	lock.Unlock()
}
func  Count1(ch chan int) {
	ch <- 1
	fmt.Println("Counting: %d",ch)
}

 func main() {
 /*	lock := &sync.Mutex{}
 	for i :=0; i < 10; i++{
 		go  Count(lock)
 	}
 	for  {

 		lock.Lock()
 		c := counter
 		lock.Unlock()
 		runtime.Gosched()
 		if c >= 10{
 			break;
 		}
 	}*/
 	chs  := make([] chan int, 10)
 	for i :=0; i < 10; i++{
 		chs[i] = make(chan int)
 		go Count1(chs[i])
 	}
 	for _, ch := range(chs){
 		<-ch
 	}
}
	 
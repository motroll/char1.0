package main 


import  "fmt"

func   calDigital(number int ,chnl chan int) {
	
	digital := 0
	for  number>0 {
		digital = number % 10
		chnl <- digital
		number /=10
	}
	close(chnl)
}
func   calSquare( number int , chnl chan int){
	sum  := 0 
	dch := make(chan int)
	go  calDigital(number,dch)
	for digital :=range dch{
		sum +=(digital*digital)
	}
	chnl <- sum


}
func  calCube(number int , chnl chan int){
	sum  := 0 
	dch := make(chan int)
	go  calDigital(number,dch)
	for digital :=range dch{
		sum +=(digital*digital*digital)
	}
	chnl <- sum
	
}

func  main(){
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go  calSquare(number,sqrch)
	go  calCube(number,cubech)
	squares,cubes := <-sqrch, <-cubech
	fmt.Println("Finnal output:",squares+cubes)

}
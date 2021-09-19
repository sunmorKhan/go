package main //ver0.3

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(149) + 2
	fmt.Println("임의로 추출된 수 :", number)

	for i := 2 ; i<number;i++{
		if number % i == 0{
			isPrime = false
			//count++ // count = count + 1 이라는 작업을 줄이겠다.
		}
	}
	if isPrime == true {
		fmt.Println(number,"는(은) 소수입니다!")
	}else{
		fmt.Println(number,"는(은) 소수가 아닙니다!")
	}

}
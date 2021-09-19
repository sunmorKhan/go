package main //ver0.4

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
	//number = 21
	fmt.Println("임의로 추출된 수 :", number)

	for i := 2 ; i<number;i++{
		if number % i == 0{
			isPrime = false
			break // 소수가 아니라는 것을 알게된 순간 break가 걸림. 소수가 아닌 경우 반복수를 줄인다.
		}
		//fmt.Print(i," ")
	}
	if isPrime{ // isPrime 자체가 bool타입으로 treu or false값만 가질 수 있음. 비교연산 제거.
		fmt.Println(number,"는(은) 소수입니다!")
	}else{
		fmt.Println(number,"는(은) 소수가 아닙니다!")
	}

}
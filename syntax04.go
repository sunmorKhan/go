package main //ver0.1

import (
	"fmt"
	"math/rand"
	"time" //시간관련
)

func main()  { // 2~150까지 소수 판별기 소수 : 1과 자기자신 이외에 나누어 떨어지지않는 수
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(149) + 2 // 2 더함으로써 0과 1 제거 2~150R까지 0~148 +2
	fmt.Println("임의로 추출된 수 :", number)

	for i := 1 ; i<=number;i++{
		if number % i == 0{
			count++
		}
	}
	if count == 2 {
		fmt.Println(number,"는(은) 소수입니다!")
	}else{
		fmt.Println(number,"는(은) 소수가 아닙니다!")
	}

}

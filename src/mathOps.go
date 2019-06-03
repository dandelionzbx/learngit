package src

import (
	"crypto/rand"
	"fmt"
	"math"
)

/*
- we start with generate two large primes p and q and p,q must st.gcd(pq,(p-1)*(q-1))=1
- compute n = p * q
- get the lowest common multiplier for (p-1) and (q-1)
- select random integer g where g \in Z_(n^2)^*
-
 */

func gcd(a int, b int)int{
	for b >0{
		a, b = b, a % b
	}
	return a
}

func lcm(a int, b int)int{
	return a*b/(gcd(a,b))
}
func GenerateG()int{
	b := make([]byte, 20)
	fmt.Println(b)

	g, err := rand.Read(b)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(g)
	return g
}

//func Lambda(p int, q int)int{
//	return lcm(p-1, q-1)
//}

func Lfunction(p int, q int)float64{
	g := GenerateG()
	miu := math.Pow(float64(g),float64(lcm(p-1,q-1)))
	miu = math.Mod(miu, float64(p*q*p*q))-1.0
	miu = miu / float64(p*q)
	math.inv
	return miu
}
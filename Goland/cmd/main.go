package main

import (
	"auth"
	"fmt"
	"time"
)

func main(){

	a := new(auth.AuthToken)
	fmt.Println(a.Genrate("aaa","ad",time.Now().Unix(),"ad"))

}

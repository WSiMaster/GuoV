package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func random(n int) int {
	rand.Seed(time.Now().Unix())

	var port int32
	port = rand.Int31n(5015-5000) + 5000
	fmt.Println(port)
	return int(port)
}

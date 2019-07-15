package main

import (
	"fmt"
	DbUtility "miniProject/pkg"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("test file main fun")
	DbUtility.InitiateDB()
	wg.Wait()
	fmt.Println("al;kdf")

}

package main

// go routine
// concurrency != parallel execution

// in golang we may achive some level of parallel execution , if you ave multicore CPUs
// mutex
import (
	"fmt"
	// "math/rand"
	"time"
	"sync"
)

// var m = sync.Mutex{}

// ReadWrite Mutex 

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal excution time: %v \n", time.Since(t0))
	fmt.Printf("\nThe result are %v \n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	// m.Lock()
	results = append(results, dbData[i])
	// m.Unlock()
	wg.Done()
}
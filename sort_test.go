package algorithms

import (
	"testing"
	"runtime"
	"sync"
	//"fmt"
	"fmt"
)

/*to be able to see results exec command below:
go test -bench .
this command runs each test file under the package within benchmark funcs those names starts with Benchmark
btw 1 billion ns equals 1 sec
*/

var globArray []int

func BenchmarkInsertionSorting(b *testing.B) {
	array := Genesis("hundred thousand")
	globArray = InsertionSort(array, 0, 100)
}

func BenchmarkSelectionSorting(b *testing.B) {
	array := Genesis("hundred thousand")
	SelectionSort(array, 0, 100)
}

func BenchmarkParallelSorting(b *testing.B) {
	array := Genesis("hundred thousand")
	//fmt.Println("parallelizing with " + strconv.Itoa(runtime.NumCPU()) + " cpus...")
	numOfCpus := runtime.NumCPU()
	increaseRate := 100 / numOfCpus
	runtime.GOMAXPROCS(numOfCpus)
	var wg sync.WaitGroup
	for i := 0; i < numOfCpus; i++ {
		wg.Add(1)
		go func(i int){
			InsertionSort(array, i * increaseRate, (i + 1) * increaseRate)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

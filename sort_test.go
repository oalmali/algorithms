package algorithms

import (
	"testing"
	"runtime"
	"sync"
	//"fmt"
)

/*to be able to see results exec command below:
go test -bench .
this command runs each test file under the package within benchmark funcs those names starts with Benchmark
btw 1 billion ns equals 1 sec
*/

var globArray []int

func BenchmarkInsertionSorting(b *testing.B) {
	array := Genesis("ten thousand")
	globArray = InsertionSort(array, 0, 100)
}

func BenchmarkSelectionSorting(b *testing.B) {
	array := Genesis("ten thousand")
	SelectionSort(array, 0, 100)
}

func BenchmarkParallelSorting(b *testing.B) {
	array := Genesis("ten thousand")
	//fmt.Println("parallelizing with " + strconv.Itoa(runtime.NumCPU()) + " cpus...")
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		InsertionSort(array, 0, 50)
		wg.Done()
	}()
	wg.Add(1)
	go func(){
		wg.Done()
	}()
	wg.Wait()
}

package p1310

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	type args struct {
		arr     []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{arr: []int{1, 3, 4, 8}, queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}}, []int{2, 7, 14, 8}},
		{"Example 2", args{arr: []int{4, 8, 2, 10}, queries: [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}}, []int{8, 0, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorQueries(tt.args.arr, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readIntSlice(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var size int
	if scanner.Scan() {
		size, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}
	res := make([]int, size)
	index := 0
	for scanner.Scan() {
		res[index], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		index++
	}
	return res
}
func writeIntSlice(filename string, arr []int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = fmt.Fprintf(writer, "%d\n", len(arr))
	if err != nil {
		log.Fatal(err)
	}

	for _, num := range arr {
		_, err = fmt.Fprintf(writer, "%d\n", num)
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()
}

func readIntMatrix(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var size int
	if scanner.Scan() {
		size, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}
	res := make([][]int, size)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")
		res[index] = make([]int, len(strs))
		for i, str := range strs {
			res[index][i], err = strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
		}
		index++
	}
	return res
}

func writeIntMatrix(filename string, arr [][]int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = fmt.Fprintf(writer, "%d\n", len(arr))
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range arr {
		s := fmt.Sprint(a)
		s = s[1 : len(s)-1]
		_, err = fmt.Fprintf(writer, "%s\n", s)
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()
}

func Test_generate(t *testing.T) {
	//now := time.Now().Unix()
	now := int64(1)
	fmt.Println(now)
	seed := rand.NewSource(now)
	random := rand.New(seed)
	lenQueries := random.Intn(3*10e4-1) + 1
	lenArr := random.Intn(3*10e4-1) + 1
	arr := make([]int, lenArr)
	for i := 0; i < lenArr; i++ {
		arr[i] = random.Intn(10e9-1) + 1
	}
	queries := make([][]int, lenQueries)
	for i := 0; i < lenQueries; i++ {
		left := random.Intn(lenArr)
		right := random.Intn(lenArr)
		if left > right {
			left, right = right, left
		}
		queries[i] = []int{left, right}
	}
	writeIntSlice("input.arr1", arr)
	writeIntMatrix("input.queries1", queries)
	//x := readIntSlice("input.arr0")
	//fmt.Println(x)
	//y := readIntMatrix("input.queries0")
	//fmt.Println(y)

}

func Benchmark_xorQueries0_0(b *testing.B) {
	arr := readIntSlice("input.arr0")
	queries := readIntMatrix("input.queries0")
	b.ResetTimer()
	b.ReportAllocs()
	b.Name()

	for i := 0; i < b.N; i++ {
		xorQueries0(arr, queries)
	}
}
func Benchmark_xorQueries0_1(b *testing.B) {
	arr := readIntSlice("input.arr1")
	queries := readIntMatrix("input.queries1")
	b.ResetTimer()
	b.ReportAllocs()
	b.Name()

	for i := 0; i < b.N; i++ {
		xorQueries0(arr, queries)
	}
}

func Benchmark_xorQueries1_0(b *testing.B) {
	arr := readIntSlice("input.arr0")
	queries := readIntMatrix("input.queries0")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		xorQueries1(arr, queries)
	}
}
func Benchmark_xorQueries1_1(b *testing.B) {
	arr := readIntSlice("input.arr1")
	queries := readIntMatrix("input.queries1")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		xorQueries1(arr, queries)
	}
}

func Benchmark_xorQueries2_0(b *testing.B) {
	arr := readIntSlice("input.arr0")
	queries := readIntMatrix("input.queries0")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		xorQueries2(arr, queries)
	}
}
func Benchmark_xorQueries2_1(b *testing.B) {
	arr := readIntSlice("input.arr1")
	queries := readIntMatrix("input.queries1")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		xorQueries2(arr, queries)
	}
}

func Benchmark_xorQueries_0(b *testing.B) {
	arr := readIntSlice("input.arr0")
	queries := readIntMatrix("input.queries0")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		xorQueries(arr, queries)
	}
}
func Benchmark_xorQueries_1(b *testing.B) {
	arr := readIntSlice("input.arr1")
	queries := readIntMatrix("input.queries1")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		xorQueries(arr, queries)
	}
}

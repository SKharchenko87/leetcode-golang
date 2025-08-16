package p1323

/*
ToDo дизасемблировать
Benchmark_maximum69Number
Benchmark_maximum69Number-12     	23099710	        51.65 ns/op	       0 B/op	       0 allocs/op
Benchmark_maximum69Number4
Benchmark_maximum69Number4-12    	20953890	        51.13 ns/op	       0 B/op	       0 allocs/op
Benchmark_maximum69Number3
Benchmark_maximum69Number3-12    	15441372	        80.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_maximum69Number2
Benchmark_maximum69Number2-12    	13262188	        91.80 ns/op	       0 B/op	       0 allocs/op
Benchmark_maximum69Number1
Benchmark_maximum69Number1-12    	100000000	        10.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_maximum69Number0
Benchmark_maximum69Number0-12    	26909026	        46.31 ns/op	       0 B/op	       0 allocs/op
*/

func maximum69Number(num int) int {
	i := 1000
	for ; i >= 1 && num/i%10 != 6; i /= 10 {
	}
	return num + 3*i
}

func maximum69Number4(num int) int {
	i := 1000
	for ; i >= 1 && num/i%10 != 6; i /= 10 {
	}
	return num + 3*i
}

func maximum69Number3(num int) int {
	last6Index := 0
	i := 1
	for {
		v := num / i
		if v == 0 {
			return num + 3*last6Index
		}
		if v%10 == 6 {
			last6Index = i
		}
		i *= 10
	}
}

func maximum69Number2(num int) int {
	for i := 1000; i >= 1; i /= 10 {
		if num%(i*10)/i == 6 {
			return num + 3*i
		}
	}
	return num
}

func maximum69Number1(num int) int {
	last6Index := 0
	tmp := num
	i := 1
	for tmp > 0 {
		if tmp%10 == 6 {
			last6Index = i
		}
		tmp /= 10
		i *= 10
	}
	return num + 3*last6Index
}

func maximum69Number0(num int) int {
	res := 0
	for i := 1000; i >= 1; i /= 10 {
		v := num / i
		if v == 6 {
			return res + i*9 + num%i
		} else if v > 0 {
			res += 9 * i
			num %= i
		}
	}
	return res
}

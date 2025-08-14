package p1352

type ProductOfNumbers struct {
	list []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{}}
}

func (this *ProductOfNumbers) Add(num int) {
	this.list = append(this.list, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	res := 1
	l := len(this.list)
	for i := l - k; i < l; i++ {
		res *= this.list[i]
	}
	return res
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

func run(commands []string, param []interface{}) []interface{} {
	l := len(commands)
	var productOfNumbers ProductOfNumbers
	res := make([]interface{}, l)
	for i := 0; i < l; i++ {
		switch commands[i] {
		case "ProductOfNumbers":
			productOfNumbers = Constructor()
		case "add":
			productOfNumbers.Add(param[i].(int))
		case "getProduct":
			res[i] = productOfNumbers.GetProduct(param[i].(int))
		}
	}
	return res
}

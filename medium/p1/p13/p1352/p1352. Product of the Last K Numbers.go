package p1352

type ProductOfNumbers struct {
	list    []int
	lastNil int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{1}, 0}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.lastNil = len(this.list)
		this.list = append(this.list, 1)
	} else {
		this.list = append(this.list, this.list[len(this.list)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {

	if this.lastNil+k >= len(this.list) {
		return 0
	}
	return this.list[len(this.list)-1] / this.list[len(this.list)-k-1]
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

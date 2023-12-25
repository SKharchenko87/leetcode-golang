package main

func distMoney(money int, children int) int {
	money = money - children
	if money < 0 {
		return -1
	}
	o := money % 7
	x := money / 7
	if x == 0 {
		return 0
	}
	if o == 3 && (children-x) == 1 {
		x--
	}
	return x
}

func main() {
	money := 20
	children := 3
	distMoney(money, children)
}

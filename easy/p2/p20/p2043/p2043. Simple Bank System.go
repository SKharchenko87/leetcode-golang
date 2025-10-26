package p2043

type Bank struct {
	balance []int64
	length  int
}

func Constructor(balance []int64) Bank {
	return Bank{balance, len(balance)}
}

func (this *Bank) valid(account int) bool {
	return 0 <= account && account < this.length
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	account1--
	account2--
	if this.valid(account1) && this.valid(account2) && this.balance[account1] >= money {
		if account1 != account2 {
			this.balance[account1], this.balance[account2] = this.balance[account1]-money, this.balance[account2]+money
		}
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	account--
	if this.valid(account) {
		this.balance[account] = this.balance[account] + money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	account--
	if this.valid(account) && this.balance[account] >= money {
		this.balance[account] = this.balance[account] - money
		return true
	}
	return false
}

func run(commands []string, params []interface{}) (res []interface{}) {
	var bank Bank
	var ok bool
	bank = Constructor(params[0].([]int64))
	res = append(res, nil)
	for i := 1; i < len(commands); i++ {
		switch commands[i] {
		case "deposit":
			depositParam := params[i].([]int64)
			ok = bank.Deposit(int(depositParam[0]), depositParam[1])
		case "withdraw":
			withdrawParam := params[i].([]int64)
			ok = bank.Withdraw(int(withdrawParam[0]), withdrawParam[1])
		case "transfer":
			transferParam := params[i].([]int64)
			ok = bank.Transfer(int(transferParam[0]), int(transferParam[1]), transferParam[2])
		}
		res = append(res, ok)
	}
	return res
}

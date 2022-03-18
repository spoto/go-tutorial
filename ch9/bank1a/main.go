package bank

type withdrawEnvelope struct {
	amount int
	reply  chan<- bool
}

var deposits = make(chan int)
var withdraw = make(chan withdrawEnvelope)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	reply := make(chan bool)
	withdraw <- withdrawEnvelope{amount, reply}
	return <-reply
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case envelope := <-withdraw:
			if balance-envelope.amount >= 0 {
				balance -= envelope.amount
				envelope.reply <- true
			} else {
				envelope.reply <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

type Currency int

const (
	USD Currency = iota
	THB
	JPY
)

var currencyData = map[Currency]struct {
	Symbol string
	Rate   int
}{
	USD: {"USD", 100},
	THB: {"THB", 3093},
	JPY: {"JPY", 15300},
}

func (c Currency) String() string {
	return currencyData[c].Symbol
}

func (c Currency) Rate() int {
	return currencyData[c].Rate
}

type Balance struct {
	Symbol string
	Cash   int
}

func (b Balance) String() string {
	return fmt.Sprintf("%s:%d", b.Symbol, b.Cash)
}

func (b Balance) IsZero() bool {
	return b.Cash == 0
}

type Account struct {
	CarNo    int
	Balances map[string]Balance
}

func NewAccount(carNo int) *Account {
	return &Account{
		CarNo:    carNo,
		Balances: make(map[string]Balance),
	}
}

func (a *Account) AddBalance(symbol string, amount int) error {
	current := a.Balances[symbol]
	newCash, err := SafeAdd(current.Cash, amount)
	if err != nil {
		return fmt.Errorf("add balance failed: %w", err)
	}
	a.Balances[symbol] = Balance{
		Symbol: symbol,
		Cash:   newCash,
	}
	return nil
}

func (a *Account) SubBalance(symbol string, amount int) error {
	current := a.Balances[symbol]
	newCash, err := SafeSub(current.Cash, amount)
	if err != nil {
		return fmt.Errorf("insufficient balance: %w", err)
	}

	if newCash == 0 {
		delete(a.Balances, symbol)
	} else {
		a.Balances[symbol] = Balance{
			Symbol: symbol,
			Cash:   newCash,
		}
	}
	return nil
}

func (a *Account) GetBalance(symbol string) int {
	return a.Balances[symbol].Cash
}

func (a *Account) String() string {
	return fmt.Sprintf("Account{CarNo:%d, Balances:%v}", a.CarNo, a.Balances)
}

type Bank struct {
	bursary  int
	accounts map[int]*Account
}

func NewBank(initialBursary int) *Bank {
	return &Bank{
		bursary:  initialBursary,
		accounts: make(map[int]*Account),
	}
}

func (b *Bank) ArmoredTruck(cash int) bool {
	if cash > 0 {
		b.bursary += cash
		return true
	}
	return false
}

func (b *Bank) AllAccounts() []*Account {
	accounts := make([]*Account, 0, len(b.accounts))
	for _, acc := range b.accounts {
		accounts = append(accounts, acc)
	}
	return accounts
}

func (b *Bank) GetAccount(carNo int) (*Account, error) {
	if acc, ok := b.accounts[carNo]; ok {
		return acc, nil
	}
	return nil, ErrAccountNotFound
}

func (b *Bank) OpenAccount() *Account {
	carNo := generateCardNumber()
	acc := NewAccount(carNo)
	b.accounts[carNo] = acc
	return acc
}

func (b *Bank) Transfer(from, to int, amount int, symbol string) error {
	fromAcc, err := b.GetAccount(from)
	if err != nil {
		return fmt.Errorf("from account: %w", err)
	}

	toAcc, err := b.GetAccount(to)
	if err != nil {
		return fmt.Errorf("to account: %w", err)
	}

	if err := fromAcc.SubBalance(symbol, amount); err != nil {
		return err
	}

	if err := toAcc.AddBalance(symbol, amount); err != nil {
		// Rollback
		fromAcc.AddBalance(symbol, amount)
		return err
	}

	return nil
}

func (b *Bank) BalanceOf(carNo int) (map[string]Balance, error) {
	acc, err := b.GetAccount(carNo)
	if err != nil {
		return nil, err
	}

	// Return a copy to prevent external modification
	balances := make(map[string]Balance)
	for k, v := range acc.Balances {
		balances[k] = v
	}
	return balances, nil
}

func (b *Bank) BalanceOfSymbol(carNo int, symbol string) (int, error) {
	acc, err := b.GetAccount(carNo)
	if err != nil {
		return 0, err
	}
	return acc.GetBalance(symbol), nil
}

func (b *Bank) Deposit(carNo int, amount int, symbol string) error {
	acc, err := b.GetAccount(carNo)
	if err != nil {
		return err
	}
	return acc.AddBalance(symbol, amount)
}

func (b *Bank) Withdraw(carNo int, amount int, symbol string) error {
	acc, err := b.GetAccount(carNo)
	if err != nil {
		return err
	}
	return acc.SubBalance(symbol, amount)
}

type BankSystem struct {
	*Bank
}

func NewBankSystem(initialBursary int) *BankSystem {
	return &BankSystem{
		Bank: NewBank(initialBursary),
	}
}

func generateCardNumber() int {
	return 100000000 + rand.Intn(899999999)
}

// ============ Errors ============
var (
	ErrAccountNotFound        = errors.New("account not found")
	ErrAdditionOverflow       = errors.New("addition overflow")
	ErrSubtractionUnderflow   = errors.New("subtraction underflow")
	ErrMultiplicationOverflow = errors.New("multiplication overflow")
	ErrDivisionByZero         = errors.New("division by zero")
)

// ============ SafeMath ============
func SafeAdd(a, b int) (int, error) {
	if b > 0 && a > math.MaxInt-b {
		return 0, ErrAdditionOverflow
	}
	if b < 0 && a < math.MinInt-b {
		return 0, ErrSubtractionUnderflow
	}
	return a + b, nil
}

func SafeSub(a, b int) (int, error) {
	if b > 0 && a < math.MinInt+b {
		return 0, ErrSubtractionUnderflow
	}
	if b < 0 && a > math.MaxInt+b {
		return 0, ErrAdditionOverflow
	}
	return a - b, nil
}

func main() {
	bs := NewBankSystem(1000000) // 10000.00 USD
	fmt.Printf("Bank initial capital: %d USD\n", bs.bursary)

	fmt.Println("\n============ open first account ============")
	a1 := bs.OpenAccount()
	bs.Deposit(a1.CarNo, 10000, "USD")   // add 100 USD
	bs.Deposit(a1.CarNo, 1530000, "JPY") // add 100 USD Equivalent to Japanese yen

	balances, _ := bs.BalanceOf(a1.CarNo)
	fmt.Printf("account %d balance: %v\n", a1.CarNo, balances)

	fmt.Println("\n============ open first account ============")
	a2 := bs.OpenAccount()
	fmt.Printf("account %d open account successful \n", a2.CarNo)

	fmt.Println("\n============ transfer ============")
	err := bs.Transfer(a1.CarNo, a2.CarNo, 1000, "USD") // 10 USD
	if err != nil {
		fmt.Printf("fail: %v\n", err)
	} else {
		fmt.Println("successful: 10 USD")
	}

	fmt.Println("\n============ account balance ============")
	b1, _ := bs.BalanceOf(a1.CarNo)
	b2, _ := bs.BalanceOf(a2.CarNo)
	fmt.Printf("account %d balance: %v\n", a1.CarNo, b1)
	fmt.Printf("account %d balance: %v\n", a2.CarNo, b2)

	fmt.Println("\n============ bank capital ============")
	fmt.Printf("bank capital: %d USD\n", bs.bursary)

	fmt.Println("\n============ all account ============")
	for _, acc := range bs.AllAccounts() {
		fmt.Printf("%v\n", acc)
	}
}

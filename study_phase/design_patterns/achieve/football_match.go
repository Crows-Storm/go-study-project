package main

// TODO: Apply what you have learned to achieve, case NO2: FootBall Match

type Match struct {
	ID   int
	Home string
	Away string
	Odds float32
	Live string // live / prematch
	// handicap float32
}

type Balance struct {
	Cash   float64 // all balance
	Freeze float64 // freeze cash
}

//type BalanceInterface interface {
//	Balance() float64
//}

func (b *Balance) Balance() float64 {
	return b.Cash - b.Freeze
}

type Player struct {
	ID   int
	Name string
	Balance
}

type Ticker struct {
	ID       int
	MatchId  int
	PlayerId int
	Stake    float64
	Odds     float32
}

type TickerInterface interface {
	Settlement(matchId int) bool
}

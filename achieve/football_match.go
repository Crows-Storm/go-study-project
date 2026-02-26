package main

// TODO: Apply what you have learned to achieve, case NO2: football ticket

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID      uuid.UUID
	Home    string
	Away    string
	Live    bool
	Markets map[int]*Market // 0: 1x2  1: handicap
	Result  [2]int          // final score
}

type Market struct {
	ID         uuid.UUID
	Type       int // 0: 1x2  1: handicap
	Selections map[string]*Selection
	Settled    bool
}

type Selection struct {
	ID   uuid.UUID
	Name string
	Odds float64
}

type Player struct {
	ID uuid.UUID
	Wallet
}

type Wallet struct {
	Cash   float64
	Freeze float64
	mu     sync.Mutex
}

type Ticket struct {
	ID        uuid.UUID
	PlayerID  uuid.UUID
	MatchID   uuid.UUID
	MarketID  uuid.UUID
	SelectKey string
	Stake     float64
	Odds      float64
	Result    string // WIN / LOSE / DRAW
}

type BetRequest struct {
	Player  *Player
	Match   *Match
	Market  int
	Select  string
	Stake   float64
	ReplyCh chan error
}

type Settlement struct {
	Match *Match
}

var (
	betCh       = make(chan BetRequest)
	settleCh    = make(chan Settlement)
	ticketStore = make(map[uuid.UUID]*Ticket)
	ticketMutex sync.Mutex
)

func register() *Player {
	return &Player{
		ID: uuid.New(),
		Wallet: Wallet{
			Cash:   10000,
			Freeze: 0,
		},
	}
}

func (w *Wallet) freeze(amount float64) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.Cash < amount {
		return fmt.Errorf("insufficient balance")
	}
	w.Cash -= amount
	w.Freeze += amount
	return nil
}

func (w *Wallet) settleWin(amount float64, stake float64) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.Freeze -= stake
	w.Cash += amount
}

func (w *Wallet) settleLose(stake float64) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.Freeze -= stake
}

func betEngine() {
	for req := range betCh {

		market, ok := req.Match.Markets[req.Market]
		if !ok {
			req.ReplyCh <- fmt.Errorf("market not found")
			continue
		}

		selection, ok := market.Selections[req.Select]
		if !ok {
			req.ReplyCh <- fmt.Errorf("selection not found")
			continue
		}

		if err := req.Player.freeze(req.Stake); err != nil {
			req.ReplyCh <- err
			continue
		}

		ticket := &Ticket{
			ID:        uuid.New(),
			PlayerID:  req.Player.ID,
			MatchID:   req.Match.ID,
			MarketID:  market.ID,
			SelectKey: req.Select,
			Stake:     req.Stake,
			Odds:      selection.Odds,
		}

		ticketMutex.Lock()
		ticketStore[ticket.ID] = ticket
		ticketMutex.Unlock()

		req.ReplyCh <- nil
	}
}

func settlementEngine() {
	for s := range settleCh {

		for _, market := range s.Match.Markets {

			if market.Settled {
				continue
			}

			for _, ticket := range ticketStore {

				if ticket.MatchID != s.Match.ID || ticket.MarketID != market.ID {
					continue
				}

				win := evaluateResult(s.Match, ticket.SelectKey)

				player := playerStore[ticket.PlayerID]

				if win {
					payout := ticket.Stake * ticket.Odds
					player.settleWin(payout, ticket.Stake)
					ticket.Result = "WIN"
				} else {
					player.settleLose(ticket.Stake)
					ticket.Result = "LOSE"
				}
			}

			market.Settled = true
		}
	}
}

func evaluateResult(match *Match, selectKey string) bool {
	home := match.Result[0]
	away := match.Result[1]

	switch selectKey {
	case "HOME":
		return home > away
	case "AWAY":
		return away > home
	case "DRAW":
		return home == away
	}
	return false
}

var playerStore = make(map[uuid.UUID]*Player)

func main() {

	go betEngine()
	go settlementEngine()

	// create player
	player := register()
	playerStore[player.ID] = player

	// create match
	match := &Match{
		ID:   uuid.New(),
		Home: "Real Madrid",
		Away: "Manchester City",
		Live: false,
		Markets: map[int]*Market{
			0: {
				ID:   uuid.New(),
				Type: 0,
				Selections: map[string]*Selection{
					"HOME": {ID: uuid.New(), Name: "HOME", Odds: 2.2},
					"DRAW": {ID: uuid.New(), Name: "DRAW", Odds: 3},
					"AWAY": {ID: uuid.New(), Name: "AWAY", Odds: 1.8},
				},
			},
		},
	}

	// place bet
	replyCh := make(chan error)

	betCh <- BetRequest{
		Player:  player,
		Match:   match,
		Market:  0,
		Select:  "HOME",
		Stake:   1000,
		ReplyCh: replyCh,
	}

	if err := <-replyCh; err != nil {
		fmt.Println("Bet failed:", err)
	} else {
		fmt.Println("Bet success")
	}

	fmt.Println("Before settle:", player.Cash, player.Freeze)

	// simulate match result
	time.Sleep(time.Second)
	match.Result = [2]int{2, 1}

	settleCh <- Settlement{Match: match}

	time.Sleep(time.Second)

	fmt.Println("After settle:", player.Cash, player.Freeze)
}

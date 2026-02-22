package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type result struct {
	id  int
	op  string
	err error
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	const routines = 10
	const inserts = routines * 2

	ch := make(chan result, inserts)

	waitInserts := inserts

	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertUser(id)

			ch <- insertTrans(id)
		}(i)
	}

	for waitInserts > 0 {
		r := <-ch
		log.Printf("N: %d ID: %d OP: %s ERR: %v", waitInserts, r.id, r.op, r.err)
		waitInserts--
	}

	log.Println("Inserts Complete")
}

func insertUser(id int) result {
	r := result{id: id,
		op: fmt.Sprintf("insert USERS value (%d)", id),
	}

	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}
	return r
}

func insertTrans(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert TRANS value (%d)", id),
	}
	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}
	return r
}

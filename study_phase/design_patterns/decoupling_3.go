package main

//import (
//	"errors"
//	"fmt"
//	"io"
//	"math/rand"
//	"time"
//)
//
//func init() {
//	rand.NewSource(time.Now().UnixNano())
//}
//
//// Data is the structure of the data we are copying.
//type Data struct {
//	Line string
//}
//
//// Xenia is a system we need to pull data from.
//type Xenia struct {
//	Host    string
//	Timeout time.Duration
//}
//
//// Pillar is a system we need to store data into.
//type Pillar struct {
//	Host    string
//	Timeout time.Duration
//}
//
//// Puller declares behavior for pulling data.
//type Puller interface {
//	Pull(d *Data) error
//}
//
//// Storer declares behavior for storing data.
//type Storer interface {
//	Store(d *Data) error
//}
//
//type PullStorer interface {
//	Puller
//	Storer
//	StudentUp
//}
//
//type Student struct {
//	Name string
//}
//type StudentUp interface { // new interface
//	Up() string
//}
//
//func (*Student) Up() string {
//	return "Uppppppp"
//}
//
//func (*Xenia) Pull(d *Data) error {
//	switch rand.Intn(10) {
//	// random number
//	case 1, 9:
//		return io.EOF
//	case 5:
//		return errors.New("Error reading data from Xenia")
//	default:
//		d.Line = "Data"
//		fmt.Println("In: ", d.Line)
//		return nil
//	}
//}
//
//func (*Pillar) Store(d *Data) error {
//	fmt.Println("Out: ", d.Line)
//	return nil
//}
//
//type System struct {
//	Xenia
//	Pillar
//	Student
//}
//
//func pull(ps Puller, data []Data) (int, error) {
//	for i := range data {
//		if err := ps.Pull(&data[i]); err != nil {
//			return i, err
//		}
//	}
//	return len(data), nil
//}
//
//func store(ps Storer, data []Data) (int, error) {
//	for i := range data {
//		if err := ps.Store(&data[i]); err != nil {
//			return i, err
//		}
//	}
//	return len(data), nil
//}
//
//func up(u StudentUp) string {
//	return u.Up()
//}
//
//func Copy(ps PullStorer, batch int) error {
//	data := make([]Data, batch)
//
//	for {
//		i, err := pull(ps, data)
//		if i > 0 {
//			if _, err := store(ps, data[:i]); err != nil {
//				return err
//			}
//		}
//		fmt.Println(up(ps)) // mock a new interface combination
//		if err != nil {
//			return err
//		}
//	}
//}
//
//func main() {
//	sys := System{
//		Xenia: Xenia{
//			Host:    "localhost:8000",
//			Timeout: time.Second,
//		},
//		Pillar: Pillar{
//			Host:    "localhost:9000",
//			Timeout: time.Second,
//		},
//	}
//
//	if err := Copy(&sys, 3); err != io.EOF {
//		fmt.Println(err)
//	}
//}

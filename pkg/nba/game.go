package nba

import (
	"fmt"
	"math/rand"
	"time"
)

var Game []GameInfoPerMin

func StartGame () {
	score:=NewGameInfo()
	start :=time.Now().Unix()
	quit := make(chan bool)
	fmt.Println("start")
	go score.Game(start,quit)
	ticker := time.NewTicker(5 * time.Second)
	i:=0
	for {
		select {
		case <- ticker.C:
			i++
			Game = append(Game,GameInfoPerMin{Minutes: i, GameInfo: score})
		case <- quit:
			ticker.Stop()
			return
		}
	}
}

func (s *GameInfo) Game(start int64,quit chan bool) {
	for {
		res:=rand.Intn(100)
		switch  {
		case res >49:
			if Attack() {
				s.Second.Score++
			}
			s.Second.Attack++

		case res<=49:
			if Attack() {
				s.First.Score++
			}
			s.First.Attack++

		}
		fmt.Println(s)
		fmt.Println("sec:",time.Now().Unix()-start)
		if time.Now().Unix()-start>240 {
			quit<-true
			return
		}

	}
}


func Attack ()bool{
	goal:=false
	rand.Seed(time.Now().Unix())
	if rand.Intn(99)> 49 {
		goal = true
	}
	f:=rand.Intn(1750)+250
	time.Sleep(time.Duration(f)*time.Millisecond)
	return goal
}

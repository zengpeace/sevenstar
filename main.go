package main

import (
	"fmt"
	"sevenstar/game"
	"time"
)

const (
	adminBackPercent   int = 2
	mainProxyPercent   int = 3
	directProxyPercent int = 5
)

var inputData = []game.OneVote{
	game.OneVote{game.OneStatic, game.VoteNumType{'X', 'X', 'X', 9, 'X'}, false, 234, "one"},
	game.OneVote{game.TwoStatic, game.VoteNumType{'X', 'X', 'X', 3, 9}, false, 678, "two"},
	game.OneVote{game.TwoAppear, game.VoteNumType{2, 3, 'X', 'X', 'X'}, false, 111, "three"},
	game.OneVote{game.FourAppear, game.VoteNumType{2, 7, 9, 1, 'X'}, false, 222, "four"},
	game.OneVote{game.TwoStaticOfFive, game.VoteNumType{2, 'X', 'X', 'X', 9}, false, 222, "five"},
}

var oddsPercent = game.VoteWinType{
	100, 200, 300, 400, 500, 600, 700, 800,
}

var finallyResult = game.VoteNumType{
	2, 0, 1, 2, 9,
}

func main() {
	//value := game.ShowOne(&finallyResult, inputData, &oddsPercent)
	//fmt.Println("total win", value)

	//res := game.ShowAll(&inputData, len(inputData), &oddsPercent, true)
	//fmt.Println(res)

	//res, value := game.StopByGivenPercent(&inputData, len(inputData), &oddsPercent, 90)
	//fmt.Println(res, value)

	//res := game.StopByPureRandom(&inputData, len(inputData), &oddsPercent)
	//fmt.Println(res)

	in := game.CreateFakeInputVoteData(2000)
	/*for i := 0; i < len(in); i++ {
		fmt.Println(i, in[i])
	}*/

	/*fmt.Println("before ShowAll:", time.Now())
	game.ShowAll(&in, len(in), &oddsPercent, true)
	fmt.Println("after ShowAll:", time.Now())*/
	/*for i := 0; i < len(in); i++ {
		fmt.Println(i, res[i])
	}*/

	/*fmt.Println("before StopByPureRandom:", time.Now())
	game.StopByPureRandom(&in, len(in), &oddsPercent)
	fmt.Println("after StopByPureRandom:", time.Now())

	fmt.Println("before StopByGivenPercent:", time.Now())
	game.StopByGivenPercent(&in, len(in), &oddsPercent, 80)
	fmt.Println("after StopByGivenPercent:", time.Now())

	fmt.Println("before StopBySetGiveNum:", time.Now())
	game.StopBySetGiveNum(&finallyResult, &in, len(in), &oddsPercent)
	fmt.Println("after StopBySetGiveNum:", time.Now())

	fmt.Println("before StopByGivenPercentAdmin:", time.Now())
	game.StopByGivenPercentAdmin(&in, len(in), &oddsPercent, 80)
	fmt.Println("after StopByGivenPercentAdmin:", time.Now())*/

	var startTime, stopTime, diffMs int64

	startTime = time.Now().UnixNano()
	game.ShowAll(&in, len(in), &oddsPercent, false)
	stopTime = time.Now().UnixNano()
	diffMs = (stopTime - startTime) / 1000000
	fmt.Println("ShowAll:", diffMs, "ms")

	startTime = time.Now().UnixNano()
	game.StopByPureRandom(&in, len(in), &oddsPercent)
	stopTime = time.Now().UnixNano()
	diffMs = (stopTime - startTime) / 1000000
	fmt.Println("StopByPureRandom:", diffMs, "ms")

	startTime = time.Now().UnixNano()
	game.StopByGivenPercent(&in, len(in), &oddsPercent, 80)
	stopTime = time.Now().UnixNano()
	diffMs = (stopTime - startTime) / 1000000
	fmt.Println("StopByGivenPercent:", diffMs, "ms")

	startTime = time.Now().UnixNano()
	game.StopByGivenPercentAdmin(&in, len(in), &oddsPercent, 80)
	stopTime = time.Now().UnixNano()
	diffMs = (stopTime - startTime) / 1000000
	fmt.Println("StopByGivenPercentAdmin:", diffMs, "ms")
}

package game

import (
	"fmt"
	"math/rand"
)

/*
Q:
由于投注量太少，给定开奖率正负3%范围内，并没有合适的结果，要怎么处理

*/

/******************************************************* private *****************************************************************/
/*********************************************************************************************************************************/
/*********************************************************************************************************************************/

const (
	numLen     int = 5
	typeNum    int = 8
	errPercent int = 3
)

func checkWin(one *OneVote, result *VoteNumType) bool {
	switch one.Types {
	case OneStatic:
		return checkStaticWin(&one.Number, result, 1)
	case TwoStatic:
		return checkStaticWin(&one.Number, result, 2)
	case TwoStaticOfFive:
		return checkStaticWin(&one.Number, result, 2)
	case ThreeStatic:
		return checkStaticWin(&one.Number, result, 3)
	case FourStatic:
		return checkStaticWin(&one.Number, result, 4)
	case TwoAppear:
		return checkAppearWin(&one.Number, *result, 2)
	case ThreeAppear:
		return checkAppearWin(&one.Number, *result, 3)
	case FourAppear:
		return checkAppearWin(&one.Number, *result, 4)
	default:
		fmt.Println("unknow voteType ", one.Types)
		return false
	}

	//return false
}

func checkStaticWin(num *VoteNumType, result *VoteNumType, needMatchNum int) bool {
	matchNum := 0
	for i := 0; i < numLen; i++ {
		if num[i] == 'X' {
			continue
		}

		if num[i] == result[i] {
			matchNum++
			if matchNum >= needMatchNum {
				return true
			}
		}
	}

	return false
}

func checkAppearWin(num *VoteNumType, result VoteNumType, needMatchNum int) bool {
	matchNum := 0
	stop := false
	for i := 0; i < numLen; i++ {
		if num[i] == 'X' {
			//continue
			return false
		}

		//only try to match the front four data
		for j := 0; j < numLen-1; j++ {
			if num[i] == result[j] {
				matchNum++
				if matchNum >= needMatchNum {
					stop = true
				} else {
					result[j] = 'X'
				}
				break
			}
		}

		if stop {
			return true
		}
	}

	return false
}

func getRandInt(value int) int {
	//rand.Seed(time.Now().Unix())
	return rand.Intn(value)
}

func showOne(result *VoteNumType, vote *[]OneVote, voteLen int, oddsPercent *VoteWinType) VoteResultType {
	total := 0
	for i := 0; i < voteLen; i++ {
		v := &(*vote)[i]
		total += v.VoteAmount
		ret := checkWin(v, result)
		v.IsWin = ret
		//fmt.Println("vote", i, "is win ?", ret)
		if ret {
			total -= v.VoteAmount * oddsPercent[v.Types] / 100
		} else {
			//
		}
	}

	//fmt.Println(*result, ":", total)
	return VoteResultType{*result, total, 0}
}

/***************************************************** public interface **********************************************************/
/*********************************************************************************************************************************/
/*********************************************************************************************************************************/

/*
ShowAll 展示所有可能五球结果对应的盈亏信息
*/
func ShowAll(vote *[]OneVote, voteLen int, oddsPercent *VoteWinType, only10k bool) []VoteResultType {
	all := make([]VoteResultType, 10000, 20000)
	for one := 0; one < 10; one++ {
		for two := 0; two < 10; two++ {
			for three := 0; three < 10; three++ {
				for four := 0; four < 10; four++ {
					if only10k {
						five := 0
						result := &VoteNumType{one, two, three, four, five}
						res := showOne(result, vote, voteLen, oddsPercent)
						all = append(all, res)
					} else {
						for five := 0; five < 10; five++ {
							result := &VoteNumType{one, two, three, four, five}
							res := showOne(result, vote, voteLen, oddsPercent)
							all = append(all, res)
						}
					}
				}
			}
		}
	}

	return all
}

/*
GetTotalBuyValue 获取总的投注额
*/
func GetTotalBuyValue(vote *[]OneVote, voteLen int) int {
	total := 0
	for i := 0; i < voteLen; i++ {
		v := &(*vote)[i]
		total += v.VoteAmount
	}

	return total
}

/*
StopByGivenPercent 管理员设置开奖比率
*/
func StopByGivenPercent(vote *[]OneVote, voteLen int, oddsPercent *VoteWinType, percent int) (bool, VoteResultType) {
	minPercent := percent - errPercent
	maxPercent := percent + errPercent
	totalBuy := GetTotalBuyValue(vote, voteLen)

	for one := 0; one < 10; one++ {
		for two := 0; two < 10; two++ {
			for three := 0; three < 10; three++ {
				for four := 0; four < 10; four++ {
					for five := 0; five < 10; five++ {
						result := &VoteNumType{one, two, three, four, five}
						res := showOne(result, vote, voteLen, oddsPercent)
						curPercent := (totalBuy - res.Value) * 100 / totalBuy
						res.Percent = curPercent
						//fmt.Println(res)
						if curPercent >= minPercent && curPercent <= maxPercent {
							return true, res
						}
					}
				}
			}
		}
	}

	return false, VoteResultType{}
}

/*
StopBySetGiveNum 计算对于特定五球结构的盈亏信息
ps:the win info of every note is written in (*vote)[i].IsWin
*/
func StopBySetGiveNum(result *VoteNumType, vote *[]OneVote, voteLen int, oddsPercent *VoteWinType) VoteResultType {
	return showOne(result, vote, voteLen, oddsPercent)
}

/*
StopByGivenPercentAdmin 设置实际占系统管理员金额开奖比率
*/
func StopByGivenPercentAdmin(vote *[]OneVote, voteLen int, oddsPercent *VoteWinType, adminBackPercent int) (bool, VoteResultType) {
	minPercent := adminBackPercent - errPercent
	maxPercent := adminBackPercent + errPercent
	totalBuy := GetTotalBuyValue(vote, voteLen)

	for one := 0; one < 10; one++ {
		for two := 0; two < 10; two++ {
			for three := 0; three < 10; three++ {
				for four := 0; four < 10; four++ {
					for five := 0; five < 10; five++ {
						result := &VoteNumType{one, two, three, four, five}
						res := showOne(result, vote, voteLen, oddsPercent)
						curPercent := (totalBuy - res.Value) * 100 / (totalBuy * adminBackPercent / 100)
						res.Percent = curPercent
						if curPercent >= minPercent && curPercent <= maxPercent {
							return true, res
						}
					}
				}
			}
		}
	}

	return false, VoteResultType{}
}

/*
StopByPureRandom 纯随机开奖
*/
func StopByPureRandom(vote *[]OneVote, voteLen int, oddsPercent *VoteWinType) VoteResultType {
	var res VoteNumType
	r := getRandInt(100000)
	res[0] = r / 10000
	res[1] = r / 1000 % 10
	res[2] = r / 100 % 10
	res[3] = r / 10 % 10
	res[4] = r % 10

	return StopBySetGiveNum(&res, vote, voteLen, oddsPercent)
}

/************************************************** for debug *************************************************/
/*********************************************************************************************************************************/
/*********************************************************************************************************************************/

/*
CreateFakeInputVoteData 创建虚拟投注输入表
*/
func CreateFakeInputVoteData(voteNum int) []OneVote {
	votes := make([]OneVote, 1, 2)

	for i := 0; i < voteNum; i++ {
		var vt VoteNumType
		for j := 0; j < numLen; j++ {
			value := getRandInt(1000) % 15
			if value > 9 {
				value = 'X'
			}

			vt[j] = value
		}

		if vt[0] == 'X' && vt[1] == 'X' && vt[2] == 'X' && vt[3] == 'X' && vt[4] == 'X' {
			vt[0] = 7
		} else if vt[0] != 'X' && vt[1] != 'X' && vt[2] != 'X' && vt[3] != 'X' && vt[4] != 'X' {
			vt[4] = 'X'
		}

		xnum := 0
		for k := 0; k < numLen; k++ {
			if vt[k] == 'X' {
				xnum++
			}
		}

		var voteTypes int
		notxnum := numLen - xnum
		switch notxnum {
		case 1:
			voteTypes = OneStatic
		case 2:
			voteTypes = TwoStaticOfFive
		case 3:
			voteTypes = ThreeStatic
		case 4:
			voteTypes = FourStatic
		}

		continuous := 0
		for k := 0; k < numLen; k++ {
			if vt[k] == 'X' {
				break
			} else {
				continuous++
			}
		}

		if notxnum == continuous {
			switch notxnum {
			case 1:
				voteTypes = OneStatic
			case 2:
				voteTypes = TwoAppear
			case 3:
				voteTypes = ThreeAppear
			case 4:
				voteTypes = FourAppear
			}

		}

		id := fmt.Sprintf("%d", i)
		v := OneVote{voteTypes, vt, false, 100, id}
		votes = append(votes, v)
	}

	return votes[1:]
}

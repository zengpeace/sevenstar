package game

/*
投注类型：共八种
一字定		X,X,X,X,7
二字定		X,X,2,6,X	(最后一位必须为X)
五位二字定	X,X,1,X,3
三字定		X,2,2,6,X	(最后一位必须为X)
四字定		1,5,2,6,X	(最后一位必须为X)
二字现		2,4,X,X,X	(最后三位必须为X)
三字现		0,2,5,X,X	(最后二位必须为X)
四字现		1,2,4,8,X	(最后一位必须为X)
*/
const (
	OneStatic = iota
	TwoStatic
	TwoStaticOfFive
	ThreeStatic
	FourStatic
	TwoAppear
	ThreeAppear
	FourAppear
)

/*
VoteWinType 每种投注类型的赔率表的类型
*/
type VoteWinType [8]int

/*
VoteNumType 每一注数字的类型
*/
type VoteNumType [5]int

/*
OneVote 每一注的所有信息

参数规则
Types			投注类型，就是上面的八种
Number			投注的球值，0~9的五个数
IsWin			这局有没有中奖
VoteAmount		投注额
ID				用户id
*/
type OneVote struct {
	Types      int
	Number     VoteNumType
	IsWin      bool
	VoteAmount int
	ID         string
}

/*
VoteResultType 对于某一个五球结果的庄家收入类型
Number 五球结果
Value  庄家收入额（不算抽水）
*/
type VoteResultType struct {
	Number  VoteNumType
	Value   int
	Percent int
}

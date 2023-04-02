package redisConstant

import "time"

const (
	LikeSetKeyPrefix      = "cache:giligili:video:likeSet:"
	LikeCountKeyPrefix    = "cache:giligili:video:likeCount:"
	DislikeSetKeyPrefix   = "cache:giligili:video:dislikeSet:"
	DislikeCountKeyPrefix = "cache:giligili:video:dislikeCount:"
	RankingListKeyPrefix  = "cache:giligili:video:rankingList:"
	BarrageListKeyPrefix  = "cache:giligili:video:barrageList:"
	BarrageListExpire     = time.Minute * 5
)

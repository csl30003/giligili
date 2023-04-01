package types

type Barrage struct {
	UserId    int64  `json:"userId"`
	VideoId   int64  `json:"videoId"`
	Text      string `json:"text"`
	Color     int64  `json:"color"`
	Type      int64  `json:"type"`
	Timestamp int64  `json:"timestamp"`
}

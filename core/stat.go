package core

type Stat struct {
	Buffer *BufferStat `json:"buffer"`
}
type BufferStat struct {
	Size    uint64 `json:"size"`
	UsedCnt uint64 `json:"used_cnt"`
	Rate    uint64 `json:"rate"`
}

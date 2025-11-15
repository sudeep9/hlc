package hlc

var globalHLC *HLC

func init() {
	globalHLC = NewHLC()
}

func Next() int64 {
	return globalHLC.Next()
}

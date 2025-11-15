package hlc

var globalHLC *HLC

func init() {
	globalHLC = NewHLC(0)
}

func Next() int64 {
	return globalHLC.Next()
}

func SetLastTS(ts int64) {
	globalHLC.SetLastTS(ts)
}

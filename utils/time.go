package utils

import "time"

func TimeJST(t time.Time) time.Time {
	return t.In(time.FixedZone("Asia/Tokyo", 9*60*60))
}

package format

import (
	"time"

	utils "github.com/cyclo9/goutils"
)

// Formats Unix time in seconds to "15:04:05"
func UnixTime(unixTimeSec int64) string {
	timeObj := time.Unix(unixTimeSec, 0)

	loc, err := time.LoadLocation("America/Los_Angeles")
	utils.HandleErr(err, "loading location")

	pacificTime := timeObj.In(loc)

	return pacificTime.Format("15:04:05 01/02/2006")
}

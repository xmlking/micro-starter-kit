package time

import (
	"time"

	"github.com/golang/protobuf/ptypes"
    "github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
    "github.com/rs/zerolog/log"
)

// GetCurrentTime
func GetCurrentTime() *timestamp.Timestamp {
    now := time.Now()
    return ToTimestamp(now)
}

// ToTimestamp returns a protobuf Timestamp from a Time object
func ToTimestamp(t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	if nil != err {
		log.Error().Err(err).Msg("Time to Timestamp error")
	}
	return ts
}

// ToTime returns a Time object from a protobuf Timestamp
func ToTime(ts *timestamp.Timestamp) time.Time {
	if nil == ts {
		return time.Time{}
	}
	t, err := ptypes.Timestamp(ts)
	if nil != err {
        log.Error().Err(err).Msg("Timestamp to Times error")
	}
	return t
}

// ToDuration returns a Time object from a protobuf Duration
func ToDuration(d *duration.Duration) time.Duration {
    dur, err := ptypes.Duration(d)
    if err != nil {
        log.Error().Err(err).Send()
    }
    return dur
}

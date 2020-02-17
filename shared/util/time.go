package util

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/xmlking/micro-starter-kit/shared/micro/logger"
)

// TimeToTimestamp returns a protobuf Timestamp from a Time object
func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	if nil != err {
		log.WithError(err, "Time to Timestamp error")
	}
	return ts
}

// TimestampToTime returns a Time object from a protobuf Timestamp
func TimestampToTime(ts *timestamp.Timestamp) time.Time {
	if nil == ts {
		return time.Time{}
	}
	t, err := ptypes.Timestamp(ts)
	if nil != err {
		log.WithError(err, "Timestamp to Times error")
	}
	return t
}

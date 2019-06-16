package util

import (
	"time"

	"github.com/micro/go-micro/util/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// TimeToTimestamp returns a protobuf Timestamp from a Time object
func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	if nil != err {
		log.Fatalf("Time to Timestamp error: %s", err.Error())
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
		log.Fatalf("Timestamp to Times error: %s", err.Error())
	}
	return t
}


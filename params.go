package sortby

import (
	"gopkg.in/mgo.v2/bson"
)

// ParamSortByID only use for function ByID
type ParamSortByID struct {
	Indicator map[bson.ObjectId]int // will order by this indicator from 0 to len(indicator)
	Value     interface{}           // pointer from slice of struct, if not pointer will panic
	Reverse   bool                  // when true will reverse order from indicator
}

type byID ParamSortByID

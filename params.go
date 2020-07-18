package sortby

import (
	"github.com/globalsign/mgo/bson"
)

// ParamMongoID will sort by mongodb object ID
type ParamMongoID struct {
	Indicator      []bson.ObjectId // will order by this indicator from 0 to len(indicator)
	Value          interface{}     // pointer from slice of struct, if not pointer will panic
	Reverse        bool            // when true will reverse order from indicator
	FieldName      string          // field that want to reference for sorting
	indicatorIndex map[bson.ObjectId]int
}

type byID ParamMongoID

// ParamsByString will sort by string index
type ParamsByString struct {
	Indicator      []string    // will order by this indicator from 0 to len(indicator)
	Value          interface{} // pointer from slice of struct, if not pointer will panic
	Reverse        bool        // when true will reverse order from indicator
	FieldName      string      // field that want to reference for sorting
	indicatorIndex map[string]int
}

type byString ParamsByString

package sortby

import (
	"github.com/globalsign/mgo/bson"
)

// ParamMongoID will sort by mongodb object ID
type paramMongoID struct {
	indicator      []bson.ObjectId // will order by this indicator from 0 to len(indicator)
	data           interface{}     // pointer from slice of struct, if not pointer will panic
	reverse        bool            // when true will reverse order from indicator
	fieldName      string          // field that want to reference for sorting
	indicatorIndex map[bson.ObjectId]int
}

type byID paramMongoID

// ParamsByString will sort by string index
type paramsByString struct {
	indicator      []string    // will order by this indicator from 0 to len(indicator)
	data           interface{} // pointer from slice of struct, if not pointer will panic
	reverse        bool        // when true will reverse order from indicator
	fieldName      string      // field that want to reference for sorting
	indicatorIndex map[string]int
}

type byString paramsByString

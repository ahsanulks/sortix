package sortby

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/globalsign/mgo/bson"
)

var referenceFieldByID = []string{"Id", "ID"}

// ByID will sort pointer slice of struct with field Id or ID. if struct doesn't have field Id or ID will give panic
func ByID(ps ParamSortByID) {
	resultv := reflect.ValueOf(ps.Value)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("sort value not pointer")
	}
	sort.Sort(byID(ps))
}

func (s byID) Len() int {
	return reflect.ValueOf(s.Value).Elem().Len()
}

func (s byID) Swap(i, j int) {
	v := reflect.ValueOf(s.Value).Elem()
	tempI := v.Index(i)
	tempJ := v.Index(j)
	temp := tempI.Interface()
	v.Index(i).Set(tempJ)
	v.Index(j).Set(reflect.ValueOf(temp))
}

func (s byID) Less(i, j int) bool {
	v := reflect.ValueOf(s.Value).Elem()
	existingField := existField(v.Index(i), referenceFieldByID)
	indexI := s.Indicator[v.Index(i).FieldByName(existingField).Interface().(bson.ObjectId)]
	indexJ := s.Indicator[v.Index(j).FieldByName(existingField).Interface().(bson.ObjectId)]
	if s.Reverse {
		return indexI > indexJ
	}
	return indexI < indexJ
}

func existField(v reflect.Value, fields []string) string {
	var validField string
	for _, field := range fields {
		getField := v.FieldByName(field)
		if getField.IsValid() {
			validField = field
			break
		}
	}
	if validField != "" {
		return validField
	}
	errMessage := fmt.Sprintf("struct doesn't have field %v", fields)
	panic(errMessage)
}

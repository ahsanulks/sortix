package sortby

import "sort"

type sortingBy interface {
	SetIndicatorIndex()
	CheckFieldName() error
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func executeSort(params sortingBy) error {
	err := params.CheckFieldName()
	if err != nil {
		return err
	}
	params.SetIndicatorIndex()
	sort.Sort(params)
	return nil
}

// func MongoDBID([]bson.ObjectID, interface{}) {

// }

// func Strings([]string, interface{}) {

// }

// func Ints([]ints, interface{}) {

// }

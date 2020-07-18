package sortby

import (
	"errors"
	"reflect"
)

//StringSort will sorting by index slice of string
func StringSort(ps ParamsByString) error {
	resultv := reflect.ValueOf(ps.Value)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return errors.New("sort value not pointer")
	}
	exec := byString(ps)
	return executeSort(&exec)
}

func (s *byString) SetIndicatorIndex() {
	indicatorIndex := make(map[string]int)
	for i, v := range s.Indicator {
		indicatorIndex[v] = i
	}
	s.indicatorIndex = indicatorIndex
}

func (s *byString) CheckFieldName() error {
	if s.FieldName == "" {
		return errors.New("FieldName is Required")
	}
	v := reflect.ValueOf(s.Value).Elem()
	getField := v.Index(0).FieldByName(s.FieldName)
	if !getField.IsValid() {
		return errors.New("field not found")
	}
	return nil
}

func (s *byString) Len() int {
	return reflect.ValueOf(s.Value).Elem().Len()
}

func (s *byString) Less(i, j int) bool {
	v := reflect.ValueOf(s.Value).Elem()
	indexI := s.indicatorIndex[v.Index(i).FieldByName(s.FieldName).Interface().(string)]
	indexJ := s.indicatorIndex[v.Index(j).FieldByName(s.FieldName).Interface().(string)]
	if s.Reverse {
		return indexI > indexJ
	}
	return indexI < indexJ
}

func (s *byString) Swap(i, j int) {
	v := reflect.ValueOf(s.Value).Elem()
	tempI := v.Index(i)
	tempJ := v.Index(j)
	temp := tempI.Interface()
	v.Index(i).Set(tempJ)
	v.Index(j).Set(reflect.ValueOf(temp))
}

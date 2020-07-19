package sortix

import (
	"errors"
	"reflect"
	"sort"
)

// Strings will sorting by index slice of string
func Strings(data interface{}, reference []string) (sortixService, error) {
	resultv := reflect.ValueOf(data)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return nil, errors.New("sort value not pointer")
	}
	params := paramsByString{
		indicator: reference,
		data:      data,
	}
	service := byString(params)
	return &service, nil
}

func (s *byString) SortBy(fieldName string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("error when sorting")
		}
	}()
	if s.Len() == 0 || s.Len() == 1 {
		return nil
	}
	s.fieldName = fieldName
	err = s.CheckFieldName()
	if err != nil {
		return err
	}
	s.SetIndicatorIndex()
	sort.Sort(s)
	return nil
}

func (s *byString) ReverseSortBy(fieldName string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("error when sorting")
		}
	}()
	if s.Len() == 0 || s.Len() == 1 {
		return nil
	}
	s.fieldName = fieldName
	err = s.CheckFieldName()
	if err != nil {
		return err
	}
	s.SetIndicatorIndex()
	s.reverse = true
	sort.Sort(s)
	return nil
}

func (s *byString) SetIndicatorIndex() {
	indicatorIndex := make(map[string]int)
	for i, v := range s.indicator {
		indicatorIndex[v] = i
	}
	s.indicatorIndex = indicatorIndex
}

func (s *byString) CheckFieldName() error {
	if s.fieldName == "" {
		return errors.New("FieldName is Required")
	}
	v := reflect.ValueOf(s.data).Elem()
	getField := v.Index(0).FieldByName(s.fieldName)
	if !getField.IsValid() {
		return errors.New("field not found")
	}
	return nil
}

func (s *byString) Len() int {
	return reflect.ValueOf(s.data).Elem().Len()
}

func (s *byString) Less(i, j int) bool {
	v := reflect.ValueOf(s.data).Elem()
	indexI := s.indicatorIndex[v.Index(i).FieldByName(s.fieldName).Interface().(string)]
	indexJ := s.indicatorIndex[v.Index(j).FieldByName(s.fieldName).Interface().(string)]
	if s.reverse {
		return indexI > indexJ
	}
	return indexI < indexJ
}

func (s *byString) Swap(i, j int) {
	v := reflect.ValueOf(s.data).Elem()
	tempI := v.Index(i)
	tempJ := v.Index(j)
	temp := tempI.Interface()
	v.Index(i).Set(tempJ)
	v.Index(j).Set(reflect.ValueOf(temp))
}

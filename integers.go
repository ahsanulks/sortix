package sortix

import (
	"errors"
	"reflect"
	"sort"
)

// Integers will sorting by index slice of int
func Integers(data interface{}, reference []int) (SortixService, error) {
	resultv := reflect.ValueOf(data)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return nil, errors.New("sort value not pointer")
	}
	params := paramsByInteger{
		indicator: reference,
		data:      data,
	}
	service := byInteger(params)
	return &service, nil
}

func (s *byInteger) SortBy(fieldName string) (err error) {
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

func (s *byInteger) ReverseSortBy(fieldName string) (err error) {
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

func (s *byInteger) SetIndicatorIndex() {
	indicatorIndex := make(map[int]int)
	for i, v := range s.indicator {
		indicatorIndex[v] = i
	}
	s.indicatorIndex = indicatorIndex
}

func (s *byInteger) CheckFieldName() error {
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

func (s *byInteger) Len() int {
	return reflect.ValueOf(s.data).Elem().Len()
}

func (s *byInteger) Less(i, j int) bool {
	v := reflect.ValueOf(s.data).Elem()
	indexI := s.indicatorIndex[v.Index(i).FieldByName(s.fieldName).Interface().(int)]
	indexJ := s.indicatorIndex[v.Index(j).FieldByName(s.fieldName).Interface().(int)]
	if s.reverse {
		return indexI > indexJ
	}
	return indexI < indexJ
}

func (s *byInteger) Swap(i, j int) {
	v := reflect.ValueOf(s.data).Elem()
	tempI := v.Index(i)
	tempJ := v.Index(j)
	temp := tempI.Interface()
	v.Index(i).Set(tempJ)
	v.Index(j).Set(reflect.ValueOf(temp))
}

package sortix

import (
	"errors"
	"reflect"
	"sort"

	"github.com/globalsign/mgo/bson"
)

// MongoID will sort pointer slice of struct based on field name. if struct doesn't have field name will sort by Id
func MongoID(data interface{}, reference []bson.ObjectId) (SortixService, error) {
	resultv := reflect.ValueOf(data)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return nil, errors.New("sort value not pointer slice")
	}
	params := paramMongoID{
		indicator: reference,
		data:      data,
	}
	service := byID(params)
	return &service, nil
}

func (s *byID) SortBy(fieldName string) (err error) {
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

func (s *byID) ReverseSortBy(fieldName string) (err error) {
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

func (s *byID) SetIndicatorIndex() {
	indicatorIndex := make(map[bson.ObjectId]int)
	for i, v := range s.indicator {
		indicatorIndex[v] = i
	}
	s.indicatorIndex = indicatorIndex
}

func (s *byID) CheckFieldName() error {
	if s.fieldName == "" {
		s.fieldName = "Id"
	}
	v := reflect.ValueOf(s.data).Elem()
	getField := v.Index(0).FieldByName(s.fieldName)
	if !getField.IsValid() {
		return errors.New("field not found")
	}
	return nil
}

func (s *byID) Len() int {
	return reflect.ValueOf(s.data).Elem().Len()
}

func (s *byID) Less(i, j int) bool {
	v := reflect.ValueOf(s.data).Elem()
	indexI := s.indicatorIndex[v.Index(i).FieldByName(s.fieldName).Interface().(bson.ObjectId)]
	indexJ := s.indicatorIndex[v.Index(j).FieldByName(s.fieldName).Interface().(bson.ObjectId)]
	if s.reverse {
		return indexI > indexJ
	}
	return indexI < indexJ
}

func (s *byID) Swap(i, j int) {
	v := reflect.ValueOf(s.data).Elem()
	tempI := v.Index(i)
	tempJ := v.Index(j)
	temp := tempI.Interface()
	v.Index(i).Set(tempJ)
	v.Index(j).Set(reflect.ValueOf(temp))
}

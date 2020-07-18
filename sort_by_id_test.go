package sortby_test

import (
	"testing"

	sortby "github.com/ahsanulks/sort-by"
	"github.com/bxcodec/faker"
	"gopkg.in/mgo.v2/bson"
)

type TestingData struct {
	Id   bson.ObjectId
	Name string
}

type InvalidTestingData struct {
	Name string
}

func createTestingData(len int) []TestingData {
	var result []TestingData
	for i := 0; i < len; i++ {
		testData := TestingData{
			Id:   bson.NewObjectId(),
			Name: faker.FirstName,
		}
		result = append(result, testData)
	}
	return result
}

func createInvalidTestingData(len int) []InvalidTestingData {
	var result []InvalidTestingData
	for i := 0; i < len; i++ {
		testData := InvalidTestingData{
			Name: faker.FirstName,
		}
		result = append(result, testData)
	}
	return result
}

func TestByID(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	data := createTestingData(len(sortingIndicatorIndex))
	sortingIndicator := make(map[bson.ObjectId]int)
	var sortedData []TestingData
	for i, v := range sortingIndicatorIndex {
		sortingIndicator[data[v].Id] = i
		sortedData = append(sortedData, data[v])
	}
	normalArgs := sortby.ParamSortByID{
		Indicator: sortingIndicator,
		Value:     &data,
		Reverse:   false,
	}
	invalidValue := sortby.ParamSortByID{
		Indicator: sortingIndicator,
		Value:     data,
		Reverse:   false,
	}
	invalidData := createInvalidTestingData(len(sortingIndicatorIndex))
	invalidNotHaveField := sortby.ParamSortByID{
		Indicator: sortingIndicator,
		Value:     &invalidData,
		Reverse:   false,
	}
	tests := []struct {
		name  string
		args  sortby.ParamSortByID
		panic bool
	}{
		// TODO: Add test cases.
		{
			name:  "success sorting with field Id",
			args:  normalArgs,
			panic: false,
		},
		{
			name:  "should be panic when not pointer",
			args:  invalidValue,
			panic: true,
		},
		{
			name:  "should be panic when doesn have field ID or Id",
			args:  invalidNotHaveField,
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				shouldPanic(t, sortby.ByID, tt.args)
				return
			}
			sortby.ByID(tt.args)
			for i, v := range data {
				if v != sortedData[i] {
					t.Errorf("data not same")
					return
				}
			}
		})
	}
}

func TestByIDReverse(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	sortingIndicator := make(map[bson.ObjectId]int)
	var sortedData []TestingData
	for i, v := range sortingIndicatorIndex {
		sortingIndicator[data[v].Id] = i
		sortedData = append(sortedData, data[v])
	}
	normalArgs := sortby.ParamSortByID{
		Indicator: sortingIndicator,
		Value:     &data,
		Reverse:   true,
	}
	tests := []struct {
		name string
		args sortby.ParamSortByID
	}{
		// TODO: Add test cases.
		{
			name: "success sorting with field Id",
			args: normalArgs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortby.ByID(tt.args)
			for i, v := range data {
				if v != sortedData[lenData-i-1] {
					t.Errorf("data not same")
					return
				}
			}
		})
	}
}

func shouldPanic(t *testing.T, f func(sortby.ParamSortByID), p sortby.ParamSortByID) {
	defer func() { recover() }()
	f(p)
	t.Errorf("should have panicked")
}

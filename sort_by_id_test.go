package sortby_test

import (
	"testing"

	sortby "github.com/ahsanulks/sort-by"
	"github.com/bxcodec/faker"
	"github.com/globalsign/mgo/bson"
)

type TestingData struct {
	Id   bson.ObjectId
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

func TestMongoID(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	data := createTestingData(len(sortingIndicatorIndex))
	var sortingIndicator []bson.ObjectId
	var sortedData []TestingData
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Id)
		sortedData = append(sortedData, data[v])
	}
	normalArgs := sortby.ParamMongoID{
		Indicator: sortingIndicator,
		Value:     &data,
		Reverse:   false,
	}
	invalidValue := sortby.ParamMongoID{
		Indicator: sortingIndicator,
		Value:     data,
		Reverse:   false,
		FieldName: "Id",
	}
	tests := []struct {
		name      string
		args      sortby.ParamMongoID
		wantError bool
	}{
		// TODO: Add test cases.
		{
			name:      "success sorting with field Id",
			args:      normalArgs,
			wantError: false,
		},
		{
			name:      "should be panic when not pointer",
			args:      invalidValue,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := sortby.MongoID(tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("should be error")
				return
			}
			for i, v := range data {
				if v != sortedData[i] {
					t.Errorf("data not same")
					return
				}
			}
		})
	}
}

func TestMongoIDReverse(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	var sortingIndicator []bson.ObjectId
	var sortedData []TestingData
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Id)
		sortedData = append(sortedData, data[v])
	}
	normalArgs := sortby.ParamMongoID{
		Indicator: sortingIndicator,
		Value:     &data,
		Reverse:   true,
		FieldName: "Id",
	}
	tests := []struct {
		name string
		args sortby.ParamMongoID
	}{
		// TODO: Add test cases.
		{
			name: "success sorting with field Id",
			args: normalArgs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortby.MongoID(tt.args)
			for i, v := range data {
				if v != sortedData[lenData-i-1] {
					t.Errorf("data not same")
					return
				}
			}
		})
	}
}

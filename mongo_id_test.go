package sortby_test

import (
	"testing"

	"github.com/ahsanulks/sortby"
	"github.com/globalsign/mgo/bson"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

type TestingData struct {
	Id    bson.ObjectId
	Name  string
	Email string
}

func createTestingData(len int) []TestingData {
	var result []TestingData
	for i := 0; i < len; i++ {
		testData := TestingData{
			Id:    bson.NewObjectId(),
			Name:  fake.CharactersN(10),
			Email: fake.EmailAddress(),
		}
		result = append(result, testData)
	}
	return result
}

func Test_byID_SortBy(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	data2 := data
	var sortingIndicator []bson.ObjectId
	var sortedData []TestingData
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Id)
		sortedData = append(sortedData, data[v])
	}
	var noData []TestingData
	type args struct {
		data      interface{}
		indicator []bson.ObjectId
		fieldName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		len0    bool
	}{
		// TODO: Add test cases.
		{
			name:    "field name not found",
			args:    args{&data, sortingIndicator, "ID"},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "use default field name",
			args:    args{&data, sortingIndicator, ""},
			wantErr: false,
			len0:    false,
		},
		{
			name:    "len data 0",
			args:    args{&noData, sortingIndicator, "Id"},
			wantErr: false,
			len0:    true,
		},
		{
			name:    "normal case",
			args:    args{&data, sortingIndicator, "Id"},
			wantErr: false,
			len0:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortService, _ := sortby.MongoID(tt.args.data, tt.args.indicator)
			err := sortService.SortBy(tt.args.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("byID.SortBy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if tt.len0 && !tt.wantErr {
				return
			}
			for i, v := range data {
				if v != sortedData[i] {
					t.Errorf("data not sorted")
					return
				}
			}
			data = data2
		})
	}
}

func Test_byID_ReverseSortBy(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	data2 := data
	var sortingIndicator []bson.ObjectId
	var sortedData []TestingData
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Id)
		sortedData = append(sortedData, data[v])
	}
	var noData []TestingData
	type args struct {
		data      interface{}
		indicator []bson.ObjectId
		fieldName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		len0    bool
	}{
		// TODO: Add test cases.
		{
			name:    "field name not found",
			args:    args{&data, sortingIndicator, "ID"},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "use default field name",
			args:    args{&data, sortingIndicator, ""},
			wantErr: false,
			len0:    false,
		},
		{
			name:    "len data 0",
			args:    args{&noData, sortingIndicator, "Id"},
			wantErr: false,
			len0:    true,
		},
		{
			name:    "normal case",
			args:    args{&data, sortingIndicator, "Id"},
			wantErr: false,
			len0:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortService, _ := sortby.MongoID(tt.args.data, tt.args.indicator)
			err := sortService.ReverseSortBy(tt.args.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("byID.RverseSortBy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if tt.len0 && !tt.wantErr {
				return
			}
			for i, v := range data {
				if v != sortedData[lenData-i-1] {
					t.Errorf("data not sorted")
					return
				}
			}
			data = data2
		})
	}
}

func TestMongoID(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	var sortingIndicator []bson.ObjectId
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Id)
	}
	type args struct {
		data      interface{}
		reference []bson.ObjectId
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantNil bool
	}{
		// TODO: Add test cases.
		{
			name:    "normal case",
			args:    args{&data, sortingIndicator},
			wantErr: false,
			wantNil: false,
		},
		{
			name:    "not pointer case",
			args:    args{data, sortingIndicator},
			wantErr: true,
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortby.MongoID(tt.args.data, tt.args.reference)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantNil {
				assert.NotNil(t, got)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

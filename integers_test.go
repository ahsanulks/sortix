package sortix_test

import (
	"testing"

	"github.com/ahsanulks/sortix"
	sortby "github.com/ahsanulks/sortix"
	"github.com/stretchr/testify/assert"
)

func Test_byInteger_SortBy(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	data2 := data
	var sortingIndicator []int
	var sortedData []TestingData
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Phone)
		sortedData = append(sortedData, data[v])
	}
	var noData []TestingData
	type args struct {
		data      interface{}
		indicator []int
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
			args:    args{&data, sortingIndicator, "Telephone"},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "when not input fieldName",
			args:    args{&data, sortingIndicator, ""},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "len data 0",
			args:    args{&noData, sortingIndicator, "Phone"},
			wantErr: false,
			len0:    true,
		},
		{
			name:    "when field not a integer",
			args:    args{&data, sortingIndicator, "Id"},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "normal case",
			args:    args{&data, sortingIndicator, "Phone"},
			wantErr: false,
			len0:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortService, _ := sortix.Integers(tt.args.data, tt.args.indicator)
			err := sortService.SortBy(tt.args.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("byIntegers.SortBy() error = %v, wantErr %v", err, tt.wantErr)
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

func Test_byInteger_ReverseSortBy(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	data2 := data
	var sortingIndicator []int
	var sortedData []TestingData
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Phone)
		sortedData = append(sortedData, data[v])
	}
	var noData []TestingData
	type args struct {
		data      interface{}
		indicator []int
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
			args:    args{&data, sortingIndicator, "Telephone"},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "when not input fieldName",
			args:    args{&data, sortingIndicator, ""},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "len data 0",
			args:    args{&noData, sortingIndicator, "Phone"},
			wantErr: false,
			len0:    true,
		},
		{
			name:    "when field not a integer",
			args:    args{&data, sortingIndicator, "Id"},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "when field not a integer",
			args:    args{&data, sortingIndicator, "Id"},
			wantErr: true,
			len0:    false,
		},
		{
			name:    "normal case",
			args:    args{&data, sortingIndicator, "Phone"},
			wantErr: false,
			len0:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortService, _ := sortby.Integers(tt.args.data, tt.args.indicator)
			err := sortService.ReverseSortBy(tt.args.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("byIntegers.ReverseSortBy() error = %v, wantErr %v", err, tt.wantErr)
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

func TestIntegers(t *testing.T) {
	sortingIndicatorIndex := []int{0, 2, 3, 4, 1}
	lenData := len(sortingIndicatorIndex)
	data := createTestingData(lenData)
	var sortingIndicator []int
	for _, v := range sortingIndicatorIndex {
		sortingIndicator = append(sortingIndicator, data[v].Phone)
	}
	type args struct {
		data      interface{}
		reference []int
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
			got, err := sortby.Integers(tt.args.data, tt.args.reference)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortby.Integers() error = %v, wantErr %v", err, tt.wantErr)
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

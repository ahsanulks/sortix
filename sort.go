package sortix

type sortingBy interface {
	SetIndicatorIndex()
	CheckFieldName() error
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// SortixService is for sort or reverse sort by field name
type sortixService interface {
	SortBy(fieldName string) error        // will sort data from reference by fieldname. fieldname is case sensitive
	ReverseSortBy(fieldName string) error // will sort data from reference by fieldname with reverse order. fieldname is case sensitive
}

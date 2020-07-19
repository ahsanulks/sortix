package sortix

type sortingBy interface {
	SetIndicatorIndex()
	CheckFieldName() error
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// SortixService is for sort or reverse sort by field name
type SortixService interface {
	SortBy(fieldName string) error
	ReverseSortBy(fieldName string) error
}

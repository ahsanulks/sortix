# Sortix
Sortix is custom sort based on [golang sort](https://golang.org/pkg/sort/#Sort) with reference indicator. You can sort slice of struct based on indicator you want.

# Install
```
go get github.com/ahsanulks/sortix
```

# Import
```
import (
  "github.com/ahsanulks/sortix"
)
```

# Test
To run project test run with
```
make test
```

# Examples

## With Reference Order
```
user := []struct {
  Id   int
  Name string
}{
  {
    Id: 5,
    Name: "John",
  },
  {
    Id: 1,
    Name: "Wing",
  },
  {
    Id: 10,
    Name: "Ali",
  },
}

sortIndicator := []int{10,5,1}
srtix, err := sortix.Integers(&user,sortIndicator)
if err != nil {
  panic(err)
}
srtix.SortBy("Id")
// Output: [{Id:10 Name:Ali} {Id:5 Name:John} {Id:1 Name:Wing}]
```

## With Reverse Reference Order
```
user := []struct {
  Id   int
  Name string
}{
  {
    Id: 5,
    Name: "John",
  },
  {
    Id: 1,
    Name: "Wing",
  },
  {
    Id: 10,
    Name: "Ali",
  },
}

indicator := []int{10,5,1}
srtix, err := sortix.Integers(&user,indicator)
if err != nil {
  panic(err)
}
srtix.ReverseSortBy("Id") // must valid field name on struct
// Output: [{Id:1 Name:Wing} {Id:5 Name:John} {Id:10 Name:Ali}]
```

# Available Sort Data
|Data type|Implementation                    |
|---------|----------------------------------|
|Integer  |`sortix.Integers(&data,indicator)`|
|String   |`sortix.Strings(&data,indicator)` |
|MongoId  |`sortix.MongoID(&data,indicator)` |

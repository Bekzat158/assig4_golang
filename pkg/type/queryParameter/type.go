package queryParameter

import (
	"Bekzat/pkg/type/pagination"
	"Bekzat/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}

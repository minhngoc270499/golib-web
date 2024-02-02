package golibweb

// PagingResource represent paging resource
type PagingResource struct {
	Total       int         `json:"total,omitempty"`
	TotalPage   int         `json:"total_page,omitempty"`
	CurrentPage int         `json:"current_page,omitempty"`
	PerPage     int         `json:"per_page,omitempty"`
	Items       interface{} `json:"items"`
}

// NewPagingResource return a new paging resource instance
func NewPagingResource() *PagingResource {
	return &PagingResource{}
}

// WithTotal set total
func (r *PagingResource) WithTotal(total int) *PagingResource {
	r.Total = total
	return r
}

// WithTotalPage set total page
func (r *PagingResource) WithTotalPage(totalPage int) *PagingResource {
	r.TotalPage = totalPage
	return r
}

// WithCurrentPage set current page
func (r *PagingResource) WithCurrentPage(currentPage int) *PagingResource {
	r.CurrentPage = currentPage
	return r
}

// WithPerPage set per page
func (r *PagingResource) WithPerPage(perPage int) *PagingResource {
	r.PerPage = perPage
	return r
}

// WithItems set items
func (r *PagingResource) WithItems(items interface{}) *PagingResource {
	r.Items = items
	return r
}

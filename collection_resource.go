package golibweb

// CollectionResource represent collection resource
type CollectionResource struct {
	Items interface{} `json:"items"`
}

// NewCollectionResource return a new collection resource instance
func NewCollectionResource(items interface{}) *CollectionResource {
	return &CollectionResource{Items: items}
}

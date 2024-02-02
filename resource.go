package golibweb

type Resource struct {
	Meta MetaResource `json:"meta"`
	Data interface{}  `json:"data,omitempty"`
}

func NewResource(meta MetaResource, data interface{}) Resource {
	return Resource{
		Meta: meta,
		Data: data,
	}
}

func NewResourceWithoutData(meta MetaResource) Resource {
	return NewResource(meta, nil)
}

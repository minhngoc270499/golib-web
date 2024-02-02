package golibweb

type MetaResource struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func NewMetaResource(code int64, message string) MetaResource {
	return MetaResource{
		Code:    code,
		Message: message,
	}
}

func NewSuccessMetaResource() MetaResource {
	return MetaResource{
		Code:    200,
		Message: "Success",
	}
}

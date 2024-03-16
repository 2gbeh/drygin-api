package pipes

type TagCreatePipe struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type TagUpdatePipe struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

type TagResponsePipe struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

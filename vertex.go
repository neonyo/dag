package dag

type vertex struct {
	Id       string
	Value    Service
	Children []*vertex
}

func newVertex(id string, service Service) *vertex {
	return &vertex{
		Id:       id,
		Value:    service,
		Children: make([]*vertex, 0),
	}
}

package response

func NewResponse(p *Pagination, rows interface{}) *Pagination {
	p.Rows = rows
	return p
}

type Responsable interface {
	ToResponse() map[string]interface{}
}

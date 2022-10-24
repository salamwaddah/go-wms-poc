package response

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type Pagination struct {
	query      echo.Context
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"data"`
}

func NewPagination(c echo.Context) *Pagination {

	p := &Pagination{
		query: c,
	}
	p.Limit = p.GetLimit()
	p.Page = p.GetPage()

	return p
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	limit, err := strconv.Atoi(p.query.QueryParam("limit"))

	if err != nil {
		limit = 10
	}

	return limit
}

func (p *Pagination) GetPage() int {
	page, err := strconv.Atoi(p.query.QueryParam("page"))

	if err != nil {
		page = 1
	}

	return page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Debug().Model(value).Scopes(scopes...).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Debug().Scopes(scopes...).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

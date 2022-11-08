package handler

import (
	"balance"
	"github.com/gin-gonic/gin"
	"strconv"
)

//GeneratePaginationFromRequest ..
func generatePaginationFromRequest(c *gin.Context) balance.Pagination {
	// Initializing default
	//	var mode string
	limit := 2
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return balance.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}

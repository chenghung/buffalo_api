package actions

import (
	"strconv"

	"github.com/chenghung/buffalo_api/models"
	"github.com/gobuffalo/buffalo"
)

func UserIndex(c buffalo.Context) error {
	page, _ := strconv.Atoi(c.Param("page"))
	perPage, _ := strconv.Atoi(c.Param("per_page"))

	c.Logger().Debug("page: ", page)
	c.Logger().Debug("per page: ", perPage)

	users := models.Users{}
	models.DB.Paginate(page, perPage).All(&users)

	c.Logger().Debug("user size: ", len(users))

	return c.Render(200, r.JSON(users))
}

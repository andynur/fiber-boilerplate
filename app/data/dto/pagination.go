package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/log"
)

func GetPagination(c *fiber.Ctx) (pageNo, pageSize int) {
	ps := c.Query("page_size")
	pn := c.Query("page")
	pageSize, pageNo = 10, 1

	if len(ps) > 0 {
		psInt, err := strconv.Atoi(ps)
		if err != nil {
			log.Error().Err(err)
		} else {
			pageSize = psInt
		}
	}

	if len(pn) > 0 {
		pnInt, err := strconv.Atoi(pn)
		if err != nil {
			log.Error().Err(err)
		} else {
			pageNo = pnInt
		}
	}

	return pageNo, pageSize
}

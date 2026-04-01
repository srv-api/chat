package history

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) GetChatHistory(c echo.Context) error {

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	userID, _ := strconv.Atoi(c.QueryParam(userid))
	receiverID, _ := strconv.Atoi(c.QueryParam("receiver_id"))

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}

	data, err := h.serviceHistory.GetHistory(userID, receiverID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

package member

import (
	"bytedance/auth"
	"bytedance/db"
	"bytedance/redis_server"
	"bytedance/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(c *gin.Context) {
	var request types.DeleteMemberRequest
	var response types.DeleteMemberResponse

	if err := c.Bind(&request); err != nil {
		response.Code = types.UnknownError
		c.JSON(http.StatusOK, response)
		return
	}

	//删除cookie
	auth.Logout(c)

	// 删除成员
	if _, errNo := db.GetMemberByID(request.UserID); errNo != types.OK {
		response.Code = errNo
		c.JSON(http.StatusOK, response)
		return
	}
	redis_server.DeleteMemberByID(request.UserID)
	db.NewDB().Exec("update member set is_deleted=1 where member_id=?", request.UserID)
	response.Code = types.OK
	c.JSON(http.StatusOK, response)
	return
}

package ginuser

import (
	"GoEx8/common"
	"GoEx8/component"
	"GoEx8/modules/user/userbiz"
	"GoEx8/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewDeleteUserBiz(store)

		if err := biz.DeleteUser(c.Request.Context(), id); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}

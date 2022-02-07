package ginuser

import (
	"GoEx8/common"
	"GoEx8/component"
	"GoEx8/modules/user/userbiz"
	"GoEx8/modules/user/usermodel"
	"GoEx8/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewUpdateUserBiz(store)

		if err := biz.UpdateUser(c.Request.Context(), id, &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}

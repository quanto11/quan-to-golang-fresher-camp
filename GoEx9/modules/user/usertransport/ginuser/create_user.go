package ginuser

import (
	"GoEx8/common"
	"GoEx8/component"
	"GoEx8/modules/user/userbiz"
	"GoEx8/modules/user/usermodel"
	"GoEx8/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := userbiz.NewCreateUser(store)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

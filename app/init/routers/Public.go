package routers

import (
	"github.com/gin-gonic/gin"
	"service-all/app/controller"
)

func InitPublicRouter(r *gin.RouterGroup) {
	c := controller.NewController()
	public := r.Group("")
	{
		public.POST("/ping", c.Ping)
		public.POST("/key", c.Key)
		public.POST("/key2/:num", c.Key2)
		public.POST("/action", c.Action)
		public.POST("/ppt/:action", c.PptAction)
		public.GET("/siganl/:role", c.WebSocketSiganl)
		public.GET("/SmallAppsiganl/:openid", c.SmallAppWebSocketSiganl)
		public.GET("/pikSiganl", c.PikWebSocket)
	}
}

package route

import (
	v1 "github.com/Arapgp/Arapgp-Server-go/route/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter is to generate & initialize router
// router returned
func InitRouter() (r *gin.Engine) {
	r = gin.New()

	// corresponding to those /api/v1 in docs
	apiv1 := r.Group("/api/v1")
	if apiv1 != nil {

		// About pubKey(REST-ful)
		apiv1.GET("/pubKey", HeaderWrapper(v1.GetPubKey))
		apiv1.POST("/pubKey", HeaderWrapper(v1.PostPutPubKey))
		apiv1.PUT("/pubKey", HeaderWrapper(v1.PostPutPubKey))
		apiv1.DELETE("/pubKey", HeaderWrapper(v1.DeletePubKey))

		// About User Profile
		apiv1.POST("/signup", HeaderWrapper(v1.Register))
		apiv1.POST("/login", HeaderWrapper(v1.Login))
		apiv1.POST("/logout", HeaderWrapper(Auth(v1.Logout)))
		apiv1.GET("/user", HeaderWrapper(v1.GetUsersByName))

		// About File(s) (REST-ful)
		apiv1.POST("/user/:username/file", HeaderWrapper(v1.PostFileByUserName))
		apiv1.PUT("/user/:username/file", HeaderWrapper(v1.PutFileByUserName))
		apiv1.GET("/user/:username/file", HeaderWrapper(v1.GetFileByUserName))
		apiv1.DELETE("/user/:username/file", HeaderWrapper(v1.DeleteFileByUserName))
		apiv1.GET("/user/:username/filesinfo", HeaderWrapper(v1.GetFilesInfoByUserName))

		// About Test
		apiv1.GET("/ping", HeaderWrapper(v1.Ping))
	}

	return
}

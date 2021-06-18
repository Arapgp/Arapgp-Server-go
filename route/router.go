package route

import (
	v1 "github.com/Arapgp/Arapgp-Server-go/route/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter is to generate & initialize router
// router returned
func InitRouter() (r *gin.Engine) {
	r = gin.New()

	// use CORS wrapper first
	// then add api router in the following lines
	r.Use(CORS())

	// corresponding to those /api/v1 in docs
	apiv1 := r.Group("/api/v1")
	if apiv1 != nil {

		// About pubKey(REST-ful)
		apiv1.GET("/pubKey", v1.GetPubKey)
		apiv1.POST("/pubKey", v1.PostPutPubKey)
		apiv1.PUT("/pubKey", v1.PostPutPubKey)
		apiv1.DELETE("/pubKey", v1.DeletePubKey)

		// About User Profile
		apiv1.POST("/signup", v1.Register)
		apiv1.POST("/login", v1.Login)
		apiv1.POST("/logout", Auth(v1.Logout))
		apiv1.GET("/user", v1.GetUsersByName)

		// About File(s) (REST-ful)
		apiv1.POST("/user/:username/file", v1.PostFileByUserName)
		apiv1.PUT("/user/:username/file", v1.PutFileByUserName)
		apiv1.GET("/user/:username/file", v1.GetFileByUserName)
		apiv1.DELETE("/user/:username/file", v1.DeleteFileByUserName)
		apiv1.GET("/user/:username/filesinfo", v1.GetFilesInfoByUserName)

		// About Test
		apiv1.GET("/ping", v1.Ping)
	}

	return
}

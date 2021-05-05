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
		apiv1.GET("/pubKey", v1.GetPubKey)
		apiv1.POST("/pubKey", v1.PostPubKey)
		apiv1.PUT("/pubKey", v1.UpdatePubKey)
		apiv1.DELETE("/pubKey", v1.DeletePubKey)

		// About User Profile
		apiv1.POST("/signup", v1.Register)
		apiv1.POST("/login", v1.Login)
		apiv1.POST("/logout", v1.Logout)
		apiv1.GET("/user", v1.GetUsersByName)

	}

	return
}

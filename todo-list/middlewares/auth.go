package middlewares

import (
	firebaseContext "context"
	"net/http"
	"todo-app/config"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		client, err := config.FirebaseApp.Auth(firebaseContext.Background())
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		token, err := client.VerifyIDToken(firebaseContext.Background(), tokenString)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		email := token.Claims["email"]

		// if tokenString == "" {
		// 	context.JSON(401, gin.H{"error": "request does not contain an access token"})
		// 	context.Abort()
		// 	return
		// }
		// if err := auth.ValidateTokenAndAddUserInfoToRequest(tokenString, context); err != nil {
		// 	context.JSON(401, gin.H{"error": err.Error()})
		// 	context.Abort()
		// 	return
		// }
		context.Request.Header.Add("email", email.(string))
		context.Next()
	}
}

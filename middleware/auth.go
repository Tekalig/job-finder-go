// middleware/auth.go
package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func Auth(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(401, gin.H{"error": "No authorization header"})
            c.Abort()
            return
        }
        
        token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwtSecret), nil
        })
        
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
            c.Set("userID", claims.UserID)
            c.Set("role", claims.Role)
            c.Next()
        } else {
            c.JSON(401, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }
    }
}

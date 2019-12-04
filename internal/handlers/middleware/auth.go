package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"oe/conf"
	"oe/internal/models"
	ocontext "oe/pkg/context"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidAuthHeader = errors.New("auth header is invalid")
	// Err = errors.New("auth header is invalid")

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = errors.New("missing exp field")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")
)

func authError(c *gin.Context, err error) {
	errKey := "message"
	errMsgHeader := "[Auth] error: "
	e := gin.H{errKey: errMsgHeader + err.Error()}
	c.AbortWithStatusJSON(http.StatusUnauthorized, e)
}

var (
	bearerRegexp = regexp.MustCompile(`^(?:B|b)earer (\S+$)`)
)

func extractBearerToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	fmt.Println(authHeader)
	if authHeader == "" {
		return "", nil
	}

	matches := bearerRegexp.FindStringSubmatch(authHeader)
	if len(matches) != 2 {

		return "", ErrInvalidAuthHeader
	}

	return matches[1], nil
}

func Middleware(config *conf.GlobalConfiguration, orm *models.ORM) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		bearerToken, err := extractBearerToken(c)

		if err != nil {
			authError(c, ErrInvalidAuthHeader)
			return
		}

		if bearerToken == "" {
			// log.Info("Making unauthenticated request")
			fmt.Println("Making unauthenticated request")
			// c.Request = addToContext(c, &ocontext.ProjectContextKeys.UserCtxKey, nil)
			c.Next()
			return
		}

		t, err := ParseToken(c, config, bearerToken)

		if claims, ok := t.Claims.(jwt.MapClaims); ok {
			if claims["exp"] != nil {
				// issuer := claims["iss"].(string)
				userid := claims["user_id"].(string)

				if user, err := orm.FindUserByJWT(userid); err != nil {

					c.Next()

				} else {
					c.Request = addToContext(c, ocontext.ProjectContextKeys.UserCtxKey, user)

					c.Next()

					return
				}
			} else {
				authError(c, ErrMissingExpField)
				return
			}
		} else {
			authError(c, err)
			return
		}

	})
}

// ParseToken parse jwt token from gin context
func ParseToken(c *gin.Context, config *conf.GlobalConfiguration, token string) (t *jwt.Token, err error) {

	if err != nil {
		return nil, err
	}
	SigningAlgorithm := "H256"
	Key := []byte("dsdsdsds")
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		// save token string if vaild
		// c.Set("JWT_TOKEN", token)
		return Key, nil
	})
}

func addToContext(c *gin.Context, key ocontext.ContextKey, value interface{}) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
}

/*
JWT claims struct
*/
type Token struct {
	UserId string
	jwt.StandardClaims
}

// EpochNow is a helper function that returns the NumericDate time value used by the spec
func EpochNow() int64 {
	return time.Now().UTC().Unix()
}

// ExpireIn is a helper function to return calculated time in the future for "exp" claim
func ExpireIn(tm time.Duration) int64 {
	return EpochNow() + int64(tm.Seconds())
}

// JWTClaims represents the JWT claims information.
type JWTClaims struct {
	Email        string                 `json:"email"`
	UserID       string                 `json:"user_id"`
	AppMetaData  map[string]interface{} `json:"app_metadata"`
	UserMetaData map[string]interface{} `json:"user_metadata"`
	jwt.StandardClaims
}

func CreateJWTToken(user models.User, seller bool) string {

	m := make(map[string]interface{})

	m["roles"] = []string{"oe.user"}

	m["scopes"] = user.GetPermissions()
	if seller {
		m["roles"] = []string{"oe.seller"}
	}

	tk := JWTClaims{
		UserID:      user.ID,
		Email:       user.Email,
		AppMetaData: m,
	}

	tk.ExpiresAt = ExpireIn(3 * time.Hour)
	tk.Issuer = "auth.oe.co.ke"
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// TODO read key from  env
	tokenString, _ := token.SignedString([]byte("dsdsdsds"))
	return tokenString

}

func GeneratePassword(password string) string {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), 11)

	if err != nil {
		return ""
	}

	return string(pw)
}

func getCurrentUser(ctx context.Context) *models.User {
	// return ctx.Value(ocontext.ProjectContextKeys.UserCtxKey).(*models.User)
	return ForContext(ctx)
}

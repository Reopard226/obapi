package jwt

import (
	"context"
	"log"
	"net/http"
	"oceanbolt.com/obapi/internal/iam/iamclient"
	"oceanbolt.com/obapi/rpc/iam"
	"strings"
	"time"

	model "oceanbolt.com/obapi/internal/echoapi/utl/model"

	"github.com/labstack/echo"

	"github.com/dgrijalva/jwt-go"
)

// New generates new JWT service necessery for auth middleware
func New(secret, algo string, d int, iam iamclient.OceanboltIAMClient) *Service {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		key:       secret,
		algo:      signingMethod,
		duration:  time.Duration(d) * time.Minute,
		iamClient: iam,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key string

	// Duration for which the jwt token is valid.
	duration time.Duration

	// JWT signing algorithm
	algo jwt.SigningMethod

	iamClient iamclient.OceanboltIAMClient
}

// MWFunc makes JWT implement the Middleware interface.
func (j *Service) MWFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := j.ParseToken(c)
			if err != nil || !token.Valid {
				return c.NoContent(http.StatusUnauthorized)
			}

			claims := token.Claims.(jwt.MapClaims)

			obkid := claims["obkid"].(string)
			userId := claims["sub"].(string)
			apikeyId := claims["obkid"].(string)

			accessType := "browser"
			if claims["ktype"] == "apikey" {
				accessType = "apikey"
				permissions, err := j.iamClient.ValidateKey(context.Background(), &iam.UserKey{
					ApikeyId: apikeyId,
					UserId:   userId,
				})
				if err != nil {
					log.Fatal(err)
				}

				if !permissions.Valid {
					return c.NoContent(http.StatusUnauthorized)
				}

				interfaces := make([]interface{}, len(permissions.Permissions))

				for k, v := range permissions.Permissions {
					interfaces[k] = v
				}

				claims["permissions"] = interfaces
			}
			//fmt.Printf("{\"time\":\"%v\",\"userId\":\"%s\",\"access_type\":\"%s\",\"url\":\"%s\"}\n", time.Now().Format(time.RFC3339), claims["sub"], accessType, c.Path())

			c.Set("user", token)
			//c.Set("permissions",claims["permissions"])
			c.Set("apikeyId", apikeyId)
			c.Set("user_id", userId)
			c.Set("obkid", obkid)
			c.Set("access_type", accessType)

			//log.Printf("%v",c.Get("permissions").([]string))
			return next(c)
		}
	}
}

// ParseToken parses token from Authorization header
func (j *Service) ParseToken(c echo.Context) (*jwt.Token, error) {

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return nil, model.ErrGeneric
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, model.ErrGeneric
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, model.ErrGeneric
		}
		pub, err := j.iamClient.GetPublicKey(context.Background(), &iam.PrivateKey{Kid: token.Header["kid"].(string)})
		if err != nil {
			log.Println(err)
		}
		parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(pub.PublicKey)
		if err != nil {
			log.Println(err)
		}
		return parsedKey, nil
	}

	return jwt.Parse(parts[1], keyFunc)

}

func (j *Service) VerifyStandardClaims() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if !strings.HasPrefix(c.Request().RequestURI, "/v1") {
				return next(c)
			}

			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			audience := claims.VerifyAudience("https://api.oceanbolt.com", false)
			if !audience {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid audience")
			}
			issuer := claims.VerifyIssuer("https://oceanbolt.eu.auth0.com/", false)
			if !issuer {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid issuer")
			}

			return next(c)
		}
	}
}

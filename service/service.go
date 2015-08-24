package service

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/vapost/mservice-dload/util"
	"github.com/vapost/mservice-dload/retriever"
	"github.com/vapost/mservice-dload/validator"
)

type Config struct {
	ServiceHost string
	Jwt map[string]string
	Endpoints map[string] string
	HotelParams map[string]string

}

type Service struct {
}

func (s *Service) Run(config Config) error {

	// jwt token generation
	jwt := jwtutil.JwtGenerator{
		Key: config.Jwt["key"],
		Consumer: config.Jwt["consumer"],
		TTL: config.Jwt["TTL"],
	}


	r := gin.Default()


	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/token", func(c *gin.Context) {
		c.String(http.StatusOK, jwt.GetToken())
	})

	r.Run(config.ServiceHost)

	return nil
}


func (s *Service) GetHotelObjects(config Config) error {
	
	// jwt token generation
	jwt := jwtutil.JwtGenerator{
		Key: config.Jwt["key"],
		Consumer: config.Jwt["consumer"],
		TTL: config.Jwt["TTL"],
	}


	hotelObjectRetriever := retriever.HotelObject{ Endpoints: config.Endpoints, Params: config.HotelParams, Token: jwt.GetToken() }


	hotels := hotelObjectRetriever.GetHotelObjects()


	offerValidator := validator.OfferValidator{}

	offerValidator.ValidateOffers(hotels)


	fmt.Println(hotels)
	return nil
}
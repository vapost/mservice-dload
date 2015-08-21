package service

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type Config struct {
	ServiceHost string
}

type StatService struct {
}

func (s *StatService) Run(config Config) error {

	r := gin.Default()


	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(config.ServiceHost)

	return nil
}


type Bla struct {
	yo string
}

func (b *Bla)goNuts() string {
	return "Going crazey"
}
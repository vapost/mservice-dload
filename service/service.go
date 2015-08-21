package service


type Config struct {
	ServiceHost string
}

type StatService struct {
}

func (s *StatService) Run(config Config) error {

	/*db, err := s.getDb(config)

	if err != nil {
		return err
	}*/

	statResource := &resource.StatResource{}

	//statResource.Db = db

	r := gin.Default()

	/*r.POST("/stats", statResource.Store)

	r.GET("/stats", statResource.Index)

	r.GET("/stats/:id", statResource.Show)
*/
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(config.ServiceHost)

	return nil
}
package retriever

import (
	"github.com/vapost/mservice-dload/resource"
	"github.com/vapost/mservice-dload/util/query-string-helper"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HotelObject struct {
	Params map[string] string
	Endpoints map[string] string
	Token string
}

func (hr *HotelObject) GetHotelObjects() []resource.Offer {

	
	endpoint := hr.GetEndpoint()

	resp, err := http.Get(endpoint)
	checkErr(err, "Remote call failed")

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "Reading body failed")

	

	var r resource.OfferResponse

	json.Unmarshal(body, &r)

	fmt.Println(r.Data.Offers[0])

	return r.Data.Offers
}

func (hr *HotelObject) GetEndpoint() string{

	querystringutil.Something.GoCrazy()

	// get the base
	base := hr.Endpoints["base"]

	// bind the base with the endpoint and corresponding booking type
	endpoint := base + fmt.Sprintf(hr.Endpoints["get_hotels"], hr.Params["type"])

	// append all params as query string
	endpoint += "?"
	for index,element := range hr.Params {
  		endpoint += fmt.Sprintf("%s=%s&", index, element)
	}
	endpoint += fmt.Sprintf("jwt_token=%s", hr.Token)

	return endpoint
}


func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

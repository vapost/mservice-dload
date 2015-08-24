package validator


import(
	"github.com/vapost/mservice-dload/resource"
	//"fmt"
)


type OfferValidator struct {

}

func (validator *OfferValidator) ValidateOffers(offers []resource.Offer) []resource.Offer {
	// fmt.Println(offers)

	for _, offer := range offers {
		
	}

	return offers
}
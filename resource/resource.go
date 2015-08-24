package resource

type OfferResponse struct {
		Data struct {
			Offers []Offer `json:"offers"`
		} `json:"data"`
	}

type Offer struct {
	Id string `json:"id"`
	DateStart string `json:"date_start"`
					DateEnd string `json:"date_end"`
					Duration int `json:"duration"`
					Price struct {
						Amount float32 `json:"amount"`
						Currency string `json:"currency"`
						UnitType string `json:"unit_type"`
					}
					TourOperator struct {
						Name string `json:"name"`
						Code string `json:"code"`
						Image struct {
							Url string `json:"url"`
						} `json:"image"`
					} `json:"tour_operator"`
					Accommodation struct {
						Type string `json:"type"`
						Board string `json:"board"`
						BookCode string `json:"book_code"`
					} `json:"accommodation"`
					Object struct {
						ID int `json:"id"`
						GiataId int `json:"giata_id"`
						BookCode string `json:"book_code"`
						Name string `json:"name"`
						Category float32 `json:"category"`
						PosterImage struct {
							Url string `json:"url"`
						} `json:"poster_image"`
						Location struct {
							Country string `json:"country"`
							City string `json:"city"`
							Region string `json:"region"`
							ProviderRegionId int `json:"provider_region_id"`
						} `json:"location"`
					} `json:"object"`
}	
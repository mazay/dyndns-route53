package ipapi

import "github.com/go-resty/resty/v2"

type IpResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

// GetIp retrieves the public IP address of the current machine by making
// a request to the ip-api.com service. It returns the IP address as a
// string if successful, or an error if the request fails.
func GetIp() (string, error) {
	result := IpResponse{}
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetResult(&result).
		Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}

	return result.Query, nil
}

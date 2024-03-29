package iplocator

import (
	"ca-amartha/businesses/iplocator"
	"context"
	"encoding/json"
	"net/http"
)

type IpAPI struct {
	httpClient http.Client
}

func NewIpAPI() iplocator.Repository {
	return &IpAPI{
		httpClient: http.Client{},
	}
}

func (ipl *IpAPI) GetLocationByIP(ctx context.Context, ip string) (iplocator.Domain, error) {
	req, _ := http.NewRequest("GET", "https://ipapi.co/"+ip+"/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := ipl.httpClient.Do(req)
	if err != nil {
		return iplocator.Domain{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return iplocator.Domain{}, err
	}

	return data.toDomain(), nil
}

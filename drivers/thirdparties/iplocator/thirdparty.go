package iplocator

import (
	"context"
	"encoding/json"
	"net/http"
)

type IPLocator struct {
	httpClient http.Client
}

func NewIPLocator() *IPLocator {
	return &IPLocator{
		httpClient: http.Client{},
	}
}

func (iplocator *IPLocator) NewsGetLocationByIP(ctx context.Context, ip string) (Response, error) {
	req, _ := http.NewRequest("GET", "https://ipapi.co/"+ip+"/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := iplocator.httpClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Response{}, err
	}

	return data, nil
}

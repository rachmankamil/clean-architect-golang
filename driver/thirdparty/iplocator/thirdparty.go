package iplocator

import (
	"ca-amartha/bussiness/news"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type IPLocator struct {
	httpClient http.Client
}

func NewIPLocator() news.IPLocatorRepository {
	return &IPLocator{
		httpClient: http.Client{},
	}
}

func (iplocator *IPLocator) getLocationByIP(ctx context.Context, ip string) Response {
	req, _ := http.NewRequest("GET", "https://ipapi.co/"+ip+"/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := iplocator.httpClient.Do(req)
	if err != nil {
		return Response{}
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Response{}
	}

	return data
}

func (iplocator *IPLocator) NewsGetLocationByIP(ctx context.Context, ip string) (news.IPStatDomain, error) {
	resp := iplocator.getLocationByIP(ctx, ip)

	if resp != (Response{}) {
		return news.IPStatDomain{}, errors.New("data not found")
	}

	return resp.ToNewsDomain(), nil
}

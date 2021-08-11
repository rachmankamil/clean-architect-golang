package iplocator

import "context"

type Domain struct {
	IP          string
	Version     string
	City        string
	Region      string
	CountryName string
}

type Repository interface {
	GetLocationByIP(ctx context.Context, ip string) (Domain, error)
}

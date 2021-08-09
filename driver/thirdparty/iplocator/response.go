package iplocator

import "ca-amartha/bussiness/news"

type Response struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryName        string  `json:"country_name"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  float64 `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

func (resp *Response) ToNewsDomain() news.IPStatDomain {
	return news.IPStatDomain{
		IP:          resp.IP,
		Version:     resp.Version,
		City:        resp.City,
		Region:      resp.Region,
		CountryName: resp.CountryName,
	}
}

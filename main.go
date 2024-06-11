package main

import (
	"fmt"
	"log"

	"github.com/hooklift/gowsdl/soap"
)

func main() {
	countryServiceClient :=  soap.NewClient("http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso")
	countryService := country.NewCountryInfoServiceSoapType(countryServiceClient)

	fullCountryInfoRequest := &country.CountriesUsingCurrency{SISOCurrencyCode:" USD"}
	fullCountryInfoResponse, err := countryService.ContriesUsingCurrency(fullCountryInfoRequest)
	if err != nil {
		log.Fatal("Error calling FullContryInfo: %v", err)

	} 

	for i := range fullCountryInfoResponse.ContriesUsingCurrencyResult.TCountryCodeAndName {
		fmt.Printf("Country Code: %s, Country Name: %s/n",
		&fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SISOCode,
		&fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SName),
	}
}
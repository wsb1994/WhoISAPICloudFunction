package info

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"example.com/module/config"
)

// Autogenerated from JSON -> Go Struct
type Info_AutoGenerated struct {
	WhoisRecord struct {
		CreatedDate string `json:"createdDate"`
		UpdatedDate string `json:"updatedDate"`
		ExpiresDate string `json:"expiresDate"`
		Registrant  struct {
			Organization string `json:"organization"`
			State        string `json:"state"`
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			RawText      string `json:"rawText"`
		} `json:"registrant"`
		AdministrativeContact struct {
			Organization string `json:"organization"`
			State        string `json:"state"`
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			RawText      string `json:"rawText"`
		} `json:"administrativeContact"`
		TechnicalContact struct {
			Organization string `json:"organization"`
			State        string `json:"state"`
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			RawText      string `json:"rawText"`
		} `json:"technicalContact"`
		DomainName  string `json:"domainName"`
		NameServers struct {
			RawText   string        `json:"rawText"`
			HostNames []string      `json:"hostNames"`
			Ips       []interface{} `json:"ips"`
		} `json:"nameServers"`
		Status       string `json:"status"`
		RawText      string `json:"rawText"`
		ParseCode    int    `json:"parseCode"`
		Header       string `json:"header"`
		StrippedText string `json:"strippedText"`
		Footer       string `json:"footer"`
		Audit        struct {
			CreatedDate string `json:"createdDate"`
			UpdatedDate string `json:"updatedDate"`
		} `json:"audit"`
		RegistrarName         string `json:"registrarName"`
		RegistrarIANAID       string `json:"registrarIANAID"`
		CreatedDateNormalized string `json:"createdDateNormalized"`
		UpdatedDateNormalized string `json:"updatedDateNormalized"`
		ExpiresDateNormalized string `json:"expiresDateNormalized"`
		RegistryData          struct {
			CreatedDate time.Time `json:"createdDate"`
			UpdatedDate time.Time `json:"updatedDate"`
			ExpiresDate time.Time `json:"expiresDate"`
			DomainName  string    `json:"domainName"`
			NameServers struct {
				RawText   string        `json:"rawText"`
				HostNames []string      `json:"hostNames"`
				Ips       []interface{} `json:"ips"`
			} `json:"nameServers"`
			Status       string    `json:"status"`
			RawText      time.Time `json:"rawText"`
			ParseCode    int       `json:"parseCode"`
			Header       string    `json:"header"`
			StrippedText time.Time `json:"strippedText"`
			Footer       string    `json:"footer"`
			Audit        struct {
				CreatedDate string `json:"createdDate"`
				UpdatedDate string `json:"updatedDate"`
			} `json:"audit"`
			RegistrarName         string `json:"registrarName"`
			RegistrarIANAID       string `json:"registrarIANAID"`
			CreatedDateNormalized string `json:"createdDateNormalized"`
			UpdatedDateNormalized string `json:"updatedDateNormalized"`
			ExpiresDateNormalized string `json:"expiresDateNormalized"`
			WhoisServer           string `json:"whoisServer"`
		} `json:"registryData"`
		ContactEmail       string `json:"contactEmail"`
		DomainNameExt      string `json:"domainNameExt"`
		EstimatedDomainAge int    `json:"estimatedDomainAge"`
	} `json:"WhoisRecord"`
}

// Generate a valid json string to return to the api caller from g after obtaining valid results.
func (g *Info_AutoGenerated) GenerateResults() (string, error) {
	bytes, err := json.Marshal(g)
	results := string(bytes)
	return results, err

}

// Generates a URI in the singular configuration I elected to provide. A better example might be to take in a list of command line
// arguments, build a vector of them as part of some native struct with them, and iterate and sprintf to the end each of the configuration steps.
// This is not a huge deal to not include, as I did not find any fun or fancy options requiring such a feature to incorporate.
func generateURI(conf config.Config, domain string) string {
	value := fmt.Sprintf("%sapiKey=%s&domainName=%s&outputFormat=%s", conf.ApiUrl, conf.ApiKey, domain, conf.Format)
	return value
}

// Alter g to represent the information associated with a domain.
func (g *Info_AutoGenerated) ObtainResults(domain string) error {
	/// generate string using sprintf to request the correct information
	conf := config.GenerateConfig()

	URI := generateURI(*conf, domain)
	resp, err := http.Get(URI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &g); err != nil {
		return err
	}
	return nil
}

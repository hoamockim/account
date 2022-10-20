package configs

import "time"

type Cert struct {
	CommonName         string        `json:"commonName" env:"cert_commonName"`
	Country            string        `json:"country" env:"cert_country"`
	Province           string        `json:"province" env:"cert_province"`
	Locality           string        `json:"locality" env:"cert_locality"`
	Organization       string        `json:"organization" env:"cert_organization"`
	OrganizationalUnit string        `json:"organizationalUnit" env:"cert_organizationalUnit"`
	Validity           time.Duration `json:"validity" env:"cert_validity"`
}

package cdnfrontdoorsecret


type CdnFrontdoorSecretSecret struct {
	// customer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_secret#customer_certificate CdnFrontdoorSecret#customer_certificate}
	CustomerCertificate interface{} `field:"required" json:"customerCertificate" yaml:"customerCertificate"`
}


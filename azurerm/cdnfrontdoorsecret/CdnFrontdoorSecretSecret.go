package cdnfrontdoorsecret


type CdnFrontdoorSecretSecret struct {
	// customer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_secret#customer_certificate CdnFrontdoorSecret#customer_certificate}
	CustomerCertificate interface{} `field:"required" json:"customerCertificate" yaml:"customerCertificate"`
}


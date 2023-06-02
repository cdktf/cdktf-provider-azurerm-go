package applicationgateway


type ApplicationGatewayBackendHttpSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#cookie_based_affinity ApplicationGateway#cookie_based_affinity}.
	CookieBasedAffinity *string `field:"required" json:"cookieBasedAffinity" yaml:"cookieBasedAffinity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#port ApplicationGateway#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#protocol ApplicationGateway#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#affinity_cookie_name ApplicationGateway#affinity_cookie_name}.
	AffinityCookieName *string `field:"optional" json:"affinityCookieName" yaml:"affinityCookieName"`
	// authentication_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#authentication_certificate ApplicationGateway#authentication_certificate}
	AuthenticationCertificate interface{} `field:"optional" json:"authenticationCertificate" yaml:"authenticationCertificate"`
	// connection_draining block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#connection_draining ApplicationGateway#connection_draining}
	ConnectionDraining *ApplicationGatewayBackendHttpSettingsConnectionDraining `field:"optional" json:"connectionDraining" yaml:"connectionDraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#host_name ApplicationGateway#host_name}.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#path ApplicationGateway#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#pick_host_name_from_backend_address ApplicationGateway#pick_host_name_from_backend_address}.
	PickHostNameFromBackendAddress interface{} `field:"optional" json:"pickHostNameFromBackendAddress" yaml:"pickHostNameFromBackendAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#probe_name ApplicationGateway#probe_name}.
	ProbeName *string `field:"optional" json:"probeName" yaml:"probeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#request_timeout ApplicationGateway#request_timeout}.
	RequestTimeout *float64 `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/application_gateway#trusted_root_certificate_names ApplicationGateway#trusted_root_certificate_names}.
	TrustedRootCertificateNames *[]*string `field:"optional" json:"trustedRootCertificateNames" yaml:"trustedRootCertificateNames"`
}


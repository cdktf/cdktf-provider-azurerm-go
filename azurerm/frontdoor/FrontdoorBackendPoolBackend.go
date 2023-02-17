package frontdoor


type FrontdoorBackendPoolBackend struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#address Frontdoor#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#host_header Frontdoor#host_header}.
	HostHeader *string `field:"required" json:"hostHeader" yaml:"hostHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#http_port Frontdoor#http_port}.
	HttpPort *float64 `field:"required" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#https_port Frontdoor#https_port}.
	HttpsPort *float64 `field:"required" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enabled Frontdoor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#priority Frontdoor#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#weight Frontdoor#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


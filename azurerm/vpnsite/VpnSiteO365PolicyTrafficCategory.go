package vpnsite


type VpnSiteO365PolicyTrafficCategory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_site#allow_endpoint_enabled VpnSite#allow_endpoint_enabled}.
	AllowEndpointEnabled interface{} `field:"optional" json:"allowEndpointEnabled" yaml:"allowEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_site#default_endpoint_enabled VpnSite#default_endpoint_enabled}.
	DefaultEndpointEnabled interface{} `field:"optional" json:"defaultEndpointEnabled" yaml:"defaultEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_site#optimize_endpoint_enabled VpnSite#optimize_endpoint_enabled}.
	OptimizeEndpointEnabled interface{} `field:"optional" json:"optimizeEndpointEnabled" yaml:"optimizeEndpointEnabled"`
}

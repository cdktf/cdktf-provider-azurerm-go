package frontdoor


type FrontdoorBackendPoolHealthProbe struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enabled Frontdoor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#interval_in_seconds Frontdoor#interval_in_seconds}.
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#path Frontdoor#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#probe_method Frontdoor#probe_method}.
	ProbeMethod *string `field:"optional" json:"probeMethod" yaml:"probeMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#protocol Frontdoor#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


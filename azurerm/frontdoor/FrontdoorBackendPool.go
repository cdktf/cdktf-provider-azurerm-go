package frontdoor


type FrontdoorBackendPool struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend Frontdoor#backend}
	Backend interface{} `field:"required" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#health_probe_name Frontdoor#health_probe_name}.
	HealthProbeName *string `field:"required" json:"healthProbeName" yaml:"healthProbeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#load_balancing_name Frontdoor#load_balancing_name}.
	LoadBalancingName *string `field:"required" json:"loadBalancingName" yaml:"loadBalancingName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

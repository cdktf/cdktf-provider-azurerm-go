package applicationgateway


type ApplicationGatewayBackendHttpSettingsConnectionDraining struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#drain_timeout_sec ApplicationGateway#drain_timeout_sec}.
	DrainTimeoutSec *float64 `field:"required" json:"drainTimeoutSec" yaml:"drainTimeoutSec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#enabled ApplicationGateway#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}


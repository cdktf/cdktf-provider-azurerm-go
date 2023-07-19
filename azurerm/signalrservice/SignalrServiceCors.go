package signalrservice


type SignalrServiceCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/signalr_service#allowed_origins SignalrService#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
}


package signalrservice


type SignalrServiceCors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/signalr_service#allowed_origins SignalrService#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
}


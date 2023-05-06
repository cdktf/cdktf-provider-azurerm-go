package appservice


type AppServiceLogsHttpLogsFileSystem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/app_service#retention_in_days AppService#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/app_service#retention_in_mb AppService#retention_in_mb}.
	RetentionInMb *float64 `field:"required" json:"retentionInMb" yaml:"retentionInMb"`
}


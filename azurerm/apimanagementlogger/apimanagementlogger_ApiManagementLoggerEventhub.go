package apimanagementlogger


type ApiManagementLoggerEventhub struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_logger#connection_string ApiManagementLogger#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management_logger#name ApiManagementLogger#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}


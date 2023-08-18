package apimanagement


type ApiManagementVirtualNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/api_management#subnet_id ApiManagement#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


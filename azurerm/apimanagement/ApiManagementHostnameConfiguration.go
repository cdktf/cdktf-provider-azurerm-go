package apimanagement


type ApiManagementHostnameConfiguration struct {
	// developer_portal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#developer_portal ApiManagement#developer_portal}
	DeveloperPortal interface{} `field:"optional" json:"developerPortal" yaml:"developerPortal"`
	// management block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#management ApiManagement#management}
	Management interface{} `field:"optional" json:"management" yaml:"management"`
	// portal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#portal ApiManagement#portal}
	Portal interface{} `field:"optional" json:"portal" yaml:"portal"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#proxy ApiManagement#proxy}
	Proxy interface{} `field:"optional" json:"proxy" yaml:"proxy"`
	// scm block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#scm ApiManagement#scm}
	Scm interface{} `field:"optional" json:"scm" yaml:"scm"`
}

package labservicelab


type LabServiceLabVirtualMachineAdminUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_lab#password LabServiceLab#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_lab#username LabServiceLab#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}


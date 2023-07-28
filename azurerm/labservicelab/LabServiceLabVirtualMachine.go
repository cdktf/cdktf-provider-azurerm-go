package labservicelab


type LabServiceLabVirtualMachine struct {
	// admin_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#admin_user LabServiceLab#admin_user}
	AdminUser *LabServiceLabVirtualMachineAdminUser `field:"required" json:"adminUser" yaml:"adminUser"`
	// image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#image_reference LabServiceLab#image_reference}
	ImageReference *LabServiceLabVirtualMachineImageReference `field:"required" json:"imageReference" yaml:"imageReference"`
	// sku block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#sku LabServiceLab#sku}
	Sku *LabServiceLabVirtualMachineSku `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#additional_capability_gpu_drivers_installed LabServiceLab#additional_capability_gpu_drivers_installed}.
	AdditionalCapabilityGpuDriversInstalled interface{} `field:"optional" json:"additionalCapabilityGpuDriversInstalled" yaml:"additionalCapabilityGpuDriversInstalled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#create_option LabServiceLab#create_option}.
	CreateOption *string `field:"optional" json:"createOption" yaml:"createOption"`
	// non_admin_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#non_admin_user LabServiceLab#non_admin_user}
	NonAdminUser *LabServiceLabVirtualMachineNonAdminUser `field:"optional" json:"nonAdminUser" yaml:"nonAdminUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#shared_password_enabled LabServiceLab#shared_password_enabled}.
	SharedPasswordEnabled interface{} `field:"optional" json:"sharedPasswordEnabled" yaml:"sharedPasswordEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/lab_service_lab#usage_quota LabServiceLab#usage_quota}.
	UsageQuota *string `field:"optional" json:"usageQuota" yaml:"usageQuota"`
}


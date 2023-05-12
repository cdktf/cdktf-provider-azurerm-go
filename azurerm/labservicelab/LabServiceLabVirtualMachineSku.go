package labservicelab


type LabServiceLabVirtualMachineSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/lab_service_lab#capacity LabServiceLab#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/lab_service_lab#name LabServiceLab#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}


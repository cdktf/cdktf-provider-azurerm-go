package springcloudaccelerator


type SpringCloudAcceleratorTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_accelerator#create SpringCloudAccelerator#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_accelerator#delete SpringCloudAccelerator#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_accelerator#read SpringCloudAccelerator#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


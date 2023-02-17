package batchpool


type BatchPoolStartTaskContainer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#image_name BatchPool#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// registry block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#registry BatchPool#registry}
	Registry interface{} `field:"optional" json:"registry" yaml:"registry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#run_options BatchPool#run_options}.
	RunOptions *string `field:"optional" json:"runOptions" yaml:"runOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#working_directory BatchPool#working_directory}.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}


package springcloudcustomizedaccelerator


type SpringCloudCustomizedAcceleratorGitRepository struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_customized_accelerator#url SpringCloudCustomizedAccelerator#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_customized_accelerator#basic_auth SpringCloudCustomizedAccelerator#basic_auth}
	BasicAuth *SpringCloudCustomizedAcceleratorGitRepositoryBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_customized_accelerator#branch SpringCloudCustomizedAccelerator#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_customized_accelerator#commit SpringCloudCustomizedAccelerator#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_customized_accelerator#git_tag SpringCloudCustomizedAccelerator#git_tag}.
	GitTag *string `field:"optional" json:"gitTag" yaml:"gitTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_customized_accelerator#interval_in_seconds SpringCloudCustomizedAccelerator#interval_in_seconds}.
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// ssh_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_customized_accelerator#ssh_auth SpringCloudCustomizedAccelerator#ssh_auth}
	SshAuth *SpringCloudCustomizedAcceleratorGitRepositorySshAuth `field:"optional" json:"sshAuth" yaml:"sshAuth"`
}

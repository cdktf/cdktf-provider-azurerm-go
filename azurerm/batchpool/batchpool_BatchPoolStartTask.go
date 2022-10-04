package batchpool


type BatchPoolStartTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#command_line BatchPool#command_line}.
	CommandLine *string `field:"required" json:"commandLine" yaml:"commandLine"`
	// user_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_identity BatchPool#user_identity}
	UserIdentity *BatchPoolStartTaskUserIdentity `field:"required" json:"userIdentity" yaml:"userIdentity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#common_environment_properties BatchPool#common_environment_properties}.
	CommonEnvironmentProperties *map[string]*string `field:"optional" json:"commonEnvironmentProperties" yaml:"commonEnvironmentProperties"`
	// resource_file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#resource_file BatchPool#resource_file}
	ResourceFile interface{} `field:"optional" json:"resourceFile" yaml:"resourceFile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#task_retry_maximum BatchPool#task_retry_maximum}.
	TaskRetryMaximum *float64 `field:"optional" json:"taskRetryMaximum" yaml:"taskRetryMaximum"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#wait_for_success BatchPool#wait_for_success}.
	WaitForSuccess interface{} `field:"optional" json:"waitForSuccess" yaml:"waitForSuccess"`
}


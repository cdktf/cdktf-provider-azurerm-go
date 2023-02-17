package containerregistrytaskschedulerunnow


type ContainerRegistryTaskScheduleRunNowTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task_schedule_run_now#create ContainerRegistryTaskScheduleRunNow#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task_schedule_run_now#delete ContainerRegistryTaskScheduleRunNow#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task_schedule_run_now#read ContainerRegistryTaskScheduleRunNow#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


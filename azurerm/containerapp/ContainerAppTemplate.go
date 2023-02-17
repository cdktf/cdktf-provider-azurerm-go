package containerapp


type ContainerAppTemplate struct {
	// container block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#container ContainerApp#container}
	Container interface{} `field:"required" json:"container" yaml:"container"`
	// The maximum number of replicas for this container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#max_replicas ContainerApp#max_replicas}
	MaxReplicas *float64 `field:"optional" json:"maxReplicas" yaml:"maxReplicas"`
	// The minimum number of replicas for this container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#min_replicas ContainerApp#min_replicas}
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// The suffix for the revision.
	//
	// This value must be unique for the lifetime of the Resource. If omitted the service will use a hash function to create one.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#revision_suffix ContainerApp#revision_suffix}
	RevisionSuffix *string `field:"optional" json:"revisionSuffix" yaml:"revisionSuffix"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#volume ContainerApp#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}


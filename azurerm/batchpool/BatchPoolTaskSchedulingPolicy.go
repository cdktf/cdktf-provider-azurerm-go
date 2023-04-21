package batchpool


type BatchPoolTaskSchedulingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/batch_pool#node_fill_type BatchPool#node_fill_type}.
	NodeFillType *string `field:"optional" json:"nodeFillType" yaml:"nodeFillType"`
}


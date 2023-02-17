package batchpool


type BatchPoolTaskSchedulingPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#node_fill_type BatchPool#node_fill_type}.
	NodeFillType *string `field:"optional" json:"nodeFillType" yaml:"nodeFillType"`
}


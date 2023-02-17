package provider


type AzurermProviderFeaturesVirtualMachineScaleSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#roll_instances_when_required AzurermProvider#roll_instances_when_required}.
	RollInstancesWhenRequired interface{} `field:"required" json:"rollInstancesWhenRequired" yaml:"rollInstancesWhenRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#force_delete AzurermProvider#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#scale_to_zero_before_deletion AzurermProvider#scale_to_zero_before_deletion}.
	ScaleToZeroBeforeDeletion interface{} `field:"optional" json:"scaleToZeroBeforeDeletion" yaml:"scaleToZeroBeforeDeletion"`
}


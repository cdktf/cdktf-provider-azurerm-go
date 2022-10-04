package databoxedgedevice


type DataboxEdgeDeviceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databox_edge_device#create DataboxEdgeDevice#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databox_edge_device#delete DataboxEdgeDevice#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databox_edge_device#read DataboxEdgeDevice#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databox_edge_device#update DataboxEdgeDevice#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


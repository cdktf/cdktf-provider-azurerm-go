package netappsnapshotpolicy


type NetappSnapshotPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_snapshot_policy#create NetappSnapshotPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_snapshot_policy#delete NetappSnapshotPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_snapshot_policy#read NetappSnapshotPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_snapshot_policy#update NetappSnapshotPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

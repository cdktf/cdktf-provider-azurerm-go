package netappvolume


type NetappVolumeDataProtectionSnapshotPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#snapshot_policy_id NetappVolume#snapshot_policy_id}.
	SnapshotPolicyId *string `field:"required" json:"snapshotPolicyId" yaml:"snapshotPolicyId"`
}


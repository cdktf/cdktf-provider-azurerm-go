package netappsnapshotpolicy


type NetappSnapshotPolicyHourlySchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_snapshot_policy#minute NetappSnapshotPolicy#minute}.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_snapshot_policy#snapshots_to_keep NetappSnapshotPolicy#snapshots_to_keep}.
	SnapshotsToKeep *float64 `field:"required" json:"snapshotsToKeep" yaml:"snapshotsToKeep"`
}


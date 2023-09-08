// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthbasecluster


type HdinsightHbaseClusterRolesWorkerNodeAutoscaleRecurrence struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/hdinsight_hbase_cluster#schedule HdinsightHbaseCluster#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/hdinsight_hbase_cluster#timezone HdinsightHbaseCluster#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}


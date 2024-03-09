// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster


type RedhatOpenshiftClusterNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/redhat_openshift_cluster#pod_cidr RedhatOpenshiftCluster#pod_cidr}.
	PodCidr *string `field:"required" json:"podCidr" yaml:"podCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/redhat_openshift_cluster#service_cidr RedhatOpenshiftCluster#service_cidr}.
	ServiceCidr *string `field:"required" json:"serviceCidr" yaml:"serviceCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/redhat_openshift_cluster#outbound_type RedhatOpenshiftCluster#outbound_type}.
	OutboundType *string `field:"optional" json:"outboundType" yaml:"outboundType"`
}


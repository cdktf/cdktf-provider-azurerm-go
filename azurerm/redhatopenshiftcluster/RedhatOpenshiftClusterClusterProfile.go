// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster


type RedhatOpenshiftClusterClusterProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/redhat_openshift_cluster#domain RedhatOpenshiftCluster#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/redhat_openshift_cluster#version RedhatOpenshiftCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/redhat_openshift_cluster#fips_enabled RedhatOpenshiftCluster#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/redhat_openshift_cluster#pull_secret RedhatOpenshiftCluster#pull_secret}.
	PullSecret *string `field:"optional" json:"pullSecret" yaml:"pullSecret"`
}


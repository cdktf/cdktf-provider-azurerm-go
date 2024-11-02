// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster


type RedhatOpenshiftClusterServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/redhat_openshift_cluster#client_id RedhatOpenshiftCluster#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/redhat_openshift_cluster#client_secret RedhatOpenshiftCluster#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}


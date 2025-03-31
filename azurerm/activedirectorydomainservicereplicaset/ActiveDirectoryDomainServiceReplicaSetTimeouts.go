// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomainservicereplicaset


type ActiveDirectoryDomainServiceReplicaSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.25.0/docs/resources/active_directory_domain_service_replica_set#create ActiveDirectoryDomainServiceReplicaSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.25.0/docs/resources/active_directory_domain_service_replica_set#delete ActiveDirectoryDomainServiceReplicaSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.25.0/docs/resources/active_directory_domain_service_replica_set#read ActiveDirectoryDomainServiceReplicaSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


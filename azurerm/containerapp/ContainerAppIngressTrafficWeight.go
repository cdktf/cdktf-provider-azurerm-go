// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppIngressTrafficWeight struct {
	// The percentage of traffic to send to this revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_app#percentage ContainerApp#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// The label to apply to the revision as a name prefix for routing traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_app#label ContainerApp#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// This traffic Weight relates to the latest stable Container Revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_app#latest_revision ContainerApp#latest_revision}
	LatestRevision interface{} `field:"optional" json:"latestRevision" yaml:"latestRevision"`
	// The suffix string to append to the revision.
	//
	// This must be unique for the Container App's lifetime. A default hash created by the service will be used if this value is omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_app#revision_suffix ContainerApp#revision_suffix}
	RevisionSuffix *string `field:"optional" json:"revisionSuffix" yaml:"revisionSuffix"`
}


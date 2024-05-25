// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationhybridrunbookworkergroup


type AutomationHybridRunbookWorkerGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/automation_hybrid_runbook_worker_group#create AutomationHybridRunbookWorkerGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/automation_hybrid_runbook_worker_group#delete AutomationHybridRunbookWorkerGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/automation_hybrid_runbook_worker_group#read AutomationHybridRunbookWorkerGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/automation_hybrid_runbook_worker_group#update AutomationHybridRunbookWorkerGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


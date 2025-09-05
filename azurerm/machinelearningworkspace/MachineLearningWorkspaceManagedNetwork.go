// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningworkspace


type MachineLearningWorkspaceManagedNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/machine_learning_workspace#isolation_mode MachineLearningWorkspace#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/machine_learning_workspace#provision_on_creation_enabled MachineLearningWorkspace#provision_on_creation_enabled}.
	ProvisionOnCreationEnabled interface{} `field:"optional" json:"provisionOnCreationEnabled" yaml:"provisionOnCreationEnabled"`
}


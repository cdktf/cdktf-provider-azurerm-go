// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package costanomalyalert


type CostAnomalyAlertTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/cost_anomaly_alert#create CostAnomalyAlert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/cost_anomaly_alert#delete CostAnomalyAlert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/cost_anomaly_alert#read CostAnomalyAlert#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/cost_anomaly_alert#update CostAnomalyAlert#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


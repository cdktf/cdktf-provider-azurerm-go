// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databricksworkspace


type DatabricksWorkspaceEnhancedSecurityCompliance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/databricks_workspace#automatic_cluster_update_enabled DatabricksWorkspace#automatic_cluster_update_enabled}.
	AutomaticClusterUpdateEnabled interface{} `field:"optional" json:"automaticClusterUpdateEnabled" yaml:"automaticClusterUpdateEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/databricks_workspace#compliance_security_profile_enabled DatabricksWorkspace#compliance_security_profile_enabled}.
	ComplianceSecurityProfileEnabled interface{} `field:"optional" json:"complianceSecurityProfileEnabled" yaml:"complianceSecurityProfileEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/databricks_workspace#compliance_security_profile_standards DatabricksWorkspace#compliance_security_profile_standards}.
	ComplianceSecurityProfileStandards *[]*string `field:"optional" json:"complianceSecurityProfileStandards" yaml:"complianceSecurityProfileStandards"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/databricks_workspace#enhanced_security_monitoring_enabled DatabricksWorkspace#enhanced_security_monitoring_enabled}.
	EnhancedSecurityMonitoringEnabled interface{} `field:"optional" json:"enhancedSecurityMonitoringEnabled" yaml:"enhancedSecurityMonitoringEnabled"`
}


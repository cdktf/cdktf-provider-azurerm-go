// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterMaintenanceWindowNodeOs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#duration KubernetesCluster#duration}.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#frequency KubernetesCluster#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#interval KubernetesCluster#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#day_of_month KubernetesCluster#day_of_month}.
	DayOfMonth *float64 `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#day_of_week KubernetesCluster#day_of_week}.
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// not_allowed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#not_allowed KubernetesCluster#not_allowed}
	NotAllowed interface{} `field:"optional" json:"notAllowed" yaml:"notAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#start_date KubernetesCluster#start_date}.
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#start_time KubernetesCluster#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#utc_offset KubernetesCluster#utc_offset}.
	UtcOffset *string `field:"optional" json:"utcOffset" yaml:"utcOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster#week_index KubernetesCluster#week_index}.
	WeekIndex *string `field:"optional" json:"weekIndex" yaml:"weekIndex"`
}


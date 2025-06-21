// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedhatOpenshiftClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// api_server_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#api_server_profile RedhatOpenshiftCluster#api_server_profile}
	ApiServerProfile *RedhatOpenshiftClusterApiServerProfile `field:"required" json:"apiServerProfile" yaml:"apiServerProfile"`
	// cluster_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#cluster_profile RedhatOpenshiftCluster#cluster_profile}
	ClusterProfile *RedhatOpenshiftClusterClusterProfile `field:"required" json:"clusterProfile" yaml:"clusterProfile"`
	// ingress_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#ingress_profile RedhatOpenshiftCluster#ingress_profile}
	IngressProfile *RedhatOpenshiftClusterIngressProfile `field:"required" json:"ingressProfile" yaml:"ingressProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#location RedhatOpenshiftCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// main_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#main_profile RedhatOpenshiftCluster#main_profile}
	MainProfile *RedhatOpenshiftClusterMainProfile `field:"required" json:"mainProfile" yaml:"mainProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#name RedhatOpenshiftCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#network_profile RedhatOpenshiftCluster#network_profile}
	NetworkProfile *RedhatOpenshiftClusterNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#resource_group_name RedhatOpenshiftCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// service_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#service_principal RedhatOpenshiftCluster#service_principal}
	ServicePrincipal *RedhatOpenshiftClusterServicePrincipal `field:"required" json:"servicePrincipal" yaml:"servicePrincipal"`
	// worker_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#worker_profile RedhatOpenshiftCluster#worker_profile}
	WorkerProfile *RedhatOpenshiftClusterWorkerProfile `field:"required" json:"workerProfile" yaml:"workerProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#id RedhatOpenshiftCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#tags RedhatOpenshiftCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/redhat_openshift_cluster#timeouts RedhatOpenshiftCluster#timeouts}
	Timeouts *RedhatOpenshiftClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


package springcloudapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#name SpringCloudApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#resource_group_name SpringCloudApp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#service_name SpringCloudApp#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#addon_json SpringCloudApp#addon_json}.
	AddonJson *string `field:"optional" json:"addonJson" yaml:"addonJson"`
	// custom_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#custom_persistent_disk SpringCloudApp#custom_persistent_disk}
	CustomPersistentDisk interface{} `field:"optional" json:"customPersistentDisk" yaml:"customPersistentDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#https_only SpringCloudApp#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#id SpringCloudApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#identity SpringCloudApp#identity}
	Identity *SpringCloudAppIdentity `field:"optional" json:"identity" yaml:"identity"`
	// ingress_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#ingress_settings SpringCloudApp#ingress_settings}
	IngressSettings *SpringCloudAppIngressSettings `field:"optional" json:"ingressSettings" yaml:"ingressSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#is_public SpringCloudApp#is_public}.
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#persistent_disk SpringCloudApp#persistent_disk}
	PersistentDisk *SpringCloudAppPersistentDisk `field:"optional" json:"persistentDisk" yaml:"persistentDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#public_endpoint_enabled SpringCloudApp#public_endpoint_enabled}.
	PublicEndpointEnabled interface{} `field:"optional" json:"publicEndpointEnabled" yaml:"publicEndpointEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#timeouts SpringCloudApp#timeouts}
	Timeouts *SpringCloudAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/spring_cloud_app#tls_enabled SpringCloudApp#tls_enabled}.
	TlsEnabled interface{} `field:"optional" json:"tlsEnabled" yaml:"tlsEnabled"`
}


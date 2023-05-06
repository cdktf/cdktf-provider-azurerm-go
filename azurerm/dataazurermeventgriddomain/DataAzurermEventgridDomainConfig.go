package dataazurermeventgriddomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermEventgridDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/eventgrid_domain#name DataAzurermEventgridDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/eventgrid_domain#resource_group_name DataAzurermEventgridDomain#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/eventgrid_domain#id DataAzurermEventgridDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/eventgrid_domain#inbound_ip_rule DataAzurermEventgridDomain#inbound_ip_rule}.
	InboundIpRule interface{} `field:"optional" json:"inboundIpRule" yaml:"inboundIpRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/eventgrid_domain#public_network_access_enabled DataAzurermEventgridDomain#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/eventgrid_domain#tags DataAzurermEventgridDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/data-sources/eventgrid_domain#timeouts DataAzurermEventgridDomain#timeouts}
	Timeouts *DataAzurermEventgridDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


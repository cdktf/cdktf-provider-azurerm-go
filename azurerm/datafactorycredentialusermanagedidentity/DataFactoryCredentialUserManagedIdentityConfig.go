// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorycredentialusermanagedidentity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryCredentialUserManagedIdentityConfig struct {
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
	// The resource ID of the parent Data Factory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/data_factory_credential_user_managed_identity#data_factory_id DataFactoryCredentialUserManagedIdentity#data_factory_id}
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// The resource ID of the User Assigned Managed Identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/data_factory_credential_user_managed_identity#identity_id DataFactoryCredentialUserManagedIdentity#identity_id}
	IdentityId *string `field:"required" json:"identityId" yaml:"identityId"`
	// The desired name of the credential resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/data_factory_credential_user_managed_identity#name DataFactoryCredentialUserManagedIdentity#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Optional) List of string annotations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/data_factory_credential_user_managed_identity#annotations DataFactoryCredentialUserManagedIdentity#annotations}
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// (Optional) Short text description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/data_factory_credential_user_managed_identity#description DataFactoryCredentialUserManagedIdentity#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/data_factory_credential_user_managed_identity#id DataFactoryCredentialUserManagedIdentity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/data_factory_credential_user_managed_identity#timeouts DataFactoryCredentialUserManagedIdentity#timeouts}
	Timeouts *DataFactoryCredentialUserManagedIdentityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


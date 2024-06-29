// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment


type NginxDeploymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/nginx_deployment#root_file NginxDeployment#root_file}.
	RootFile *string `field:"required" json:"rootFile" yaml:"rootFile"`
	// config_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/nginx_deployment#config_file NginxDeployment#config_file}
	ConfigFile interface{} `field:"optional" json:"configFile" yaml:"configFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/nginx_deployment#package_data NginxDeployment#package_data}.
	PackageData *string `field:"optional" json:"packageData" yaml:"packageData"`
	// protected_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/nginx_deployment#protected_file NginxDeployment#protected_file}
	ProtectedFile interface{} `field:"optional" json:"protectedFile" yaml:"protectedFile"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redislinkedserver


type RedisLinkedServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/redis_linked_server#create RedisLinkedServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/redis_linked_server#delete RedisLinkedServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/redis_linked_server#read RedisLinkedServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


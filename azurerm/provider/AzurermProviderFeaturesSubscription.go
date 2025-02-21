// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesSubscription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs#prevent_cancellation_on_destroy AzurermProvider#prevent_cancellation_on_destroy}.
	PreventCancellationOnDestroy interface{} `field:"optional" json:"preventCancellationOnDestroy" yaml:"preventCancellationOnDestroy"`
}


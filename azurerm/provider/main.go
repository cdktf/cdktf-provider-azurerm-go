package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		reflect.TypeOf((*AzurermProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "auxiliaryTenantIds", GoGetter: "AuxiliaryTenantIds"},
			_jsii_.MemberProperty{JsiiProperty: "auxiliaryTenantIdsInput", GoGetter: "AuxiliaryTenantIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificate", GoGetter: "ClientCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificateInput", GoGetter: "ClientCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePassword", GoGetter: "ClientCertificatePassword"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePasswordInput", GoGetter: "ClientCertificatePasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePath", GoGetter: "ClientCertificatePath"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePathInput", GoGetter: "ClientCertificatePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "disableCorrelationRequestId", GoGetter: "DisableCorrelationRequestId"},
			_jsii_.MemberProperty{JsiiProperty: "disableCorrelationRequestIdInput", GoGetter: "DisableCorrelationRequestIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableTerraformPartnerId", GoGetter: "DisableTerraformPartnerId"},
			_jsii_.MemberProperty{JsiiProperty: "disableTerraformPartnerIdInput", GoGetter: "DisableTerraformPartnerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "environmentInput", GoGetter: "EnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "features", GoGetter: "Features"},
			_jsii_.MemberProperty{JsiiProperty: "featuresInput", GoGetter: "FeaturesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHost", GoGetter: "MetadataHost"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHostInput", GoGetter: "MetadataHostInput"},
			_jsii_.MemberProperty{JsiiProperty: "msiEndpoint", GoGetter: "MsiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "msiEndpointInput", GoGetter: "MsiEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestToken", GoGetter: "OidcRequestToken"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestTokenInput", GoGetter: "OidcRequestTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestUrl", GoGetter: "OidcRequestUrl"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestUrlInput", GoGetter: "OidcRequestUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcToken", GoGetter: "OidcToken"},
			_jsii_.MemberProperty{JsiiProperty: "oidcTokenFilePath", GoGetter: "OidcTokenFilePath"},
			_jsii_.MemberProperty{JsiiProperty: "oidcTokenFilePathInput", GoGetter: "OidcTokenFilePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcTokenInput", GoGetter: "OidcTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "partnerId", GoGetter: "PartnerId"},
			_jsii_.MemberProperty{JsiiProperty: "partnerIdInput", GoGetter: "PartnerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuxiliaryTenantIds", GoMethod: "ResetAuxiliaryTenantIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCertificate", GoMethod: "ResetClientCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCertificatePassword", GoMethod: "ResetClientCertificatePassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCertificatePath", GoMethod: "ResetClientCertificatePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecret", GoMethod: "ResetClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableCorrelationRequestId", GoMethod: "ResetDisableCorrelationRequestId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableTerraformPartnerId", GoMethod: "ResetDisableTerraformPartnerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironment", GoMethod: "ResetEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataHost", GoMethod: "ResetMetadataHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetMsiEndpoint", GoMethod: "ResetMsiEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcRequestToken", GoMethod: "ResetOidcRequestToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcRequestUrl", GoMethod: "ResetOidcRequestUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcToken", GoMethod: "ResetOidcToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcTokenFilePath", GoMethod: "ResetOidcTokenFilePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPartnerId", GoMethod: "ResetPartnerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipProviderRegistration", GoMethod: "ResetSkipProviderRegistration"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageUseAzuread", GoMethod: "ResetStorageUseAzuread"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubscriptionId", GoMethod: "ResetSubscriptionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTenantId", GoMethod: "ResetTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseCli", GoMethod: "ResetUseCli"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseMsi", GoMethod: "ResetUseMsi"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseOidc", GoMethod: "ResetUseOidc"},
			_jsii_.MemberProperty{JsiiProperty: "skipProviderRegistration", GoGetter: "SkipProviderRegistration"},
			_jsii_.MemberProperty{JsiiProperty: "skipProviderRegistrationInput", GoGetter: "SkipProviderRegistrationInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageUseAzuread", GoGetter: "StorageUseAzuread"},
			_jsii_.MemberProperty{JsiiProperty: "storageUseAzureadInput", GoGetter: "StorageUseAzureadInput"},
			_jsii_.MemberProperty{JsiiProperty: "subscriptionId", GoGetter: "SubscriptionId"},
			_jsii_.MemberProperty{JsiiProperty: "subscriptionIdInput", GoGetter: "SubscriptionIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "tenantIdInput", GoGetter: "TenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useCli", GoGetter: "UseCli"},
			_jsii_.MemberProperty{JsiiProperty: "useCliInput", GoGetter: "UseCliInput"},
			_jsii_.MemberProperty{JsiiProperty: "useMsi", GoGetter: "UseMsi"},
			_jsii_.MemberProperty{JsiiProperty: "useMsiInput", GoGetter: "UseMsiInput"},
			_jsii_.MemberProperty{JsiiProperty: "useOidc", GoGetter: "UseOidc"},
			_jsii_.MemberProperty{JsiiProperty: "useOidcInput", GoGetter: "UseOidcInput"},
		},
		func() interface{} {
			j := jsiiProxy_AzurermProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderConfig",
		reflect.TypeOf((*AzurermProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeatures",
		reflect.TypeOf((*AzurermProviderFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesApiManagement",
		reflect.TypeOf((*AzurermProviderFeaturesApiManagement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesAppConfiguration",
		reflect.TypeOf((*AzurermProviderFeaturesAppConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesApplicationInsights",
		reflect.TypeOf((*AzurermProviderFeaturesApplicationInsights)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesCognitiveAccount",
		reflect.TypeOf((*AzurermProviderFeaturesCognitiveAccount)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesKeyVault",
		reflect.TypeOf((*AzurermProviderFeaturesKeyVault)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesLogAnalyticsWorkspace",
		reflect.TypeOf((*AzurermProviderFeaturesLogAnalyticsWorkspace)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesManagedDisk",
		reflect.TypeOf((*AzurermProviderFeaturesManagedDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesNetwork",
		reflect.TypeOf((*AzurermProviderFeaturesNetwork)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesResourceGroup",
		reflect.TypeOf((*AzurermProviderFeaturesResourceGroup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesTemplateDeployment",
		reflect.TypeOf((*AzurermProviderFeaturesTemplateDeployment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesVirtualMachine",
		reflect.TypeOf((*AzurermProviderFeaturesVirtualMachine)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.provider.AzurermProviderFeaturesVirtualMachineScaleSet",
		reflect.TypeOf((*AzurermProviderFeaturesVirtualMachineScaleSet)(nil)).Elem(),
	)
}

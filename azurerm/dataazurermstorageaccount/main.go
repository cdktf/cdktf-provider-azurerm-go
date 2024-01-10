// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstorageaccount

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		reflect.TypeOf((*DataAzurermStorageAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessTier", GoGetter: "AccessTier"},
			_jsii_.MemberProperty{JsiiProperty: "accountKind", GoGetter: "AccountKind"},
			_jsii_.MemberProperty{JsiiProperty: "accountReplicationType", GoGetter: "AccountReplicationType"},
			_jsii_.MemberProperty{JsiiProperty: "accountTier", GoGetter: "AccountTier"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowNestedItemsToBePublic", GoGetter: "AllowNestedItemsToBePublic"},
			_jsii_.MemberProperty{JsiiProperty: "azureFilesAuthentication", GoGetter: "AzureFilesAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customDomain", GoGetter: "CustomDomain"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableHttpsTrafficOnly", GoGetter: "EnableHttpsTrafficOnly"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identity", GoGetter: "Identity"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureEncryptionEnabled", GoGetter: "InfrastructureEncryptionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isHnsEnabled", GoGetter: "IsHnsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "minTlsVersion", GoGetter: "MinTlsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minTlsVersionInput", GoGetter: "MinTlsVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "nfsv3Enabled", GoGetter: "Nfsv3Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "primaryAccessKey", GoGetter: "PrimaryAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobConnectionString", GoGetter: "PrimaryBlobConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobEndpoint", GoGetter: "PrimaryBlobEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobHost", GoGetter: "PrimaryBlobHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobInternetEndpoint", GoGetter: "PrimaryBlobInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobInternetHost", GoGetter: "PrimaryBlobInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobMicrosoftEndpoint", GoGetter: "PrimaryBlobMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobMicrosoftHost", GoGetter: "PrimaryBlobMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryConnectionString", GoGetter: "PrimaryConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsEndpoint", GoGetter: "PrimaryDfsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsHost", GoGetter: "PrimaryDfsHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsInternetEndpoint", GoGetter: "PrimaryDfsInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsInternetHost", GoGetter: "PrimaryDfsInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsMicrosoftEndpoint", GoGetter: "PrimaryDfsMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsMicrosoftHost", GoGetter: "PrimaryDfsMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileEndpoint", GoGetter: "PrimaryFileEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileHost", GoGetter: "PrimaryFileHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileInternetEndpoint", GoGetter: "PrimaryFileInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileInternetHost", GoGetter: "PrimaryFileInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileMicrosoftEndpoint", GoGetter: "PrimaryFileMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileMicrosoftHost", GoGetter: "PrimaryFileMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryLocation", GoGetter: "PrimaryLocation"},
			_jsii_.MemberProperty{JsiiProperty: "primaryQueueEndpoint", GoGetter: "PrimaryQueueEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryQueueHost", GoGetter: "PrimaryQueueHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryQueueMicrosoftEndpoint", GoGetter: "PrimaryQueueMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryQueueMicrosoftHost", GoGetter: "PrimaryQueueMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryTableEndpoint", GoGetter: "PrimaryTableEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryTableHost", GoGetter: "PrimaryTableHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryTableMicrosoftEndpoint", GoGetter: "PrimaryTableMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryTableMicrosoftHost", GoGetter: "PrimaryTableMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebEndpoint", GoGetter: "PrimaryWebEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebHost", GoGetter: "PrimaryWebHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebInternetEndpoint", GoGetter: "PrimaryWebInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebInternetHost", GoGetter: "PrimaryWebInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebMicrosoftEndpoint", GoGetter: "PrimaryWebMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebMicrosoftHost", GoGetter: "PrimaryWebMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "queueEncryptionKeyType", GoGetter: "QueueEncryptionKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTlsVersion", GoMethod: "ResetMinTlsVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryAccessKey", GoGetter: "SecondaryAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobConnectionString", GoGetter: "SecondaryBlobConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobEndpoint", GoGetter: "SecondaryBlobEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobHost", GoGetter: "SecondaryBlobHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobInternetEndpoint", GoGetter: "SecondaryBlobInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobInternetHost", GoGetter: "SecondaryBlobInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobMicrosoftEndpoint", GoGetter: "SecondaryBlobMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobMicrosoftHost", GoGetter: "SecondaryBlobMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryConnectionString", GoGetter: "SecondaryConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsEndpoint", GoGetter: "SecondaryDfsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsHost", GoGetter: "SecondaryDfsHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsInternetEndpoint", GoGetter: "SecondaryDfsInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsInternetHost", GoGetter: "SecondaryDfsInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsMicrosoftEndpoint", GoGetter: "SecondaryDfsMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsMicrosoftHost", GoGetter: "SecondaryDfsMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileEndpoint", GoGetter: "SecondaryFileEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileHost", GoGetter: "SecondaryFileHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileInternetEndpoint", GoGetter: "SecondaryFileInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileInternetHost", GoGetter: "SecondaryFileInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileMicrosoftEndpoint", GoGetter: "SecondaryFileMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileMicrosoftHost", GoGetter: "SecondaryFileMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryLocation", GoGetter: "SecondaryLocation"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryQueueEndpoint", GoGetter: "SecondaryQueueEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryQueueHost", GoGetter: "SecondaryQueueHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryQueueMicrosoftEndpoint", GoGetter: "SecondaryQueueMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryQueueMicrosoftHost", GoGetter: "SecondaryQueueMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryTableEndpoint", GoGetter: "SecondaryTableEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryTableHost", GoGetter: "SecondaryTableHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryTableMicrosoftEndpoint", GoGetter: "SecondaryTableMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryTableMicrosoftHost", GoGetter: "SecondaryTableMicrosoftHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebEndpoint", GoGetter: "SecondaryWebEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebHost", GoGetter: "SecondaryWebHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebInternetEndpoint", GoGetter: "SecondaryWebInternetEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebInternetHost", GoGetter: "SecondaryWebInternetHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebMicrosoftEndpoint", GoGetter: "SecondaryWebMicrosoftEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebMicrosoftHost", GoGetter: "SecondaryWebMicrosoftHost"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tableEncryptionKeyType", GoGetter: "TableEncryptionKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccount{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountAzureFilesAuthentication",
		reflect.TypeOf((*DataAzurermStorageAccountAzureFilesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectory",
		reflect.TypeOf((*DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectory)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectoryList",
		reflect.TypeOf((*DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectoryList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectoryList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference",
		reflect.TypeOf((*DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainGuid", GoGetter: "DomainGuid"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainSid", GoGetter: "DomainSid"},
			_jsii_.MemberProperty{JsiiProperty: "forestName", GoGetter: "ForestName"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "netbiosDomainName", GoGetter: "NetbiosDomainName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageSid", GoGetter: "StorageSid"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountAzureFilesAuthenticationActiveDirectoryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountAzureFilesAuthenticationList",
		reflect.TypeOf((*DataAzurermStorageAccountAzureFilesAuthenticationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountAzureFilesAuthenticationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountAzureFilesAuthenticationOutputReference",
		reflect.TypeOf((*DataAzurermStorageAccountAzureFilesAuthenticationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeDirectory", GoGetter: "ActiveDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "directoryType", GoGetter: "DirectoryType"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountAzureFilesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountConfig",
		reflect.TypeOf((*DataAzurermStorageAccountConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountCustomDomain",
		reflect.TypeOf((*DataAzurermStorageAccountCustomDomain)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountCustomDomainList",
		reflect.TypeOf((*DataAzurermStorageAccountCustomDomainList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountCustomDomainList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountCustomDomainOutputReference",
		reflect.TypeOf((*DataAzurermStorageAccountCustomDomainOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountCustomDomainOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountIdentity",
		reflect.TypeOf((*DataAzurermStorageAccountIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountIdentityList",
		reflect.TypeOf((*DataAzurermStorageAccountIdentityList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountIdentityList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountIdentityOutputReference",
		reflect.TypeOf((*DataAzurermStorageAccountIdentityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "identityIds", GoGetter: "IdentityIds"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountIdentityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountTimeouts",
		reflect.TypeOf((*DataAzurermStorageAccountTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountTimeoutsOutputReference",
		reflect.TypeOf((*DataAzurermStorageAccountTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

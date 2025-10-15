// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermoracleautonomousdatabaseclonefromdatabase

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
		reflect.TypeOf((*DataAzurermOracleAutonomousDatabaseCloneFromDatabase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actualUsedDataStorageSizeInTb", GoGetter: "ActualUsedDataStorageSizeInTb"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allocatedStorageSizeInTb", GoGetter: "AllocatedStorageSizeInTb"},
			_jsii_.MemberProperty{JsiiProperty: "allowedIpAddresses", GoGetter: "AllowedIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingEnabled", GoGetter: "AutoScalingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingForStorageEnabled", GoGetter: "AutoScalingForStorageEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "availableUpgradeVersions", GoGetter: "AvailableUpgradeVersions"},
			_jsii_.MemberProperty{JsiiProperty: "backupRetentionPeriodInDays", GoGetter: "BackupRetentionPeriodInDays"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "characterSet", GoGetter: "CharacterSet"},
			_jsii_.MemberProperty{JsiiProperty: "computeCount", GoGetter: "ComputeCount"},
			_jsii_.MemberProperty{JsiiProperty: "computeModel", GoGetter: "ComputeModel"},
			_jsii_.MemberProperty{JsiiProperty: "connectionStrings", GoGetter: "ConnectionStrings"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCount", GoGetter: "CpuCoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "customerContacts", GoGetter: "CustomerContacts"},
			_jsii_.MemberProperty{JsiiProperty: "databaseVersion", GoGetter: "DatabaseVersion"},
			_jsii_.MemberProperty{JsiiProperty: "databaseWorkload", GoGetter: "DatabaseWorkload"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInGb", GoGetter: "DataStorageSizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInTb", GoGetter: "DataStorageSizeInTb"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "failedDataRecoveryInSeconds", GoGetter: "FailedDataRecoveryInSeconds"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "inMemoryAreaInGb", GoGetter: "InMemoryAreaInGb"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleDetails", GoGetter: "LifecycleDetails"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleState", GoGetter: "LifecycleState"},
			_jsii_.MemberProperty{JsiiProperty: "localAdgAutoFailoverMaxDataLossLimitInSeconds", GoGetter: "LocalAdgAutoFailoverMaxDataLossLimitInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "localDataGuardEnabled", GoGetter: "LocalDataGuardEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "longTermBackupSchedule", GoGetter: "LongTermBackupSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "memoryPerOracleComputeUnitInGb", GoGetter: "MemoryPerOracleComputeUnitInGb"},
			_jsii_.MemberProperty{JsiiProperty: "mtlsConnectionRequired", GoGetter: "MtlsConnectionRequired"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "nationalCharacterSet", GoGetter: "NationalCharacterSet"},
			_jsii_.MemberProperty{JsiiProperty: "nextLongTermBackupTimestamp", GoGetter: "NextLongTermBackupTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ocid", GoGetter: "Ocid"},
			_jsii_.MemberProperty{JsiiProperty: "ociUrl", GoGetter: "OciUrl"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "peerDatabaseIds", GoGetter: "PeerDatabaseIds"},
			_jsii_.MemberProperty{JsiiProperty: "preview", GoGetter: "Preview"},
			_jsii_.MemberProperty{JsiiProperty: "previewVersionWithServiceTermsAccepted", GoGetter: "PreviewVersionWithServiceTermsAccepted"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointIp", GoGetter: "PrivateEndpointIp"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointLabel", GoGetter: "PrivateEndpointLabel"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointUrl", GoGetter: "PrivateEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisionableCpus", GoGetter: "ProvisionableCpus"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "reconnectCloneEnabled", GoGetter: "ReconnectCloneEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "refreshableClone", GoGetter: "RefreshableClone"},
			_jsii_.MemberProperty{JsiiProperty: "refreshableStatus", GoGetter: "RefreshableStatus"},
			_jsii_.MemberProperty{JsiiProperty: "remoteDataGuardEnabled", GoGetter: "RemoteDataGuardEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceConsoleUrl", GoGetter: "ServiceConsoleUrl"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseId", GoGetter: "SourceAutonomousDatabaseId"},
			_jsii_.MemberProperty{JsiiProperty: "sqlWebDeveloperUrl", GoGetter: "SqlWebDeveloperUrl"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "supportedRegionsToCloneTo", GoGetter: "SupportedRegionsToCloneTo"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeCreatedInUtc", GoGetter: "TimeCreatedInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeDataGuardRoleChangedInUtc", GoGetter: "TimeDataGuardRoleChangedInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeDeletionOfFreeAutonomousDatabaseInUtc", GoGetter: "TimeDeletionOfFreeAutonomousDatabaseInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeLocalDataGuardEnabledInUtc", GoGetter: "TimeLocalDataGuardEnabledInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeMaintenanceBeginInUtc", GoGetter: "TimeMaintenanceBeginInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeMaintenanceEndInUtc", GoGetter: "TimeMaintenanceEndInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfLastFailoverInUtc", GoGetter: "TimeOfLastFailoverInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfLastRefreshInUtc", GoGetter: "TimeOfLastRefreshInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfLastRefreshPointInUtc", GoGetter: "TimeOfLastRefreshPointInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfLastSwitchoverInUtc", GoGetter: "TimeOfLastSwitchoverInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeReclamationOfFreeAutonomousDatabaseInUtc", GoGetter: "TimeReclamationOfFreeAutonomousDatabaseInUtc"},
			_jsii_.MemberProperty{JsiiProperty: "timeUntilReconnectInUtc", GoGetter: "TimeUntilReconnectInUtc"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "usedDataStorageSizeInGb", GoGetter: "UsedDataStorageSizeInGb"},
			_jsii_.MemberProperty{JsiiProperty: "usedDataStorageSizeInTb", GoGetter: "UsedDataStorageSizeInTb"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabaseConfig",
		reflect.TypeOf((*DataAzurermOracleAutonomousDatabaseCloneFromDatabaseConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupSchedule",
		reflect.TypeOf((*DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupSchedule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleList",
		reflect.TypeOf((*DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleList)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleOutputReference",
		reflect.TypeOf((*DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
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
			_jsii_.MemberProperty{JsiiProperty: "repeatCadence", GoGetter: "RepeatCadence"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionPeriodInDays", GoGetter: "RetentionPeriodInDays"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfBackupInUtc", GoGetter: "TimeOfBackupInUtc"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeouts",
		reflect.TypeOf((*DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference",
		reflect.TypeOf((*DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

package mssqldatabase

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		reflect.TypeOf((*MssqlDatabase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoPauseDelayInMinutes", GoGetter: "AutoPauseDelayInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "autoPauseDelayInMinutesInput", GoGetter: "AutoPauseDelayInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "collation", GoGetter: "Collation"},
			_jsii_.MemberProperty{JsiiProperty: "collationInput", GoGetter: "CollationInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createMode", GoGetter: "CreateMode"},
			_jsii_.MemberProperty{JsiiProperty: "createModeInput", GoGetter: "CreateModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationSourceDatabaseId", GoGetter: "CreationSourceDatabaseId"},
			_jsii_.MemberProperty{JsiiProperty: "creationSourceDatabaseIdInput", GoGetter: "CreationSourceDatabaseIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "elasticPoolId", GoGetter: "ElasticPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "elasticPoolIdInput", GoGetter: "ElasticPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "geoBackupEnabled", GoGetter: "GeoBackupEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "geoBackupEnabledInput", GoGetter: "GeoBackupEnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "import", GoGetter: "Import"},
			_jsii_.MemberProperty{JsiiProperty: "importInput", GoGetter: "ImportInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ledgerEnabled", GoGetter: "LedgerEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ledgerEnabledInput", GoGetter: "LedgerEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "licenseType", GoGetter: "LicenseType"},
			_jsii_.MemberProperty{JsiiProperty: "licenseTypeInput", GoGetter: "LicenseTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "longTermRetentionPolicy", GoGetter: "LongTermRetentionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "longTermRetentionPolicyInput", GoGetter: "LongTermRetentionPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceConfigurationName", GoGetter: "MaintenanceConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceConfigurationNameInput", GoGetter: "MaintenanceConfigurationNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxSizeGb", GoGetter: "MaxSizeGb"},
			_jsii_.MemberProperty{JsiiProperty: "maxSizeGbInput", GoGetter: "MaxSizeGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "minCapacity", GoGetter: "MinCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "minCapacityInput", GoGetter: "MinCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putImport", GoMethod: "PutImport"},
			_jsii_.MemberMethod{JsiiMethod: "putLongTermRetentionPolicy", GoMethod: "PutLongTermRetentionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putShortTermRetentionPolicy", GoMethod: "PutShortTermRetentionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putThreatDetectionPolicy", GoMethod: "PutThreatDetectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readReplicaCount", GoGetter: "ReadReplicaCount"},
			_jsii_.MemberProperty{JsiiProperty: "readReplicaCountInput", GoGetter: "ReadReplicaCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "readScale", GoGetter: "ReadScale"},
			_jsii_.MemberProperty{JsiiProperty: "readScaleInput", GoGetter: "ReadScaleInput"},
			_jsii_.MemberProperty{JsiiProperty: "recoverDatabaseId", GoGetter: "RecoverDatabaseId"},
			_jsii_.MemberProperty{JsiiProperty: "recoverDatabaseIdInput", GoGetter: "RecoverDatabaseIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoPauseDelayInMinutes", GoMethod: "ResetAutoPauseDelayInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetCollation", GoMethod: "ResetCollation"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateMode", GoMethod: "ResetCreateMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreationSourceDatabaseId", GoMethod: "ResetCreationSourceDatabaseId"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticPoolId", GoMethod: "ResetElasticPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeoBackupEnabled", GoMethod: "ResetGeoBackupEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetImport", GoMethod: "ResetImport"},
			_jsii_.MemberMethod{JsiiMethod: "resetLedgerEnabled", GoMethod: "ResetLedgerEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseType", GoMethod: "ResetLicenseType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLongTermRetentionPolicy", GoMethod: "ResetLongTermRetentionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceConfigurationName", GoMethod: "ResetMaintenanceConfigurationName"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxSizeGb", GoMethod: "ResetMaxSizeGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinCapacity", GoMethod: "ResetMinCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadReplicaCount", GoMethod: "ResetReadReplicaCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadScale", GoMethod: "ResetReadScale"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecoverDatabaseId", GoMethod: "ResetRecoverDatabaseId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestoreDroppedDatabaseId", GoMethod: "ResetRestoreDroppedDatabaseId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestorePointInTime", GoMethod: "ResetRestorePointInTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetSampleName", GoMethod: "ResetSampleName"},
			_jsii_.MemberMethod{JsiiMethod: "resetShortTermRetentionPolicy", GoMethod: "ResetShortTermRetentionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkuName", GoMethod: "ResetSkuName"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageAccountType", GoMethod: "ResetStorageAccountType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreatDetectionPolicy", GoMethod: "ResetThreatDetectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransparentDataEncryptionEnabled", GoMethod: "ResetTransparentDataEncryptionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneRedundant", GoMethod: "ResetZoneRedundant"},
			_jsii_.MemberProperty{JsiiProperty: "restoreDroppedDatabaseId", GoGetter: "RestoreDroppedDatabaseId"},
			_jsii_.MemberProperty{JsiiProperty: "restoreDroppedDatabaseIdInput", GoGetter: "RestoreDroppedDatabaseIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "restorePointInTime", GoGetter: "RestorePointInTime"},
			_jsii_.MemberProperty{JsiiProperty: "restorePointInTimeInput", GoGetter: "RestorePointInTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sampleName", GoGetter: "SampleName"},
			_jsii_.MemberProperty{JsiiProperty: "sampleNameInput", GoGetter: "SampleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "serverId", GoGetter: "ServerId"},
			_jsii_.MemberProperty{JsiiProperty: "serverIdInput", GoGetter: "ServerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "shortTermRetentionPolicy", GoGetter: "ShortTermRetentionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "shortTermRetentionPolicyInput", GoGetter: "ShortTermRetentionPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "skuName", GoGetter: "SkuName"},
			_jsii_.MemberProperty{JsiiProperty: "skuNameInput", GoGetter: "SkuNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountType", GoGetter: "StorageAccountType"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountTypeInput", GoGetter: "StorageAccountTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "threatDetectionPolicy", GoGetter: "ThreatDetectionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "threatDetectionPolicyInput", GoGetter: "ThreatDetectionPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "transparentDataEncryptionEnabled", GoGetter: "TransparentDataEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "transparentDataEncryptionEnabledInput", GoGetter: "TransparentDataEncryptionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "zoneRedundant", GoGetter: "ZoneRedundant"},
			_jsii_.MemberProperty{JsiiProperty: "zoneRedundantInput", GoGetter: "ZoneRedundantInput"},
		},
		func() interface{} {
			j := jsiiProxy_MssqlDatabase{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseConfig",
		reflect.TypeOf((*MssqlDatabaseConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseImport",
		reflect.TypeOf((*MssqlDatabaseImport)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseImportOutputReference",
		reflect.TypeOf((*MssqlDatabaseImportOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "administratorLogin", GoGetter: "AdministratorLogin"},
			_jsii_.MemberProperty{JsiiProperty: "administratorLoginInput", GoGetter: "AdministratorLoginInput"},
			_jsii_.MemberProperty{JsiiProperty: "administratorLoginPassword", GoGetter: "AdministratorLoginPassword"},
			_jsii_.MemberProperty{JsiiProperty: "administratorLoginPasswordInput", GoGetter: "AdministratorLoginPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationType", GoGetter: "AuthenticationType"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationTypeInput", GoGetter: "AuthenticationTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetStorageAccountId", GoMethod: "ResetStorageAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountId", GoGetter: "StorageAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountIdInput", GoGetter: "StorageAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageKey", GoGetter: "StorageKey"},
			_jsii_.MemberProperty{JsiiProperty: "storageKeyInput", GoGetter: "StorageKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageKeyType", GoGetter: "StorageKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "storageKeyTypeInput", GoGetter: "StorageKeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageUri", GoGetter: "StorageUri"},
			_jsii_.MemberProperty{JsiiProperty: "storageUriInput", GoGetter: "StorageUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MssqlDatabaseImportOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseLongTermRetentionPolicy",
		reflect.TypeOf((*MssqlDatabaseLongTermRetentionPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseLongTermRetentionPolicyOutputReference",
		reflect.TypeOf((*MssqlDatabaseLongTermRetentionPolicyOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "monthlyRetention", GoGetter: "MonthlyRetention"},
			_jsii_.MemberProperty{JsiiProperty: "monthlyRetentionInput", GoGetter: "MonthlyRetentionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMonthlyRetention", GoMethod: "ResetMonthlyRetention"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeeklyRetention", GoMethod: "ResetWeeklyRetention"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeekOfYear", GoMethod: "ResetWeekOfYear"},
			_jsii_.MemberMethod{JsiiMethod: "resetYearlyRetention", GoMethod: "ResetYearlyRetention"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyRetention", GoGetter: "WeeklyRetention"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyRetentionInput", GoGetter: "WeeklyRetentionInput"},
			_jsii_.MemberProperty{JsiiProperty: "weekOfYear", GoGetter: "WeekOfYear"},
			_jsii_.MemberProperty{JsiiProperty: "weekOfYearInput", GoGetter: "WeekOfYearInput"},
			_jsii_.MemberProperty{JsiiProperty: "yearlyRetention", GoGetter: "YearlyRetention"},
			_jsii_.MemberProperty{JsiiProperty: "yearlyRetentionInput", GoGetter: "YearlyRetentionInput"},
		},
		func() interface{} {
			j := jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseShortTermRetentionPolicy",
		reflect.TypeOf((*MssqlDatabaseShortTermRetentionPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseShortTermRetentionPolicyOutputReference",
		reflect.TypeOf((*MssqlDatabaseShortTermRetentionPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backupIntervalInHours", GoGetter: "BackupIntervalInHours"},
			_jsii_.MemberProperty{JsiiProperty: "backupIntervalInHoursInput", GoGetter: "BackupIntervalInHoursInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBackupIntervalInHours", GoMethod: "ResetBackupIntervalInHours"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDays", GoGetter: "RetentionDays"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDaysInput", GoGetter: "RetentionDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseThreatDetectionPolicy",
		reflect.TypeOf((*MssqlDatabaseThreatDetectionPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseThreatDetectionPolicyOutputReference",
		reflect.TypeOf((*MssqlDatabaseThreatDetectionPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disabledAlerts", GoGetter: "DisabledAlerts"},
			_jsii_.MemberProperty{JsiiProperty: "disabledAlertsInput", GoGetter: "DisabledAlertsInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailAccountAdmins", GoGetter: "EmailAccountAdmins"},
			_jsii_.MemberProperty{JsiiProperty: "emailAccountAdminsInput", GoGetter: "EmailAccountAdminsInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailAddresses", GoGetter: "EmailAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "emailAddressesInput", GoGetter: "EmailAddressesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDisabledAlerts", GoMethod: "ResetDisabledAlerts"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailAccountAdmins", GoMethod: "ResetEmailAccountAdmins"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailAddresses", GoMethod: "ResetEmailAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetentionDays", GoMethod: "ResetRetentionDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetState", GoMethod: "ResetState"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageAccountAccessKey", GoMethod: "ResetStorageAccountAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageEndpoint", GoMethod: "ResetStorageEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDays", GoGetter: "RetentionDays"},
			_jsii_.MemberProperty{JsiiProperty: "retentionDaysInput", GoGetter: "RetentionDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "stateInput", GoGetter: "StateInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountAccessKey", GoGetter: "StorageAccountAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountAccessKeyInput", GoGetter: "StorageAccountAccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageEndpoint", GoGetter: "StorageEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "storageEndpointInput", GoGetter: "StorageEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseTimeouts",
		reflect.TypeOf((*MssqlDatabaseTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseTimeoutsOutputReference",
		reflect.TypeOf((*MssqlDatabaseTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_MssqlDatabaseTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
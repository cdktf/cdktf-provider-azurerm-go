// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/mssqldatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_database azurerm_mssql_database}.
type MssqlDatabase interface {
	cdktf.TerraformResource
	AutoPauseDelayInMinutes() *float64
	SetAutoPauseDelayInMinutes(val *float64)
	AutoPauseDelayInMinutesInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Collation() *string
	SetCollation(val *string)
	CollationInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateMode() *string
	SetCreateMode(val *string)
	CreateModeInput() *string
	CreationSourceDatabaseId() *string
	SetCreationSourceDatabaseId(val *string)
	CreationSourceDatabaseIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ElasticPoolId() *string
	SetElasticPoolId(val *string)
	ElasticPoolIdInput() *string
	EnclaveType() *string
	SetEnclaveType(val *string)
	EnclaveTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoBackupEnabled() interface{}
	SetGeoBackupEnabled(val interface{})
	GeoBackupEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() MssqlDatabaseIdentityOutputReference
	IdentityInput() *MssqlDatabaseIdentity
	IdInput() *string
	Import() MssqlDatabaseImportOutputReference
	ImportInput() *MssqlDatabaseImport
	LedgerEnabled() interface{}
	SetLedgerEnabled(val interface{})
	LedgerEnabledInput() interface{}
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LongTermRetentionPolicy() MssqlDatabaseLongTermRetentionPolicyOutputReference
	LongTermRetentionPolicyInput() *MssqlDatabaseLongTermRetentionPolicy
	MaintenanceConfigurationName() *string
	SetMaintenanceConfigurationName(val *string)
	MaintenanceConfigurationNameInput() *string
	MaxSizeGb() *float64
	SetMaxSizeGb(val *float64)
	MaxSizeGbInput() *float64
	MinCapacity() *float64
	SetMinCapacity(val *float64)
	MinCapacityInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReadReplicaCount() *float64
	SetReadReplicaCount(val *float64)
	ReadReplicaCountInput() *float64
	ReadScale() interface{}
	SetReadScale(val interface{})
	ReadScaleInput() interface{}
	RecoverDatabaseId() *string
	SetRecoverDatabaseId(val *string)
	RecoverDatabaseIdInput() *string
	RecoveryPointId() *string
	SetRecoveryPointId(val *string)
	RecoveryPointIdInput() *string
	RestoreDroppedDatabaseId() *string
	SetRestoreDroppedDatabaseId(val *string)
	RestoreDroppedDatabaseIdInput() *string
	RestoreLongTermRetentionBackupId() *string
	SetRestoreLongTermRetentionBackupId(val *string)
	RestoreLongTermRetentionBackupIdInput() *string
	RestorePointInTime() *string
	SetRestorePointInTime(val *string)
	RestorePointInTimeInput() *string
	SampleName() *string
	SetSampleName(val *string)
	SampleNameInput() *string
	SecondaryType() *string
	SetSecondaryType(val *string)
	SecondaryTypeInput() *string
	ServerId() *string
	SetServerId(val *string)
	ServerIdInput() *string
	ShortTermRetentionPolicy() MssqlDatabaseShortTermRetentionPolicyOutputReference
	ShortTermRetentionPolicyInput() *MssqlDatabaseShortTermRetentionPolicy
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreatDetectionPolicy() MssqlDatabaseThreatDetectionPolicyOutputReference
	ThreatDetectionPolicyInput() *MssqlDatabaseThreatDetectionPolicy
	Timeouts() MssqlDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransparentDataEncryptionEnabled() interface{}
	SetTransparentDataEncryptionEnabled(val interface{})
	TransparentDataEncryptionEnabledInput() interface{}
	TransparentDataEncryptionKeyAutomaticRotationEnabled() interface{}
	SetTransparentDataEncryptionKeyAutomaticRotationEnabled(val interface{})
	TransparentDataEncryptionKeyAutomaticRotationEnabledInput() interface{}
	TransparentDataEncryptionKeyVaultKeyId() *string
	SetTransparentDataEncryptionKeyVaultKeyId(val *string)
	TransparentDataEncryptionKeyVaultKeyIdInput() *string
	ZoneRedundant() interface{}
	SetZoneRedundant(val interface{})
	ZoneRedundantInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutIdentity(value *MssqlDatabaseIdentity)
	PutImport(value *MssqlDatabaseImport)
	PutLongTermRetentionPolicy(value *MssqlDatabaseLongTermRetentionPolicy)
	PutShortTermRetentionPolicy(value *MssqlDatabaseShortTermRetentionPolicy)
	PutThreatDetectionPolicy(value *MssqlDatabaseThreatDetectionPolicy)
	PutTimeouts(value *MssqlDatabaseTimeouts)
	ResetAutoPauseDelayInMinutes()
	ResetCollation()
	ResetCreateMode()
	ResetCreationSourceDatabaseId()
	ResetElasticPoolId()
	ResetEnclaveType()
	ResetGeoBackupEnabled()
	ResetId()
	ResetIdentity()
	ResetImport()
	ResetLedgerEnabled()
	ResetLicenseType()
	ResetLongTermRetentionPolicy()
	ResetMaintenanceConfigurationName()
	ResetMaxSizeGb()
	ResetMinCapacity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadReplicaCount()
	ResetReadScale()
	ResetRecoverDatabaseId()
	ResetRecoveryPointId()
	ResetRestoreDroppedDatabaseId()
	ResetRestoreLongTermRetentionBackupId()
	ResetRestorePointInTime()
	ResetSampleName()
	ResetSecondaryType()
	ResetShortTermRetentionPolicy()
	ResetSkuName()
	ResetStorageAccountType()
	ResetTags()
	ResetThreatDetectionPolicy()
	ResetTimeouts()
	ResetTransparentDataEncryptionEnabled()
	ResetTransparentDataEncryptionKeyAutomaticRotationEnabled()
	ResetTransparentDataEncryptionKeyVaultKeyId()
	ResetZoneRedundant()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MssqlDatabase
type jsiiProxy_MssqlDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MssqlDatabase) AutoPauseDelayInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoPauseDelayInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) AutoPauseDelayInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoPauseDelayInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) CreateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) CreateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) CreationSourceDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationSourceDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) CreationSourceDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationSourceDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ElasticPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ElasticPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) EnclaveType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enclaveType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) EnclaveTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enclaveTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) GeoBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) GeoBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Identity() MssqlDatabaseIdentityOutputReference {
	var returns MssqlDatabaseIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) IdentityInput() *MssqlDatabaseIdentity {
	var returns *MssqlDatabaseIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Import() MssqlDatabaseImportOutputReference {
	var returns MssqlDatabaseImportOutputReference
	_jsii_.Get(
		j,
		"import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ImportInput() *MssqlDatabaseImport {
	var returns *MssqlDatabaseImport
	_jsii_.Get(
		j,
		"importInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) LedgerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ledgerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) LedgerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ledgerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) LongTermRetentionPolicy() MssqlDatabaseLongTermRetentionPolicyOutputReference {
	var returns MssqlDatabaseLongTermRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"longTermRetentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) LongTermRetentionPolicyInput() *MssqlDatabaseLongTermRetentionPolicy {
	var returns *MssqlDatabaseLongTermRetentionPolicy
	_jsii_.Get(
		j,
		"longTermRetentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) MaintenanceConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) MaintenanceConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) MaxSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) MaxSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) MinCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ReadReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readReplicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ReadReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readReplicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ReadScale() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ReadScaleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RecoverDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoverDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RecoverDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoverDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RecoveryPointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RecoveryPointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RestoreDroppedDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDroppedDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RestoreDroppedDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDroppedDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RestoreLongTermRetentionBackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreLongTermRetentionBackupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RestoreLongTermRetentionBackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreLongTermRetentionBackupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RestorePointInTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restorePointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) RestorePointInTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restorePointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) SampleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) SampleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) SecondaryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) SecondaryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ShortTermRetentionPolicy() MssqlDatabaseShortTermRetentionPolicyOutputReference {
	var returns MssqlDatabaseShortTermRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"shortTermRetentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ShortTermRetentionPolicyInput() *MssqlDatabaseShortTermRetentionPolicy {
	var returns *MssqlDatabaseShortTermRetentionPolicy
	_jsii_.Get(
		j,
		"shortTermRetentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ThreatDetectionPolicy() MssqlDatabaseThreatDetectionPolicyOutputReference {
	var returns MssqlDatabaseThreatDetectionPolicyOutputReference
	_jsii_.Get(
		j,
		"threatDetectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ThreatDetectionPolicyInput() *MssqlDatabaseThreatDetectionPolicy {
	var returns *MssqlDatabaseThreatDetectionPolicy
	_jsii_.Get(
		j,
		"threatDetectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) Timeouts() MssqlDatabaseTimeoutsOutputReference {
	var returns MssqlDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TransparentDataEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transparentDataEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TransparentDataEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transparentDataEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TransparentDataEncryptionKeyAutomaticRotationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transparentDataEncryptionKeyAutomaticRotationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TransparentDataEncryptionKeyAutomaticRotationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transparentDataEncryptionKeyAutomaticRotationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TransparentDataEncryptionKeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentDataEncryptionKeyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) TransparentDataEncryptionKeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentDataEncryptionKeyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ZoneRedundant() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabase) ZoneRedundantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundantInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_database azurerm_mssql_database} Resource.
func NewMssqlDatabase(scope constructs.Construct, id *string, config *MssqlDatabaseConfig) MssqlDatabase {
	_init_.Initialize()

	if err := validateNewMssqlDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_database azurerm_mssql_database} Resource.
func NewMssqlDatabase_Override(m MssqlDatabase, scope constructs.Construct, id *string, config *MssqlDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetAutoPauseDelayInMinutes(val *float64) {
	if err := j.validateSetAutoPauseDelayInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoPauseDelayInMinutes",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetCollation(val *string) {
	if err := j.validateSetCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetCreateMode(val *string) {
	if err := j.validateSetCreateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetCreationSourceDatabaseId(val *string) {
	if err := j.validateSetCreationSourceDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationSourceDatabaseId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetElasticPoolId(val *string) {
	if err := j.validateSetElasticPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticPoolId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetEnclaveType(val *string) {
	if err := j.validateSetEnclaveTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enclaveType",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetGeoBackupEnabled(val interface{}) {
	if err := j.validateSetGeoBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetLedgerEnabled(val interface{}) {
	if err := j.validateSetLedgerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ledgerEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetMaintenanceConfigurationName(val *string) {
	if err := j.validateSetMaintenanceConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceConfigurationName",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetMaxSizeGb(val *float64) {
	if err := j.validateSetMaxSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSizeGb",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetMinCapacity(val *float64) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetReadReplicaCount(val *float64) {
	if err := j.validateSetReadReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readReplicaCount",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetReadScale(val interface{}) {
	if err := j.validateSetReadScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readScale",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetRecoverDatabaseId(val *string) {
	if err := j.validateSetRecoverDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoverDatabaseId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetRecoveryPointId(val *string) {
	if err := j.validateSetRecoveryPointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetRestoreDroppedDatabaseId(val *string) {
	if err := j.validateSetRestoreDroppedDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreDroppedDatabaseId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetRestoreLongTermRetentionBackupId(val *string) {
	if err := j.validateSetRestoreLongTermRetentionBackupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreLongTermRetentionBackupId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetRestorePointInTime(val *string) {
	if err := j.validateSetRestorePointInTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restorePointInTime",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetSampleName(val *string) {
	if err := j.validateSetSampleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleName",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetSecondaryType(val *string) {
	if err := j.validateSetSecondaryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryType",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetServerId(val *string) {
	if err := j.validateSetServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetStorageAccountType(val *string) {
	if err := j.validateSetStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetTransparentDataEncryptionEnabled(val interface{}) {
	if err := j.validateSetTransparentDataEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentDataEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetTransparentDataEncryptionKeyAutomaticRotationEnabled(val interface{}) {
	if err := j.validateSetTransparentDataEncryptionKeyAutomaticRotationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentDataEncryptionKeyAutomaticRotationEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetTransparentDataEncryptionKeyVaultKeyId(val *string) {
	if err := j.validateSetTransparentDataEncryptionKeyVaultKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentDataEncryptionKeyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase)SetZoneRedundant(val interface{}) {
	if err := j.validateSetZoneRedundantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneRedundant",
		val,
	)
}

// Generates CDKTF code for importing a MssqlDatabase resource upon running "cdktf plan <stack-name>".
func MssqlDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMssqlDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func MssqlDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MssqlDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MssqlDatabase) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MssqlDatabase) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MssqlDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MssqlDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlDatabase) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MssqlDatabase) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutIdentity(value *MssqlDatabaseIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutImport(value *MssqlDatabaseImport) {
	if err := m.validatePutImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putImport",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutLongTermRetentionPolicy(value *MssqlDatabaseLongTermRetentionPolicy) {
	if err := m.validatePutLongTermRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLongTermRetentionPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutShortTermRetentionPolicy(value *MssqlDatabaseShortTermRetentionPolicy) {
	if err := m.validatePutShortTermRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putShortTermRetentionPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutThreatDetectionPolicy(value *MssqlDatabaseThreatDetectionPolicy) {
	if err := m.validatePutThreatDetectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putThreatDetectionPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutTimeouts(value *MssqlDatabaseTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetAutoPauseDelayInMinutes() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoPauseDelayInMinutes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetCollation() {
	_jsii_.InvokeVoid(
		m,
		"resetCollation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetCreateMode() {
	_jsii_.InvokeVoid(
		m,
		"resetCreateMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetCreationSourceDatabaseId() {
	_jsii_.InvokeVoid(
		m,
		"resetCreationSourceDatabaseId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetElasticPoolId() {
	_jsii_.InvokeVoid(
		m,
		"resetElasticPoolId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetEnclaveType() {
	_jsii_.InvokeVoid(
		m,
		"resetEnclaveType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetGeoBackupEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetGeoBackupEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetImport() {
	_jsii_.InvokeVoid(
		m,
		"resetImport",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetLedgerEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetLedgerEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetLicenseType() {
	_jsii_.InvokeVoid(
		m,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetLongTermRetentionPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetLongTermRetentionPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetMaintenanceConfigurationName() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceConfigurationName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetMaxSizeGb() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxSizeGb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetMinCapacity() {
	_jsii_.InvokeVoid(
		m,
		"resetMinCapacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetReadReplicaCount() {
	_jsii_.InvokeVoid(
		m,
		"resetReadReplicaCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetReadScale() {
	_jsii_.InvokeVoid(
		m,
		"resetReadScale",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetRecoverDatabaseId() {
	_jsii_.InvokeVoid(
		m,
		"resetRecoverDatabaseId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetRecoveryPointId() {
	_jsii_.InvokeVoid(
		m,
		"resetRecoveryPointId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetRestoreDroppedDatabaseId() {
	_jsii_.InvokeVoid(
		m,
		"resetRestoreDroppedDatabaseId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetRestoreLongTermRetentionBackupId() {
	_jsii_.InvokeVoid(
		m,
		"resetRestoreLongTermRetentionBackupId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetRestorePointInTime() {
	_jsii_.InvokeVoid(
		m,
		"resetRestorePointInTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetSampleName() {
	_jsii_.InvokeVoid(
		m,
		"resetSampleName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetSecondaryType() {
	_jsii_.InvokeVoid(
		m,
		"resetSecondaryType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetShortTermRetentionPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetShortTermRetentionPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetSkuName() {
	_jsii_.InvokeVoid(
		m,
		"resetSkuName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetStorageAccountType() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetThreatDetectionPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetThreatDetectionPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetTransparentDataEncryptionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetTransparentDataEncryptionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetTransparentDataEncryptionKeyAutomaticRotationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetTransparentDataEncryptionKeyAutomaticRotationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetTransparentDataEncryptionKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetTransparentDataEncryptionKeyVaultKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) ResetZoneRedundant() {
	_jsii_.InvokeVoid(
		m,
		"resetZoneRedundant",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


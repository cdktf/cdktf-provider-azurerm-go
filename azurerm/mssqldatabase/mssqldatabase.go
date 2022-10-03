package mssqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/mssqldatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database azurerm_mssql_database}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	IdInput() *string
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
	RestoreDroppedDatabaseId() *string
	SetRestoreDroppedDatabaseId(val *string)
	RestoreDroppedDatabaseIdInput() *string
	RestorePointInTime() *string
	SetRestorePointInTime(val *string)
	RestorePointInTimeInput() *string
	SampleName() *string
	SetSampleName(val *string)
	SampleNameInput() *string
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
	ZoneRedundant() interface{}
	SetZoneRedundant(val interface{})
	ZoneRedundantInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutLongTermRetentionPolicy(value *MssqlDatabaseLongTermRetentionPolicy)
	PutShortTermRetentionPolicy(value *MssqlDatabaseShortTermRetentionPolicy)
	PutThreatDetectionPolicy(value *MssqlDatabaseThreatDetectionPolicy)
	PutTimeouts(value *MssqlDatabaseTimeouts)
	ResetAutoPauseDelayInMinutes()
	ResetCollation()
	ResetCreateMode()
	ResetCreationSourceDatabaseId()
	ResetElasticPoolId()
	ResetGeoBackupEnabled()
	ResetId()
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
	ResetRestoreDroppedDatabaseId()
	ResetRestorePointInTime()
	ResetSampleName()
	ResetShortTermRetentionPolicy()
	ResetSkuName()
	ResetStorageAccountType()
	ResetTags()
	ResetThreatDetectionPolicy()
	ResetTimeouts()
	ResetTransparentDataEncryptionEnabled()
	ResetZoneRedundant()
	SynthesizeAttributes() *map[string]interface{}
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

func (j *jsiiProxy_MssqlDatabase) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_MssqlDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database azurerm_mssql_database} Resource.
func NewMssqlDatabase(scope constructs.Construct, id *string, config *MssqlDatabaseConfig) MssqlDatabase {
	_init_.Initialize()

	j := jsiiProxy_MssqlDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database azurerm_mssql_database} Resource.
func NewMssqlDatabase_Override(m MssqlDatabase, scope constructs.Construct, id *string, config *MssqlDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetAutoPauseDelayInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"autoPauseDelayInMinutes",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetCollation(val *string) {
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetCreateMode(val *string) {
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetCreationSourceDatabaseId(val *string) {
	_jsii_.Set(
		j,
		"creationSourceDatabaseId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetElasticPoolId(val *string) {
	_jsii_.Set(
		j,
		"elasticPoolId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetGeoBackupEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"geoBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetLedgerEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"ledgerEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetLicenseType(val *string) {
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetMaintenanceConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"maintenanceConfigurationName",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetMaxSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"maxSizeGb",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetMinCapacity(val *float64) {
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetReadReplicaCount(val *float64) {
	_jsii_.Set(
		j,
		"readReplicaCount",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetReadScale(val interface{}) {
	_jsii_.Set(
		j,
		"readScale",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetRecoverDatabaseId(val *string) {
	_jsii_.Set(
		j,
		"recoverDatabaseId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetRestoreDroppedDatabaseId(val *string) {
	_jsii_.Set(
		j,
		"restoreDroppedDatabaseId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetRestorePointInTime(val *string) {
	_jsii_.Set(
		j,
		"restorePointInTime",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetSampleName(val *string) {
	_jsii_.Set(
		j,
		"sampleName",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetServerId(val *string) {
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetSkuName(val *string) {
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetStorageAccountType(val *string) {
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetTransparentDataEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"transparentDataEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabase) SetZoneRedundant(val interface{}) {
	_jsii_.Set(
		j,
		"zoneRedundant",
		val,
	)
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabase",
		"isConstruct",
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

func (m *jsiiProxy_MssqlDatabase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MssqlDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutLongTermRetentionPolicy(value *MssqlDatabaseLongTermRetentionPolicy) {
	_jsii_.InvokeVoid(
		m,
		"putLongTermRetentionPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutShortTermRetentionPolicy(value *MssqlDatabaseShortTermRetentionPolicy) {
	_jsii_.InvokeVoid(
		m,
		"putShortTermRetentionPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutThreatDetectionPolicy(value *MssqlDatabaseThreatDetectionPolicy) {
	_jsii_.InvokeVoid(
		m,
		"putThreatDetectionPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlDatabase) PutTimeouts(value *MssqlDatabaseTimeouts) {
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

func (m *jsiiProxy_MssqlDatabase) ResetRestoreDroppedDatabaseId() {
	_jsii_.InvokeVoid(
		m,
		"resetRestoreDroppedDatabaseId",
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

type MssqlDatabaseConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#name MssqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#server_id MssqlDatabase#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#auto_pause_delay_in_minutes MssqlDatabase#auto_pause_delay_in_minutes}.
	AutoPauseDelayInMinutes *float64 `field:"optional" json:"autoPauseDelayInMinutes" yaml:"autoPauseDelayInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#collation MssqlDatabase#collation}.
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#create_mode MssqlDatabase#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#creation_source_database_id MssqlDatabase#creation_source_database_id}.
	CreationSourceDatabaseId *string `field:"optional" json:"creationSourceDatabaseId" yaml:"creationSourceDatabaseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#elastic_pool_id MssqlDatabase#elastic_pool_id}.
	ElasticPoolId *string `field:"optional" json:"elasticPoolId" yaml:"elasticPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#geo_backup_enabled MssqlDatabase#geo_backup_enabled}.
	GeoBackupEnabled interface{} `field:"optional" json:"geoBackupEnabled" yaml:"geoBackupEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#id MssqlDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#ledger_enabled MssqlDatabase#ledger_enabled}.
	LedgerEnabled interface{} `field:"optional" json:"ledgerEnabled" yaml:"ledgerEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#license_type MssqlDatabase#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// long_term_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#long_term_retention_policy MssqlDatabase#long_term_retention_policy}
	LongTermRetentionPolicy *MssqlDatabaseLongTermRetentionPolicy `field:"optional" json:"longTermRetentionPolicy" yaml:"longTermRetentionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#maintenance_configuration_name MssqlDatabase#maintenance_configuration_name}.
	MaintenanceConfigurationName *string `field:"optional" json:"maintenanceConfigurationName" yaml:"maintenanceConfigurationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#max_size_gb MssqlDatabase#max_size_gb}.
	MaxSizeGb *float64 `field:"optional" json:"maxSizeGb" yaml:"maxSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#min_capacity MssqlDatabase#min_capacity}.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#read_replica_count MssqlDatabase#read_replica_count}.
	ReadReplicaCount *float64 `field:"optional" json:"readReplicaCount" yaml:"readReplicaCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#read_scale MssqlDatabase#read_scale}.
	ReadScale interface{} `field:"optional" json:"readScale" yaml:"readScale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#recover_database_id MssqlDatabase#recover_database_id}.
	RecoverDatabaseId *string `field:"optional" json:"recoverDatabaseId" yaml:"recoverDatabaseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#restore_dropped_database_id MssqlDatabase#restore_dropped_database_id}.
	RestoreDroppedDatabaseId *string `field:"optional" json:"restoreDroppedDatabaseId" yaml:"restoreDroppedDatabaseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#restore_point_in_time MssqlDatabase#restore_point_in_time}.
	RestorePointInTime *string `field:"optional" json:"restorePointInTime" yaml:"restorePointInTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#sample_name MssqlDatabase#sample_name}.
	SampleName *string `field:"optional" json:"sampleName" yaml:"sampleName"`
	// short_term_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#short_term_retention_policy MssqlDatabase#short_term_retention_policy}
	ShortTermRetentionPolicy *MssqlDatabaseShortTermRetentionPolicy `field:"optional" json:"shortTermRetentionPolicy" yaml:"shortTermRetentionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#sku_name MssqlDatabase#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#storage_account_type MssqlDatabase#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#tags MssqlDatabase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// threat_detection_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#threat_detection_policy MssqlDatabase#threat_detection_policy}
	ThreatDetectionPolicy *MssqlDatabaseThreatDetectionPolicy `field:"optional" json:"threatDetectionPolicy" yaml:"threatDetectionPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#timeouts MssqlDatabase#timeouts}
	Timeouts *MssqlDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#transparent_data_encryption_enabled MssqlDatabase#transparent_data_encryption_enabled}.
	TransparentDataEncryptionEnabled interface{} `field:"optional" json:"transparentDataEncryptionEnabled" yaml:"transparentDataEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#zone_redundant MssqlDatabase#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

type MssqlDatabaseLongTermRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#monthly_retention MssqlDatabase#monthly_retention}.
	MonthlyRetention *string `field:"optional" json:"monthlyRetention" yaml:"monthlyRetention"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#weekly_retention MssqlDatabase#weekly_retention}.
	WeeklyRetention *string `field:"optional" json:"weeklyRetention" yaml:"weeklyRetention"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#week_of_year MssqlDatabase#week_of_year}.
	WeekOfYear *float64 `field:"optional" json:"weekOfYear" yaml:"weekOfYear"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#yearly_retention MssqlDatabase#yearly_retention}.
	YearlyRetention *string `field:"optional" json:"yearlyRetention" yaml:"yearlyRetention"`
}

type MssqlDatabaseLongTermRetentionPolicyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlDatabaseLongTermRetentionPolicy
	SetInternalValue(val *MssqlDatabaseLongTermRetentionPolicy)
	MonthlyRetention() *string
	SetMonthlyRetention(val *string)
	MonthlyRetentionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklyRetention() *string
	SetWeeklyRetention(val *string)
	WeeklyRetentionInput() *string
	WeekOfYear() *float64
	SetWeekOfYear(val *float64)
	WeekOfYearInput() *float64
	YearlyRetention() *string
	SetYearlyRetention(val *string)
	YearlyRetentionInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMonthlyRetention()
	ResetWeeklyRetention()
	ResetWeekOfYear()
	ResetYearlyRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlDatabaseLongTermRetentionPolicyOutputReference
type jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) InternalValue() *MssqlDatabaseLongTermRetentionPolicy {
	var returns *MssqlDatabaseLongTermRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) MonthlyRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) MonthlyRetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeeklyRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeeklyRetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeekOfYear() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) WeekOfYearInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) YearlyRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yearlyRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) YearlyRetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yearlyRetentionInput",
		&returns,
	)
	return returns
}


func NewMssqlDatabaseLongTermRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlDatabaseLongTermRetentionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseLongTermRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlDatabaseLongTermRetentionPolicyOutputReference_Override(m MssqlDatabaseLongTermRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseLongTermRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetInternalValue(val *MssqlDatabaseLongTermRetentionPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetMonthlyRetention(val *string) {
	_jsii_.Set(
		j,
		"monthlyRetention",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetWeeklyRetention(val *string) {
	_jsii_.Set(
		j,
		"weeklyRetention",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetWeekOfYear(val *float64) {
	_jsii_.Set(
		j,
		"weekOfYear",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) SetYearlyRetention(val *string) {
	_jsii_.Set(
		j,
		"yearlyRetention",
		val,
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetMonthlyRetention() {
	_jsii_.InvokeVoid(
		m,
		"resetMonthlyRetention",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetWeeklyRetention() {
	_jsii_.InvokeVoid(
		m,
		"resetWeeklyRetention",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetWeekOfYear() {
	_jsii_.InvokeVoid(
		m,
		"resetWeekOfYear",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ResetYearlyRetention() {
	_jsii_.InvokeVoid(
		m,
		"resetYearlyRetention",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseLongTermRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlDatabaseShortTermRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#retention_days MssqlDatabase#retention_days}.
	RetentionDays *float64 `field:"required" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#backup_interval_in_hours MssqlDatabase#backup_interval_in_hours}.
	BackupIntervalInHours *float64 `field:"optional" json:"backupIntervalInHours" yaml:"backupIntervalInHours"`
}

type MssqlDatabaseShortTermRetentionPolicyOutputReference interface {
	cdktf.ComplexObject
	BackupIntervalInHours() *float64
	SetBackupIntervalInHours(val *float64)
	BackupIntervalInHoursInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlDatabaseShortTermRetentionPolicy
	SetInternalValue(val *MssqlDatabaseShortTermRetentionPolicy)
	RetentionDays() *float64
	SetRetentionDays(val *float64)
	RetentionDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBackupIntervalInHours()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlDatabaseShortTermRetentionPolicyOutputReference
type jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) BackupIntervalInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupIntervalInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) BackupIntervalInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupIntervalInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) InternalValue() *MssqlDatabaseShortTermRetentionPolicy {
	var returns *MssqlDatabaseShortTermRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) RetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlDatabaseShortTermRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlDatabaseShortTermRetentionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseShortTermRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlDatabaseShortTermRetentionPolicyOutputReference_Override(m MssqlDatabaseShortTermRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseShortTermRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) SetBackupIntervalInHours(val *float64) {
	_jsii_.Set(
		j,
		"backupIntervalInHours",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) SetInternalValue(val *MssqlDatabaseShortTermRetentionPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) SetRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) ResetBackupIntervalInHours() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupIntervalInHours",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseShortTermRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlDatabaseThreatDetectionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#disabled_alerts MssqlDatabase#disabled_alerts}.
	DisabledAlerts *[]*string `field:"optional" json:"disabledAlerts" yaml:"disabledAlerts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#email_account_admins MssqlDatabase#email_account_admins}.
	EmailAccountAdmins *string `field:"optional" json:"emailAccountAdmins" yaml:"emailAccountAdmins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#email_addresses MssqlDatabase#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#retention_days MssqlDatabase#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#state MssqlDatabase#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#storage_account_access_key MssqlDatabase#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#storage_endpoint MssqlDatabase#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
}

type MssqlDatabaseThreatDetectionPolicyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisabledAlerts() *[]*string
	SetDisabledAlerts(val *[]*string)
	DisabledAlertsInput() *[]*string
	EmailAccountAdmins() *string
	SetEmailAccountAdmins(val *string)
	EmailAccountAdminsInput() *string
	EmailAddresses() *[]*string
	SetEmailAddresses(val *[]*string)
	EmailAddressesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlDatabaseThreatDetectionPolicy
	SetInternalValue(val *MssqlDatabaseThreatDetectionPolicy)
	RetentionDays() *float64
	SetRetentionDays(val *float64)
	RetentionDaysInput() *float64
	State() *string
	SetState(val *string)
	StateInput() *string
	StorageAccountAccessKey() *string
	SetStorageAccountAccessKey(val *string)
	StorageAccountAccessKeyInput() *string
	StorageEndpoint() *string
	SetStorageEndpoint(val *string)
	StorageEndpointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDisabledAlerts()
	ResetEmailAccountAdmins()
	ResetEmailAddresses()
	ResetRetentionDays()
	ResetState()
	ResetStorageAccountAccessKey()
	ResetStorageEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlDatabaseThreatDetectionPolicyOutputReference
type jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) DisabledAlerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) DisabledAlertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) EmailAccountAdmins() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAccountAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) EmailAccountAdminsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAccountAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) EmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) EmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) InternalValue() *MssqlDatabaseThreatDetectionPolicy {
	var returns *MssqlDatabaseThreatDetectionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) RetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) StorageEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) StorageEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlDatabaseThreatDetectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlDatabaseThreatDetectionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlDatabaseThreatDetectionPolicyOutputReference_Override(m MssqlDatabaseThreatDetectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetDisabledAlerts(val *[]*string) {
	_jsii_.Set(
		j,
		"disabledAlerts",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetEmailAccountAdmins(val *string) {
	_jsii_.Set(
		j,
		"emailAccountAdmins",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetEmailAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"emailAddresses",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetInternalValue(val *MssqlDatabaseThreatDetectionPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetStorageAccountAccessKey(val *string) {
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetStorageEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageEndpoint",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ResetDisabledAlerts() {
	_jsii_.InvokeVoid(
		m,
		"resetDisabledAlerts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ResetEmailAccountAdmins() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailAccountAdmins",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ResetEmailAddresses() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailAddresses",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ResetRetentionDays() {
	_jsii_.InvokeVoid(
		m,
		"resetRetentionDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		m,
		"resetState",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ResetStorageEndpoint() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageEndpoint",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseThreatDetectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#create MssqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#delete MssqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#read MssqlDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#update MssqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type MssqlDatabaseTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlDatabaseTimeoutsOutputReference
type jsiiProxy_MssqlDatabaseTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewMssqlDatabaseTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlDatabaseTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlDatabaseTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlDatabaseTimeoutsOutputReference_Override(m MssqlDatabaseTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlDatabase.MssqlDatabaseTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


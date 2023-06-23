package sqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/sqldatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/sql_database azurerm_sql_database}.
type SqlDatabase interface {
	cdktf.TerraformResource
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
	CreationDate() *string
	DefaultSecondaryLocation() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	ElasticPoolName() *string
	SetElasticPoolName(val *string)
	ElasticPoolNameInput() *string
	Encryption() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Import() SqlDatabaseImportOutputReference
	ImportInput() *SqlDatabaseImport
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxSizeBytes() *string
	SetMaxSizeBytes(val *string)
	MaxSizeBytesInput() *string
	MaxSizeGb() *string
	SetMaxSizeGb(val *string)
	MaxSizeGbInput() *string
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
	ReadScale() interface{}
	SetReadScale(val interface{})
	ReadScaleInput() interface{}
	RequestedServiceObjectiveId() *string
	SetRequestedServiceObjectiveId(val *string)
	RequestedServiceObjectiveIdInput() *string
	RequestedServiceObjectiveName() *string
	SetRequestedServiceObjectiveName(val *string)
	RequestedServiceObjectiveNameInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RestorePointInTime() *string
	SetRestorePointInTime(val *string)
	RestorePointInTimeInput() *string
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
	SourceDatabaseDeletionDate() *string
	SetSourceDatabaseDeletionDate(val *string)
	SourceDatabaseDeletionDateInput() *string
	SourceDatabaseId() *string
	SetSourceDatabaseId(val *string)
	SourceDatabaseIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreatDetectionPolicy() SqlDatabaseThreatDetectionPolicyOutputReference
	ThreatDetectionPolicyInput() *SqlDatabaseThreatDetectionPolicy
	Timeouts() SqlDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutImport(value *SqlDatabaseImport)
	PutThreatDetectionPolicy(value *SqlDatabaseThreatDetectionPolicy)
	PutTimeouts(value *SqlDatabaseTimeouts)
	ResetCollation()
	ResetCreateMode()
	ResetEdition()
	ResetElasticPoolName()
	ResetId()
	ResetImport()
	ResetMaxSizeBytes()
	ResetMaxSizeGb()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadScale()
	ResetRequestedServiceObjectiveId()
	ResetRequestedServiceObjectiveName()
	ResetRestorePointInTime()
	ResetSourceDatabaseDeletionDate()
	ResetSourceDatabaseId()
	ResetTags()
	ResetThreatDetectionPolicy()
	ResetTimeouts()
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

// The jsii proxy struct for SqlDatabase
type jsiiProxy_SqlDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SqlDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) CreateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) CreateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) DefaultSecondaryLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ElasticPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ElasticPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Encryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Import() SqlDatabaseImportOutputReference {
	var returns SqlDatabaseImportOutputReference
	_jsii_.Get(
		j,
		"import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ImportInput() *SqlDatabaseImport {
	var returns *SqlDatabaseImport
	_jsii_.Get(
		j,
		"importInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) MaxSizeBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) MaxSizeBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) MaxSizeGb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) MaxSizeGbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ReadScale() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ReadScaleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) RequestedServiceObjectiveId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedServiceObjectiveId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) RequestedServiceObjectiveIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedServiceObjectiveIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) RequestedServiceObjectiveName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedServiceObjectiveName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) RequestedServiceObjectiveNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedServiceObjectiveNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) RestorePointInTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restorePointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) RestorePointInTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restorePointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) SourceDatabaseDeletionDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDatabaseDeletionDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) SourceDatabaseDeletionDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDatabaseDeletionDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) SourceDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) SourceDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ThreatDetectionPolicy() SqlDatabaseThreatDetectionPolicyOutputReference {
	var returns SqlDatabaseThreatDetectionPolicyOutputReference
	_jsii_.Get(
		j,
		"threatDetectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ThreatDetectionPolicyInput() *SqlDatabaseThreatDetectionPolicy {
	var returns *SqlDatabaseThreatDetectionPolicy
	_jsii_.Get(
		j,
		"threatDetectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) Timeouts() SqlDatabaseTimeoutsOutputReference {
	var returns SqlDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ZoneRedundant() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabase) ZoneRedundantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundantInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/sql_database azurerm_sql_database} Resource.
func NewSqlDatabase(scope constructs.Construct, id *string, config *SqlDatabaseConfig) SqlDatabase {
	_init_.Initialize()

	if err := validateNewSqlDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/sql_database azurerm_sql_database} Resource.
func NewSqlDatabase_Override(s SqlDatabase, scope constructs.Construct, id *string, config *SqlDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SqlDatabase)SetCollation(val *string) {
	if err := j.validateSetCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetCreateMode(val *string) {
	if err := j.validateSetCreateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetElasticPoolName(val *string) {
	if err := j.validateSetElasticPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticPoolName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetMaxSizeBytes(val *string) {
	if err := j.validateSetMaxSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSizeBytes",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetMaxSizeGb(val *string) {
	if err := j.validateSetMaxSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSizeGb",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetReadScale(val interface{}) {
	if err := j.validateSetReadScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readScale",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetRequestedServiceObjectiveId(val *string) {
	if err := j.validateSetRequestedServiceObjectiveIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedServiceObjectiveId",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetRequestedServiceObjectiveName(val *string) {
	if err := j.validateSetRequestedServiceObjectiveNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedServiceObjectiveName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetRestorePointInTime(val *string) {
	if err := j.validateSetRestorePointInTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restorePointInTime",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetSourceDatabaseDeletionDate(val *string) {
	if err := j.validateSetSourceDatabaseDeletionDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDatabaseDeletionDate",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetSourceDatabaseId(val *string) {
	if err := j.validateSetSourceDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDatabaseId",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase)SetZoneRedundant(val interface{}) {
	if err := j.validateSetZoneRedundantParameters(val); err != nil {
		panic(err)
	}
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
func SqlDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SqlDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SqlDatabase) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SqlDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SqlDatabase) PutImport(value *SqlDatabaseImport) {
	if err := s.validatePutImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putImport",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabase) PutThreatDetectionPolicy(value *SqlDatabaseThreatDetectionPolicy) {
	if err := s.validatePutThreatDetectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThreatDetectionPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabase) PutTimeouts(value *SqlDatabaseTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabase) ResetCollation() {
	_jsii_.InvokeVoid(
		s,
		"resetCollation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetCreateMode() {
	_jsii_.InvokeVoid(
		s,
		"resetCreateMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetEdition() {
	_jsii_.InvokeVoid(
		s,
		"resetEdition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetElasticPoolName() {
	_jsii_.InvokeVoid(
		s,
		"resetElasticPoolName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetImport() {
	_jsii_.InvokeVoid(
		s,
		"resetImport",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetMaxSizeBytes() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxSizeBytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetMaxSizeGb() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxSizeGb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetReadScale() {
	_jsii_.InvokeVoid(
		s,
		"resetReadScale",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetRequestedServiceObjectiveId() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestedServiceObjectiveId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetRequestedServiceObjectiveName() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestedServiceObjectiveName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetRestorePointInTime() {
	_jsii_.InvokeVoid(
		s,
		"resetRestorePointInTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetSourceDatabaseDeletionDate() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceDatabaseDeletionDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetSourceDatabaseId() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceDatabaseId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetThreatDetectionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatDetectionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) ResetZoneRedundant() {
	_jsii_.InvokeVoid(
		s,
		"resetZoneRedundant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


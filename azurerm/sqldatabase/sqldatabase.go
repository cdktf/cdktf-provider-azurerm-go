package sqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/sqldatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database azurerm_sql_database}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_SqlDatabase) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database azurerm_sql_database} Resource.
func NewSqlDatabase(scope constructs.Construct, id *string, config *SqlDatabaseConfig) SqlDatabase {
	_init_.Initialize()

	j := jsiiProxy_SqlDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database azurerm_sql_database} Resource.
func NewSqlDatabase_Override(s SqlDatabase, scope constructs.Construct, id *string, config *SqlDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SqlDatabase) SetCollation(val *string) {
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetCreateMode(val *string) {
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetEdition(val *string) {
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetElasticPoolName(val *string) {
	_jsii_.Set(
		j,
		"elasticPoolName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetMaxSizeBytes(val *string) {
	_jsii_.Set(
		j,
		"maxSizeBytes",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetMaxSizeGb(val *string) {
	_jsii_.Set(
		j,
		"maxSizeGb",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetReadScale(val interface{}) {
	_jsii_.Set(
		j,
		"readScale",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetRequestedServiceObjectiveId(val *string) {
	_jsii_.Set(
		j,
		"requestedServiceObjectiveId",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetRequestedServiceObjectiveName(val *string) {
	_jsii_.Set(
		j,
		"requestedServiceObjectiveName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetRestorePointInTime(val *string) {
	_jsii_.Set(
		j,
		"restorePointInTime",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetServerName(val *string) {
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetSourceDatabaseDeletionDate(val *string) {
	_jsii_.Set(
		j,
		"sourceDatabaseDeletionDate",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetSourceDatabaseId(val *string) {
	_jsii_.Set(
		j,
		"sourceDatabaseId",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SqlDatabase) SetZoneRedundant(val interface{}) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabase",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SqlDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SqlDatabase) PutImport(value *SqlDatabaseImport) {
	_jsii_.InvokeVoid(
		s,
		"putImport",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabase) PutThreatDetectionPolicy(value *SqlDatabaseThreatDetectionPolicy) {
	_jsii_.InvokeVoid(
		s,
		"putThreatDetectionPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabase) PutTimeouts(value *SqlDatabaseTimeouts) {
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

type SqlDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#location SqlDatabase#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#name SqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#resource_group_name SqlDatabase#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#server_name SqlDatabase#server_name}.
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#collation SqlDatabase#collation}.
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#create_mode SqlDatabase#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#edition SqlDatabase#edition}.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#elastic_pool_name SqlDatabase#elastic_pool_name}.
	ElasticPoolName *string `field:"optional" json:"elasticPoolName" yaml:"elasticPoolName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#id SqlDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// import block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#import SqlDatabase#import}
	Import *SqlDatabaseImport `field:"optional" json:"import" yaml:"import"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#max_size_bytes SqlDatabase#max_size_bytes}.
	MaxSizeBytes *string `field:"optional" json:"maxSizeBytes" yaml:"maxSizeBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#max_size_gb SqlDatabase#max_size_gb}.
	MaxSizeGb *string `field:"optional" json:"maxSizeGb" yaml:"maxSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#read_scale SqlDatabase#read_scale}.
	ReadScale interface{} `field:"optional" json:"readScale" yaml:"readScale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#requested_service_objective_id SqlDatabase#requested_service_objective_id}.
	RequestedServiceObjectiveId *string `field:"optional" json:"requestedServiceObjectiveId" yaml:"requestedServiceObjectiveId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#requested_service_objective_name SqlDatabase#requested_service_objective_name}.
	RequestedServiceObjectiveName *string `field:"optional" json:"requestedServiceObjectiveName" yaml:"requestedServiceObjectiveName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#restore_point_in_time SqlDatabase#restore_point_in_time}.
	RestorePointInTime *string `field:"optional" json:"restorePointInTime" yaml:"restorePointInTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#source_database_deletion_date SqlDatabase#source_database_deletion_date}.
	SourceDatabaseDeletionDate *string `field:"optional" json:"sourceDatabaseDeletionDate" yaml:"sourceDatabaseDeletionDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#source_database_id SqlDatabase#source_database_id}.
	SourceDatabaseId *string `field:"optional" json:"sourceDatabaseId" yaml:"sourceDatabaseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#tags SqlDatabase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// threat_detection_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#threat_detection_policy SqlDatabase#threat_detection_policy}
	ThreatDetectionPolicy *SqlDatabaseThreatDetectionPolicy `field:"optional" json:"threatDetectionPolicy" yaml:"threatDetectionPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#timeouts SqlDatabase#timeouts}
	Timeouts *SqlDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#zone_redundant SqlDatabase#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

type SqlDatabaseImport struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#administrator_login SqlDatabase#administrator_login}.
	AdministratorLogin *string `field:"required" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#administrator_login_password SqlDatabase#administrator_login_password}.
	AdministratorLoginPassword *string `field:"required" json:"administratorLoginPassword" yaml:"administratorLoginPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#authentication_type SqlDatabase#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#storage_key SqlDatabase#storage_key}.
	StorageKey *string `field:"required" json:"storageKey" yaml:"storageKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#storage_key_type SqlDatabase#storage_key_type}.
	StorageKeyType *string `field:"required" json:"storageKeyType" yaml:"storageKeyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#storage_uri SqlDatabase#storage_uri}.
	StorageUri *string `field:"required" json:"storageUri" yaml:"storageUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#operation_mode SqlDatabase#operation_mode}.
	OperationMode *string `field:"optional" json:"operationMode" yaml:"operationMode"`
}

type SqlDatabaseImportOutputReference interface {
	cdktf.ComplexObject
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorLoginPassword() *string
	SetAdministratorLoginPassword(val *string)
	AdministratorLoginPasswordInput() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
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
	InternalValue() *SqlDatabaseImport
	SetInternalValue(val *SqlDatabaseImport)
	OperationMode() *string
	SetOperationMode(val *string)
	OperationModeInput() *string
	StorageKey() *string
	SetStorageKey(val *string)
	StorageKeyInput() *string
	StorageKeyType() *string
	SetStorageKeyType(val *string)
	StorageKeyTypeInput() *string
	StorageUri() *string
	SetStorageUri(val *string)
	StorageUriInput() *string
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
	ResetOperationMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SqlDatabaseImportOutputReference
type jsiiProxy_SqlDatabaseImportOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) AdministratorLoginPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) AdministratorLoginPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) InternalValue() *SqlDatabaseImport {
	var returns *SqlDatabaseImport
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) OperationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) OperationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) StorageKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) StorageKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) StorageKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) StorageKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) StorageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) StorageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSqlDatabaseImportOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SqlDatabaseImportOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SqlDatabaseImportOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseImportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlDatabaseImportOutputReference_Override(s SqlDatabaseImportOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseImportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetAdministratorLogin(val *string) {
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetAdministratorLoginPassword(val *string) {
	_jsii_.Set(
		j,
		"administratorLoginPassword",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetAuthenticationType(val *string) {
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetInternalValue(val *SqlDatabaseImport) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetOperationMode(val *string) {
	_jsii_.Set(
		j,
		"operationMode",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetStorageKey(val *string) {
	_jsii_.Set(
		j,
		"storageKey",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetStorageKeyType(val *string) {
	_jsii_.Set(
		j,
		"storageKeyType",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetStorageUri(val *string) {
	_jsii_.Set(
		j,
		"storageUri",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseImportOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) ResetOperationMode() {
	_jsii_.InvokeVoid(
		s,
		"resetOperationMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseImportOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SqlDatabaseThreatDetectionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#disabled_alerts SqlDatabase#disabled_alerts}.
	DisabledAlerts *[]*string `field:"optional" json:"disabledAlerts" yaml:"disabledAlerts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#email_account_admins SqlDatabase#email_account_admins}.
	EmailAccountAdmins *string `field:"optional" json:"emailAccountAdmins" yaml:"emailAccountAdmins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#email_addresses SqlDatabase#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#retention_days SqlDatabase#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#state SqlDatabase#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#storage_account_access_key SqlDatabase#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#storage_endpoint SqlDatabase#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
}

type SqlDatabaseThreatDetectionPolicyOutputReference interface {
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
	InternalValue() *SqlDatabaseThreatDetectionPolicy
	SetInternalValue(val *SqlDatabaseThreatDetectionPolicy)
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

// The jsii proxy struct for SqlDatabaseThreatDetectionPolicyOutputReference
type jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) DisabledAlerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) DisabledAlertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAccountAdmins() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAccountAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAccountAdminsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAccountAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) InternalValue() *SqlDatabaseThreatDetectionPolicy {
	var returns *SqlDatabaseThreatDetectionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) RetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSqlDatabaseThreatDetectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SqlDatabaseThreatDetectionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlDatabaseThreatDetectionPolicyOutputReference_Override(s SqlDatabaseThreatDetectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetDisabledAlerts(val *[]*string) {
	_jsii_.Set(
		j,
		"disabledAlerts",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetEmailAccountAdmins(val *string) {
	_jsii_.Set(
		j,
		"emailAccountAdmins",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetEmailAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"emailAddresses",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetInternalValue(val *SqlDatabaseThreatDetectionPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetStorageAccountAccessKey(val *string) {
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetStorageEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageEndpoint",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetDisabledAlerts() {
	_jsii_.InvokeVoid(
		s,
		"resetDisabledAlerts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetEmailAccountAdmins() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailAccountAdmins",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetEmailAddresses() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailAddresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetRetentionDays() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		s,
		"resetState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetStorageEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#create SqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#delete SqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#read SqlDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_database#update SqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type SqlDatabaseTimeoutsOutputReference interface {
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

// The jsii proxy struct for SqlDatabaseTimeoutsOutputReference
type jsiiProxy_SqlDatabaseTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewSqlDatabaseTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SqlDatabaseTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SqlDatabaseTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlDatabaseTimeoutsOutputReference_Override(s SqlDatabaseTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package postgresqlflexibleserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/postgresqlflexibleserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server azurerm_postgresql_flexible_server}.
type PostgresqlFlexibleServer interface {
	cdktf.TerraformResource
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorPassword() *string
	SetAdministratorPassword(val *string)
	AdministratorPasswordInput() *string
	BackupRetentionDays() *float64
	SetBackupRetentionDays(val *float64)
	BackupRetentionDaysInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DelegatedSubnetId() *string
	SetDelegatedSubnetId(val *string)
	DelegatedSubnetIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoRedundantBackupEnabled() interface{}
	SetGeoRedundantBackupEnabled(val interface{})
	GeoRedundantBackupEnabledInput() interface{}
	HighAvailability() PostgresqlFlexibleServerHighAvailabilityOutputReference
	HighAvailabilityInput() *PostgresqlFlexibleServerHighAvailability
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() PostgresqlFlexibleServerMaintenanceWindowOutputReference
	MaintenanceWindowInput() *PostgresqlFlexibleServerMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PointInTimeRestoreTimeInUtc() *string
	SetPointInTimeRestoreTimeInUtc(val *string)
	PointInTimeRestoreTimeInUtcInput() *string
	PrivateDnsZoneId() *string
	SetPrivateDnsZoneId(val *string)
	PrivateDnsZoneIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccessEnabled() cdktf.IResolvable
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	SourceServerId() *string
	SetSourceServerId(val *string)
	SourceServerIdInput() *string
	StorageMb() *float64
	SetStorageMb(val *float64)
	StorageMbInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PostgresqlFlexibleServerTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutHighAvailability(value *PostgresqlFlexibleServerHighAvailability)
	PutMaintenanceWindow(value *PostgresqlFlexibleServerMaintenanceWindow)
	PutTimeouts(value *PostgresqlFlexibleServerTimeouts)
	ResetAdministratorLogin()
	ResetAdministratorPassword()
	ResetBackupRetentionDays()
	ResetCreateMode()
	ResetDelegatedSubnetId()
	ResetGeoRedundantBackupEnabled()
	ResetHighAvailability()
	ResetId()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRestoreTimeInUtc()
	ResetPrivateDnsZoneId()
	ResetSkuName()
	ResetSourceServerId()
	ResetStorageMb()
	ResetTags()
	ResetTimeouts()
	ResetVersion()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PostgresqlFlexibleServer
type jsiiProxy_PostgresqlFlexibleServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) BackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) BackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CreateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CreateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) DelegatedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) DelegatedSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) GeoRedundantBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) GeoRedundantBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) HighAvailability() PostgresqlFlexibleServerHighAvailabilityOutputReference {
	var returns PostgresqlFlexibleServerHighAvailabilityOutputReference
	_jsii_.Get(
		j,
		"highAvailability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) HighAvailabilityInput() *PostgresqlFlexibleServerHighAvailability {
	var returns *PostgresqlFlexibleServerHighAvailability
	_jsii_.Get(
		j,
		"highAvailabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) MaintenanceWindow() PostgresqlFlexibleServerMaintenanceWindowOutputReference {
	var returns PostgresqlFlexibleServerMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) MaintenanceWindowInput() *PostgresqlFlexibleServerMaintenanceWindow {
	var returns *PostgresqlFlexibleServerMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PointInTimeRestoreTimeInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRestoreTimeInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PointInTimeRestoreTimeInUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRestoreTimeInUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PrivateDnsZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PrivateDnsZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PublicNetworkAccessEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SourceServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SourceServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) StorageMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) StorageMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Timeouts() PostgresqlFlexibleServerTimeoutsOutputReference {
	var returns PostgresqlFlexibleServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server azurerm_postgresql_flexible_server} Resource.
func NewPostgresqlFlexibleServer(scope constructs.Construct, id *string, config *PostgresqlFlexibleServerConfig) PostgresqlFlexibleServer {
	_init_.Initialize()

	j := jsiiProxy_PostgresqlFlexibleServer{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server azurerm_postgresql_flexible_server} Resource.
func NewPostgresqlFlexibleServer_Override(p PostgresqlFlexibleServer, scope constructs.Construct, id *string, config *PostgresqlFlexibleServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetAdministratorLogin(val *string) {
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetAdministratorPassword(val *string) {
	_jsii_.Set(
		j,
		"administratorPassword",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetBackupRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetCreateMode(val *string) {
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetDelegatedSubnetId(val *string) {
	_jsii_.Set(
		j,
		"delegatedSubnetId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetGeoRedundantBackupEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"geoRedundantBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetPointInTimeRestoreTimeInUtc(val *string) {
	_jsii_.Set(
		j,
		"pointInTimeRestoreTimeInUtc",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetPrivateDnsZoneId(val *string) {
	_jsii_.Set(
		j,
		"privateDnsZoneId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetSkuName(val *string) {
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetSourceServerId(val *string) {
	_jsii_.Set(
		j,
		"sourceServerId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetStorageMb(val *float64) {
	_jsii_.Set(
		j,
		"storageMb",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SetZone(val *string) {
	_jsii_.Set(
		j,
		"zone",
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
func PostgresqlFlexibleServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PostgresqlFlexibleServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutHighAvailability(value *PostgresqlFlexibleServerHighAvailability) {
	_jsii_.InvokeVoid(
		p,
		"putHighAvailability",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutMaintenanceWindow(value *PostgresqlFlexibleServerMaintenanceWindow) {
	_jsii_.InvokeVoid(
		p,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutTimeouts(value *PostgresqlFlexibleServerTimeouts) {
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAdministratorLogin() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorLogin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAdministratorPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetBackupRetentionDays() {
	_jsii_.InvokeVoid(
		p,
		"resetBackupRetentionDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetCreateMode() {
	_jsii_.InvokeVoid(
		p,
		"resetCreateMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetDelegatedSubnetId() {
	_jsii_.InvokeVoid(
		p,
		"resetDelegatedSubnetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetGeoRedundantBackupEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetGeoRedundantBackupEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetHighAvailability() {
	_jsii_.InvokeVoid(
		p,
		"resetHighAvailability",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		p,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetPointInTimeRestoreTimeInUtc() {
	_jsii_.InvokeVoid(
		p,
		"resetPointInTimeRestoreTimeInUtc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetPrivateDnsZoneId() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateDnsZoneId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetSkuName() {
	_jsii_.InvokeVoid(
		p,
		"resetSkuName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetSourceServerId() {
	_jsii_.InvokeVoid(
		p,
		"resetSourceServerId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetStorageMb() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageMb",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetZone() {
	_jsii_.InvokeVoid(
		p,
		"resetZone",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PostgresqlFlexibleServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#location PostgresqlFlexibleServer#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#name PostgresqlFlexibleServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#resource_group_name PostgresqlFlexibleServer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#administrator_login PostgresqlFlexibleServer#administrator_login}.
	AdministratorLogin *string `field:"optional" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#administrator_password PostgresqlFlexibleServer#administrator_password}.
	AdministratorPassword *string `field:"optional" json:"administratorPassword" yaml:"administratorPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#backup_retention_days PostgresqlFlexibleServer#backup_retention_days}.
	BackupRetentionDays *float64 `field:"optional" json:"backupRetentionDays" yaml:"backupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#create_mode PostgresqlFlexibleServer#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#delegated_subnet_id PostgresqlFlexibleServer#delegated_subnet_id}.
	DelegatedSubnetId *string `field:"optional" json:"delegatedSubnetId" yaml:"delegatedSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#geo_redundant_backup_enabled PostgresqlFlexibleServer#geo_redundant_backup_enabled}.
	GeoRedundantBackupEnabled interface{} `field:"optional" json:"geoRedundantBackupEnabled" yaml:"geoRedundantBackupEnabled"`
	// high_availability block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#high_availability PostgresqlFlexibleServer#high_availability}
	HighAvailability *PostgresqlFlexibleServerHighAvailability `field:"optional" json:"highAvailability" yaml:"highAvailability"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#id PostgresqlFlexibleServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#maintenance_window PostgresqlFlexibleServer#maintenance_window}
	MaintenanceWindow *PostgresqlFlexibleServerMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#point_in_time_restore_time_in_utc PostgresqlFlexibleServer#point_in_time_restore_time_in_utc}.
	PointInTimeRestoreTimeInUtc *string `field:"optional" json:"pointInTimeRestoreTimeInUtc" yaml:"pointInTimeRestoreTimeInUtc"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#private_dns_zone_id PostgresqlFlexibleServer#private_dns_zone_id}.
	PrivateDnsZoneId *string `field:"optional" json:"privateDnsZoneId" yaml:"privateDnsZoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#sku_name PostgresqlFlexibleServer#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#source_server_id PostgresqlFlexibleServer#source_server_id}.
	SourceServerId *string `field:"optional" json:"sourceServerId" yaml:"sourceServerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#storage_mb PostgresqlFlexibleServer#storage_mb}.
	StorageMb *float64 `field:"optional" json:"storageMb" yaml:"storageMb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#tags PostgresqlFlexibleServer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#timeouts PostgresqlFlexibleServer#timeouts}
	Timeouts *PostgresqlFlexibleServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#version PostgresqlFlexibleServer#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#zone PostgresqlFlexibleServer#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

type PostgresqlFlexibleServerHighAvailability struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#mode PostgresqlFlexibleServer#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#standby_availability_zone PostgresqlFlexibleServer#standby_availability_zone}.
	StandbyAvailabilityZone *string `field:"optional" json:"standbyAvailabilityZone" yaml:"standbyAvailabilityZone"`
}

type PostgresqlFlexibleServerHighAvailabilityOutputReference interface {
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
	InternalValue() *PostgresqlFlexibleServerHighAvailability
	SetInternalValue(val *PostgresqlFlexibleServerHighAvailability)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	StandbyAvailabilityZone() *string
	SetStandbyAvailabilityZone(val *string)
	StandbyAvailabilityZoneInput() *string
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
	ResetStandbyAvailabilityZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresqlFlexibleServerHighAvailabilityOutputReference
type jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) InternalValue() *PostgresqlFlexibleServerHighAvailability {
	var returns *PostgresqlFlexibleServerHighAvailability
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) StandbyAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) StandbyAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresqlFlexibleServerHighAvailabilityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PostgresqlFlexibleServerHighAvailabilityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServerHighAvailabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresqlFlexibleServerHighAvailabilityOutputReference_Override(p PostgresqlFlexibleServerHighAvailabilityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServerHighAvailabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) SetInternalValue(val *PostgresqlFlexibleServerHighAvailability) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) SetStandbyAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"standbyAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) ResetStandbyAvailabilityZone() {
	_jsii_.InvokeVoid(
		p,
		"resetStandbyAvailabilityZone",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerHighAvailabilityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PostgresqlFlexibleServerMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#day_of_week PostgresqlFlexibleServer#day_of_week}.
	DayOfWeek *float64 `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#start_hour PostgresqlFlexibleServer#start_hour}.
	StartHour *float64 `field:"optional" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#start_minute PostgresqlFlexibleServer#start_minute}.
	StartMinute *float64 `field:"optional" json:"startMinute" yaml:"startMinute"`
}

type PostgresqlFlexibleServerMaintenanceWindowOutputReference interface {
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
	DayOfWeek() *float64
	SetDayOfWeek(val *float64)
	DayOfWeekInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *PostgresqlFlexibleServerMaintenanceWindow
	SetInternalValue(val *PostgresqlFlexibleServerMaintenanceWindow)
	StartHour() *float64
	SetStartHour(val *float64)
	StartHourInput() *float64
	StartMinute() *float64
	SetStartMinute(val *float64)
	StartMinuteInput() *float64
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
	ResetDayOfWeek()
	ResetStartHour()
	ResetStartMinute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresqlFlexibleServerMaintenanceWindowOutputReference
type jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) DayOfWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) DayOfWeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) InternalValue() *PostgresqlFlexibleServerMaintenanceWindow {
	var returns *PostgresqlFlexibleServerMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) StartHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) StartHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) StartMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) StartMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresqlFlexibleServerMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PostgresqlFlexibleServerMaintenanceWindowOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServerMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresqlFlexibleServerMaintenanceWindowOutputReference_Override(p PostgresqlFlexibleServerMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServerMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetDayOfWeek(val *float64) {
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetInternalValue(val *PostgresqlFlexibleServerMaintenanceWindow) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetStartHour(val *float64) {
	_jsii_.Set(
		j,
		"startHour",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetStartMinute(val *float64) {
	_jsii_.Set(
		j,
		"startMinute",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) ResetDayOfWeek() {
	_jsii_.InvokeVoid(
		p,
		"resetDayOfWeek",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) ResetStartHour() {
	_jsii_.InvokeVoid(
		p,
		"resetStartHour",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) ResetStartMinute() {
	_jsii_.InvokeVoid(
		p,
		"resetStartMinute",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PostgresqlFlexibleServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#create PostgresqlFlexibleServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#delete PostgresqlFlexibleServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#read PostgresqlFlexibleServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#update PostgresqlFlexibleServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type PostgresqlFlexibleServerTimeoutsOutputReference interface {
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

// The jsii proxy struct for PostgresqlFlexibleServerTimeoutsOutputReference
type jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewPostgresqlFlexibleServerTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PostgresqlFlexibleServerTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServerTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresqlFlexibleServerTimeoutsOutputReference_Override(p PostgresqlFlexibleServerTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServerTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		p,
		"resetCreate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		p,
		"resetDelete",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		p,
		"resetRead",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		p,
		"resetUpdate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServerTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


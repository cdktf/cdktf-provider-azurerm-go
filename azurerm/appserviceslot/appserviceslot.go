package appserviceslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/appserviceslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot azurerm_app_service_slot}.
type AppServiceSlot interface {
	cdktf.TerraformResource
	AppServiceName() *string
	SetAppServiceName(val *string)
	AppServiceNameInput() *string
	AppServicePlanId() *string
	SetAppServicePlanId(val *string)
	AppServicePlanIdInput() *string
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() AppServiceSlotAuthSettingsOutputReference
	AuthSettingsInput() *AppServiceSlotAuthSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientAffinityEnabled() interface{}
	SetClientAffinityEnabled(val interface{})
	ClientAffinityEnabledInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionString() AppServiceSlotConnectionStringList
	ConnectionStringInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultSiteHostname() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsOnly() interface{}
	SetHttpsOnly(val interface{})
	HttpsOnlyInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() AppServiceSlotIdentityOutputReference
	IdentityInput() *AppServiceSlotIdentity
	IdInput() *string
	KeyVaultReferenceIdentityId() *string
	SetKeyVaultReferenceIdentityId(val *string)
	KeyVaultReferenceIdentityIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Logs() AppServiceSlotLogsOutputReference
	LogsInput() *AppServiceSlotLogs
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SiteConfig() AppServiceSlotSiteConfigOutputReference
	SiteConfigInput() *AppServiceSlotSiteConfig
	SiteCredential() AppServiceSlotSiteCredentialList
	StorageAccount() AppServiceSlotStorageAccountList
	StorageAccountInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AppServiceSlotTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAuthSettings(value *AppServiceSlotAuthSettings)
	PutConnectionString(value interface{})
	PutIdentity(value *AppServiceSlotIdentity)
	PutLogs(value *AppServiceSlotLogs)
	PutSiteConfig(value *AppServiceSlotSiteConfig)
	PutStorageAccount(value interface{})
	PutTimeouts(value *AppServiceSlotTimeouts)
	ResetAppSettings()
	ResetAuthSettings()
	ResetClientAffinityEnabled()
	ResetConnectionString()
	ResetEnabled()
	ResetHttpsOnly()
	ResetId()
	ResetIdentity()
	ResetKeyVaultReferenceIdentityId()
	ResetLogs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSiteConfig()
	ResetStorageAccount()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppServiceSlot
type jsiiProxy_AppServiceSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppServiceSlot) AppServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) AppServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) AppServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) AppServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) AuthSettings() AppServiceSlotAuthSettingsOutputReference {
	var returns AppServiceSlotAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) AuthSettingsInput() *AppServiceSlotAuthSettings {
	var returns *AppServiceSlotAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ClientAffinityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAffinityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ClientAffinityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAffinityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ConnectionString() AppServiceSlotConnectionStringList {
	var returns AppServiceSlotConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) DefaultSiteHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSiteHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Identity() AppServiceSlotIdentityOutputReference {
	var returns AppServiceSlotIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) IdentityInput() *AppServiceSlotIdentity {
	var returns *AppServiceSlotIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) KeyVaultReferenceIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) KeyVaultReferenceIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Logs() AppServiceSlotLogsOutputReference {
	var returns AppServiceSlotLogsOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) LogsInput() *AppServiceSlotLogs {
	var returns *AppServiceSlotLogs
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) SiteConfig() AppServiceSlotSiteConfigOutputReference {
	var returns AppServiceSlotSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) SiteConfigInput() *AppServiceSlotSiteConfig {
	var returns *AppServiceSlotSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) SiteCredential() AppServiceSlotSiteCredentialList {
	var returns AppServiceSlotSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) StorageAccount() AppServiceSlotStorageAccountList {
	var returns AppServiceSlotStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) Timeouts() AppServiceSlotTimeoutsOutputReference {
	var returns AppServiceSlotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot azurerm_app_service_slot} Resource.
func NewAppServiceSlot(scope constructs.Construct, id *string, config *AppServiceSlotConfig) AppServiceSlot {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlot{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot azurerm_app_service_slot} Resource.
func NewAppServiceSlot_Override(a AppServiceSlot, scope constructs.Construct, id *string, config *AppServiceSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlot",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetAppServiceName(val *string) {
	_jsii_.Set(
		j,
		"appServiceName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetAppServicePlanId(val *string) {
	_jsii_.Set(
		j,
		"appServicePlanId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetAppSettings(val *map[string]*string) {
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetClientAffinityEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"clientAffinityEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetHttpsOnly(val interface{}) {
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetKeyVaultReferenceIdentityId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlot) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
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
func AppServiceSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppServiceSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppServiceSlot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppServiceSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppServiceSlot) PutAuthSettings(value *AppServiceSlotAuthSettings) {
	_jsii_.InvokeVoid(
		a,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlot) PutConnectionString(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlot) PutIdentity(value *AppServiceSlotIdentity) {
	_jsii_.InvokeVoid(
		a,
		"putIdentity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlot) PutLogs(value *AppServiceSlotLogs) {
	_jsii_.InvokeVoid(
		a,
		"putLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlot) PutSiteConfig(value *AppServiceSlotSiteConfig) {
	_jsii_.InvokeVoid(
		a,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlot) PutStorageAccount(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlot) PutTimeouts(value *AppServiceSlotTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetClientAffinityEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetClientAffinityEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetConnectionString() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetKeyVaultReferenceIdentityId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultReferenceIdentityId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetSiteConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSiteConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotAuthSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#enabled AppServiceSlot#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#active_directory AppServiceSlot#active_directory}
	ActiveDirectory *AppServiceSlotAuthSettingsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#additional_login_params AppServiceSlot#additional_login_params}.
	AdditionalLoginParams *map[string]*string `field:"optional" json:"additionalLoginParams" yaml:"additionalLoginParams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#allowed_external_redirect_urls AppServiceSlot#allowed_external_redirect_urls}.
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#default_provider AppServiceSlot#default_provider}.
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// facebook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#facebook AppServiceSlot#facebook}
	Facebook *AppServiceSlotAuthSettingsFacebook `field:"optional" json:"facebook" yaml:"facebook"`
	// google block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#google AppServiceSlot#google}
	Google *AppServiceSlotAuthSettingsGoogle `field:"optional" json:"google" yaml:"google"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#issuer AppServiceSlot#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// microsoft block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#microsoft AppServiceSlot#microsoft}
	Microsoft *AppServiceSlotAuthSettingsMicrosoft `field:"optional" json:"microsoft" yaml:"microsoft"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#runtime_version AppServiceSlot#runtime_version}.
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#token_refresh_extension_hours AppServiceSlot#token_refresh_extension_hours}.
	TokenRefreshExtensionHours *float64 `field:"optional" json:"tokenRefreshExtensionHours" yaml:"tokenRefreshExtensionHours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#token_store_enabled AppServiceSlot#token_store_enabled}.
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#twitter AppServiceSlot#twitter}
	Twitter *AppServiceSlotAuthSettingsTwitter `field:"optional" json:"twitter" yaml:"twitter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#unauthenticated_client_action AppServiceSlot#unauthenticated_client_action}.
	UnauthenticatedClientAction *string `field:"optional" json:"unauthenticatedClientAction" yaml:"unauthenticatedClientAction"`
}

type AppServiceSlotAuthSettingsActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#client_id AppServiceSlot#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#allowed_audiences AppServiceSlot#allowed_audiences}.
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#client_secret AppServiceSlot#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

type AppServiceSlotAuthSettingsActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	AllowedAudiences() *[]*string
	SetAllowedAudiences(val *[]*string)
	AllowedAudiencesInput() *[]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	InternalValue() *AppServiceSlotAuthSettingsActiveDirectory
	SetInternalValue(val *AppServiceSlotAuthSettingsActiveDirectory)
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
	ResetAllowedAudiences()
	ResetClientSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotAuthSettingsActiveDirectoryOutputReference
type jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) AllowedAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) AllowedAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) InternalValue() *AppServiceSlotAuthSettingsActiveDirectory {
	var returns *AppServiceSlotAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotAuthSettingsActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotAuthSettingsActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotAuthSettingsActiveDirectoryOutputReference_Override(a AppServiceSlotAuthSettingsActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetAllowedAudiences(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedAudiences",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetInternalValue(val *AppServiceSlotAuthSettingsActiveDirectory) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ResetAllowedAudiences() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedAudiences",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotAuthSettingsFacebook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#app_id AppServiceSlot#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#app_secret AppServiceSlot#app_secret}.
	AppSecret *string `field:"required" json:"appSecret" yaml:"appSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#oauth_scopes AppServiceSlot#oauth_scopes}.
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

type AppServiceSlotAuthSettingsFacebookOutputReference interface {
	cdktf.ComplexObject
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	AppSecret() *string
	SetAppSecret(val *string)
	AppSecretInput() *string
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
	InternalValue() *AppServiceSlotAuthSettingsFacebook
	SetInternalValue(val *AppServiceSlotAuthSettingsFacebook)
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
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
	ResetOauthScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotAuthSettingsFacebookOutputReference
type jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) AppSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) AppSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) InternalValue() *AppServiceSlotAuthSettingsFacebook {
	var returns *AppServiceSlotAuthSettingsFacebook
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotAuthSettingsFacebookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotAuthSettingsFacebookOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsFacebookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotAuthSettingsFacebookOutputReference_Override(a AppServiceSlotAuthSettingsFacebookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsFacebookOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetAppSecret(val *string) {
	_jsii_.Set(
		j,
		"appSecret",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetInternalValue(val *AppServiceSlotAuthSettingsFacebook) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsFacebookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotAuthSettingsGoogle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#client_id AppServiceSlot#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#client_secret AppServiceSlot#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#oauth_scopes AppServiceSlot#oauth_scopes}.
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

type AppServiceSlotAuthSettingsGoogleOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	InternalValue() *AppServiceSlotAuthSettingsGoogle
	SetInternalValue(val *AppServiceSlotAuthSettingsGoogle)
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
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
	ResetOauthScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotAuthSettingsGoogleOutputReference
type jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) InternalValue() *AppServiceSlotAuthSettingsGoogle {
	var returns *AppServiceSlotAuthSettingsGoogle
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotAuthSettingsGoogleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotAuthSettingsGoogleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsGoogleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotAuthSettingsGoogleOutputReference_Override(a AppServiceSlotAuthSettingsGoogleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsGoogleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetInternalValue(val *AppServiceSlotAuthSettingsGoogle) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsGoogleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotAuthSettingsMicrosoft struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#client_id AppServiceSlot#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#client_secret AppServiceSlot#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#oauth_scopes AppServiceSlot#oauth_scopes}.
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

type AppServiceSlotAuthSettingsMicrosoftOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	InternalValue() *AppServiceSlotAuthSettingsMicrosoft
	SetInternalValue(val *AppServiceSlotAuthSettingsMicrosoft)
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
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
	ResetOauthScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotAuthSettingsMicrosoftOutputReference
type jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) InternalValue() *AppServiceSlotAuthSettingsMicrosoft {
	var returns *AppServiceSlotAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotAuthSettingsMicrosoftOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotAuthSettingsMicrosoftOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsMicrosoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotAuthSettingsMicrosoftOutputReference_Override(a AppServiceSlotAuthSettingsMicrosoftOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsMicrosoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetInternalValue(val *AppServiceSlotAuthSettingsMicrosoft) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetOauthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsMicrosoftOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotAuthSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() AppServiceSlotAuthSettingsActiveDirectoryOutputReference
	ActiveDirectoryInput() *AppServiceSlotAuthSettingsActiveDirectory
	AdditionalLoginParams() *map[string]*string
	SetAdditionalLoginParams(val *map[string]*string)
	AdditionalLoginParamsInput() *map[string]*string
	AllowedExternalRedirectUrls() *[]*string
	SetAllowedExternalRedirectUrls(val *[]*string)
	AllowedExternalRedirectUrlsInput() *[]*string
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
	DefaultProvider() *string
	SetDefaultProvider(val *string)
	DefaultProviderInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Facebook() AppServiceSlotAuthSettingsFacebookOutputReference
	FacebookInput() *AppServiceSlotAuthSettingsFacebook
	// Experimental.
	Fqn() *string
	Google() AppServiceSlotAuthSettingsGoogleOutputReference
	GoogleInput() *AppServiceSlotAuthSettingsGoogle
	InternalValue() *AppServiceSlotAuthSettings
	SetInternalValue(val *AppServiceSlotAuthSettings)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	Microsoft() AppServiceSlotAuthSettingsMicrosoftOutputReference
	MicrosoftInput() *AppServiceSlotAuthSettingsMicrosoft
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenRefreshExtensionHours() *float64
	SetTokenRefreshExtensionHours(val *float64)
	TokenRefreshExtensionHoursInput() *float64
	TokenStoreEnabled() interface{}
	SetTokenStoreEnabled(val interface{})
	TokenStoreEnabledInput() interface{}
	Twitter() AppServiceSlotAuthSettingsTwitterOutputReference
	TwitterInput() *AppServiceSlotAuthSettingsTwitter
	UnauthenticatedClientAction() *string
	SetUnauthenticatedClientAction(val *string)
	UnauthenticatedClientActionInput() *string
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
	PutActiveDirectory(value *AppServiceSlotAuthSettingsActiveDirectory)
	PutFacebook(value *AppServiceSlotAuthSettingsFacebook)
	PutGoogle(value *AppServiceSlotAuthSettingsGoogle)
	PutMicrosoft(value *AppServiceSlotAuthSettingsMicrosoft)
	PutTwitter(value *AppServiceSlotAuthSettingsTwitter)
	ResetActiveDirectory()
	ResetAdditionalLoginParams()
	ResetAllowedExternalRedirectUrls()
	ResetDefaultProvider()
	ResetFacebook()
	ResetGoogle()
	ResetIssuer()
	ResetMicrosoft()
	ResetRuntimeVersion()
	ResetTokenRefreshExtensionHours()
	ResetTokenStoreEnabled()
	ResetTwitter()
	ResetUnauthenticatedClientAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotAuthSettingsOutputReference
type jsiiProxy_AppServiceSlotAuthSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ActiveDirectory() AppServiceSlotAuthSettingsActiveDirectoryOutputReference {
	var returns AppServiceSlotAuthSettingsActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ActiveDirectoryInput() *AppServiceSlotAuthSettingsActiveDirectory {
	var returns *AppServiceSlotAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AdditionalLoginParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AdditionalLoginParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Facebook() AppServiceSlotAuthSettingsFacebookOutputReference {
	var returns AppServiceSlotAuthSettingsFacebookOutputReference
	_jsii_.Get(
		j,
		"facebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) FacebookInput() *AppServiceSlotAuthSettingsFacebook {
	var returns *AppServiceSlotAuthSettingsFacebook
	_jsii_.Get(
		j,
		"facebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Google() AppServiceSlotAuthSettingsGoogleOutputReference {
	var returns AppServiceSlotAuthSettingsGoogleOutputReference
	_jsii_.Get(
		j,
		"google",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GoogleInput() *AppServiceSlotAuthSettingsGoogle {
	var returns *AppServiceSlotAuthSettingsGoogle
	_jsii_.Get(
		j,
		"googleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) InternalValue() *AppServiceSlotAuthSettings {
	var returns *AppServiceSlotAuthSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Microsoft() AppServiceSlotAuthSettingsMicrosoftOutputReference {
	var returns AppServiceSlotAuthSettingsMicrosoftOutputReference
	_jsii_.Get(
		j,
		"microsoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) MicrosoftInput() *AppServiceSlotAuthSettingsMicrosoft {
	var returns *AppServiceSlotAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"microsoftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenRefreshExtensionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenRefreshExtensionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Twitter() AppServiceSlotAuthSettingsTwitterOutputReference {
	var returns AppServiceSlotAuthSettingsTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TwitterInput() *AppServiceSlotAuthSettingsTwitter {
	var returns *AppServiceSlotAuthSettingsTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) UnauthenticatedClientAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) UnauthenticatedClientActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientActionInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotAuthSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotAuthSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotAuthSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotAuthSettingsOutputReference_Override(a AppServiceSlotAuthSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetAdditionalLoginParams(val *map[string]*string) {
	_jsii_.Set(
		j,
		"additionalLoginParams",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetAllowedExternalRedirectUrls(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetDefaultProvider(val *string) {
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetInternalValue(val *AppServiceSlotAuthSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetIssuer(val *string) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetTokenRefreshExtensionHours(val *float64) {
	_jsii_.Set(
		j,
		"tokenRefreshExtensionHours",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetTokenStoreEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) SetUnauthenticatedClientAction(val *string) {
	_jsii_.Set(
		j,
		"unauthenticatedClientAction",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutActiveDirectory(value *AppServiceSlotAuthSettingsActiveDirectory) {
	_jsii_.InvokeVoid(
		a,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutFacebook(value *AppServiceSlotAuthSettingsFacebook) {
	_jsii_.InvokeVoid(
		a,
		"putFacebook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutGoogle(value *AppServiceSlotAuthSettingsGoogle) {
	_jsii_.InvokeVoid(
		a,
		"putGoogle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutMicrosoft(value *AppServiceSlotAuthSettingsMicrosoft) {
	_jsii_.InvokeVoid(
		a,
		"putMicrosoft",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutTwitter(value *AppServiceSlotAuthSettingsTwitter) {
	_jsii_.InvokeVoid(
		a,
		"putTwitter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetAdditionalLoginParams() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalLoginParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetFacebook() {
	_jsii_.InvokeVoid(
		a,
		"resetFacebook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetGoogle() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetMicrosoft() {
	_jsii_.InvokeVoid(
		a,
		"resetMicrosoft",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetTokenRefreshExtensionHours() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenRefreshExtensionHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		a,
		"resetTwitter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetUnauthenticatedClientAction() {
	_jsii_.InvokeVoid(
		a,
		"resetUnauthenticatedClientAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotAuthSettingsTwitter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#consumer_key AppServiceSlot#consumer_key}.
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#consumer_secret AppServiceSlot#consumer_secret}.
	ConsumerSecret *string `field:"required" json:"consumerSecret" yaml:"consumerSecret"`
}

type AppServiceSlotAuthSettingsTwitterOutputReference interface {
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
	ConsumerKey() *string
	SetConsumerKey(val *string)
	ConsumerKeyInput() *string
	ConsumerSecret() *string
	SetConsumerSecret(val *string)
	ConsumerSecretInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *AppServiceSlotAuthSettingsTwitter
	SetInternalValue(val *AppServiceSlotAuthSettingsTwitter)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotAuthSettingsTwitterOutputReference
type jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ConsumerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ConsumerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ConsumerSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ConsumerSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) InternalValue() *AppServiceSlotAuthSettingsTwitter {
	var returns *AppServiceSlotAuthSettingsTwitter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotAuthSettingsTwitterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotAuthSettingsTwitterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsTwitterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotAuthSettingsTwitterOutputReference_Override(a AppServiceSlotAuthSettingsTwitterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsTwitterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) SetConsumerKey(val *string) {
	_jsii_.Set(
		j,
		"consumerKey",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) SetConsumerSecret(val *string) {
	_jsii_.Set(
		j,
		"consumerSecret",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) SetInternalValue(val *AppServiceSlotAuthSettingsTwitter) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsTwitterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#app_service_name AppServiceSlot#app_service_name}.
	AppServiceName *string `field:"required" json:"appServiceName" yaml:"appServiceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#app_service_plan_id AppServiceSlot#app_service_plan_id}.
	AppServicePlanId *string `field:"required" json:"appServicePlanId" yaml:"appServicePlanId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#location AppServiceSlot#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#name AppServiceSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#resource_group_name AppServiceSlot#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#app_settings AppServiceSlot#app_settings}.
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#auth_settings AppServiceSlot#auth_settings}
	AuthSettings *AppServiceSlotAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#client_affinity_enabled AppServiceSlot#client_affinity_enabled}.
	ClientAffinityEnabled interface{} `field:"optional" json:"clientAffinityEnabled" yaml:"clientAffinityEnabled"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#connection_string AppServiceSlot#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#enabled AppServiceSlot#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#https_only AppServiceSlot#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#id AppServiceSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#identity AppServiceSlot#identity}
	Identity *AppServiceSlotIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#key_vault_reference_identity_id AppServiceSlot#key_vault_reference_identity_id}.
	KeyVaultReferenceIdentityId *string `field:"optional" json:"keyVaultReferenceIdentityId" yaml:"keyVaultReferenceIdentityId"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#logs AppServiceSlot#logs}
	Logs *AppServiceSlotLogs `field:"optional" json:"logs" yaml:"logs"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#site_config AppServiceSlot#site_config}
	SiteConfig *AppServiceSlotSiteConfig `field:"optional" json:"siteConfig" yaml:"siteConfig"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#storage_account AppServiceSlot#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#tags AppServiceSlot#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#timeouts AppServiceSlot#timeouts}
	Timeouts *AppServiceSlotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type AppServiceSlotConnectionString struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#name AppServiceSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#type AppServiceSlot#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#value AppServiceSlot#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

type AppServiceSlotConnectionStringList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppServiceSlotConnectionStringOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotConnectionStringList
type jsiiProxy_AppServiceSlotConnectionStringList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppServiceSlotConnectionStringList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppServiceSlotConnectionStringList {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotConnectionStringList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotConnectionStringList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotConnectionStringList_Override(a AppServiceSlotConnectionStringList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotConnectionStringList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotConnectionStringList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringList) Get(index *float64) AppServiceSlotConnectionStringOutputReference {
	var returns AppServiceSlotConnectionStringOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotConnectionStringOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotConnectionStringOutputReference
type jsiiProxy_AppServiceSlotConnectionStringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotConnectionStringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppServiceSlotConnectionStringOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotConnectionStringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotConnectionStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotConnectionStringOutputReference_Override(a AppServiceSlotConnectionStringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotConnectionStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotConnectionStringOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotConnectionStringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#type AppServiceSlot#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#identity_ids AppServiceSlot#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type AppServiceSlotIdentityOutputReference interface {
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
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *AppServiceSlotIdentity
	SetInternalValue(val *AppServiceSlotIdentity)
	PrincipalId() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotIdentityOutputReference
type jsiiProxy_AppServiceSlotIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) InternalValue() *AppServiceSlotIdentity {
	var returns *AppServiceSlotIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotIdentityOutputReference_Override(a AppServiceSlotIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) SetInternalValue(val *AppServiceSlotIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotLogs struct {
	// application_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#application_logs AppServiceSlot#application_logs}
	ApplicationLogs *AppServiceSlotLogsApplicationLogs `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#detailed_error_messages_enabled AppServiceSlot#detailed_error_messages_enabled}.
	DetailedErrorMessagesEnabled interface{} `field:"optional" json:"detailedErrorMessagesEnabled" yaml:"detailedErrorMessagesEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#failed_request_tracing_enabled AppServiceSlot#failed_request_tracing_enabled}.
	FailedRequestTracingEnabled interface{} `field:"optional" json:"failedRequestTracingEnabled" yaml:"failedRequestTracingEnabled"`
	// http_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#http_logs AppServiceSlot#http_logs}
	HttpLogs *AppServiceSlotLogsHttpLogs `field:"optional" json:"httpLogs" yaml:"httpLogs"`
}

type AppServiceSlotLogsApplicationLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#azure_blob_storage AppServiceSlot#azure_blob_storage}
	AzureBlobStorage *AppServiceSlotLogsApplicationLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#file_system_level AppServiceSlot#file_system_level}.
	FileSystemLevel *string `field:"optional" json:"fileSystemLevel" yaml:"fileSystemLevel"`
}

type AppServiceSlotLogsApplicationLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#level AppServiceSlot#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#retention_in_days AppServiceSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#sas_url AppServiceSlot#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
}

type AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference interface {
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
	InternalValue() *AppServiceSlotLogsApplicationLogsAzureBlobStorage
	SetInternalValue(val *AppServiceSlotLogsApplicationLogsAzureBlobStorage)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	SasUrl() *string
	SetSasUrl(val *string)
	SasUrlInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference
type jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) InternalValue() *AppServiceSlotLogsApplicationLogsAzureBlobStorage {
	var returns *AppServiceSlotLogsApplicationLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SasUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SasUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference_Override(a AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetInternalValue(val *AppServiceSlotLogsApplicationLogsAzureBlobStorage) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetLevel(val *string) {
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetSasUrl(val *string) {
	_jsii_.Set(
		j,
		"sasUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotLogsApplicationLogsOutputReference interface {
	cdktf.ComplexObject
	AzureBlobStorage() AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference
	AzureBlobStorageInput() *AppServiceSlotLogsApplicationLogsAzureBlobStorage
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
	FileSystemLevel() *string
	SetFileSystemLevel(val *string)
	FileSystemLevelInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AppServiceSlotLogsApplicationLogs
	SetInternalValue(val *AppServiceSlotLogsApplicationLogs)
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
	PutAzureBlobStorage(value *AppServiceSlotLogsApplicationLogsAzureBlobStorage)
	ResetAzureBlobStorage()
	ResetFileSystemLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotLogsApplicationLogsOutputReference
type jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) AzureBlobStorage() AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference {
	var returns AppServiceSlotLogsApplicationLogsAzureBlobStorageOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) AzureBlobStorageInput() *AppServiceSlotLogsApplicationLogsAzureBlobStorage {
	var returns *AppServiceSlotLogsApplicationLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"azureBlobStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) FileSystemLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) FileSystemLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) InternalValue() *AppServiceSlotLogsApplicationLogs {
	var returns *AppServiceSlotLogsApplicationLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotLogsApplicationLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotLogsApplicationLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsApplicationLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotLogsApplicationLogsOutputReference_Override(a AppServiceSlotLogsApplicationLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsApplicationLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) SetFileSystemLevel(val *string) {
	_jsii_.Set(
		j,
		"fileSystemLevel",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) SetInternalValue(val *AppServiceSlotLogsApplicationLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) PutAzureBlobStorage(value *AppServiceSlotLogsApplicationLogsAzureBlobStorage) {
	_jsii_.InvokeVoid(
		a,
		"putAzureBlobStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) ResetAzureBlobStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureBlobStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) ResetFileSystemLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystemLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsApplicationLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotLogsHttpLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#azure_blob_storage AppServiceSlot#azure_blob_storage}
	AzureBlobStorage *AppServiceSlotLogsHttpLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// file_system block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#file_system AppServiceSlot#file_system}
	FileSystem *AppServiceSlotLogsHttpLogsFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

type AppServiceSlotLogsHttpLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#retention_in_days AppServiceSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#sas_url AppServiceSlot#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
}

type AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference interface {
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
	InternalValue() *AppServiceSlotLogsHttpLogsAzureBlobStorage
	SetInternalValue(val *AppServiceSlotLogsHttpLogsAzureBlobStorage)
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	SasUrl() *string
	SetSasUrl(val *string)
	SasUrlInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference
type jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) InternalValue() *AppServiceSlotLogsHttpLogsAzureBlobStorage {
	var returns *AppServiceSlotLogsHttpLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SasUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SasUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference_Override(a AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SetInternalValue(val *AppServiceSlotLogsHttpLogsAzureBlobStorage) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SetRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SetSasUrl(val *string) {
	_jsii_.Set(
		j,
		"sasUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotLogsHttpLogsFileSystem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#retention_in_days AppServiceSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#retention_in_mb AppServiceSlot#retention_in_mb}.
	RetentionInMb *float64 `field:"required" json:"retentionInMb" yaml:"retentionInMb"`
}

type AppServiceSlotLogsHttpLogsFileSystemOutputReference interface {
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
	InternalValue() *AppServiceSlotLogsHttpLogsFileSystem
	SetInternalValue(val *AppServiceSlotLogsHttpLogsFileSystem)
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	RetentionInMb() *float64
	SetRetentionInMb(val *float64)
	RetentionInMbInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotLogsHttpLogsFileSystemOutputReference
type jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) InternalValue() *AppServiceSlotLogsHttpLogsFileSystem {
	var returns *AppServiceSlotLogsHttpLogsFileSystem
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) RetentionInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) RetentionInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotLogsHttpLogsFileSystemOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotLogsHttpLogsFileSystemOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsHttpLogsFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotLogsHttpLogsFileSystemOutputReference_Override(a AppServiceSlotLogsHttpLogsFileSystemOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsHttpLogsFileSystemOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) SetInternalValue(val *AppServiceSlotLogsHttpLogsFileSystem) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) SetRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) SetRetentionInMb(val *float64) {
	_jsii_.Set(
		j,
		"retentionInMb",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsFileSystemOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotLogsHttpLogsOutputReference interface {
	cdktf.ComplexObject
	AzureBlobStorage() AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference
	AzureBlobStorageInput() *AppServiceSlotLogsHttpLogsAzureBlobStorage
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
	FileSystem() AppServiceSlotLogsHttpLogsFileSystemOutputReference
	FileSystemInput() *AppServiceSlotLogsHttpLogsFileSystem
	// Experimental.
	Fqn() *string
	InternalValue() *AppServiceSlotLogsHttpLogs
	SetInternalValue(val *AppServiceSlotLogsHttpLogs)
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
	PutAzureBlobStorage(value *AppServiceSlotLogsHttpLogsAzureBlobStorage)
	PutFileSystem(value *AppServiceSlotLogsHttpLogsFileSystem)
	ResetAzureBlobStorage()
	ResetFileSystem()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotLogsHttpLogsOutputReference
type jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) AzureBlobStorage() AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference {
	var returns AppServiceSlotLogsHttpLogsAzureBlobStorageOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) AzureBlobStorageInput() *AppServiceSlotLogsHttpLogsAzureBlobStorage {
	var returns *AppServiceSlotLogsHttpLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"azureBlobStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) FileSystem() AppServiceSlotLogsHttpLogsFileSystemOutputReference {
	var returns AppServiceSlotLogsHttpLogsFileSystemOutputReference
	_jsii_.Get(
		j,
		"fileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) FileSystemInput() *AppServiceSlotLogsHttpLogsFileSystem {
	var returns *AppServiceSlotLogsHttpLogsFileSystem
	_jsii_.Get(
		j,
		"fileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) InternalValue() *AppServiceSlotLogsHttpLogs {
	var returns *AppServiceSlotLogsHttpLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotLogsHttpLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotLogsHttpLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsHttpLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotLogsHttpLogsOutputReference_Override(a AppServiceSlotLogsHttpLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsHttpLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) SetInternalValue(val *AppServiceSlotLogsHttpLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) PutAzureBlobStorage(value *AppServiceSlotLogsHttpLogsAzureBlobStorage) {
	_jsii_.InvokeVoid(
		a,
		"putAzureBlobStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) PutFileSystem(value *AppServiceSlotLogsHttpLogsFileSystem) {
	_jsii_.InvokeVoid(
		a,
		"putFileSystem",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) ResetAzureBlobStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureBlobStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) ResetFileSystem() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystem",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsHttpLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotLogsOutputReference interface {
	cdktf.ComplexObject
	ApplicationLogs() AppServiceSlotLogsApplicationLogsOutputReference
	ApplicationLogsInput() *AppServiceSlotLogsApplicationLogs
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
	DetailedErrorMessagesEnabled() interface{}
	SetDetailedErrorMessagesEnabled(val interface{})
	DetailedErrorMessagesEnabledInput() interface{}
	FailedRequestTracingEnabled() interface{}
	SetFailedRequestTracingEnabled(val interface{})
	FailedRequestTracingEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HttpLogs() AppServiceSlotLogsHttpLogsOutputReference
	HttpLogsInput() *AppServiceSlotLogsHttpLogs
	InternalValue() *AppServiceSlotLogs
	SetInternalValue(val *AppServiceSlotLogs)
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
	PutApplicationLogs(value *AppServiceSlotLogsApplicationLogs)
	PutHttpLogs(value *AppServiceSlotLogsHttpLogs)
	ResetApplicationLogs()
	ResetDetailedErrorMessagesEnabled()
	ResetFailedRequestTracingEnabled()
	ResetHttpLogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotLogsOutputReference
type jsiiProxy_AppServiceSlotLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ApplicationLogs() AppServiceSlotLogsApplicationLogsOutputReference {
	var returns AppServiceSlotLogsApplicationLogsOutputReference
	_jsii_.Get(
		j,
		"applicationLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ApplicationLogsInput() *AppServiceSlotLogsApplicationLogs {
	var returns *AppServiceSlotLogsApplicationLogs
	_jsii_.Get(
		j,
		"applicationLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) DetailedErrorMessagesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedErrorMessagesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) DetailedErrorMessagesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedErrorMessagesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) FailedRequestTracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedRequestTracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) FailedRequestTracingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedRequestTracingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) HttpLogs() AppServiceSlotLogsHttpLogsOutputReference {
	var returns AppServiceSlotLogsHttpLogsOutputReference
	_jsii_.Get(
		j,
		"httpLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) HttpLogsInput() *AppServiceSlotLogsHttpLogs {
	var returns *AppServiceSlotLogsHttpLogs
	_jsii_.Get(
		j,
		"httpLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) InternalValue() *AppServiceSlotLogs {
	var returns *AppServiceSlotLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotLogsOutputReference_Override(a AppServiceSlotLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) SetDetailedErrorMessagesEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"detailedErrorMessagesEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) SetFailedRequestTracingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"failedRequestTracingEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) SetInternalValue(val *AppServiceSlotLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) PutApplicationLogs(value *AppServiceSlotLogsApplicationLogs) {
	_jsii_.InvokeVoid(
		a,
		"putApplicationLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) PutHttpLogs(value *AppServiceSlotLogsHttpLogs) {
	_jsii_.InvokeVoid(
		a,
		"putHttpLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetApplicationLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetDetailedErrorMessagesEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedErrorMessagesEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetFailedRequestTracingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFailedRequestTracingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ResetHttpLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#acr_use_managed_identity_credentials AppServiceSlot#acr_use_managed_identity_credentials}.
	AcrUseManagedIdentityCredentials interface{} `field:"optional" json:"acrUseManagedIdentityCredentials" yaml:"acrUseManagedIdentityCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#acr_user_managed_identity_client_id AppServiceSlot#acr_user_managed_identity_client_id}.
	AcrUserManagedIdentityClientId *string `field:"optional" json:"acrUserManagedIdentityClientId" yaml:"acrUserManagedIdentityClientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#always_on AppServiceSlot#always_on}.
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#app_command_line AppServiceSlot#app_command_line}.
	AppCommandLine *string `field:"optional" json:"appCommandLine" yaml:"appCommandLine"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#auto_swap_slot_name AppServiceSlot#auto_swap_slot_name}.
	AutoSwapSlotName *string `field:"optional" json:"autoSwapSlotName" yaml:"autoSwapSlotName"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#cors AppServiceSlot#cors}
	Cors *AppServiceSlotSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#default_documents AppServiceSlot#default_documents}.
	DefaultDocuments *[]*string `field:"optional" json:"defaultDocuments" yaml:"defaultDocuments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#dotnet_framework_version AppServiceSlot#dotnet_framework_version}.
	DotnetFrameworkVersion *string `field:"optional" json:"dotnetFrameworkVersion" yaml:"dotnetFrameworkVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#ftps_state AppServiceSlot#ftps_state}.
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#health_check_path AppServiceSlot#health_check_path}.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#http2_enabled AppServiceSlot#http2_enabled}.
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#ip_restriction AppServiceSlot#ip_restriction}.
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#java_container AppServiceSlot#java_container}.
	JavaContainer *string `field:"optional" json:"javaContainer" yaml:"javaContainer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#java_container_version AppServiceSlot#java_container_version}.
	JavaContainerVersion *string `field:"optional" json:"javaContainerVersion" yaml:"javaContainerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#java_version AppServiceSlot#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#linux_fx_version AppServiceSlot#linux_fx_version}.
	LinuxFxVersion *string `field:"optional" json:"linuxFxVersion" yaml:"linuxFxVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#local_mysql_enabled AppServiceSlot#local_mysql_enabled}.
	LocalMysqlEnabled interface{} `field:"optional" json:"localMysqlEnabled" yaml:"localMysqlEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#managed_pipeline_mode AppServiceSlot#managed_pipeline_mode}.
	ManagedPipelineMode *string `field:"optional" json:"managedPipelineMode" yaml:"managedPipelineMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#min_tls_version AppServiceSlot#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#number_of_workers AppServiceSlot#number_of_workers}.
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#php_version AppServiceSlot#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#python_version AppServiceSlot#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#remote_debugging_enabled AppServiceSlot#remote_debugging_enabled}.
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#remote_debugging_version AppServiceSlot#remote_debugging_version}.
	RemoteDebuggingVersion *string `field:"optional" json:"remoteDebuggingVersion" yaml:"remoteDebuggingVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#scm_ip_restriction AppServiceSlot#scm_ip_restriction}.
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#scm_type AppServiceSlot#scm_type}.
	ScmType *string `field:"optional" json:"scmType" yaml:"scmType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#scm_use_main_ip_restriction AppServiceSlot#scm_use_main_ip_restriction}.
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#use_32_bit_worker_process AppServiceSlot#use_32_bit_worker_process}.
	Use32BitWorkerProcess interface{} `field:"optional" json:"use32BitWorkerProcess" yaml:"use32BitWorkerProcess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#vnet_route_all_enabled AppServiceSlot#vnet_route_all_enabled}.
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#websockets_enabled AppServiceSlot#websockets_enabled}.
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#windows_fx_version AppServiceSlot#windows_fx_version}.
	WindowsFxVersion *string `field:"optional" json:"windowsFxVersion" yaml:"windowsFxVersion"`
}

type AppServiceSlotSiteConfigCors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#allowed_origins AppServiceSlot#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#support_credentials AppServiceSlot#support_credentials}.
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

type AppServiceSlotSiteConfigCorsOutputReference interface {
	cdktf.ComplexObject
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	InternalValue() *AppServiceSlotSiteConfigCors
	SetInternalValue(val *AppServiceSlotSiteConfigCors)
	SupportCredentials() interface{}
	SetSupportCredentials(val interface{})
	SupportCredentialsInput() interface{}
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
	ResetSupportCredentials()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigCorsOutputReference
type jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) InternalValue() *AppServiceSlotSiteConfigCors {
	var returns *AppServiceSlotSiteConfigCors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SupportCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SupportCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigCorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotSiteConfigCorsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigCorsOutputReference_Override(a AppServiceSlotSiteConfigCorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SetAllowedOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SetInternalValue(val *AppServiceSlotSiteConfigCors) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SetSupportCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"supportCredentials",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) ResetSupportCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigCorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigIpRestriction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#action AppServiceSlot#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#headers AppServiceSlot#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#ip_address AppServiceSlot#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#name AppServiceSlot#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#priority AppServiceSlot#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#service_tag AppServiceSlot#service_tag}.
	ServiceTag *string `field:"optional" json:"serviceTag" yaml:"serviceTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#virtual_network_subnet_id AppServiceSlot#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
}

type AppServiceSlotSiteConfigIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_azure_fdid AppServiceSlot#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_fd_health_probe AppServiceSlot#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_forwarded_for AppServiceSlot#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_forwarded_host AppServiceSlot#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}

type AppServiceSlotSiteConfigIpRestrictionHeadersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigIpRestrictionHeadersList
type jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigIpRestrictionHeadersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppServiceSlotSiteConfigIpRestrictionHeadersList {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigIpRestrictionHeadersList_Override(a AppServiceSlotSiteConfigIpRestrictionHeadersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) Get(index *float64) AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference {
	var returns AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XAzureFdid() *[]*string
	SetXAzureFdid(val *[]*string)
	XAzureFdidInput() *[]*string
	XFdHealthProbe() *[]*string
	SetXFdHealthProbe(val *[]*string)
	XFdHealthProbeInput() *[]*string
	XForwardedFor() *[]*string
	SetXForwardedFor(val *[]*string)
	XForwardedForInput() *[]*string
	XForwardedHost() *[]*string
	SetXForwardedHost(val *[]*string)
	XForwardedHostInput() *[]*string
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
	ResetXAzureFdid()
	ResetXFdHealthProbe()
	ResetXForwardedFor()
	ResetXForwardedHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference
type jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XAzureFdid() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XAzureFdidInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XFdHealthProbe() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XFdHealthProbeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XForwardedFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XForwardedForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XForwardedHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) XForwardedHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHostInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigIpRestrictionHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigIpRestrictionHeadersOutputReference_Override(a AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetXAzureFdid(val *[]*string) {
	_jsii_.Set(
		j,
		"xAzureFdid",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetXFdHealthProbe(val *[]*string) {
	_jsii_.Set(
		j,
		"xFdHealthProbe",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetXForwardedFor(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedFor",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) SetXForwardedHost(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedHost",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ResetXAzureFdid() {
	_jsii_.InvokeVoid(
		a,
		"resetXAzureFdid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ResetXFdHealthProbe() {
	_jsii_.InvokeVoid(
		a,
		"resetXFdHealthProbe",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ResetXForwardedFor() {
	_jsii_.InvokeVoid(
		a,
		"resetXForwardedFor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ResetXForwardedHost() {
	_jsii_.InvokeVoid(
		a,
		"resetXForwardedHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigIpRestrictionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppServiceSlotSiteConfigIpRestrictionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigIpRestrictionList
type jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigIpRestrictionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppServiceSlotSiteConfigIpRestrictionList {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigIpRestrictionList_Override(a AppServiceSlotSiteConfigIpRestrictionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) Get(index *float64) AppServiceSlotSiteConfigIpRestrictionOutputReference {
	var returns AppServiceSlotSiteConfigIpRestrictionOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigIpRestrictionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Headers() AppServiceSlotSiteConfigIpRestrictionHeadersList
	HeadersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ServiceTag() *string
	SetServiceTag(val *string)
	ServiceTagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
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
	PutHeaders(value interface{})
	ResetAction()
	ResetHeaders()
	ResetIpAddress()
	ResetName()
	ResetPriority()
	ResetServiceTag()
	ResetVirtualNetworkSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigIpRestrictionOutputReference
type jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) Headers() AppServiceSlotSiteConfigIpRestrictionHeadersList {
	var returns AppServiceSlotSiteConfigIpRestrictionHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ServiceTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ServiceTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigIpRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppServiceSlotSiteConfigIpRestrictionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigIpRestrictionOutputReference_Override(a AppServiceSlotSiteConfigIpRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetServiceTag(val *string) {
	_jsii_.Set(
		j,
		"serviceTag",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) SetVirtualNetworkSubnetId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) PutHeaders(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		a,
		"resetAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ResetServiceTag() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigIpRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigOutputReference interface {
	cdktf.ComplexObject
	AcrUseManagedIdentityCredentials() interface{}
	SetAcrUseManagedIdentityCredentials(val interface{})
	AcrUseManagedIdentityCredentialsInput() interface{}
	AcrUserManagedIdentityClientId() *string
	SetAcrUserManagedIdentityClientId(val *string)
	AcrUserManagedIdentityClientIdInput() *string
	AlwaysOn() interface{}
	SetAlwaysOn(val interface{})
	AlwaysOnInput() interface{}
	AppCommandLine() *string
	SetAppCommandLine(val *string)
	AppCommandLineInput() *string
	AutoSwapSlotName() *string
	SetAutoSwapSlotName(val *string)
	AutoSwapSlotNameInput() *string
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
	Cors() AppServiceSlotSiteConfigCorsOutputReference
	CorsInput() *AppServiceSlotSiteConfigCors
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultDocuments() *[]*string
	SetDefaultDocuments(val *[]*string)
	DefaultDocumentsInput() *[]*string
	DotnetFrameworkVersion() *string
	SetDotnetFrameworkVersion(val *string)
	DotnetFrameworkVersionInput() *string
	// Experimental.
	Fqn() *string
	FtpsState() *string
	SetFtpsState(val *string)
	FtpsStateInput() *string
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	Http2Enabled() interface{}
	SetHttp2Enabled(val interface{})
	Http2EnabledInput() interface{}
	InternalValue() *AppServiceSlotSiteConfig
	SetInternalValue(val *AppServiceSlotSiteConfig)
	IpRestriction() AppServiceSlotSiteConfigIpRestrictionList
	IpRestrictionInput() interface{}
	JavaContainer() *string
	SetJavaContainer(val *string)
	JavaContainerInput() *string
	JavaContainerVersion() *string
	SetJavaContainerVersion(val *string)
	JavaContainerVersionInput() *string
	JavaVersion() *string
	SetJavaVersion(val *string)
	JavaVersionInput() *string
	LinuxFxVersion() *string
	SetLinuxFxVersion(val *string)
	LinuxFxVersionInput() *string
	LocalMysqlEnabled() interface{}
	SetLocalMysqlEnabled(val interface{})
	LocalMysqlEnabledInput() interface{}
	ManagedPipelineMode() *string
	SetManagedPipelineMode(val *string)
	ManagedPipelineModeInput() *string
	MinTlsVersion() *string
	SetMinTlsVersion(val *string)
	MinTlsVersionInput() *string
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	NumberOfWorkersInput() *float64
	PhpVersion() *string
	SetPhpVersion(val *string)
	PhpVersionInput() *string
	PythonVersion() *string
	SetPythonVersion(val *string)
	PythonVersionInput() *string
	RemoteDebuggingEnabled() interface{}
	SetRemoteDebuggingEnabled(val interface{})
	RemoteDebuggingEnabledInput() interface{}
	RemoteDebuggingVersion() *string
	SetRemoteDebuggingVersion(val *string)
	RemoteDebuggingVersionInput() *string
	ScmIpRestriction() AppServiceSlotSiteConfigScmIpRestrictionList
	ScmIpRestrictionInput() interface{}
	ScmType() *string
	SetScmType(val *string)
	ScmTypeInput() *string
	ScmUseMainIpRestriction() interface{}
	SetScmUseMainIpRestriction(val interface{})
	ScmUseMainIpRestrictionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Use32BitWorkerProcess() interface{}
	SetUse32BitWorkerProcess(val interface{})
	Use32BitWorkerProcessInput() interface{}
	VnetRouteAllEnabled() interface{}
	SetVnetRouteAllEnabled(val interface{})
	VnetRouteAllEnabledInput() interface{}
	WebsocketsEnabled() interface{}
	SetWebsocketsEnabled(val interface{})
	WebsocketsEnabledInput() interface{}
	WindowsFxVersion() *string
	SetWindowsFxVersion(val *string)
	WindowsFxVersionInput() *string
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
	PutCors(value *AppServiceSlotSiteConfigCors)
	PutIpRestriction(value interface{})
	PutScmIpRestriction(value interface{})
	ResetAcrUseManagedIdentityCredentials()
	ResetAcrUserManagedIdentityClientId()
	ResetAlwaysOn()
	ResetAppCommandLine()
	ResetAutoSwapSlotName()
	ResetCors()
	ResetDefaultDocuments()
	ResetDotnetFrameworkVersion()
	ResetFtpsState()
	ResetHealthCheckPath()
	ResetHttp2Enabled()
	ResetIpRestriction()
	ResetJavaContainer()
	ResetJavaContainerVersion()
	ResetJavaVersion()
	ResetLinuxFxVersion()
	ResetLocalMysqlEnabled()
	ResetManagedPipelineMode()
	ResetMinTlsVersion()
	ResetNumberOfWorkers()
	ResetPhpVersion()
	ResetPythonVersion()
	ResetRemoteDebuggingEnabled()
	ResetRemoteDebuggingVersion()
	ResetScmIpRestriction()
	ResetScmType()
	ResetScmUseMainIpRestriction()
	ResetUse32BitWorkerProcess()
	ResetVnetRouteAllEnabled()
	ResetWebsocketsEnabled()
	ResetWindowsFxVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigOutputReference
type jsiiProxy_AppServiceSlotSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AcrUseManagedIdentityCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrUseManagedIdentityCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AcrUseManagedIdentityCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrUseManagedIdentityCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AcrUserManagedIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acrUserManagedIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AcrUserManagedIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acrUserManagedIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AlwaysOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AlwaysOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AppCommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AppCommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AutoSwapSlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) AutoSwapSlotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) Cors() AppServiceSlotSiteConfigCorsOutputReference {
	var returns AppServiceSlotSiteConfigCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) CorsInput() *AppServiceSlotSiteConfigCors {
	var returns *AppServiceSlotSiteConfigCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) DefaultDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) DefaultDocumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) DotnetFrameworkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetFrameworkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) DotnetFrameworkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetFrameworkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) FtpsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) Http2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) Http2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) InternalValue() *AppServiceSlotSiteConfig {
	var returns *AppServiceSlotSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) IpRestriction() AppServiceSlotSiteConfigIpRestrictionList {
	var returns AppServiceSlotSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) IpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) JavaContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) JavaContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) JavaContainerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) JavaContainerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaContainerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) JavaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) LinuxFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) LinuxFxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) LocalMysqlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localMysqlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) LocalMysqlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localMysqlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ManagedPipelineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ManagedPipelineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) NumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) PhpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) PhpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) RemoteDebuggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) RemoteDebuggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebuggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) RemoteDebuggingVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) RemoteDebuggingVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ScmIpRestriction() AppServiceSlotSiteConfigScmIpRestrictionList {
	var returns AppServiceSlotSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ScmIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ScmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ScmUseMainIpRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ScmUseMainIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) Use32BitWorkerProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) Use32BitWorkerProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) VnetRouteAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) VnetRouteAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) WebsocketsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) WebsocketsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) WindowsFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) WindowsFxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsFxVersionInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotSiteConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigOutputReference_Override(a AppServiceSlotSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetAcrUseManagedIdentityCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"acrUseManagedIdentityCredentials",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetAcrUserManagedIdentityClientId(val *string) {
	_jsii_.Set(
		j,
		"acrUserManagedIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetAlwaysOn(val interface{}) {
	_jsii_.Set(
		j,
		"alwaysOn",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetAppCommandLine(val *string) {
	_jsii_.Set(
		j,
		"appCommandLine",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetAutoSwapSlotName(val *string) {
	_jsii_.Set(
		j,
		"autoSwapSlotName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetDefaultDocuments(val *[]*string) {
	_jsii_.Set(
		j,
		"defaultDocuments",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetDotnetFrameworkVersion(val *string) {
	_jsii_.Set(
		j,
		"dotnetFrameworkVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetFtpsState(val *string) {
	_jsii_.Set(
		j,
		"ftpsState",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetHttp2Enabled(val interface{}) {
	_jsii_.Set(
		j,
		"http2Enabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetInternalValue(val *AppServiceSlotSiteConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetJavaContainer(val *string) {
	_jsii_.Set(
		j,
		"javaContainer",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetJavaContainerVersion(val *string) {
	_jsii_.Set(
		j,
		"javaContainerVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetJavaVersion(val *string) {
	_jsii_.Set(
		j,
		"javaVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetLinuxFxVersion(val *string) {
	_jsii_.Set(
		j,
		"linuxFxVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetLocalMysqlEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"localMysqlEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetManagedPipelineMode(val *string) {
	_jsii_.Set(
		j,
		"managedPipelineMode",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetMinTlsVersion(val *string) {
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetPhpVersion(val *string) {
	_jsii_.Set(
		j,
		"phpVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetPythonVersion(val *string) {
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetRemoteDebuggingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"remoteDebuggingEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetRemoteDebuggingVersion(val *string) {
	_jsii_.Set(
		j,
		"remoteDebuggingVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetScmType(val *string) {
	_jsii_.Set(
		j,
		"scmType",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetScmUseMainIpRestriction(val interface{}) {
	_jsii_.Set(
		j,
		"scmUseMainIpRestriction",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetUse32BitWorkerProcess(val interface{}) {
	_jsii_.Set(
		j,
		"use32BitWorkerProcess",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetVnetRouteAllEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"vnetRouteAllEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetWebsocketsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"websocketsEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigOutputReference) SetWindowsFxVersion(val *string) {
	_jsii_.Set(
		j,
		"windowsFxVersion",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) PutCors(value *AppServiceSlotSiteConfigCors) {
	_jsii_.InvokeVoid(
		a,
		"putCors",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) PutIpRestriction(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putIpRestriction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) PutScmIpRestriction(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putScmIpRestriction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetAcrUseManagedIdentityCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetAcrUseManagedIdentityCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetAcrUserManagedIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetAcrUserManagedIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetAlwaysOn() {
	_jsii_.InvokeVoid(
		a,
		"resetAlwaysOn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetAppCommandLine() {
	_jsii_.InvokeVoid(
		a,
		"resetAppCommandLine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetAutoSwapSlotName() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSwapSlotName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		a,
		"resetCors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetDefaultDocuments() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultDocuments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetDotnetFrameworkVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDotnetFrameworkVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetFtpsState() {
	_jsii_.InvokeVoid(
		a,
		"resetFtpsState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetHttp2Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetIpRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetIpRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetJavaContainer() {
	_jsii_.InvokeVoid(
		a,
		"resetJavaContainer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetJavaContainerVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetJavaContainerVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetJavaVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetJavaVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetLinuxFxVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetLinuxFxVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetLocalMysqlEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalMysqlEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetManagedPipelineMode() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedPipelineMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetNumberOfWorkers() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberOfWorkers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetPhpVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPhpVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPythonVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetRemoteDebuggingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteDebuggingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetRemoteDebuggingVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteDebuggingVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetScmIpRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetScmIpRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetScmType() {
	_jsii_.InvokeVoid(
		a,
		"resetScmType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetScmUseMainIpRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetScmUseMainIpRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetUse32BitWorkerProcess() {
	_jsii_.InvokeVoid(
		a,
		"resetUse32BitWorkerProcess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetVnetRouteAllEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetVnetRouteAllEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetWebsocketsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetWebsocketsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ResetWindowsFxVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetWindowsFxVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigScmIpRestriction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#action AppServiceSlot#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#headers AppServiceSlot#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#ip_address AppServiceSlot#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#name AppServiceSlot#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#priority AppServiceSlot#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#service_tag AppServiceSlot#service_tag}.
	ServiceTag *string `field:"optional" json:"serviceTag" yaml:"serviceTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#virtual_network_subnet_id AppServiceSlot#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
}

type AppServiceSlotSiteConfigScmIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_azure_fdid AppServiceSlot#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_fd_health_probe AppServiceSlot#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_forwarded_for AppServiceSlot#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#x_forwarded_host AppServiceSlot#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}

type AppServiceSlotSiteConfigScmIpRestrictionHeadersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigScmIpRestrictionHeadersList
type jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigScmIpRestrictionHeadersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppServiceSlotSiteConfigScmIpRestrictionHeadersList {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigScmIpRestrictionHeadersList_Override(a AppServiceSlotSiteConfigScmIpRestrictionHeadersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) Get(index *float64) AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference {
	var returns AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XAzureFdid() *[]*string
	SetXAzureFdid(val *[]*string)
	XAzureFdidInput() *[]*string
	XFdHealthProbe() *[]*string
	SetXFdHealthProbe(val *[]*string)
	XFdHealthProbeInput() *[]*string
	XForwardedFor() *[]*string
	SetXForwardedFor(val *[]*string)
	XForwardedForInput() *[]*string
	XForwardedHost() *[]*string
	SetXForwardedHost(val *[]*string)
	XForwardedHostInput() *[]*string
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
	ResetXAzureFdid()
	ResetXFdHealthProbe()
	ResetXForwardedFor()
	ResetXForwardedHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference
type jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdid() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdidInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbe() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHostInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference_Override(a AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetXAzureFdid(val *[]*string) {
	_jsii_.Set(
		j,
		"xAzureFdid",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetXFdHealthProbe(val *[]*string) {
	_jsii_.Set(
		j,
		"xFdHealthProbe",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetXForwardedFor(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedFor",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) SetXForwardedHost(val *[]*string) {
	_jsii_.Set(
		j,
		"xForwardedHost",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXAzureFdid() {
	_jsii_.InvokeVoid(
		a,
		"resetXAzureFdid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXFdHealthProbe() {
	_jsii_.InvokeVoid(
		a,
		"resetXFdHealthProbe",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedFor() {
	_jsii_.InvokeVoid(
		a,
		"resetXForwardedFor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedHost() {
	_jsii_.InvokeVoid(
		a,
		"resetXForwardedHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigScmIpRestrictionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppServiceSlotSiteConfigScmIpRestrictionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigScmIpRestrictionList
type jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigScmIpRestrictionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppServiceSlotSiteConfigScmIpRestrictionList {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigScmIpRestrictionList_Override(a AppServiceSlotSiteConfigScmIpRestrictionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) Get(index *float64) AppServiceSlotSiteConfigScmIpRestrictionOutputReference {
	var returns AppServiceSlotSiteConfigScmIpRestrictionOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteConfigScmIpRestrictionOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Headers() AppServiceSlotSiteConfigScmIpRestrictionHeadersList
	HeadersInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ServiceTag() *string
	SetServiceTag(val *string)
	ServiceTagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkSubnetId() *string
	SetVirtualNetworkSubnetId(val *string)
	VirtualNetworkSubnetIdInput() *string
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
	PutHeaders(value interface{})
	ResetAction()
	ResetHeaders()
	ResetIpAddress()
	ResetName()
	ResetPriority()
	ResetServiceTag()
	ResetVirtualNetworkSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteConfigScmIpRestrictionOutputReference
type jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) Headers() AppServiceSlotSiteConfigScmIpRestrictionHeadersList {
	var returns AppServiceSlotSiteConfigScmIpRestrictionHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ServiceTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ServiceTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) VirtualNetworkSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) VirtualNetworkSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteConfigScmIpRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppServiceSlotSiteConfigScmIpRestrictionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteConfigScmIpRestrictionOutputReference_Override(a AppServiceSlotSiteConfigScmIpRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteConfigScmIpRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetServiceTag(val *string) {
	_jsii_.Set(
		j,
		"serviceTag",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) SetVirtualNetworkSubnetId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkSubnetId",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) PutHeaders(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		a,
		"resetAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ResetServiceTag() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ResetVirtualNetworkSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNetworkSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteConfigScmIpRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteCredential struct {
}

type AppServiceSlotSiteCredentialList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppServiceSlotSiteCredentialOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteCredentialList
type jsiiProxy_AppServiceSlotSiteCredentialList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteCredentialList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppServiceSlotSiteCredentialList {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteCredentialList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteCredentialList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteCredentialList_Override(a AppServiceSlotSiteCredentialList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteCredentialList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialList) Get(index *float64) AppServiceSlotSiteCredentialOutputReference {
	var returns AppServiceSlotSiteCredentialOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotSiteCredentialOutputReference interface {
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
	InternalValue() *AppServiceSlotSiteCredential
	SetInternalValue(val *AppServiceSlotSiteCredential)
	Password() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotSiteCredentialOutputReference
type jsiiProxy_AppServiceSlotSiteCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) InternalValue() *AppServiceSlotSiteCredential {
	var returns *AppServiceSlotSiteCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


func NewAppServiceSlotSiteCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppServiceSlotSiteCredentialOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotSiteCredentialOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotSiteCredentialOutputReference_Override(a AppServiceSlotSiteCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotSiteCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) SetInternalValue(val *AppServiceSlotSiteCredential) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotSiteCredentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#access_key AppServiceSlot#access_key}.
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#account_name AppServiceSlot#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#name AppServiceSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#share_name AppServiceSlot#share_name}.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#type AppServiceSlot#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#mount_path AppServiceSlot#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

type AppServiceSlotStorageAccountList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppServiceSlotStorageAccountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotStorageAccountList
type jsiiProxy_AppServiceSlotStorageAccountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppServiceSlotStorageAccountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppServiceSlotStorageAccountList {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotStorageAccountList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotStorageAccountList_Override(a AppServiceSlotStorageAccountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotStorageAccountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountList) Get(index *float64) AppServiceSlotStorageAccountOutputReference {
	var returns AppServiceSlotStorageAccountOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotStorageAccountOutputReference interface {
	cdktf.ComplexObject
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountPath() *string
	SetMountPath(val *string)
	MountPathInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetMountPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSlotStorageAccountOutputReference
type jsiiProxy_AppServiceSlotStorageAccountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) MountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) MountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotStorageAccountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppServiceSlotStorageAccountOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotStorageAccountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppServiceSlotStorageAccountOutputReference_Override(a AppServiceSlotStorageAccountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetMountPath(val *string) {
	_jsii_.Set(
		j,
		"mountPath",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetShareName(val *string) {
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotStorageAccountOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) ResetMountPath() {
	_jsii_.InvokeVoid(
		a,
		"resetMountPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotStorageAccountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppServiceSlotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#create AppServiceSlot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#delete AppServiceSlot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#read AppServiceSlot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service_slot#update AppServiceSlot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type AppServiceSlotTimeoutsOutputReference interface {
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

// The jsii proxy struct for AppServiceSlotTimeoutsOutputReference
type jsiiProxy_AppServiceSlotTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppServiceSlotTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotTimeoutsOutputReference_Override(a AppServiceSlotTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		a,
		"resetRead",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


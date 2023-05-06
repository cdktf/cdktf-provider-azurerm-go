package functionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/functionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/function_app azurerm_function_app}.
type FunctionApp interface {
	cdktf.TerraformResource
	AppServicePlanId() *string
	SetAppServicePlanId(val *string)
	AppServicePlanIdInput() *string
	AppSettings() *map[string]*string
	SetAppSettings(val *map[string]*string)
	AppSettingsInput() *map[string]*string
	AuthSettings() FunctionAppAuthSettingsOutputReference
	AuthSettingsInput() *FunctionAppAuthSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertMode() *string
	SetClientCertMode(val *string)
	ClientCertModeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionString() FunctionAppConnectionStringList
	ConnectionStringInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomDomainVerificationId() *string
	DailyMemoryTimeQuota() *float64
	SetDailyMemoryTimeQuota(val *float64)
	DailyMemoryTimeQuotaInput() *float64
	DefaultHostname() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableBuiltinLogging() interface{}
	SetEnableBuiltinLogging(val interface{})
	EnableBuiltinLoggingInput() interface{}
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
	Identity() FunctionAppIdentityOutputReference
	IdentityInput() *FunctionAppIdentity
	IdInput() *string
	KeyVaultReferenceIdentityId() *string
	SetKeyVaultReferenceIdentityId(val *string)
	KeyVaultReferenceIdentityIdInput() *string
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	OutboundIpAddresses() *string
	PossibleOutboundIpAddresses() *string
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
	SiteConfig() FunctionAppSiteConfigOutputReference
	SiteConfigInput() *FunctionAppSiteConfig
	SiteCredential() FunctionAppSiteCredentialList
	SourceControl() FunctionAppSourceControlOutputReference
	SourceControlInput() *FunctionAppSourceControl
	StorageAccountAccessKey() *string
	SetStorageAccountAccessKey(val *string)
	StorageAccountAccessKeyInput() *string
	StorageAccountName() *string
	SetStorageAccountName(val *string)
	StorageAccountNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FunctionAppTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutAuthSettings(value *FunctionAppAuthSettings)
	PutConnectionString(value interface{})
	PutIdentity(value *FunctionAppIdentity)
	PutSiteConfig(value *FunctionAppSiteConfig)
	PutSourceControl(value *FunctionAppSourceControl)
	PutTimeouts(value *FunctionAppTimeouts)
	ResetAppSettings()
	ResetAuthSettings()
	ResetClientCertMode()
	ResetConnectionString()
	ResetDailyMemoryTimeQuota()
	ResetEnableBuiltinLogging()
	ResetEnabled()
	ResetHttpsOnly()
	ResetId()
	ResetIdentity()
	ResetKeyVaultReferenceIdentityId()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSiteConfig()
	ResetSourceControl()
	ResetTags()
	ResetTimeouts()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FunctionApp
type jsiiProxy_FunctionApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FunctionApp) AppServicePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServicePlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) AppServicePlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServicePlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) AppSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) AppSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"appSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) AuthSettings() FunctionAppAuthSettingsOutputReference {
	var returns FunctionAppAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"authSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) AuthSettingsInput() *FunctionAppAuthSettings {
	var returns *FunctionAppAuthSettings
	_jsii_.Get(
		j,
		"authSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ClientCertMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ClientCertModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ConnectionString() FunctionAppConnectionStringList {
	var returns FunctionAppConnectionStringList
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ConnectionStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) CustomDomainVerificationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainVerificationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) DailyMemoryTimeQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) DailyMemoryTimeQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyMemoryTimeQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) DefaultHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) EnableBuiltinLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBuiltinLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) EnableBuiltinLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBuiltinLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Identity() FunctionAppIdentityOutputReference {
	var returns FunctionAppIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) IdentityInput() *FunctionAppIdentity {
	var returns *FunctionAppIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) KeyVaultReferenceIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) KeyVaultReferenceIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultReferenceIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) OutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) PossibleOutboundIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"possibleOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) SiteConfig() FunctionAppSiteConfigOutputReference {
	var returns FunctionAppSiteConfigOutputReference
	_jsii_.Get(
		j,
		"siteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) SiteConfigInput() *FunctionAppSiteConfig {
	var returns *FunctionAppSiteConfig
	_jsii_.Get(
		j,
		"siteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) SiteCredential() FunctionAppSiteCredentialList {
	var returns FunctionAppSiteCredentialList
	_jsii_.Get(
		j,
		"siteCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) SourceControl() FunctionAppSourceControlOutputReference {
	var returns FunctionAppSourceControlOutputReference
	_jsii_.Get(
		j,
		"sourceControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) SourceControlInput() *FunctionAppSourceControl {
	var returns *FunctionAppSourceControl
	_jsii_.Get(
		j,
		"sourceControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Timeouts() FunctionAppTimeoutsOutputReference {
	var returns FunctionAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionApp) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/function_app azurerm_function_app} Resource.
func NewFunctionApp(scope constructs.Construct, id *string, config *FunctionAppConfig) FunctionApp {
	_init_.Initialize()

	if err := validateNewFunctionAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionApp{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/function_app azurerm_function_app} Resource.
func NewFunctionApp_Override(f FunctionApp, scope constructs.Construct, id *string, config *FunctionAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionApp",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FunctionApp)SetAppServicePlanId(val *string) {
	if err := j.validateSetAppServicePlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appServicePlanId",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetAppSettings(val *map[string]*string) {
	if err := j.validateSetAppSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettings",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetClientCertMode(val *string) {
	if err := j.validateSetClientCertModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertMode",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetDailyMemoryTimeQuota(val *float64) {
	if err := j.validateSetDailyMemoryTimeQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyMemoryTimeQuota",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetEnableBuiltinLogging(val interface{}) {
	if err := j.validateSetEnableBuiltinLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBuiltinLogging",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetKeyVaultReferenceIdentityId(val *string) {
	if err := j.validateSetKeyVaultReferenceIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultReferenceIdentityId",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetStorageAccountName(val *string) {
	if err := j.validateSetStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FunctionApp)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
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
func FunctionApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.functionApp.FunctionApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.functionApp.FunctionApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.functionApp.FunctionApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FunctionApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.functionApp.FunctionApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FunctionApp) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FunctionApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FunctionApp) PutAuthSettings(value *FunctionAppAuthSettings) {
	if err := f.validatePutAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAuthSettings",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionApp) PutConnectionString(value interface{}) {
	if err := f.validatePutConnectionStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putConnectionString",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionApp) PutIdentity(value *FunctionAppIdentity) {
	if err := f.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIdentity",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionApp) PutSiteConfig(value *FunctionAppSiteConfig) {
	if err := f.validatePutSiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSiteConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionApp) PutSourceControl(value *FunctionAppSourceControl) {
	if err := f.validatePutSourceControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSourceControl",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionApp) PutTimeouts(value *FunctionAppTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionApp) ResetAppSettings() {
	_jsii_.InvokeVoid(
		f,
		"resetAppSettings",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetAuthSettings() {
	_jsii_.InvokeVoid(
		f,
		"resetAuthSettings",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetClientCertMode() {
	_jsii_.InvokeVoid(
		f,
		"resetClientCertMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetConnectionString() {
	_jsii_.InvokeVoid(
		f,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetDailyMemoryTimeQuota() {
	_jsii_.InvokeVoid(
		f,
		"resetDailyMemoryTimeQuota",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetEnableBuiltinLogging() {
	_jsii_.InvokeVoid(
		f,
		"resetEnableBuiltinLogging",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		f,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetIdentity() {
	_jsii_.InvokeVoid(
		f,
		"resetIdentity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetKeyVaultReferenceIdentityId() {
	_jsii_.InvokeVoid(
		f,
		"resetKeyVaultReferenceIdentityId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetOsType() {
	_jsii_.InvokeVoid(
		f,
		"resetOsType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetSiteConfig() {
	_jsii_.InvokeVoid(
		f,
		"resetSiteConfig",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetSourceControl() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceControl",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) ResetVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


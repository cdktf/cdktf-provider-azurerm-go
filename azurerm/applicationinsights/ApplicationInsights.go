package applicationinsights

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/applicationinsights/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/application_insights azurerm_application_insights}.
type ApplicationInsights interface {
	cdktf.TerraformResource
	AppId() *string
	ApplicationType() *string
	SetApplicationType(val *string)
	ApplicationTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionString() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DailyDataCapInGb() *float64
	SetDailyDataCapInGb(val *float64)
	DailyDataCapInGbInput() *float64
	DailyDataCapNotificationsDisabled() interface{}
	SetDailyDataCapNotificationsDisabled(val interface{})
	DailyDataCapNotificationsDisabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableIpMasking() interface{}
	SetDisableIpMasking(val interface{})
	DisableIpMaskingInput() interface{}
	ForceCustomerStorageForProfiler() interface{}
	SetForceCustomerStorageForProfiler(val interface{})
	ForceCustomerStorageForProfilerInput() interface{}
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
	InstrumentationKey() *string
	InternetIngestionEnabled() interface{}
	SetInternetIngestionEnabled(val interface{})
	InternetIngestionEnabledInput() interface{}
	InternetQueryEnabled() interface{}
	SetInternetQueryEnabled(val interface{})
	InternetQueryEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthenticationDisabled() interface{}
	SetLocalAuthenticationDisabled(val interface{})
	LocalAuthenticationDisabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	SamplingPercentage() *float64
	SetSamplingPercentage(val *float64)
	SamplingPercentageInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApplicationInsightsTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutTimeouts(value *ApplicationInsightsTimeouts)
	ResetDailyDataCapInGb()
	ResetDailyDataCapNotificationsDisabled()
	ResetDisableIpMasking()
	ResetForceCustomerStorageForProfiler()
	ResetId()
	ResetInternetIngestionEnabled()
	ResetInternetQueryEnabled()
	ResetLocalAuthenticationDisabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetentionInDays()
	ResetSamplingPercentage()
	ResetTags()
	ResetTimeouts()
	ResetWorkspaceId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApplicationInsights
type jsiiProxy_ApplicationInsights struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationInsights) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ApplicationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ApplicationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) DailyDataCapInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyDataCapInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) DailyDataCapInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyDataCapInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) DailyDataCapNotificationsDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyDataCapNotificationsDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) DailyDataCapNotificationsDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyDataCapNotificationsDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) DisableIpMasking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIpMasking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) DisableIpMaskingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIpMaskingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ForceCustomerStorageForProfiler() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceCustomerStorageForProfiler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ForceCustomerStorageForProfilerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceCustomerStorageForProfilerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) InstrumentationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instrumentationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) InternetIngestionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetIngestionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) InternetIngestionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetIngestionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) InternetQueryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetQueryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) InternetQueryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetQueryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) LocalAuthenticationDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) LocalAuthenticationDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) SamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) SamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) Timeouts() ApplicationInsightsTimeoutsOutputReference {
	var returns ApplicationInsightsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationInsights) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/application_insights azurerm_application_insights} Resource.
func NewApplicationInsights(scope constructs.Construct, id *string, config *ApplicationInsightsConfig) ApplicationInsights {
	_init_.Initialize()

	if err := validateNewApplicationInsightsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationInsights{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationInsights.ApplicationInsights",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/application_insights azurerm_application_insights} Resource.
func NewApplicationInsights_Override(a ApplicationInsights, scope constructs.Construct, id *string, config *ApplicationInsightsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationInsights.ApplicationInsights",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetApplicationType(val *string) {
	if err := j.validateSetApplicationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationType",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetDailyDataCapInGb(val *float64) {
	if err := j.validateSetDailyDataCapInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyDataCapInGb",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetDailyDataCapNotificationsDisabled(val interface{}) {
	if err := j.validateSetDailyDataCapNotificationsDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyDataCapNotificationsDisabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetDisableIpMasking(val interface{}) {
	if err := j.validateSetDisableIpMaskingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableIpMasking",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetForceCustomerStorageForProfiler(val interface{}) {
	if err := j.validateSetForceCustomerStorageForProfilerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceCustomerStorageForProfiler",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetInternetIngestionEnabled(val interface{}) {
	if err := j.validateSetInternetIngestionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetIngestionEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetInternetQueryEnabled(val interface{}) {
	if err := j.validateSetInternetQueryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetQueryEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetLocalAuthenticationDisabled(val interface{}) {
	if err := j.validateSetLocalAuthenticationDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthenticationDisabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetSamplingPercentage(val *float64) {
	if err := j.validateSetSamplingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingPercentage",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApplicationInsights)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
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
func ApplicationInsights_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationInsights_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.applicationInsights.ApplicationInsights",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationInsights_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationInsights_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.applicationInsights.ApplicationInsights",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationInsights_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationInsights_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.applicationInsights.ApplicationInsights",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationInsights_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.applicationInsights.ApplicationInsights",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationInsights) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationInsights) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationInsights) PutTimeouts(value *ApplicationInsightsTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetDailyDataCapInGb() {
	_jsii_.InvokeVoid(
		a,
		"resetDailyDataCapInGb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetDailyDataCapNotificationsDisabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDailyDataCapNotificationsDisabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetDisableIpMasking() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableIpMasking",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetForceCustomerStorageForProfiler() {
	_jsii_.InvokeVoid(
		a,
		"resetForceCustomerStorageForProfiler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetInternetIngestionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetInternetIngestionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetInternetQueryEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetInternetQueryEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetLocalAuthenticationDisabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalAuthenticationDisabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetSamplingPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetSamplingPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationInsights) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationInsights) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


package iotsecuritysolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/iotsecuritysolution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution azurerm_iot_security_solution}.
type IotSecuritySolution interface {
	cdktf.TerraformResource
	AdditionalWorkspace() IotSecuritySolutionAdditionalWorkspaceList
	AdditionalWorkspaceInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisabledDataSources() *[]*string
	SetDisabledDataSources(val *[]*string)
	DisabledDataSourcesInput() *[]*string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventsToExport() *[]*string
	SetEventsToExport(val *[]*string)
	EventsToExportInput() *[]*string
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
	IothubIds() *[]*string
	SetIothubIds(val *[]*string)
	IothubIdsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	LogUnmaskedIpsEnabled() interface{}
	SetLogUnmaskedIpsEnabled(val interface{})
	LogUnmaskedIpsEnabledInput() interface{}
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
	QueryForResources() *string
	SetQueryForResources(val *string)
	QueryForResourcesInput() *string
	QuerySubscriptionIds() *[]*string
	SetQuerySubscriptionIds(val *[]*string)
	QuerySubscriptionIdsInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RecommendationsEnabled() IotSecuritySolutionRecommendationsEnabledOutputReference
	RecommendationsEnabledInput() *IotSecuritySolutionRecommendationsEnabled
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IotSecuritySolutionTimeoutsOutputReference
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
	PutAdditionalWorkspace(value interface{})
	PutRecommendationsEnabled(value *IotSecuritySolutionRecommendationsEnabled)
	PutTimeouts(value *IotSecuritySolutionTimeouts)
	ResetAdditionalWorkspace()
	ResetDisabledDataSources()
	ResetEnabled()
	ResetEventsToExport()
	ResetId()
	ResetLogAnalyticsWorkspaceId()
	ResetLogUnmaskedIpsEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryForResources()
	ResetQuerySubscriptionIds()
	ResetRecommendationsEnabled()
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

// The jsii proxy struct for IotSecuritySolution
type jsiiProxy_IotSecuritySolution struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotSecuritySolution) AdditionalWorkspace() IotSecuritySolutionAdditionalWorkspaceList {
	var returns IotSecuritySolutionAdditionalWorkspaceList
	_jsii_.Get(
		j,
		"additionalWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) AdditionalWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) DisabledDataSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) DisabledDataSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) EventsToExport() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsToExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) EventsToExportInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsToExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) IothubIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iothubIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) IothubIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iothubIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) LogUnmaskedIpsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logUnmaskedIpsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) LogUnmaskedIpsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logUnmaskedIpsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) QueryForResources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryForResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) QueryForResourcesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryForResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) QuerySubscriptionIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"querySubscriptionIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) QuerySubscriptionIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"querySubscriptionIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) RecommendationsEnabled() IotSecuritySolutionRecommendationsEnabledOutputReference {
	var returns IotSecuritySolutionRecommendationsEnabledOutputReference
	_jsii_.Get(
		j,
		"recommendationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) RecommendationsEnabledInput() *IotSecuritySolutionRecommendationsEnabled {
	var returns *IotSecuritySolutionRecommendationsEnabled
	_jsii_.Get(
		j,
		"recommendationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) Timeouts() IotSecuritySolutionTimeoutsOutputReference {
	var returns IotSecuritySolutionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolution) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution azurerm_iot_security_solution} Resource.
func NewIotSecuritySolution(scope constructs.Construct, id *string, config *IotSecuritySolutionConfig) IotSecuritySolution {
	_init_.Initialize()

	j := jsiiProxy_IotSecuritySolution{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution azurerm_iot_security_solution} Resource.
func NewIotSecuritySolution_Override(i IotSecuritySolution, scope constructs.Construct, id *string, config *IotSecuritySolutionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetDisabledDataSources(val *[]*string) {
	_jsii_.Set(
		j,
		"disabledDataSources",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetEventsToExport(val *[]*string) {
	_jsii_.Set(
		j,
		"eventsToExport",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetIothubIds(val *[]*string) {
	_jsii_.Set(
		j,
		"iothubIds",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetLogUnmaskedIpsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"logUnmaskedIpsEnabled",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetQueryForResources(val *string) {
	_jsii_.Set(
		j,
		"queryForResources",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetQuerySubscriptionIds(val *[]*string) {
	_jsii_.Set(
		j,
		"querySubscriptionIds",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution) SetTags(val *map[string]*string) {
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
func IotSecuritySolution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotSecuritySolution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotSecuritySolution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotSecuritySolution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotSecuritySolution) PutAdditionalWorkspace(value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"putAdditionalWorkspace",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecuritySolution) PutRecommendationsEnabled(value *IotSecuritySolutionRecommendationsEnabled) {
	_jsii_.InvokeVoid(
		i,
		"putRecommendationsEnabled",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecuritySolution) PutTimeouts(value *IotSecuritySolutionTimeouts) {
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetAdditionalWorkspace() {
	_jsii_.InvokeVoid(
		i,
		"resetAdditionalWorkspace",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetDisabledDataSources() {
	_jsii_.InvokeVoid(
		i,
		"resetDisabledDataSources",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetEventsToExport() {
	_jsii_.InvokeVoid(
		i,
		"resetEventsToExport",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetLogAnalyticsWorkspaceId() {
	_jsii_.InvokeVoid(
		i,
		"resetLogAnalyticsWorkspaceId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetLogUnmaskedIpsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetLogUnmaskedIpsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetQueryForResources() {
	_jsii_.InvokeVoid(
		i,
		"resetQueryForResources",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetQuerySubscriptionIds() {
	_jsii_.InvokeVoid(
		i,
		"resetQuerySubscriptionIds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetRecommendationsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetRecommendationsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotSecuritySolutionAdditionalWorkspace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#data_types IotSecuritySolution#data_types}.
	DataTypes *[]*string `field:"required" json:"dataTypes" yaml:"dataTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#workspace_id IotSecuritySolution#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

type IotSecuritySolutionAdditionalWorkspaceList interface {
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
	Get(index *float64) IotSecuritySolutionAdditionalWorkspaceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotSecuritySolutionAdditionalWorkspaceList
type jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewIotSecuritySolutionAdditionalWorkspaceList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) IotSecuritySolutionAdditionalWorkspaceList {
	_init_.Initialize()

	j := jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionAdditionalWorkspaceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewIotSecuritySolutionAdditionalWorkspaceList_Override(i IotSecuritySolutionAdditionalWorkspaceList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionAdditionalWorkspaceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) Get(index *float64) IotSecuritySolutionAdditionalWorkspaceOutputReference {
	var returns IotSecuritySolutionAdditionalWorkspaceOutputReference

	_jsii_.Invoke(
		i,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotSecuritySolutionAdditionalWorkspaceOutputReference interface {
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
	DataTypes() *[]*string
	SetDataTypes(val *[]*string)
	DataTypesInput() *[]*string
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
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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

// The jsii proxy struct for IotSecuritySolutionAdditionalWorkspaceOutputReference
type jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) DataTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) DataTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


func NewIotSecuritySolutionAdditionalWorkspaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotSecuritySolutionAdditionalWorkspaceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionAdditionalWorkspaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotSecuritySolutionAdditionalWorkspaceOutputReference_Override(i IotSecuritySolutionAdditionalWorkspaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionAdditionalWorkspaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) SetDataTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"dataTypes",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) SetWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionAdditionalWorkspaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotSecuritySolutionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#display_name IotSecuritySolution#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#iothub_ids IotSecuritySolution#iothub_ids}.
	IothubIds *[]*string `field:"required" json:"iothubIds" yaml:"iothubIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#location IotSecuritySolution#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#name IotSecuritySolution#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#resource_group_name IotSecuritySolution#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// additional_workspace block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#additional_workspace IotSecuritySolution#additional_workspace}
	AdditionalWorkspace interface{} `field:"optional" json:"additionalWorkspace" yaml:"additionalWorkspace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#disabled_data_sources IotSecuritySolution#disabled_data_sources}.
	DisabledDataSources *[]*string `field:"optional" json:"disabledDataSources" yaml:"disabledDataSources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#enabled IotSecuritySolution#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#events_to_export IotSecuritySolution#events_to_export}.
	EventsToExport *[]*string `field:"optional" json:"eventsToExport" yaml:"eventsToExport"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#id IotSecuritySolution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#log_analytics_workspace_id IotSecuritySolution#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"optional" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#log_unmasked_ips_enabled IotSecuritySolution#log_unmasked_ips_enabled}.
	LogUnmaskedIpsEnabled interface{} `field:"optional" json:"logUnmaskedIpsEnabled" yaml:"logUnmaskedIpsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#query_for_resources IotSecuritySolution#query_for_resources}.
	QueryForResources *string `field:"optional" json:"queryForResources" yaml:"queryForResources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#query_subscription_ids IotSecuritySolution#query_subscription_ids}.
	QuerySubscriptionIds *[]*string `field:"optional" json:"querySubscriptionIds" yaml:"querySubscriptionIds"`
	// recommendations_enabled block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#recommendations_enabled IotSecuritySolution#recommendations_enabled}
	RecommendationsEnabled *IotSecuritySolutionRecommendationsEnabled `field:"optional" json:"recommendationsEnabled" yaml:"recommendationsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#tags IotSecuritySolution#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#timeouts IotSecuritySolution#timeouts}
	Timeouts *IotSecuritySolutionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type IotSecuritySolutionRecommendationsEnabled struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#acr_authentication IotSecuritySolution#acr_authentication}.
	AcrAuthentication interface{} `field:"optional" json:"acrAuthentication" yaml:"acrAuthentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#agent_send_unutilized_msg IotSecuritySolution#agent_send_unutilized_msg}.
	AgentSendUnutilizedMsg interface{} `field:"optional" json:"agentSendUnutilizedMsg" yaml:"agentSendUnutilizedMsg"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#baseline IotSecuritySolution#baseline}.
	Baseline interface{} `field:"optional" json:"baseline" yaml:"baseline"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#edge_hub_mem_optimize IotSecuritySolution#edge_hub_mem_optimize}.
	EdgeHubMemOptimize interface{} `field:"optional" json:"edgeHubMemOptimize" yaml:"edgeHubMemOptimize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#edge_logging_option IotSecuritySolution#edge_logging_option}.
	EdgeLoggingOption interface{} `field:"optional" json:"edgeLoggingOption" yaml:"edgeLoggingOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#inconsistent_module_settings IotSecuritySolution#inconsistent_module_settings}.
	InconsistentModuleSettings interface{} `field:"optional" json:"inconsistentModuleSettings" yaml:"inconsistentModuleSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#install_agent IotSecuritySolution#install_agent}.
	InstallAgent interface{} `field:"optional" json:"installAgent" yaml:"installAgent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#ip_filter_deny_all IotSecuritySolution#ip_filter_deny_all}.
	IpFilterDenyAll interface{} `field:"optional" json:"ipFilterDenyAll" yaml:"ipFilterDenyAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#ip_filter_permissive_rule IotSecuritySolution#ip_filter_permissive_rule}.
	IpFilterPermissiveRule interface{} `field:"optional" json:"ipFilterPermissiveRule" yaml:"ipFilterPermissiveRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#open_ports IotSecuritySolution#open_ports}.
	OpenPorts interface{} `field:"optional" json:"openPorts" yaml:"openPorts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#permissive_firewall_policy IotSecuritySolution#permissive_firewall_policy}.
	PermissiveFirewallPolicy interface{} `field:"optional" json:"permissiveFirewallPolicy" yaml:"permissiveFirewallPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#permissive_input_firewall_rules IotSecuritySolution#permissive_input_firewall_rules}.
	PermissiveInputFirewallRules interface{} `field:"optional" json:"permissiveInputFirewallRules" yaml:"permissiveInputFirewallRules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#permissive_output_firewall_rules IotSecuritySolution#permissive_output_firewall_rules}.
	PermissiveOutputFirewallRules interface{} `field:"optional" json:"permissiveOutputFirewallRules" yaml:"permissiveOutputFirewallRules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#privileged_docker_options IotSecuritySolution#privileged_docker_options}.
	PrivilegedDockerOptions interface{} `field:"optional" json:"privilegedDockerOptions" yaml:"privilegedDockerOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#shared_credentials IotSecuritySolution#shared_credentials}.
	SharedCredentials interface{} `field:"optional" json:"sharedCredentials" yaml:"sharedCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#vulnerable_tls_cipher_suite IotSecuritySolution#vulnerable_tls_cipher_suite}.
	VulnerableTlsCipherSuite interface{} `field:"optional" json:"vulnerableTlsCipherSuite" yaml:"vulnerableTlsCipherSuite"`
}

type IotSecuritySolutionRecommendationsEnabledOutputReference interface {
	cdktf.ComplexObject
	AcrAuthentication() interface{}
	SetAcrAuthentication(val interface{})
	AcrAuthenticationInput() interface{}
	AgentSendUnutilizedMsg() interface{}
	SetAgentSendUnutilizedMsg(val interface{})
	AgentSendUnutilizedMsgInput() interface{}
	Baseline() interface{}
	SetBaseline(val interface{})
	BaselineInput() interface{}
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
	EdgeHubMemOptimize() interface{}
	SetEdgeHubMemOptimize(val interface{})
	EdgeHubMemOptimizeInput() interface{}
	EdgeLoggingOption() interface{}
	SetEdgeLoggingOption(val interface{})
	EdgeLoggingOptionInput() interface{}
	// Experimental.
	Fqn() *string
	InconsistentModuleSettings() interface{}
	SetInconsistentModuleSettings(val interface{})
	InconsistentModuleSettingsInput() interface{}
	InstallAgent() interface{}
	SetInstallAgent(val interface{})
	InstallAgentInput() interface{}
	InternalValue() *IotSecuritySolutionRecommendationsEnabled
	SetInternalValue(val *IotSecuritySolutionRecommendationsEnabled)
	IpFilterDenyAll() interface{}
	SetIpFilterDenyAll(val interface{})
	IpFilterDenyAllInput() interface{}
	IpFilterPermissiveRule() interface{}
	SetIpFilterPermissiveRule(val interface{})
	IpFilterPermissiveRuleInput() interface{}
	OpenPorts() interface{}
	SetOpenPorts(val interface{})
	OpenPortsInput() interface{}
	PermissiveFirewallPolicy() interface{}
	SetPermissiveFirewallPolicy(val interface{})
	PermissiveFirewallPolicyInput() interface{}
	PermissiveInputFirewallRules() interface{}
	SetPermissiveInputFirewallRules(val interface{})
	PermissiveInputFirewallRulesInput() interface{}
	PermissiveOutputFirewallRules() interface{}
	SetPermissiveOutputFirewallRules(val interface{})
	PermissiveOutputFirewallRulesInput() interface{}
	PrivilegedDockerOptions() interface{}
	SetPrivilegedDockerOptions(val interface{})
	PrivilegedDockerOptionsInput() interface{}
	SharedCredentials() interface{}
	SetSharedCredentials(val interface{})
	SharedCredentialsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VulnerableTlsCipherSuite() interface{}
	SetVulnerableTlsCipherSuite(val interface{})
	VulnerableTlsCipherSuiteInput() interface{}
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
	ResetAcrAuthentication()
	ResetAgentSendUnutilizedMsg()
	ResetBaseline()
	ResetEdgeHubMemOptimize()
	ResetEdgeLoggingOption()
	ResetInconsistentModuleSettings()
	ResetInstallAgent()
	ResetIpFilterDenyAll()
	ResetIpFilterPermissiveRule()
	ResetOpenPorts()
	ResetPermissiveFirewallPolicy()
	ResetPermissiveInputFirewallRules()
	ResetPermissiveOutputFirewallRules()
	ResetPrivilegedDockerOptions()
	ResetSharedCredentials()
	ResetVulnerableTlsCipherSuite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotSecuritySolutionRecommendationsEnabledOutputReference
type jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AcrAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AcrAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AgentSendUnutilizedMsg() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentSendUnutilizedMsg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AgentSendUnutilizedMsgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentSendUnutilizedMsgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) Baseline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) BaselineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeHubMemOptimize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeHubMemOptimize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeHubMemOptimizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeHubMemOptimizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeLoggingOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeLoggingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeLoggingOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeLoggingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InconsistentModuleSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inconsistentModuleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InconsistentModuleSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inconsistentModuleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InstallAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InstallAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InternalValue() *IotSecuritySolutionRecommendationsEnabled {
	var returns *IotSecuritySolutionRecommendationsEnabled
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterDenyAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterDenyAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterDenyAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterDenyAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterPermissiveRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterPermissiveRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterPermissiveRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterPermissiveRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) OpenPorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) OpenPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveFirewallPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveFirewallPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveFirewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveInputFirewallRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveInputFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveInputFirewallRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveInputFirewallRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveOutputFirewallRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveOutputFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveOutputFirewallRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveOutputFirewallRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PrivilegedDockerOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedDockerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PrivilegedDockerOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedDockerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SharedCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SharedCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) VulnerableTlsCipherSuite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerableTlsCipherSuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) VulnerableTlsCipherSuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerableTlsCipherSuiteInput",
		&returns,
	)
	return returns
}


func NewIotSecuritySolutionRecommendationsEnabledOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotSecuritySolutionRecommendationsEnabledOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotSecuritySolutionRecommendationsEnabledOutputReference_Override(i IotSecuritySolutionRecommendationsEnabledOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetAcrAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"acrAuthentication",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetAgentSendUnutilizedMsg(val interface{}) {
	_jsii_.Set(
		j,
		"agentSendUnutilizedMsg",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetBaseline(val interface{}) {
	_jsii_.Set(
		j,
		"baseline",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetEdgeHubMemOptimize(val interface{}) {
	_jsii_.Set(
		j,
		"edgeHubMemOptimize",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetEdgeLoggingOption(val interface{}) {
	_jsii_.Set(
		j,
		"edgeLoggingOption",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetInconsistentModuleSettings(val interface{}) {
	_jsii_.Set(
		j,
		"inconsistentModuleSettings",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetInstallAgent(val interface{}) {
	_jsii_.Set(
		j,
		"installAgent",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetInternalValue(val *IotSecuritySolutionRecommendationsEnabled) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetIpFilterDenyAll(val interface{}) {
	_jsii_.Set(
		j,
		"ipFilterDenyAll",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetIpFilterPermissiveRule(val interface{}) {
	_jsii_.Set(
		j,
		"ipFilterPermissiveRule",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetOpenPorts(val interface{}) {
	_jsii_.Set(
		j,
		"openPorts",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetPermissiveFirewallPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"permissiveFirewallPolicy",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetPermissiveInputFirewallRules(val interface{}) {
	_jsii_.Set(
		j,
		"permissiveInputFirewallRules",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetPermissiveOutputFirewallRules(val interface{}) {
	_jsii_.Set(
		j,
		"permissiveOutputFirewallRules",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetPrivilegedDockerOptions(val interface{}) {
	_jsii_.Set(
		j,
		"privilegedDockerOptions",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetSharedCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"sharedCredentials",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SetVulnerableTlsCipherSuite(val interface{}) {
	_jsii_.Set(
		j,
		"vulnerableTlsCipherSuite",
		val,
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetAcrAuthentication() {
	_jsii_.InvokeVoid(
		i,
		"resetAcrAuthentication",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetAgentSendUnutilizedMsg() {
	_jsii_.InvokeVoid(
		i,
		"resetAgentSendUnutilizedMsg",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetBaseline() {
	_jsii_.InvokeVoid(
		i,
		"resetBaseline",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetEdgeHubMemOptimize() {
	_jsii_.InvokeVoid(
		i,
		"resetEdgeHubMemOptimize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetEdgeLoggingOption() {
	_jsii_.InvokeVoid(
		i,
		"resetEdgeLoggingOption",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetInconsistentModuleSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetInconsistentModuleSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetInstallAgent() {
	_jsii_.InvokeVoid(
		i,
		"resetInstallAgent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetIpFilterDenyAll() {
	_jsii_.InvokeVoid(
		i,
		"resetIpFilterDenyAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetIpFilterPermissiveRule() {
	_jsii_.InvokeVoid(
		i,
		"resetIpFilterPermissiveRule",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetOpenPorts() {
	_jsii_.InvokeVoid(
		i,
		"resetOpenPorts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPermissiveFirewallPolicy() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveFirewallPolicy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPermissiveInputFirewallRules() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveInputFirewallRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPermissiveOutputFirewallRules() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveOutputFirewallRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPrivilegedDockerOptions() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivilegedDockerOptions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetSharedCredentials() {
	_jsii_.InvokeVoid(
		i,
		"resetSharedCredentials",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetVulnerableTlsCipherSuite() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerableTlsCipherSuite",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotSecuritySolutionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#create IotSecuritySolution#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#delete IotSecuritySolution#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#read IotSecuritySolution#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_solution#update IotSecuritySolution#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type IotSecuritySolutionTimeoutsOutputReference interface {
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

// The jsii proxy struct for IotSecuritySolutionTimeoutsOutputReference
type jsiiProxy_IotSecuritySolutionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewIotSecuritySolutionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotSecuritySolutionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_IotSecuritySolutionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotSecuritySolutionTimeoutsOutputReference_Override(i IotSecuritySolutionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		i,
		"resetCreate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		i,
		"resetDelete",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		i,
		"resetRead",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


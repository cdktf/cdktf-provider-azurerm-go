// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritysolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/iotsecuritysolution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/iot_security_solution azurerm_iot_security_solution}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

func (j *jsiiProxy_IotSecuritySolution) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/iot_security_solution azurerm_iot_security_solution} Resource.
func NewIotSecuritySolution(scope constructs.Construct, id *string, config *IotSecuritySolutionConfig) IotSecuritySolution {
	_init_.Initialize()

	if err := validateNewIotSecuritySolutionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecuritySolution{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/iot_security_solution azurerm_iot_security_solution} Resource.
func NewIotSecuritySolution_Override(i IotSecuritySolution, scope constructs.Construct, id *string, config *IotSecuritySolutionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetDisabledDataSources(val *[]*string) {
	if err := j.validateSetDisabledDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledDataSources",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetEventsToExport(val *[]*string) {
	if err := j.validateSetEventsToExportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsToExport",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetIothubIds(val *[]*string) {
	if err := j.validateSetIothubIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iothubIds",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetLogUnmaskedIpsEnabled(val interface{}) {
	if err := j.validateSetLogUnmaskedIpsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logUnmaskedIpsEnabled",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetQueryForResources(val *string) {
	if err := j.validateSetQueryForResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryForResources",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetQuerySubscriptionIds(val *[]*string) {
	if err := j.validateSetQuerySubscriptionIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"querySubscriptionIds",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a IotSecuritySolution resource upon running "cdktf plan <stack-name>".
func IotSecuritySolution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotSecuritySolution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
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
func IotSecuritySolution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotSecuritySolution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotSecuritySolution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotSecuritySolution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotSecuritySolution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotSecuritySolution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolution",
		"isTerraformResource",
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

func (i *jsiiProxy_IotSecuritySolution) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotSecuritySolution) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotSecuritySolution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotSecuritySolution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotSecuritySolution) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotSecuritySolution) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotSecuritySolution) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotSecuritySolution) PutAdditionalWorkspace(value interface{}) {
	if err := i.validatePutAdditionalWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAdditionalWorkspace",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecuritySolution) PutRecommendationsEnabled(value *IotSecuritySolutionRecommendationsEnabled) {
	if err := i.validatePutRecommendationsEnabledParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRecommendationsEnabled",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecuritySolution) PutTimeouts(value *IotSecuritySolutionTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (i *jsiiProxy_IotSecuritySolution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
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


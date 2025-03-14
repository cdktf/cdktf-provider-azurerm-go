// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegrouptemplatedeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/resourcegrouptemplatedeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/resource_group_template_deployment azurerm_resource_group_template_deployment}.
type ResourceGroupTemplateDeployment interface {
	cdktf.TerraformResource
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
	DebugLevel() *string
	SetDebugLevel(val *string)
	DebugLevelInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentMode() *string
	SetDeploymentMode(val *string)
	DeploymentModeInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputContent() *string
	ParametersContent() *string
	SetParametersContent(val *string)
	ParametersContentInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TemplateContent() *string
	SetTemplateContent(val *string)
	TemplateContentInput() *string
	TemplateSpecVersionId() *string
	SetTemplateSpecVersionId(val *string)
	TemplateSpecVersionIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ResourceGroupTemplateDeploymentTimeoutsOutputReference
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
	PutTimeouts(value *ResourceGroupTemplateDeploymentTimeouts)
	ResetDebugLevel()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParametersContent()
	ResetTags()
	ResetTemplateContent()
	ResetTemplateSpecVersionId()
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

// The jsii proxy struct for ResourceGroupTemplateDeployment
type jsiiProxy_ResourceGroupTemplateDeployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) DebugLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"debugLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) DebugLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"debugLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) DeploymentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) DeploymentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) OutputContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) ParametersContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) ParametersContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TemplateContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TemplateContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TemplateSpecVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSpecVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TemplateSpecVersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSpecVersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) Timeouts() ResourceGroupTemplateDeploymentTimeoutsOutputReference {
	var returns ResourceGroupTemplateDeploymentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/resource_group_template_deployment azurerm_resource_group_template_deployment} Resource.
func NewResourceGroupTemplateDeployment(scope constructs.Construct, id *string, config *ResourceGroupTemplateDeploymentConfig) ResourceGroupTemplateDeployment {
	_init_.Initialize()

	if err := validateNewResourceGroupTemplateDeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceGroupTemplateDeployment{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupTemplateDeployment.ResourceGroupTemplateDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/resource_group_template_deployment azurerm_resource_group_template_deployment} Resource.
func NewResourceGroupTemplateDeployment_Override(r ResourceGroupTemplateDeployment, scope constructs.Construct, id *string, config *ResourceGroupTemplateDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupTemplateDeployment.ResourceGroupTemplateDeployment",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetDebugLevel(val *string) {
	if err := j.validateSetDebugLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debugLevel",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetDeploymentMode(val *string) {
	if err := j.validateSetDeploymentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMode",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetParametersContent(val *string) {
	if err := j.validateSetParametersContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parametersContent",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetTemplateContent(val *string) {
	if err := j.validateSetTemplateContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateContent",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupTemplateDeployment)SetTemplateSpecVersionId(val *string) {
	if err := j.validateSetTemplateSpecVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateSpecVersionId",
		val,
	)
}

// Generates CDKTF code for importing a ResourceGroupTemplateDeployment resource upon running "cdktf plan <stack-name>".
func ResourceGroupTemplateDeployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateResourceGroupTemplateDeployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupTemplateDeployment.ResourceGroupTemplateDeployment",
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
func ResourceGroupTemplateDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceGroupTemplateDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupTemplateDeployment.ResourceGroupTemplateDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceGroupTemplateDeployment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceGroupTemplateDeployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupTemplateDeployment.ResourceGroupTemplateDeployment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceGroupTemplateDeployment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceGroupTemplateDeployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupTemplateDeployment.ResourceGroupTemplateDeployment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourceGroupTemplateDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.resourceGroupTemplateDeployment.ResourceGroupTemplateDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) PutTimeouts(value *ResourceGroupTemplateDeploymentTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetDebugLevel() {
	_jsii_.InvokeVoid(
		r,
		"resetDebugLevel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetParametersContent() {
	_jsii_.InvokeVoid(
		r,
		"resetParametersContent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetTemplateContent() {
	_jsii_.InvokeVoid(
		r,
		"resetTemplateContent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetTemplateSpecVersionId() {
	_jsii_.InvokeVoid(
		r,
		"resetTemplateSpecVersionId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupTemplateDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


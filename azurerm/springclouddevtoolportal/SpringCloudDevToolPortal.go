// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springclouddevtoolportal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/springclouddevtoolportal/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/spring_cloud_dev_tool_portal azurerm_spring_cloud_dev_tool_portal}.
type SpringCloudDevToolPortal interface {
	cdktf.TerraformResource
	ApplicationAcceleratorEnabled() interface{}
	SetApplicationAcceleratorEnabled(val interface{})
	ApplicationAcceleratorEnabledInput() interface{}
	ApplicationLiveViewEnabled() interface{}
	SetApplicationLiveViewEnabled(val interface{})
	ApplicationLiveViewEnabledInput() interface{}
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
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SpringCloudServiceId() *string
	SetSpringCloudServiceId(val *string)
	SpringCloudServiceIdInput() *string
	Sso() SpringCloudDevToolPortalSsoOutputReference
	SsoInput() *SpringCloudDevToolPortalSso
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SpringCloudDevToolPortalTimeoutsOutputReference
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
	PutSso(value *SpringCloudDevToolPortalSso)
	PutTimeouts(value *SpringCloudDevToolPortalTimeouts)
	ResetApplicationAcceleratorEnabled()
	ResetApplicationLiveViewEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetSso()
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

// The jsii proxy struct for SpringCloudDevToolPortal
type jsiiProxy_SpringCloudDevToolPortal struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpringCloudDevToolPortal) ApplicationAcceleratorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationAcceleratorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) ApplicationAcceleratorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationAcceleratorEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) ApplicationLiveViewEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationLiveViewEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) ApplicationLiveViewEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationLiveViewEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) SpringCloudServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) SpringCloudServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"springCloudServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Sso() SpringCloudDevToolPortalSsoOutputReference {
	var returns SpringCloudDevToolPortalSsoOutputReference
	_jsii_.Get(
		j,
		"sso",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) SsoInput() *SpringCloudDevToolPortalSso {
	var returns *SpringCloudDevToolPortalSso
	_jsii_.Get(
		j,
		"ssoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) Timeouts() SpringCloudDevToolPortalTimeoutsOutputReference {
	var returns SpringCloudDevToolPortalTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudDevToolPortal) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/spring_cloud_dev_tool_portal azurerm_spring_cloud_dev_tool_portal} Resource.
func NewSpringCloudDevToolPortal(scope constructs.Construct, id *string, config *SpringCloudDevToolPortalConfig) SpringCloudDevToolPortal {
	_init_.Initialize()

	if err := validateNewSpringCloudDevToolPortalParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudDevToolPortal{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudDevToolPortal.SpringCloudDevToolPortal",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/spring_cloud_dev_tool_portal azurerm_spring_cloud_dev_tool_portal} Resource.
func NewSpringCloudDevToolPortal_Override(s SpringCloudDevToolPortal, scope constructs.Construct, id *string, config *SpringCloudDevToolPortalConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudDevToolPortal.SpringCloudDevToolPortal",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetApplicationAcceleratorEnabled(val interface{}) {
	if err := j.validateSetApplicationAcceleratorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationAcceleratorEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetApplicationLiveViewEnabled(val interface{}) {
	if err := j.validateSetApplicationLiveViewEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationLiveViewEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudDevToolPortal)SetSpringCloudServiceId(val *string) {
	if err := j.validateSetSpringCloudServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"springCloudServiceId",
		val,
	)
}

// Generates CDKTF code for importing a SpringCloudDevToolPortal resource upon running "cdktf plan <stack-name>".
func SpringCloudDevToolPortal_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSpringCloudDevToolPortal_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudDevToolPortal.SpringCloudDevToolPortal",
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
func SpringCloudDevToolPortal_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudDevToolPortal_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudDevToolPortal.SpringCloudDevToolPortal",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudDevToolPortal_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudDevToolPortal_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudDevToolPortal.SpringCloudDevToolPortal",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpringCloudDevToolPortal_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpringCloudDevToolPortal_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.springCloudDevToolPortal.SpringCloudDevToolPortal",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpringCloudDevToolPortal_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.springCloudDevToolPortal.SpringCloudDevToolPortal",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpringCloudDevToolPortal) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudDevToolPortal) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) PutSso(value *SpringCloudDevToolPortalSso) {
	if err := s.validatePutSsoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSso",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) PutTimeouts(value *SpringCloudDevToolPortalTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ResetApplicationAcceleratorEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationAcceleratorEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ResetApplicationLiveViewEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationLiveViewEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ResetSso() {
	_jsii_.InvokeVoid(
		s,
		"resetSso",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudDevToolPortal) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudDevToolPortal) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudDevToolPortal) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


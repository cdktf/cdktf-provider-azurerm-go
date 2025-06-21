// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoororigin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/cdnfrontdoororigin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cdn_frontdoor_origin azurerm_cdn_frontdoor_origin}.
type CdnFrontdoorOrigin interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnFrontdoorOriginGroupId() *string
	SetCdnFrontdoorOriginGroupId(val *string)
	CdnFrontdoorOriginGroupIdInput() *string
	CertificateNameCheckEnabled() interface{}
	SetCertificateNameCheckEnabled(val interface{})
	CertificateNameCheckEnabledInput() interface{}
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
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	HttpsPort() *float64
	SetHttpsPort(val *float64)
	HttpsPortInput() *float64
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
	OriginHostHeader() *string
	SetOriginHostHeader(val *string)
	OriginHostHeaderInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	PrivateLink() CdnFrontdoorOriginPrivateLinkOutputReference
	PrivateLinkInput() *CdnFrontdoorOriginPrivateLink
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CdnFrontdoorOriginTimeoutsOutputReference
	TimeoutsInput() interface{}
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	PutPrivateLink(value *CdnFrontdoorOriginPrivateLink)
	PutTimeouts(value *CdnFrontdoorOriginTimeouts)
	ResetEnabled()
	ResetHttpPort()
	ResetHttpsPort()
	ResetId()
	ResetOriginHostHeader()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPriority()
	ResetPrivateLink()
	ResetTimeouts()
	ResetWeight()
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

// The jsii proxy struct for CdnFrontdoorOrigin
type jsiiProxy_CdnFrontdoorOrigin struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CdnFrontdoorOrigin) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) CdnFrontdoorOriginGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) CdnFrontdoorOriginGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnFrontdoorOriginGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) CertificateNameCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateNameCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) CertificateNameCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateNameCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) OriginHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) OriginHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) PrivateLink() CdnFrontdoorOriginPrivateLinkOutputReference {
	var returns CdnFrontdoorOriginPrivateLinkOutputReference
	_jsii_.Get(
		j,
		"privateLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) PrivateLinkInput() *CdnFrontdoorOriginPrivateLink {
	var returns *CdnFrontdoorOriginPrivateLink
	_jsii_.Get(
		j,
		"privateLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Timeouts() CdnFrontdoorOriginTimeoutsOutputReference {
	var returns CdnFrontdoorOriginTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorOrigin) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cdn_frontdoor_origin azurerm_cdn_frontdoor_origin} Resource.
func NewCdnFrontdoorOrigin(scope constructs.Construct, id *string, config *CdnFrontdoorOriginConfig) CdnFrontdoorOrigin {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorOriginParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorOrigin{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorOrigin.CdnFrontdoorOrigin",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/cdn_frontdoor_origin azurerm_cdn_frontdoor_origin} Resource.
func NewCdnFrontdoorOrigin_Override(c CdnFrontdoorOrigin, scope constructs.Construct, id *string, config *CdnFrontdoorOriginConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorOrigin.CdnFrontdoorOrigin",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetCdnFrontdoorOriginGroupId(val *string) {
	if err := j.validateSetCdnFrontdoorOriginGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnFrontdoorOriginGroupId",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetCertificateNameCheckEnabled(val interface{}) {
	if err := j.validateSetCertificateNameCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateNameCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetHttpPort(val *float64) {
	if err := j.validateSetHttpPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetHttpsPort(val *float64) {
	if err := j.validateSetHttpsPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetOriginHostHeader(val *string) {
	if err := j.validateSetOriginHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originHostHeader",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorOrigin)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

// Generates CDKTF code for importing a CdnFrontdoorOrigin resource upon running "cdktf plan <stack-name>".
func CdnFrontdoorOrigin_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCdnFrontdoorOrigin_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorOrigin.CdnFrontdoorOrigin",
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
func CdnFrontdoorOrigin_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnFrontdoorOrigin_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorOrigin.CdnFrontdoorOrigin",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CdnFrontdoorOrigin_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnFrontdoorOrigin_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorOrigin.CdnFrontdoorOrigin",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CdnFrontdoorOrigin_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdnFrontdoorOrigin_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cdnFrontdoorOrigin.CdnFrontdoorOrigin",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CdnFrontdoorOrigin_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cdnFrontdoorOrigin.CdnFrontdoorOrigin",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) PutPrivateLink(value *CdnFrontdoorOriginPrivateLink) {
	if err := c.validatePutPrivateLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPrivateLink",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) PutTimeouts(value *CdnFrontdoorOriginTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetHttpPort() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetHttpsPort() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetOriginHostHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginHostHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetPriority() {
	_jsii_.InvokeVoid(
		c,
		"resetPriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetPrivateLink() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ResetWeight() {
	_jsii_.InvokeVoid(
		c,
		"resetWeight",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorOrigin) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorOrigin) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


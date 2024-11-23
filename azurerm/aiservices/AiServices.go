// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aiservices

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/aiservices/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/ai_services azurerm_ai_services}.
type AiServices interface {
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
	CustomerManagedKey() AiServicesCustomerManagedKeyOutputReference
	CustomerManagedKeyInput() *AiServicesCustomerManagedKey
	CustomSubdomainName() *string
	SetCustomSubdomainName(val *string)
	CustomSubdomainNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdns() *[]*string
	SetFqdns(val *[]*string)
	FqdnsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	Identity() AiServicesIdentityOutputReference
	IdentityInput() *AiServicesIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthenticationEnabled() interface{}
	SetLocalAuthenticationEnabled(val interface{})
	LocalAuthenticationEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAcls() AiServicesNetworkAclsOutputReference
	NetworkAclsInput() *AiServicesNetworkAcls
	// The tree node.
	Node() constructs.Node
	OutboundNetworkAccessRestricted() interface{}
	SetOutboundNetworkAccessRestricted(val interface{})
	OutboundNetworkAccessRestrictedInput() interface{}
	PrimaryAccessKey() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccess() *string
	SetPublicNetworkAccess(val *string)
	PublicNetworkAccessInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	Storage() AiServicesStorageList
	StorageInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AiServicesTimeoutsOutputReference
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
	PutCustomerManagedKey(value *AiServicesCustomerManagedKey)
	PutIdentity(value *AiServicesIdentity)
	PutNetworkAcls(value *AiServicesNetworkAcls)
	PutStorage(value interface{})
	PutTimeouts(value *AiServicesTimeouts)
	ResetCustomerManagedKey()
	ResetCustomSubdomainName()
	ResetFqdns()
	ResetId()
	ResetIdentity()
	ResetLocalAuthenticationEnabled()
	ResetNetworkAcls()
	ResetOutboundNetworkAccessRestricted()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccess()
	ResetStorage()
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

// The jsii proxy struct for AiServices
type jsiiProxy_AiServices struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AiServices) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) CustomerManagedKey() AiServicesCustomerManagedKeyOutputReference {
	var returns AiServicesCustomerManagedKeyOutputReference
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) CustomerManagedKeyInput() *AiServicesCustomerManagedKey {
	var returns *AiServicesCustomerManagedKey
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) CustomSubdomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSubdomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) CustomSubdomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSubdomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Fqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) FqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Identity() AiServicesIdentityOutputReference {
	var returns AiServicesIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) IdentityInput() *AiServicesIdentity {
	var returns *AiServicesIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) LocalAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) LocalAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) NetworkAcls() AiServicesNetworkAclsOutputReference {
	var returns AiServicesNetworkAclsOutputReference
	_jsii_.Get(
		j,
		"networkAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) NetworkAclsInput() *AiServicesNetworkAcls {
	var returns *AiServicesNetworkAcls
	_jsii_.Get(
		j,
		"networkAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) OutboundNetworkAccessRestricted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundNetworkAccessRestricted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) OutboundNetworkAccessRestrictedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundNetworkAccessRestrictedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) PublicNetworkAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) PublicNetworkAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Storage() AiServicesStorageList {
	var returns AiServicesStorageList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) Timeouts() AiServicesTimeoutsOutputReference {
	var returns AiServicesTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiServices) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/ai_services azurerm_ai_services} Resource.
func NewAiServices(scope constructs.Construct, id *string, config *AiServicesConfig) AiServices {
	_init_.Initialize()

	if err := validateNewAiServicesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiServices{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.aiServices.AiServices",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/ai_services azurerm_ai_services} Resource.
func NewAiServices_Override(a AiServices, scope constructs.Construct, id *string, config *AiServicesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.aiServices.AiServices",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AiServices)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetCustomSubdomainName(val *string) {
	if err := j.validateSetCustomSubdomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSubdomainName",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetFqdns(val *[]*string) {
	if err := j.validateSetFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fqdns",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetLocalAuthenticationEnabled(val interface{}) {
	if err := j.validateSetLocalAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetOutboundNetworkAccessRestricted(val interface{}) {
	if err := j.validateSetOutboundNetworkAccessRestrictedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundNetworkAccessRestricted",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetPublicNetworkAccess(val *string) {
	if err := j.validateSetPublicNetworkAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccess",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_AiServices)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a AiServices resource upon running "cdktf plan <stack-name>".
func AiServices_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAiServices_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.aiServices.AiServices",
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
func AiServices_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiServices_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.aiServices.AiServices",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiServices_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiServices_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.aiServices.AiServices",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiServices_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiServices_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.aiServices.AiServices",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AiServices_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.aiServices.AiServices",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AiServices) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AiServices) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AiServices) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiServices) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiServices) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiServices) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiServices) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiServices) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiServices) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiServices) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiServices) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiServices) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiServices) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AiServices) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiServices) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiServices) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AiServices) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiServices) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AiServices) PutCustomerManagedKey(value *AiServicesCustomerManagedKey) {
	if err := a.validatePutCustomerManagedKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiServices) PutIdentity(value *AiServicesIdentity) {
	if err := a.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiServices) PutNetworkAcls(value *AiServicesNetworkAcls) {
	if err := a.validatePutNetworkAclsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkAcls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiServices) PutStorage(value interface{}) {
	if err := a.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiServices) PutTimeouts(value *AiServicesTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiServices) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetCustomSubdomainName() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSubdomainName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetFqdns() {
	_jsii_.InvokeVoid(
		a,
		"resetFqdns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetLocalAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalAuthenticationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetNetworkAcls() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkAcls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetOutboundNetworkAccessRestricted() {
	_jsii_.InvokeVoid(
		a,
		"resetOutboundNetworkAccessRestricted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetPublicNetworkAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicNetworkAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiServices) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiServices) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiServices) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiServices) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiServices) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiServices) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgriddomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/eventgriddomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/eventgrid_domain azurerm_eventgrid_domain}.
type EventgridDomain interface {
	cdktf.TerraformResource
	AutoCreateTopicWithFirstSubscription() interface{}
	SetAutoCreateTopicWithFirstSubscription(val interface{})
	AutoCreateTopicWithFirstSubscriptionInput() interface{}
	AutoDeleteTopicWithLastSubscription() interface{}
	SetAutoDeleteTopicWithLastSubscription(val interface{})
	AutoDeleteTopicWithLastSubscriptionInput() interface{}
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
	Endpoint() *string
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
	Identity() EventgridDomainIdentityOutputReference
	IdentityInput() *EventgridDomainIdentity
	IdInput() *string
	InboundIpRule() EventgridDomainInboundIpRuleList
	InboundIpRuleInput() interface{}
	InputMappingDefaultValues() EventgridDomainInputMappingDefaultValuesOutputReference
	InputMappingDefaultValuesInput() *EventgridDomainInputMappingDefaultValues
	InputMappingFields() EventgridDomainInputMappingFieldsOutputReference
	InputMappingFieldsInput() *EventgridDomainInputMappingFields
	InputSchema() *string
	SetInputSchema(val *string)
	InputSchemaInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthEnabled() interface{}
	SetLocalAuthEnabled(val interface{})
	LocalAuthEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrimaryAccessKey() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EventgridDomainTimeoutsOutputReference
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
	PutIdentity(value *EventgridDomainIdentity)
	PutInboundIpRule(value interface{})
	PutInputMappingDefaultValues(value *EventgridDomainInputMappingDefaultValues)
	PutInputMappingFields(value *EventgridDomainInputMappingFields)
	PutTimeouts(value *EventgridDomainTimeouts)
	ResetAutoCreateTopicWithFirstSubscription()
	ResetAutoDeleteTopicWithLastSubscription()
	ResetId()
	ResetIdentity()
	ResetInboundIpRule()
	ResetInputMappingDefaultValues()
	ResetInputMappingFields()
	ResetInputSchema()
	ResetLocalAuthEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
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

// The jsii proxy struct for EventgridDomain
type jsiiProxy_EventgridDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EventgridDomain) AutoCreateTopicWithFirstSubscription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicWithFirstSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) AutoCreateTopicWithFirstSubscriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicWithFirstSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) AutoDeleteTopicWithLastSubscription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteTopicWithLastSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) AutoDeleteTopicWithLastSubscriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteTopicWithLastSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Identity() EventgridDomainIdentityOutputReference {
	var returns EventgridDomainIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) IdentityInput() *EventgridDomainIdentity {
	var returns *EventgridDomainIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InboundIpRule() EventgridDomainInboundIpRuleList {
	var returns EventgridDomainInboundIpRuleList
	_jsii_.Get(
		j,
		"inboundIpRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InboundIpRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundIpRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InputMappingDefaultValues() EventgridDomainInputMappingDefaultValuesOutputReference {
	var returns EventgridDomainInputMappingDefaultValuesOutputReference
	_jsii_.Get(
		j,
		"inputMappingDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InputMappingDefaultValuesInput() *EventgridDomainInputMappingDefaultValues {
	var returns *EventgridDomainInputMappingDefaultValues
	_jsii_.Get(
		j,
		"inputMappingDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InputMappingFields() EventgridDomainInputMappingFieldsOutputReference {
	var returns EventgridDomainInputMappingFieldsOutputReference
	_jsii_.Get(
		j,
		"inputMappingFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InputMappingFieldsInput() *EventgridDomainInputMappingFields {
	var returns *EventgridDomainInputMappingFields
	_jsii_.Get(
		j,
		"inputMappingFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InputSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) InputSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) LocalAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) LocalAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) Timeouts() EventgridDomainTimeoutsOutputReference {
	var returns EventgridDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/eventgrid_domain azurerm_eventgrid_domain} Resource.
func NewEventgridDomain(scope constructs.Construct, id *string, config *EventgridDomainConfig) EventgridDomain {
	_init_.Initialize()

	if err := validateNewEventgridDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridDomain{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridDomain.EventgridDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/eventgrid_domain azurerm_eventgrid_domain} Resource.
func NewEventgridDomain_Override(e EventgridDomain, scope constructs.Construct, id *string, config *EventgridDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridDomain.EventgridDomain",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EventgridDomain)SetAutoCreateTopicWithFirstSubscription(val interface{}) {
	if err := j.validateSetAutoCreateTopicWithFirstSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreateTopicWithFirstSubscription",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetAutoDeleteTopicWithLastSubscription(val interface{}) {
	if err := j.validateSetAutoDeleteTopicWithLastSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeleteTopicWithLastSubscription",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetInputSchema(val *string) {
	if err := j.validateSetInputSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputSchema",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetLocalAuthEnabled(val interface{}) {
	if err := j.validateSetLocalAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_EventgridDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a EventgridDomain resource upon running "cdktf plan <stack-name>".
func EventgridDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEventgridDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridDomain.EventgridDomain",
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
func EventgridDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridDomain.EventgridDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridDomain.EventgridDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridDomain.EventgridDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EventgridDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.eventgridDomain.EventgridDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EventgridDomain) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EventgridDomain) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EventgridDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EventgridDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventgridDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EventgridDomain) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventgridDomain) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EventgridDomain) PutIdentity(value *EventgridDomainIdentity) {
	if err := e.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridDomain) PutInboundIpRule(value interface{}) {
	if err := e.validatePutInboundIpRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInboundIpRule",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridDomain) PutInputMappingDefaultValues(value *EventgridDomainInputMappingDefaultValues) {
	if err := e.validatePutInputMappingDefaultValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInputMappingDefaultValues",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridDomain) PutInputMappingFields(value *EventgridDomainInputMappingFields) {
	if err := e.validatePutInputMappingFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInputMappingFields",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridDomain) PutTimeouts(value *EventgridDomainTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridDomain) ResetAutoCreateTopicWithFirstSubscription() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoCreateTopicWithFirstSubscription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetAutoDeleteTopicWithLastSubscription() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoDeleteTopicWithLastSubscription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetInboundIpRule() {
	_jsii_.InvokeVoid(
		e,
		"resetInboundIpRule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetInputMappingDefaultValues() {
	_jsii_.InvokeVoid(
		e,
		"resetInputMappingDefaultValues",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetInputMappingFields() {
	_jsii_.InvokeVoid(
		e,
		"resetInputMappingFields",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetInputSchema() {
	_jsii_.InvokeVoid(
		e,
		"resetInputSchema",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetLocalAuthEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetLocalAuthEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


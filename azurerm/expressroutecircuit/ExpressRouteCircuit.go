package expressroutecircuit

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/expressroutecircuit/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/express_route_circuit azurerm_express_route_circuit}.
type ExpressRouteCircuit interface {
	cdktf.TerraformResource
	AllowClassicOperations() interface{}
	SetAllowClassicOperations(val interface{})
	AllowClassicOperationsInput() interface{}
	AuthorizationKey() *string
	SetAuthorizationKey(val *string)
	AuthorizationKeyInput() *string
	BandwidthInGbps() *float64
	SetBandwidthInGbps(val *float64)
	BandwidthInGbpsInput() *float64
	BandwidthInMbps() *float64
	SetBandwidthInMbps(val *float64)
	BandwidthInMbpsInput() *float64
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
	ExpressRoutePortId() *string
	SetExpressRoutePortId(val *string)
	ExpressRoutePortIdInput() *string
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeeringLocation() *string
	SetPeeringLocation(val *string)
	PeeringLocationInput() *string
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
	ServiceKey() *string
	ServiceProviderName() *string
	SetServiceProviderName(val *string)
	ServiceProviderNameInput() *string
	ServiceProviderProvisioningState() *string
	Sku() ExpressRouteCircuitSkuOutputReference
	SkuInput() *ExpressRouteCircuitSku
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ExpressRouteCircuitTimeoutsOutputReference
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
	PutSku(value *ExpressRouteCircuitSku)
	PutTimeouts(value *ExpressRouteCircuitTimeouts)
	ResetAllowClassicOperations()
	ResetAuthorizationKey()
	ResetBandwidthInGbps()
	ResetBandwidthInMbps()
	ResetExpressRoutePortId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeeringLocation()
	ResetServiceProviderName()
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

// The jsii proxy struct for ExpressRouteCircuit
type jsiiProxy_ExpressRouteCircuit struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExpressRouteCircuit) AllowClassicOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) AllowClassicOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) AuthorizationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) AuthorizationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) BandwidthInGbps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) BandwidthInGbpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInGbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) BandwidthInMbps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) BandwidthInMbpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ExpressRoutePortId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRoutePortId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ExpressRoutePortIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRoutePortIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) PeeringLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) PeeringLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ServiceProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ServiceProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) ServiceProviderProvisioningState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceProviderProvisioningState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Sku() ExpressRouteCircuitSkuOutputReference {
	var returns ExpressRouteCircuitSkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) SkuInput() *ExpressRouteCircuitSku {
	var returns *ExpressRouteCircuitSku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) Timeouts() ExpressRouteCircuitTimeoutsOutputReference {
	var returns ExpressRouteCircuitTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuit) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/express_route_circuit azurerm_express_route_circuit} Resource.
func NewExpressRouteCircuit(scope constructs.Construct, id *string, config *ExpressRouteCircuitConfig) ExpressRouteCircuit {
	_init_.Initialize()

	if err := validateNewExpressRouteCircuitParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExpressRouteCircuit{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuit.ExpressRouteCircuit",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/express_route_circuit azurerm_express_route_circuit} Resource.
func NewExpressRouteCircuit_Override(e ExpressRouteCircuit, scope constructs.Construct, id *string, config *ExpressRouteCircuitConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRouteCircuit.ExpressRouteCircuit",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetAllowClassicOperations(val interface{}) {
	if err := j.validateSetAllowClassicOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClassicOperations",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetAuthorizationKey(val *string) {
	if err := j.validateSetAuthorizationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationKey",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetBandwidthInGbps(val *float64) {
	if err := j.validateSetBandwidthInGbpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthInGbps",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetBandwidthInMbps(val *float64) {
	if err := j.validateSetBandwidthInMbpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthInMbps",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetExpressRoutePortId(val *string) {
	if err := j.validateSetExpressRoutePortIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressRoutePortId",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetPeeringLocation(val *string) {
	if err := j.validateSetPeeringLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peeringLocation",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetServiceProviderName(val *string) {
	if err := j.validateSetServiceProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceProviderName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuit)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
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
func ExpressRouteCircuit_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExpressRouteCircuit_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.expressRouteCircuit.ExpressRouteCircuit",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExpressRouteCircuit_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExpressRouteCircuit_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.expressRouteCircuit.ExpressRouteCircuit",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExpressRouteCircuit_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExpressRouteCircuit_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.expressRouteCircuit.ExpressRouteCircuit",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExpressRouteCircuit_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.expressRouteCircuit.ExpressRouteCircuit",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExpressRouteCircuit) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ExpressRouteCircuit) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ExpressRouteCircuit) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExpressRouteCircuit) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) PutSku(value *ExpressRouteCircuitSku) {
	if err := e.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSku",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) PutTimeouts(value *ExpressRouteCircuitTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetAllowClassicOperations() {
	_jsii_.InvokeVoid(
		e,
		"resetAllowClassicOperations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetAuthorizationKey() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthorizationKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetBandwidthInGbps() {
	_jsii_.InvokeVoid(
		e,
		"resetBandwidthInGbps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetBandwidthInMbps() {
	_jsii_.InvokeVoid(
		e,
		"resetBandwidthInMbps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetExpressRoutePortId() {
	_jsii_.InvokeVoid(
		e,
		"resetExpressRoutePortId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetPeeringLocation() {
	_jsii_.InvokeVoid(
		e,
		"resetPeeringLocation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetServiceProviderName() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceProviderName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuit) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuit) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuit) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuit) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


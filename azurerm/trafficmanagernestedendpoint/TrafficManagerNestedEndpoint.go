package trafficmanagernestedendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/trafficmanagernestedendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint azurerm_traffic_manager_nested_endpoint}.
type TrafficManagerNestedEndpoint interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CustomHeader() TrafficManagerNestedEndpointCustomHeaderList
	CustomHeaderInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EndpointLocation() *string
	SetEndpointLocation(val *string)
	EndpointLocationInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoMappings() *[]*string
	SetGeoMappings(val *[]*string)
	GeoMappingsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinimumChildEndpoints() *float64
	SetMinimumChildEndpoints(val *float64)
	MinimumChildEndpointsInput() *float64
	MinimumRequiredChildEndpointsIpv4() *float64
	SetMinimumRequiredChildEndpointsIpv4(val *float64)
	MinimumRequiredChildEndpointsIpv4Input() *float64
	MinimumRequiredChildEndpointsIpv6() *float64
	SetMinimumRequiredChildEndpointsIpv6(val *float64)
	MinimumRequiredChildEndpointsIpv6Input() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ProfileId() *string
	SetProfileId(val *string)
	ProfileIdInput() *string
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
	Subnet() TrafficManagerNestedEndpointSubnetList
	SubnetInput() interface{}
	TargetResourceId() *string
	SetTargetResourceId(val *string)
	TargetResourceIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() TrafficManagerNestedEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	PutCustomHeader(value interface{})
	PutSubnet(value interface{})
	PutTimeouts(value *TrafficManagerNestedEndpointTimeouts)
	ResetCustomHeader()
	ResetEnabled()
	ResetEndpointLocation()
	ResetGeoMappings()
	ResetId()
	ResetMinimumRequiredChildEndpointsIpv4()
	ResetMinimumRequiredChildEndpointsIpv6()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPriority()
	ResetSubnet()
	ResetTimeouts()
	ResetWeight()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TrafficManagerNestedEndpoint
type jsiiProxy_TrafficManagerNestedEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) CustomHeader() TrafficManagerNestedEndpointCustomHeaderList {
	var returns TrafficManagerNestedEndpointCustomHeaderList
	_jsii_.Get(
		j,
		"customHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) CustomHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) EndpointLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) EndpointLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) GeoMappings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geoMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) GeoMappingsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"geoMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) MinimumChildEndpoints() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumChildEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) MinimumChildEndpointsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumChildEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) MinimumRequiredChildEndpointsIpv4() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRequiredChildEndpointsIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) MinimumRequiredChildEndpointsIpv4Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRequiredChildEndpointsIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) MinimumRequiredChildEndpointsIpv6() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRequiredChildEndpointsIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) MinimumRequiredChildEndpointsIpv6Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRequiredChildEndpointsIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) ProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) ProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Subnet() TrafficManagerNestedEndpointSubnetList {
	var returns TrafficManagerNestedEndpointSubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) SubnetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) TargetResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) TargetResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Timeouts() TrafficManagerNestedEndpointTimeoutsOutputReference {
	var returns TrafficManagerNestedEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint azurerm_traffic_manager_nested_endpoint} Resource.
func NewTrafficManagerNestedEndpoint(scope constructs.Construct, id *string, config *TrafficManagerNestedEndpointConfig) TrafficManagerNestedEndpoint {
	_init_.Initialize()

	if err := validateNewTrafficManagerNestedEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TrafficManagerNestedEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.trafficManagerNestedEndpoint.TrafficManagerNestedEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint azurerm_traffic_manager_nested_endpoint} Resource.
func NewTrafficManagerNestedEndpoint_Override(t TrafficManagerNestedEndpoint, scope constructs.Construct, id *string, config *TrafficManagerNestedEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.trafficManagerNestedEndpoint.TrafficManagerNestedEndpoint",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetEndpointLocation(val *string) {
	if err := j.validateSetEndpointLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointLocation",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetGeoMappings(val *[]*string) {
	if err := j.validateSetGeoMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoMappings",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetMinimumChildEndpoints(val *float64) {
	if err := j.validateSetMinimumChildEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumChildEndpoints",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetMinimumRequiredChildEndpointsIpv4(val *float64) {
	if err := j.validateSetMinimumRequiredChildEndpointsIpv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumRequiredChildEndpointsIpv4",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetMinimumRequiredChildEndpointsIpv6(val *float64) {
	if err := j.validateSetMinimumRequiredChildEndpointsIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumRequiredChildEndpointsIpv6",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetProfileId(val *string) {
	if err := j.validateSetProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileId",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetTargetResourceId(val *string) {
	if err := j.validateSetTargetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResourceId",
		val,
	)
}

func (j *jsiiProxy_TrafficManagerNestedEndpoint)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
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
func TrafficManagerNestedEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTrafficManagerNestedEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.trafficManagerNestedEndpoint.TrafficManagerNestedEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TrafficManagerNestedEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTrafficManagerNestedEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.trafficManagerNestedEndpoint.TrafficManagerNestedEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TrafficManagerNestedEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTrafficManagerNestedEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.trafficManagerNestedEndpoint.TrafficManagerNestedEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TrafficManagerNestedEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.trafficManagerNestedEndpoint.TrafficManagerNestedEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) PutCustomHeader(value interface{}) {
	if err := t.validatePutCustomHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) PutSubnet(value interface{}) {
	if err := t.validatePutSubnetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSubnet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) PutTimeouts(value *TrafficManagerNestedEndpointTimeouts) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetCustomHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetEndpointLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetGeoMappings() {
	_jsii_.InvokeVoid(
		t,
		"resetGeoMappings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetMinimumRequiredChildEndpointsIpv4() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumRequiredChildEndpointsIpv4",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetMinimumRequiredChildEndpointsIpv6() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumRequiredChildEndpointsIpv6",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetPriority() {
	_jsii_.InvokeVoid(
		t,
		"resetPriority",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetSubnet() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ResetWeight() {
	_jsii_.InvokeVoid(
		t,
		"resetWeight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficManagerNestedEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


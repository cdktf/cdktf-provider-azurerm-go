package servicefabriccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/servicefabriccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceFabricClusterNodeTypeOutputReference interface {
	cdktf.ComplexObject
	ApplicationPorts() ServiceFabricClusterNodeTypeApplicationPortsOutputReference
	ApplicationPortsInput() *ServiceFabricClusterNodeTypeApplicationPorts
	Capacities() *map[string]*string
	SetCapacities(val *map[string]*string)
	CapacitiesInput() *map[string]*string
	ClientEndpointPort() *float64
	SetClientEndpointPort(val *float64)
	ClientEndpointPortInput() *float64
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
	DurabilityLevel() *string
	SetDurabilityLevel(val *string)
	DurabilityLevelInput() *string
	EphemeralPorts() ServiceFabricClusterNodeTypeEphemeralPortsOutputReference
	EphemeralPortsInput() *ServiceFabricClusterNodeTypeEphemeralPorts
	// Experimental.
	Fqn() *string
	HttpEndpointPort() *float64
	SetHttpEndpointPort(val *float64)
	HttpEndpointPortInput() *float64
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsPrimary() interface{}
	SetIsPrimary(val interface{})
	IsPrimaryInput() interface{}
	IsStateless() interface{}
	SetIsStateless(val interface{})
	IsStatelessInput() interface{}
	MultipleAvailabilityZones() interface{}
	SetMultipleAvailabilityZones(val interface{})
	MultipleAvailabilityZonesInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PlacementProperties() *map[string]*string
	SetPlacementProperties(val *map[string]*string)
	PlacementPropertiesInput() *map[string]*string
	ReverseProxyEndpointPort() *float64
	SetReverseProxyEndpointPort(val *float64)
	ReverseProxyEndpointPortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutApplicationPorts(value *ServiceFabricClusterNodeTypeApplicationPorts)
	PutEphemeralPorts(value *ServiceFabricClusterNodeTypeEphemeralPorts)
	ResetApplicationPorts()
	ResetCapacities()
	ResetDurabilityLevel()
	ResetEphemeralPorts()
	ResetIsStateless()
	ResetMultipleAvailabilityZones()
	ResetPlacementProperties()
	ResetReverseProxyEndpointPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterNodeTypeOutputReference
type jsiiProxy_ServiceFabricClusterNodeTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ApplicationPorts() ServiceFabricClusterNodeTypeApplicationPortsOutputReference {
	var returns ServiceFabricClusterNodeTypeApplicationPortsOutputReference
	_jsii_.Get(
		j,
		"applicationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ApplicationPortsInput() *ServiceFabricClusterNodeTypeApplicationPorts {
	var returns *ServiceFabricClusterNodeTypeApplicationPorts
	_jsii_.Get(
		j,
		"applicationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Capacities() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) CapacitiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ClientEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ClientEndpointPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEndpointPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) DurabilityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durabilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) DurabilityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durabilityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) EphemeralPorts() ServiceFabricClusterNodeTypeEphemeralPortsOutputReference {
	var returns ServiceFabricClusterNodeTypeEphemeralPortsOutputReference
	_jsii_.Get(
		j,
		"ephemeralPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) EphemeralPortsInput() *ServiceFabricClusterNodeTypeEphemeralPorts {
	var returns *ServiceFabricClusterNodeTypeEphemeralPorts
	_jsii_.Get(
		j,
		"ephemeralPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) HttpEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) HttpEndpointPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpEndpointPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsPrimary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsPrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsStateless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStateless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsStatelessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStatelessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) MultipleAvailabilityZones() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) MultipleAvailabilityZonesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PlacementProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PlacementPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ReverseProxyEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reverseProxyEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ReverseProxyEndpointPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reverseProxyEndpointPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterNodeTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricClusterNodeTypeOutputReference {
	_init_.Initialize()

	if err := validateNewServiceFabricClusterNodeTypeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceFabricClusterNodeTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterNodeTypeOutputReference_Override(s ServiceFabricClusterNodeTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetCapacities(val *map[string]*string) {
	if err := j.validateSetCapacitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacities",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetClientEndpointPort(val *float64) {
	if err := j.validateSetClientEndpointPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientEndpointPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetDurabilityLevel(val *string) {
	if err := j.validateSetDurabilityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durabilityLevel",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetHttpEndpointPort(val *float64) {
	if err := j.validateSetHttpEndpointPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpEndpointPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetIsPrimary(val interface{}) {
	if err := j.validateSetIsPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPrimary",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetIsStateless(val interface{}) {
	if err := j.validateSetIsStatelessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isStateless",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetMultipleAvailabilityZones(val interface{}) {
	if err := j.validateSetMultipleAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multipleAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetPlacementProperties(val *map[string]*string) {
	if err := j.validateSetPlacementPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementProperties",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetReverseProxyEndpointPort(val *float64) {
	if err := j.validateSetReverseProxyEndpointPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reverseProxyEndpointPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PutApplicationPorts(value *ServiceFabricClusterNodeTypeApplicationPorts) {
	if err := s.validatePutApplicationPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putApplicationPorts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PutEphemeralPorts(value *ServiceFabricClusterNodeTypeEphemeralPorts) {
	if err := s.validatePutEphemeralPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEphemeralPorts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetApplicationPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetCapacities() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetDurabilityLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetDurabilityLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetEphemeralPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetEphemeralPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetIsStateless() {
	_jsii_.InvokeVoid(
		s,
		"resetIsStateless",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetMultipleAvailabilityZones() {
	_jsii_.InvokeVoid(
		s,
		"resetMultipleAvailabilityZones",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetPlacementProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetReverseProxyEndpointPort() {
	_jsii_.InvokeVoid(
		s,
		"resetReverseProxyEndpointPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


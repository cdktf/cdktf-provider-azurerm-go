package servicefabriccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/servicefabriccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	SetInternalValue(val *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy)
	MaxDeltaUnhealthyApplicationsPercent() *float64
	SetMaxDeltaUnhealthyApplicationsPercent(val *float64)
	MaxDeltaUnhealthyApplicationsPercentInput() *float64
	MaxDeltaUnhealthyNodesPercent() *float64
	SetMaxDeltaUnhealthyNodesPercent(val *float64)
	MaxDeltaUnhealthyNodesPercentInput() *float64
	MaxUpgradeDomainDeltaUnhealthyNodesPercent() *float64
	SetMaxUpgradeDomainDeltaUnhealthyNodesPercent(val *float64)
	MaxUpgradeDomainDeltaUnhealthyNodesPercentInput() *float64
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
	ResetMaxDeltaUnhealthyApplicationsPercent()
	ResetMaxDeltaUnhealthyNodesPercent()
	ResetMaxUpgradeDomainDeltaUnhealthyNodesPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference
type jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) InternalValue() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy {
	var returns *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyApplicationsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyApplicationsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyApplicationsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyApplicationsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyNodesPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyNodesPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyNodesPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyNodesPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxUpgradeDomainDeltaUnhealthyNodesPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpgradeDomainDeltaUnhealthyNodesPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxUpgradeDomainDeltaUnhealthyNodesPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpgradeDomainDeltaUnhealthyNodesPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference_Override(s ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetInternalValue(val *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetMaxDeltaUnhealthyApplicationsPercent(val *float64) {
	if err := j.validateSetMaxDeltaUnhealthyApplicationsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDeltaUnhealthyApplicationsPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetMaxDeltaUnhealthyNodesPercent(val *float64) {
	if err := j.validateSetMaxDeltaUnhealthyNodesPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDeltaUnhealthyNodesPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetMaxUpgradeDomainDeltaUnhealthyNodesPercent(val *float64) {
	if err := j.validateSetMaxUpgradeDomainDeltaUnhealthyNodesPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUpgradeDomainDeltaUnhealthyNodesPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ResetMaxDeltaUnhealthyApplicationsPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxDeltaUnhealthyApplicationsPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ResetMaxDeltaUnhealthyNodesPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxDeltaUnhealthyNodesPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ResetMaxUpgradeDomainDeltaUnhealthyNodesPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxUpgradeDomainDeltaUnhealthyNodesPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package servicefabriccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/servicefabriccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceFabricClusterUpgradePolicyOutputReference interface {
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
	DeltaHealthPolicy() ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference
	DeltaHealthPolicyInput() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	ForceRestartEnabled() interface{}
	SetForceRestartEnabled(val interface{})
	ForceRestartEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HealthCheckRetryTimeout() *string
	SetHealthCheckRetryTimeout(val *string)
	HealthCheckRetryTimeoutInput() *string
	HealthCheckStableDuration() *string
	SetHealthCheckStableDuration(val *string)
	HealthCheckStableDurationInput() *string
	HealthCheckWaitDuration() *string
	SetHealthCheckWaitDuration(val *string)
	HealthCheckWaitDurationInput() *string
	HealthPolicy() ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference
	HealthPolicyInput() *ServiceFabricClusterUpgradePolicyHealthPolicy
	InternalValue() *ServiceFabricClusterUpgradePolicy
	SetInternalValue(val *ServiceFabricClusterUpgradePolicy)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpgradeDomainTimeout() *string
	SetUpgradeDomainTimeout(val *string)
	UpgradeDomainTimeoutInput() *string
	UpgradeReplicaSetCheckTimeout() *string
	SetUpgradeReplicaSetCheckTimeout(val *string)
	UpgradeReplicaSetCheckTimeoutInput() *string
	UpgradeTimeout() *string
	SetUpgradeTimeout(val *string)
	UpgradeTimeoutInput() *string
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
	PutDeltaHealthPolicy(value *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy)
	PutHealthPolicy(value *ServiceFabricClusterUpgradePolicyHealthPolicy)
	ResetDeltaHealthPolicy()
	ResetForceRestartEnabled()
	ResetHealthCheckRetryTimeout()
	ResetHealthCheckStableDuration()
	ResetHealthCheckWaitDuration()
	ResetHealthPolicy()
	ResetUpgradeDomainTimeout()
	ResetUpgradeReplicaSetCheckTimeout()
	ResetUpgradeTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterUpgradePolicyOutputReference
type jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) DeltaHealthPolicy() ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference {
	var returns ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference
	_jsii_.Get(
		j,
		"deltaHealthPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) DeltaHealthPolicyInput() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy {
	var returns *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	_jsii_.Get(
		j,
		"deltaHealthPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ForceRestartEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRestartEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ForceRestartEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRestartEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckRetryTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckRetryTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckStableDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckStableDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckStableDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckStableDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckWaitDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckWaitDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckWaitDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckWaitDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthPolicy() ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference {
	var returns ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference
	_jsii_.Get(
		j,
		"healthPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthPolicyInput() *ServiceFabricClusterUpgradePolicyHealthPolicy {
	var returns *ServiceFabricClusterUpgradePolicyHealthPolicy
	_jsii_.Get(
		j,
		"healthPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) InternalValue() *ServiceFabricClusterUpgradePolicy {
	var returns *ServiceFabricClusterUpgradePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeDomainTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeDomainTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeDomainTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeDomainTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeReplicaSetCheckTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeReplicaSetCheckTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeReplicaSetCheckTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeReplicaSetCheckTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeTimeoutInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterUpgradePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterUpgradePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewServiceFabricClusterUpgradePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterUpgradePolicyOutputReference_Override(s ServiceFabricClusterUpgradePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetForceRestartEnabled(val interface{}) {
	if err := j.validateSetForceRestartEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRestartEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetHealthCheckRetryTimeout(val *string) {
	if err := j.validateSetHealthCheckRetryTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckRetryTimeout",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetHealthCheckStableDuration(val *string) {
	if err := j.validateSetHealthCheckStableDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckStableDuration",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetHealthCheckWaitDuration(val *string) {
	if err := j.validateSetHealthCheckWaitDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckWaitDuration",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetInternalValue(val *ServiceFabricClusterUpgradePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetUpgradeDomainTimeout(val *string) {
	if err := j.validateSetUpgradeDomainTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeDomainTimeout",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetUpgradeReplicaSetCheckTimeout(val *string) {
	if err := j.validateSetUpgradeReplicaSetCheckTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeReplicaSetCheckTimeout",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference)SetUpgradeTimeout(val *string) {
	if err := j.validateSetUpgradeTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeTimeout",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) PutDeltaHealthPolicy(value *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy) {
	if err := s.validatePutDeltaHealthPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeltaHealthPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) PutHealthPolicy(value *ServiceFabricClusterUpgradePolicyHealthPolicy) {
	if err := s.validatePutHealthPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHealthPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetDeltaHealthPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeltaHealthPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetForceRestartEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetForceRestartEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthCheckRetryTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckRetryTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthCheckStableDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckStableDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthCheckWaitDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckWaitDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetUpgradeDomainTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeDomainTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetUpgradeReplicaSetCheckTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeReplicaSetCheckTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetUpgradeTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


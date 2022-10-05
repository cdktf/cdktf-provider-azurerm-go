package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterNetworkProfileOutputReference interface {
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
	DnsServiceIp() *string
	SetDnsServiceIp(val *string)
	DnsServiceIpInput() *string
	DockerBridgeCidr() *string
	SetDockerBridgeCidr(val *string)
	DockerBridgeCidrInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterNetworkProfile
	SetInternalValue(val *KubernetesClusterNetworkProfile)
	IpVersions() *[]*string
	SetIpVersions(val *[]*string)
	IpVersionsInput() *[]*string
	LoadBalancerProfile() KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference
	LoadBalancerProfileInput() *KubernetesClusterNetworkProfileLoadBalancerProfile
	LoadBalancerSku() *string
	SetLoadBalancerSku(val *string)
	LoadBalancerSkuInput() *string
	NatGatewayProfile() KubernetesClusterNetworkProfileNatGatewayProfileOutputReference
	NatGatewayProfileInput() *KubernetesClusterNetworkProfileNatGatewayProfile
	NetworkMode() *string
	SetNetworkMode(val *string)
	NetworkModeInput() *string
	NetworkPlugin() *string
	SetNetworkPlugin(val *string)
	NetworkPluginInput() *string
	NetworkPolicy() *string
	SetNetworkPolicy(val *string)
	NetworkPolicyInput() *string
	OutboundType() *string
	SetOutboundType(val *string)
	OutboundTypeInput() *string
	PodCidr() *string
	SetPodCidr(val *string)
	PodCidrInput() *string
	ServiceCidr() *string
	SetServiceCidr(val *string)
	ServiceCidrInput() *string
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
	PutLoadBalancerProfile(value *KubernetesClusterNetworkProfileLoadBalancerProfile)
	PutNatGatewayProfile(value *KubernetesClusterNetworkProfileNatGatewayProfile)
	ResetDnsServiceIp()
	ResetDockerBridgeCidr()
	ResetIpVersions()
	ResetLoadBalancerProfile()
	ResetLoadBalancerSku()
	ResetNatGatewayProfile()
	ResetNetworkMode()
	ResetNetworkPolicy()
	ResetOutboundType()
	ResetPodCidr()
	ResetServiceCidr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterNetworkProfileOutputReference
type jsiiProxy_KubernetesClusterNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DnsServiceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServiceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DnsServiceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServiceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DockerBridgeCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerBridgeCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DockerBridgeCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerBridgeCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) InternalValue() *KubernetesClusterNetworkProfile {
	var returns *KubernetesClusterNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) IpVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) IpVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerProfile() KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference {
	var returns KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference
	_jsii_.Get(
		j,
		"loadBalancerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerProfileInput() *KubernetesClusterNetworkProfileLoadBalancerProfile {
	var returns *KubernetesClusterNetworkProfileLoadBalancerProfile
	_jsii_.Get(
		j,
		"loadBalancerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NatGatewayProfile() KubernetesClusterNetworkProfileNatGatewayProfileOutputReference {
	var returns KubernetesClusterNetworkProfileNatGatewayProfileOutputReference
	_jsii_.Get(
		j,
		"natGatewayProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NatGatewayProfileInput() *KubernetesClusterNetworkProfileNatGatewayProfile {
	var returns *KubernetesClusterNetworkProfileNatGatewayProfile
	_jsii_.Get(
		j,
		"natGatewayProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPlugin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPlugin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPluginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPluginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) OutboundType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) OutboundTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PodCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PodCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ServiceCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ServiceCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNetworkProfileOutputReference_Override(k KubernetesClusterNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetDnsServiceIp(val *string) {
	if err := j.validateSetDnsServiceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServiceIp",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetDockerBridgeCidr(val *string) {
	if err := j.validateSetDockerBridgeCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerBridgeCidr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetInternalValue(val *KubernetesClusterNetworkProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetIpVersions(val *[]*string) {
	if err := j.validateSetIpVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipVersions",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetLoadBalancerSku(val *string) {
	if err := j.validateSetLoadBalancerSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerSku",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetNetworkMode(val *string) {
	if err := j.validateSetNetworkModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkMode",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetNetworkPlugin(val *string) {
	if err := j.validateSetNetworkPluginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPlugin",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetNetworkPolicy(val *string) {
	if err := j.validateSetNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetOutboundType(val *string) {
	if err := j.validateSetOutboundTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetPodCidr(val *string) {
	if err := j.validateSetPodCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podCidr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetServiceCidr(val *string) {
	if err := j.validateSetServiceCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceCidr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PutLoadBalancerProfile(value *KubernetesClusterNetworkProfileLoadBalancerProfile) {
	if err := k.validatePutLoadBalancerProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLoadBalancerProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PutNatGatewayProfile(value *KubernetesClusterNetworkProfileNatGatewayProfile) {
	if err := k.validatePutNatGatewayProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNatGatewayProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetDnsServiceIp() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsServiceIp",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetDockerBridgeCidr() {
	_jsii_.InvokeVoid(
		k,
		"resetDockerBridgeCidr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetIpVersions() {
	_jsii_.InvokeVoid(
		k,
		"resetIpVersions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetLoadBalancerProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetLoadBalancerProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetLoadBalancerSku() {
	_jsii_.InvokeVoid(
		k,
		"resetLoadBalancerSku",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetNatGatewayProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetNatGatewayProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetNetworkMode() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetOutboundType() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetPodCidr() {
	_jsii_.InvokeVoid(
		k,
		"resetPodCidr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetServiceCidr() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceCidr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


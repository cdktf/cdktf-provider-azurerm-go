package vpngatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/vpngatewayconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnGatewayConnectionVpnLinkOutputReference interface {
	cdktf.ComplexObject
	BandwidthMbps() *float64
	SetBandwidthMbps(val *float64)
	BandwidthMbpsInput() *float64
	BgpEnabled() interface{}
	SetBgpEnabled(val interface{})
	BgpEnabledInput() interface{}
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
	ConnectionMode() *string
	SetConnectionMode(val *string)
	ConnectionModeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomBgpAddress() VpnGatewayConnectionVpnLinkCustomBgpAddressList
	CustomBgpAddressInput() interface{}
	EgressNatRuleIds() *[]*string
	SetEgressNatRuleIds(val *[]*string)
	EgressNatRuleIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	IngressNatRuleIds() *[]*string
	SetIngressNatRuleIds(val *[]*string)
	IngressNatRuleIdsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpsecPolicy() VpnGatewayConnectionVpnLinkIpsecPolicyList
	IpsecPolicyInput() interface{}
	LocalAzureIpAddressEnabled() interface{}
	SetLocalAzureIpAddressEnabled(val interface{})
	LocalAzureIpAddressEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PolicyBasedTrafficSelectorEnabled() interface{}
	SetPolicyBasedTrafficSelectorEnabled(val interface{})
	PolicyBasedTrafficSelectorEnabledInput() interface{}
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RatelimitEnabled() interface{}
	SetRatelimitEnabled(val interface{})
	RatelimitEnabledInput() interface{}
	RouteWeight() *float64
	SetRouteWeight(val *float64)
	RouteWeightInput() *float64
	SharedKey() *string
	SetSharedKey(val *string)
	SharedKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpnSiteLinkId() *string
	SetVpnSiteLinkId(val *string)
	VpnSiteLinkIdInput() *string
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
	PutCustomBgpAddress(value interface{})
	PutIpsecPolicy(value interface{})
	ResetBandwidthMbps()
	ResetBgpEnabled()
	ResetConnectionMode()
	ResetCustomBgpAddress()
	ResetEgressNatRuleIds()
	ResetIngressNatRuleIds()
	ResetIpsecPolicy()
	ResetLocalAzureIpAddressEnabled()
	ResetPolicyBasedTrafficSelectorEnabled()
	ResetProtocol()
	ResetRatelimitEnabled()
	ResetRouteWeight()
	ResetSharedKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionVpnLinkOutputReference
type jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BandwidthMbps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BandwidthMbpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthMbpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BgpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) BgpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ConnectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) CustomBgpAddress() VpnGatewayConnectionVpnLinkCustomBgpAddressList {
	var returns VpnGatewayConnectionVpnLinkCustomBgpAddressList
	_jsii_.Get(
		j,
		"customBgpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) CustomBgpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customBgpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) EgressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) EgressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IngressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IngressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IpsecPolicy() VpnGatewayConnectionVpnLinkIpsecPolicyList {
	var returns VpnGatewayConnectionVpnLinkIpsecPolicyList
	_jsii_.Get(
		j,
		"ipsecPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) IpsecPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipsecPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) LocalAzureIpAddressEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) LocalAzureIpAddressEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PolicyBasedTrafficSelectorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyBasedTrafficSelectorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PolicyBasedTrafficSelectorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyBasedTrafficSelectorEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RatelimitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ratelimitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RatelimitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ratelimitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RouteWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routeWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) RouteWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routeWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) SharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) VpnSiteLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnSiteLinkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) VpnSiteLinkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnSiteLinkIdInput",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnGatewayConnectionVpnLinkOutputReference {
	_init_.Initialize()

	if err := validateNewVpnGatewayConnectionVpnLinkOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkOutputReference_Override(v VpnGatewayConnectionVpnLinkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetBandwidthMbps(val *float64) {
	if err := j.validateSetBandwidthMbpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthMbps",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetBgpEnabled(val interface{}) {
	if err := j.validateSetBgpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetConnectionMode(val *string) {
	if err := j.validateSetConnectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetEgressNatRuleIds(val *[]*string) {
	if err := j.validateSetEgressNatRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetIngressNatRuleIds(val *[]*string) {
	if err := j.validateSetIngressNatRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetLocalAzureIpAddressEnabled(val interface{}) {
	if err := j.validateSetLocalAzureIpAddressEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAzureIpAddressEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetPolicyBasedTrafficSelectorEnabled(val interface{}) {
	if err := j.validateSetPolicyBasedTrafficSelectorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyBasedTrafficSelectorEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetRatelimitEnabled(val interface{}) {
	if err := j.validateSetRatelimitEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ratelimitEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetRouteWeight(val *float64) {
	if err := j.validateSetRouteWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeWeight",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetSharedKey(val *string) {
	if err := j.validateSetSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedKey",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference)SetVpnSiteLinkId(val *string) {
	if err := j.validateSetVpnSiteLinkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnSiteLinkId",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PutCustomBgpAddress(value interface{}) {
	if err := v.validatePutCustomBgpAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomBgpAddress",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) PutIpsecPolicy(value interface{}) {
	if err := v.validatePutIpsecPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpsecPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetBandwidthMbps() {
	_jsii_.InvokeVoid(
		v,
		"resetBandwidthMbps",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetBgpEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetBgpEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetConnectionMode() {
	_jsii_.InvokeVoid(
		v,
		"resetConnectionMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetCustomBgpAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomBgpAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetEgressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetEgressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetIngressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetIngressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetIpsecPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetIpsecPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetLocalAzureIpAddressEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalAzureIpAddressEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetPolicyBasedTrafficSelectorEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetPolicyBasedTrafficSelectorEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetRatelimitEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetRatelimitEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetRouteWeight() {
	_jsii_.InvokeVoid(
		v,
		"resetRouteWeight",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ResetSharedKey() {
	_jsii_.InvokeVoid(
		v,
		"resetSharedKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


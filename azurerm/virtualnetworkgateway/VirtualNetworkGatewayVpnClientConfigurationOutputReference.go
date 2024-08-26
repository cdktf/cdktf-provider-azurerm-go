// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/virtualnetworkgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualNetworkGatewayVpnClientConfigurationOutputReference interface {
	cdktf.ComplexObject
	AadAudience() *string
	SetAadAudience(val *string)
	AadAudienceInput() *string
	AadIssuer() *string
	SetAadIssuer(val *string)
	AadIssuerInput() *string
	AadTenant() *string
	SetAadTenant(val *string)
	AadTenantInput() *string
	AddressSpace() *[]*string
	SetAddressSpace(val *[]*string)
	AddressSpaceInput() *[]*string
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
	InternalValue() *VirtualNetworkGatewayVpnClientConfiguration
	SetInternalValue(val *VirtualNetworkGatewayVpnClientConfiguration)
	IpsecPolicy() VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference
	IpsecPolicyInput() *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy
	RadiusServer() VirtualNetworkGatewayVpnClientConfigurationRadiusServerList
	RadiusServerAddress() *string
	SetRadiusServerAddress(val *string)
	RadiusServerAddressInput() *string
	RadiusServerInput() interface{}
	RadiusServerSecret() *string
	SetRadiusServerSecret(val *string)
	RadiusServerSecretInput() *string
	RevokedCertificate() VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList
	RevokedCertificateInput() interface{}
	RootCertificate() VirtualNetworkGatewayVpnClientConfigurationRootCertificateList
	RootCertificateInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkGatewayClientConnection() VirtualNetworkGatewayVpnClientConfigurationVirtualNetworkGatewayClientConnectionList
	VirtualNetworkGatewayClientConnectionInput() interface{}
	VpnAuthTypes() *[]*string
	SetVpnAuthTypes(val *[]*string)
	VpnAuthTypesInput() *[]*string
	VpnClientProtocols() *[]*string
	SetVpnClientProtocols(val *[]*string)
	VpnClientProtocolsInput() *[]*string
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
	PutIpsecPolicy(value *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy)
	PutRadiusServer(value interface{})
	PutRevokedCertificate(value interface{})
	PutRootCertificate(value interface{})
	PutVirtualNetworkGatewayClientConnection(value interface{})
	ResetAadAudience()
	ResetAadIssuer()
	ResetAadTenant()
	ResetIpsecPolicy()
	ResetRadiusServer()
	ResetRadiusServerAddress()
	ResetRadiusServerSecret()
	ResetRevokedCertificate()
	ResetRootCertificate()
	ResetVirtualNetworkGatewayClientConnection()
	ResetVpnAuthTypes()
	ResetVpnClientProtocols()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayVpnClientConfigurationOutputReference
type jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadTenant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AadTenantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aadTenantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AddressSpace() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) AddressSpaceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) InternalValue() *VirtualNetworkGatewayVpnClientConfiguration {
	var returns *VirtualNetworkGatewayVpnClientConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) IpsecPolicy() VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference {
	var returns VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference
	_jsii_.Get(
		j,
		"ipsecPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) IpsecPolicyInput() *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy {
	var returns *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy
	_jsii_.Get(
		j,
		"ipsecPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServer() VirtualNetworkGatewayVpnClientConfigurationRadiusServerList {
	var returns VirtualNetworkGatewayVpnClientConfigurationRadiusServerList
	_jsii_.Get(
		j,
		"radiusServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"radiusServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RadiusServerSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiusServerSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RevokedCertificate() VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList {
	var returns VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList
	_jsii_.Get(
		j,
		"revokedCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RevokedCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokedCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RootCertificate() VirtualNetworkGatewayVpnClientConfigurationRootCertificateList {
	var returns VirtualNetworkGatewayVpnClientConfigurationRootCertificateList
	_jsii_.Get(
		j,
		"rootCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) RootCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VirtualNetworkGatewayClientConnection() VirtualNetworkGatewayVpnClientConfigurationVirtualNetworkGatewayClientConnectionList {
	var returns VirtualNetworkGatewayVpnClientConfigurationVirtualNetworkGatewayClientConnectionList
	_jsii_.Get(
		j,
		"virtualNetworkGatewayClientConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VirtualNetworkGatewayClientConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualNetworkGatewayClientConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnAuthTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnAuthTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnAuthTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnAuthTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnClientProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnClientProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) VpnClientProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnClientProtocolsInput",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayVpnClientConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayVpnClientConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualNetworkGatewayVpnClientConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayVpnClientConfigurationOutputReference_Override(v VirtualNetworkGatewayVpnClientConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetAadAudience(val *string) {
	if err := j.validateSetAadAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aadAudience",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetAadIssuer(val *string) {
	if err := j.validateSetAadIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aadIssuer",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetAadTenant(val *string) {
	if err := j.validateSetAadTenantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aadTenant",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetAddressSpace(val *[]*string) {
	if err := j.validateSetAddressSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSpace",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetInternalValue(val *VirtualNetworkGatewayVpnClientConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetRadiusServerAddress(val *string) {
	if err := j.validateSetRadiusServerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radiusServerAddress",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetRadiusServerSecret(val *string) {
	if err := j.validateSetRadiusServerSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radiusServerSecret",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetVpnAuthTypes(val *[]*string) {
	if err := j.validateSetVpnAuthTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnAuthTypes",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference)SetVpnClientProtocols(val *[]*string) {
	if err := j.validateSetVpnClientProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnClientProtocols",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) PutIpsecPolicy(value *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy) {
	if err := v.validatePutIpsecPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpsecPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) PutRadiusServer(value interface{}) {
	if err := v.validatePutRadiusServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRadiusServer",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) PutRevokedCertificate(value interface{}) {
	if err := v.validatePutRevokedCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRevokedCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) PutRootCertificate(value interface{}) {
	if err := v.validatePutRootCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRootCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) PutVirtualNetworkGatewayClientConnection(value interface{}) {
	if err := v.validatePutVirtualNetworkGatewayClientConnectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putVirtualNetworkGatewayClientConnection",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetAadAudience() {
	_jsii_.InvokeVoid(
		v,
		"resetAadAudience",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetAadIssuer() {
	_jsii_.InvokeVoid(
		v,
		"resetAadIssuer",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetAadTenant() {
	_jsii_.InvokeVoid(
		v,
		"resetAadTenant",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetIpsecPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetIpsecPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRadiusServer() {
	_jsii_.InvokeVoid(
		v,
		"resetRadiusServer",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRadiusServerAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetRadiusServerAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRadiusServerSecret() {
	_jsii_.InvokeVoid(
		v,
		"resetRadiusServerSecret",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRevokedCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetRevokedCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetRootCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetRootCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetVirtualNetworkGatewayClientConnection() {
	_jsii_.InvokeVoid(
		v,
		"resetVirtualNetworkGatewayClientConnection",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetVpnAuthTypes() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnAuthTypes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ResetVpnClientProtocols() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnClientProtocols",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


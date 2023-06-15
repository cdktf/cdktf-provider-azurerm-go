package keyvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/keyvault/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultNetworkAclsOutputReference interface {
	cdktf.ComplexObject
	Bypass() *string
	SetBypass(val *string)
	BypassInput() *string
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
	DefaultAction() *string
	SetDefaultAction(val *string)
	DefaultActionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KeyVaultNetworkAcls
	SetInternalValue(val *KeyVaultNetworkAcls)
	IpRules() *[]*string
	SetIpRules(val *[]*string)
	IpRulesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkSubnetIds() *[]*string
	SetVirtualNetworkSubnetIds(val *[]*string)
	VirtualNetworkSubnetIdsInput() *[]*string
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
	ResetIpRules()
	ResetVirtualNetworkSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyVaultNetworkAclsOutputReference
type jsiiProxy_KeyVaultNetworkAclsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) Bypass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) BypassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) DefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) DefaultActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) InternalValue() *KeyVaultNetworkAcls {
	var returns *KeyVaultNetworkAcls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) IpRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) IpRulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) VirtualNetworkSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference) VirtualNetworkSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualNetworkSubnetIdsInput",
		&returns,
	)
	return returns
}


func NewKeyVaultNetworkAclsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KeyVaultNetworkAclsOutputReference {
	_init_.Initialize()

	if err := validateNewKeyVaultNetworkAclsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultNetworkAclsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVault.KeyVaultNetworkAclsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKeyVaultNetworkAclsOutputReference_Override(k KeyVaultNetworkAclsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVault.KeyVaultNetworkAclsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetBypass(val *string) {
	if err := j.validateSetBypassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypass",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetDefaultAction(val *string) {
	if err := j.validateSetDefaultActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAction",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetInternalValue(val *KeyVaultNetworkAcls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetIpRules(val *[]*string) {
	if err := j.validateSetIpRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRules",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KeyVaultNetworkAclsOutputReference)SetVirtualNetworkSubnetIds(val *[]*string) {
	if err := j.validateSetVirtualNetworkSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkSubnetIds",
		val,
	)
}

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) ResetIpRules() {
	_jsii_.InvokeVoid(
		k,
		"resetIpRules",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) ResetVirtualNetworkSubnetIds() {
	_jsii_.InvokeVoid(
		k,
		"resetVirtualNetworkSubnetIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KeyVaultNetworkAclsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


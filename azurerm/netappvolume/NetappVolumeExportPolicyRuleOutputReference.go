package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/netappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeExportPolicyRuleOutputReference interface {
	cdktf.ComplexObject
	AllowedClients() *[]*string
	SetAllowedClients(val *[]*string)
	AllowedClientsInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProtocolsEnabled() *[]*string
	SetProtocolsEnabled(val *[]*string)
	ProtocolsEnabledInput() *[]*string
	RootAccessEnabled() interface{}
	SetRootAccessEnabled(val interface{})
	RootAccessEnabledInput() interface{}
	RuleIndex() *float64
	SetRuleIndex(val *float64)
	RuleIndexInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnixReadOnly() interface{}
	SetUnixReadOnly(val interface{})
	UnixReadOnlyInput() interface{}
	UnixReadWrite() interface{}
	SetUnixReadWrite(val interface{})
	UnixReadWriteInput() interface{}
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
	ResetProtocolsEnabled()
	ResetRootAccessEnabled()
	ResetUnixReadOnly()
	ResetUnixReadWrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeExportPolicyRuleOutputReference
type jsiiProxy_NetappVolumeExportPolicyRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) AllowedClients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) AllowedClientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ProtocolsEnabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ProtocolsEnabledInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RootAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RootAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RuleIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RuleIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadWriteInput",
		&returns,
	)
	return returns
}


func NewNetappVolumeExportPolicyRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetappVolumeExportPolicyRuleOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeExportPolicyRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeExportPolicyRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeExportPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetappVolumeExportPolicyRuleOutputReference_Override(n NetappVolumeExportPolicyRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeExportPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetAllowedClients(val *[]*string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetProtocolsEnabled(val *[]*string) {
	if err := j.validateSetProtocolsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolsEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetRootAccessEnabled(val interface{}) {
	if err := j.validateSetRootAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetRuleIndex(val *float64) {
	if err := j.validateSetRuleIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetUnixReadOnly(val interface{}) {
	if err := j.validateSetUnixReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unixReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference)SetUnixReadWrite(val interface{}) {
	if err := j.validateSetUnixReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unixReadWrite",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetProtocolsEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocolsEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetRootAccessEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetRootAccessEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetUnixReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetUnixReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetUnixReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetUnixReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


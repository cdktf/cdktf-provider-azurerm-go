package firewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/firewallpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallPolicyIntrusionDetectionOutputReference interface {
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
	InternalValue() *FirewallPolicyIntrusionDetection
	SetInternalValue(val *FirewallPolicyIntrusionDetection)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	PrivateRanges() *[]*string
	SetPrivateRanges(val *[]*string)
	PrivateRangesInput() *[]*string
	SignatureOverrides() FirewallPolicyIntrusionDetectionSignatureOverridesList
	SignatureOverridesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficBypass() FirewallPolicyIntrusionDetectionTrafficBypassList
	TrafficBypassInput() interface{}
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
	PutSignatureOverrides(value interface{})
	PutTrafficBypass(value interface{})
	ResetMode()
	ResetPrivateRanges()
	ResetSignatureOverrides()
	ResetTrafficBypass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyIntrusionDetectionOutputReference
type jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) InternalValue() *FirewallPolicyIntrusionDetection {
	var returns *FirewallPolicyIntrusionDetection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) PrivateRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) PrivateRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) SignatureOverrides() FirewallPolicyIntrusionDetectionSignatureOverridesList {
	var returns FirewallPolicyIntrusionDetectionSignatureOverridesList
	_jsii_.Get(
		j,
		"signatureOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) SignatureOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signatureOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) TrafficBypass() FirewallPolicyIntrusionDetectionTrafficBypassList {
	var returns FirewallPolicyIntrusionDetectionTrafficBypassList
	_jsii_.Get(
		j,
		"trafficBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) TrafficBypassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficBypassInput",
		&returns,
	)
	return returns
}


func NewFirewallPolicyIntrusionDetectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FirewallPolicyIntrusionDetectionOutputReference {
	_init_.Initialize()

	if err := validateNewFirewallPolicyIntrusionDetectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicyIntrusionDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFirewallPolicyIntrusionDetectionOutputReference_Override(f FirewallPolicyIntrusionDetectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicy.FirewallPolicyIntrusionDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference)SetInternalValue(val *FirewallPolicyIntrusionDetection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference)SetPrivateRanges(val *[]*string) {
	if err := j.validateSetPrivateRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateRanges",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) PutSignatureOverrides(value interface{}) {
	if err := f.validatePutSignatureOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSignatureOverrides",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) PutTrafficBypass(value interface{}) {
	if err := f.validatePutTrafficBypassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTrafficBypass",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		f,
		"resetMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ResetPrivateRanges() {
	_jsii_.InvokeVoid(
		f,
		"resetPrivateRanges",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ResetSignatureOverrides() {
	_jsii_.InvokeVoid(
		f,
		"resetSignatureOverrides",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ResetTrafficBypass() {
	_jsii_.InvokeVoid(
		f,
		"resetTrafficBypass",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyIntrusionDetectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


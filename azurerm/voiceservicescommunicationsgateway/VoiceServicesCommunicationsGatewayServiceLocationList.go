package voiceservicescommunicationsgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/voiceservicescommunicationsgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VoiceServicesCommunicationsGatewayServiceLocationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) VoiceServicesCommunicationsGatewayServiceLocationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VoiceServicesCommunicationsGatewayServiceLocationList
type jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVoiceServicesCommunicationsGatewayServiceLocationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VoiceServicesCommunicationsGatewayServiceLocationList {
	_init_.Initialize()

	if err := validateNewVoiceServicesCommunicationsGatewayServiceLocationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.voiceServicesCommunicationsGateway.VoiceServicesCommunicationsGatewayServiceLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVoiceServicesCommunicationsGatewayServiceLocationList_Override(v VoiceServicesCommunicationsGatewayServiceLocationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.voiceServicesCommunicationsGateway.VoiceServicesCommunicationsGatewayServiceLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) Get(index *float64) VoiceServicesCommunicationsGatewayServiceLocationOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VoiceServicesCommunicationsGatewayServiceLocationOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


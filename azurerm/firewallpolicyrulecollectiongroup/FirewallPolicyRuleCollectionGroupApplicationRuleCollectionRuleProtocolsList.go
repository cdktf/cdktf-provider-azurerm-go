package firewallpolicyrulecollectiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/firewallpolicyrulecollectiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList interface {
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
	Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList
type jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList {
	_init_.Initialize()

	if err := validateNewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList_Override(f FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) Get(index *float64) FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference {
	if err := f.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocolsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

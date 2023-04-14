package sentinelalertruleanomalybuiltin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/sentinelalertruleanomalybuiltin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) SentinelAlertRuleAnomalyBuiltInMultiSelectObservationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList
type jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSentinelAlertRuleAnomalyBuiltInMultiSelectObservationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList {
	_init_.Initialize()

	if err := validateNewSentinelAlertRuleAnomalyBuiltInMultiSelectObservationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyBuiltIn.SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSentinelAlertRuleAnomalyBuiltInMultiSelectObservationList_Override(s SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelAlertRuleAnomalyBuiltIn.SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) Get(index *float64) SentinelAlertRuleAnomalyBuiltInMultiSelectObservationOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns SentinelAlertRuleAnomalyBuiltInMultiSelectObservationOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SentinelAlertRuleAnomalyBuiltInMultiSelectObservationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


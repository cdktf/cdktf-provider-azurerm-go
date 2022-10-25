//go:build !no_runtime_type_checking

package eventgridsystemtopiceventsubscription

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutBoolEqualsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEquals:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEquals)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEquals:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEquals)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEquals; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutIsNotNullParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNull:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNull)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNull:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNull)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNull; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutIsNullOrUndefinedParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefined:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefined)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefined:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefined)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefined; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberGreaterThanParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThan:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThan)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThan:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThan)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThan; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberGreaterThanOrEqualsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberIn:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberIn:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberInRangeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRange:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRange)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRange:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRange)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRange; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberLessThanParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThan:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThan)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThan:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThan)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThan; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberLessThanOrEqualsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEquals:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEquals)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEquals:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEquals)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEquals; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberNotInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotIn:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotIn:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutNumberNotInRangeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRange:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRange)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRange:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRange)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRange; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringBeginsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWith:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWith:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringContainsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringContains:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringContains)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringContains:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringContains)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringContains; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringEndsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWith:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWith:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringIn:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringIn:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotBeginsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWith:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWith:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotContainsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContains:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContains)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContains:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContains)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContains; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotEndsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWith:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWith:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotIn:
		value := value.(*[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotIn:
		value_ := value.([]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateSetInternalValueParameters(val *EventgridSystemTopicEventSubscriptionAdvancedFilter) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEventgridSystemTopicEventSubscriptionAdvancedFilterOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}


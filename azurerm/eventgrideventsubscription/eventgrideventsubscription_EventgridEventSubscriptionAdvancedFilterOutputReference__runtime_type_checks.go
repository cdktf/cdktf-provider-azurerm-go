//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package eventgrideventsubscription

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutBoolEqualsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterBoolEquals:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterBoolEquals)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterBoolEquals:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterBoolEquals)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterBoolEquals; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutIsNotNullParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterIsNotNull:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterIsNotNull)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterIsNotNull:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterIsNotNull)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterIsNotNull; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutIsNullOrUndefinedParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterIsNullOrUndefined:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterIsNullOrUndefined)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterIsNullOrUndefined:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterIsNullOrUndefined)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterIsNullOrUndefined; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberGreaterThanParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThan:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThan)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberGreaterThan:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThan)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThan; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberGreaterThanOrEqualsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberIn:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberIn:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberInRangeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberInRange:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberInRange)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberInRange:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberInRange)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberInRange; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberLessThanParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberLessThan:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberLessThan)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberLessThan:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberLessThan)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberLessThan; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberLessThanOrEqualsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEquals:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEquals)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEquals:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEquals)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEquals; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberNotInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberNotIn:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberNotIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberNotIn:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberNotIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberNotIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutNumberNotInRangeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterNumberNotInRange:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterNumberNotInRange)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterNumberNotInRange:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterNumberNotInRange)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterNumberNotInRange; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringBeginsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringBeginsWith:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringBeginsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringBeginsWith:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringBeginsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringBeginsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringContainsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringContains:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringContains)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringContains:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringContains)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringContains; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringEndsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringEndsWith:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringEndsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringEndsWith:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringEndsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringEndsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringIn:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringIn:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotBeginsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringNotBeginsWith:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringNotBeginsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringNotBeginsWith:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringNotBeginsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringNotBeginsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotContainsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringNotContains:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringNotContains)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringNotContains:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringNotContains)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringNotContains; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotEndsWithParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringNotEndsWith:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringNotEndsWith)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringNotEndsWith:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringNotEndsWith)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringNotEndsWith; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validatePutStringNotInParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*EventgridEventSubscriptionAdvancedFilterStringNotIn:
		value := value.(*[]*EventgridEventSubscriptionAdvancedFilterStringNotIn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*EventgridEventSubscriptionAdvancedFilterStringNotIn:
		value_ := value.([]*EventgridEventSubscriptionAdvancedFilterStringNotIn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*EventgridEventSubscriptionAdvancedFilterStringNotIn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateSetInternalValueParameters(val *EventgridEventSubscriptionAdvancedFilter) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEventgridEventSubscriptionAdvancedFilterOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}


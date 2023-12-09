// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package paloaltovirtualnetworkappliance

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateAddMoveTargetParameters(moveTarget *string) error {
	if moveTarget == nil {
		return fmt.Errorf("parameter moveTarget is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateImportFromParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateMoveFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateMoveToParameters(moveTarget *string, index interface{}) error {
	if moveTarget == nil {
		return fmt.Errorf("parameter moveTarget is required, but nil was provided")
	}

	switch index.(type) {
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
		return fmt.Errorf("parameter index must be one of the allowed types: *string, *float64; received %#v (a %T)", index, index)
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateMoveToIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoVirtualNetworkAppliance) validatePutTimeoutsParameters(value *PaloAltoVirtualNetworkApplianceTimeouts) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func validatePaloAltoVirtualNetworkAppliance_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if importToId == nil {
		return fmt.Errorf("parameter importToId is required, but nil was provided")
	}

	if importFromId == nil {
		return fmt.Errorf("parameter importFromId is required, but nil was provided")
	}

	return nil
}

func validatePaloAltoVirtualNetworkAppliance_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePaloAltoVirtualNetworkAppliance_IsTerraformElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePaloAltoVirtualNetworkAppliance_IsTerraformResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateSetConnectionParameters(val interface{}) error {
	switch val.(type) {
	case *cdktf.SSHProvisionerConnection:
		val := val.(*cdktf.SSHProvisionerConnection)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktf.SSHProvisionerConnection:
		val_ := val.(cdktf.SSHProvisionerConnection)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case *cdktf.WinrmProvisionerConnection:
		val := val.(*cdktf.WinrmProvisionerConnection)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktf.WinrmProvisionerConnection:
		val_ := val.(cdktf.WinrmProvisionerConnection)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *cdktf.SSHProvisionerConnection, *cdktf.WinrmProvisionerConnection; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateSetCountParameters(val interface{}) error {
	switch val.(type) {
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
	case cdktf.TerraformCount:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *float64, cdktf.TerraformCount; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateSetIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateSetProvisionersParameters(val *[]interface{}) error {
	for idx_97dfc6, v := range *val {
		switch v.(type) {
		case *cdktf.FileProvisioner:
			v := v.(*cdktf.FileProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktf.FileProvisioner:
			v_ := v.(cdktf.FileProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case *cdktf.LocalExecProvisioner:
			v := v.(*cdktf.LocalExecProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktf.LocalExecProvisioner:
			v_ := v.(cdktf.LocalExecProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case *cdktf.RemoteExecProvisioner:
			v := v.(*cdktf.RemoteExecProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktf.RemoteExecProvisioner:
			v_ := v.(cdktf.RemoteExecProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *cdktf.FileProvisioner, *cdktf.LocalExecProvisioner, *cdktf.RemoteExecProvisioner; received %#v (a %T)", idx_97dfc6, v, v)
			}
		}
	}

	return nil
}

func (j *jsiiProxy_PaloAltoVirtualNetworkAppliance) validateSetVirtualHubIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPaloAltoVirtualNetworkApplianceParameters(scope constructs.Construct, id *string, config *PaloAltoVirtualNetworkApplianceConfig) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package hdinsightinteractivequerycluster

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActions:
		val := val.(*[]*HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActions)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActions:
		val_ := val.([]*HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActions)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActions; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewHdinsightInteractiveQueryClusterRolesWorkerNodeScriptActionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}


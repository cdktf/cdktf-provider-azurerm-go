// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package firewallpolicyrulecollectiongroup

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule:
		val := val.(*[]*FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule:
		val_ := val.([]*FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRule; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewFirewallPolicyRuleCollectionGroupNetworkRuleCollectionRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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


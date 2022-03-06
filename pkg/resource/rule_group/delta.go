// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package rule_group

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.Capacity, b.ko.Spec.Capacity) {
		delta.Add("Spec.Capacity", a.ko.Spec.Capacity, b.ko.Spec.Capacity)
	} else if a.ko.Spec.Capacity != nil && b.ko.Spec.Capacity != nil {
		if *a.ko.Spec.Capacity != *b.ko.Spec.Capacity {
			delta.Add("Spec.Capacity", a.ko.Spec.Capacity, b.ko.Spec.Capacity)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DryRun, b.ko.Spec.DryRun) {
		delta.Add("Spec.DryRun", a.ko.Spec.DryRun, b.ko.Spec.DryRun)
	} else if a.ko.Spec.DryRun != nil && b.ko.Spec.DryRun != nil {
		if *a.ko.Spec.DryRun != *b.ko.Spec.DryRun {
			delta.Add("Spec.DryRun", a.ko.Spec.DryRun, b.ko.Spec.DryRun)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup, b.ko.Spec.RuleGroup) {
		delta.Add("Spec.RuleGroup", a.ko.Spec.RuleGroup, b.ko.Spec.RuleGroup)
	} else if a.ko.Spec.RuleGroup != nil && b.ko.Spec.RuleGroup != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RuleVariables, b.ko.Spec.RuleGroup.RuleVariables) {
			delta.Add("Spec.RuleGroup.RuleVariables", a.ko.Spec.RuleGroup.RuleVariables, b.ko.Spec.RuleGroup.RuleVariables)
		} else if a.ko.Spec.RuleGroup.RuleVariables != nil && b.ko.Spec.RuleGroup.RuleVariables != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RuleVariables.IPSets, b.ko.Spec.RuleGroup.RuleVariables.IPSets) {
				delta.Add("Spec.RuleGroup.RuleVariables.IPSets", a.ko.Spec.RuleGroup.RuleVariables.IPSets, b.ko.Spec.RuleGroup.RuleVariables.IPSets)
			} else if a.ko.Spec.RuleGroup.RuleVariables.IPSets != nil && b.ko.Spec.RuleGroup.RuleVariables.IPSets != nil {
				if !reflect.DeepEqual(a.ko.Spec.RuleGroup.RuleVariables.IPSets, b.ko.Spec.RuleGroup.RuleVariables.IPSets) {
					delta.Add("Spec.RuleGroup.RuleVariables.IPSets", a.ko.Spec.RuleGroup.RuleVariables.IPSets, b.ko.Spec.RuleGroup.RuleVariables.IPSets)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RuleVariables.PortSets, b.ko.Spec.RuleGroup.RuleVariables.PortSets) {
				delta.Add("Spec.RuleGroup.RuleVariables.PortSets", a.ko.Spec.RuleGroup.RuleVariables.PortSets, b.ko.Spec.RuleGroup.RuleVariables.PortSets)
			} else if a.ko.Spec.RuleGroup.RuleVariables.PortSets != nil && b.ko.Spec.RuleGroup.RuleVariables.PortSets != nil {
				if !reflect.DeepEqual(a.ko.Spec.RuleGroup.RuleVariables.PortSets, b.ko.Spec.RuleGroup.RuleVariables.PortSets) {
					delta.Add("Spec.RuleGroup.RuleVariables.PortSets", a.ko.Spec.RuleGroup.RuleVariables.PortSets, b.ko.Spec.RuleGroup.RuleVariables.PortSets)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RulesSource, b.ko.Spec.RuleGroup.RulesSource) {
			delta.Add("Spec.RuleGroup.RulesSource", a.ko.Spec.RuleGroup.RulesSource, b.ko.Spec.RuleGroup.RulesSource)
		} else if a.ko.Spec.RuleGroup.RulesSource != nil && b.ko.Spec.RuleGroup.RulesSource != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RulesSource.RulesSourceList, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList) {
				delta.Add("Spec.RuleGroup.RulesSource.RulesSourceList", a.ko.Spec.RuleGroup.RulesSource.RulesSourceList, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList)
			} else if a.ko.Spec.RuleGroup.RulesSource.RulesSourceList != nil && b.ko.Spec.RuleGroup.RulesSource.RulesSourceList != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType) {
					delta.Add("Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType", a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType)
				} else if a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType != nil && b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType != nil {
					if *a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType != *b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType {
						delta.Add("Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType", a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.GeneratedRulesType)
					}
				}
				if !ackcompare.SliceStringPEqual(a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.TargetTypes, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.TargetTypes) {
					delta.Add("Spec.RuleGroup.RulesSource.RulesSourceList.TargetTypes", a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.TargetTypes, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.TargetTypes)
				}
				if !ackcompare.SliceStringPEqual(a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.Targets, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.Targets) {
					delta.Add("Spec.RuleGroup.RulesSource.RulesSourceList.Targets", a.ko.Spec.RuleGroup.RulesSource.RulesSourceList.Targets, b.ko.Spec.RuleGroup.RulesSource.RulesSourceList.Targets)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RulesSource.RulesString, b.ko.Spec.RuleGroup.RulesSource.RulesString) {
				delta.Add("Spec.RuleGroup.RulesSource.RulesString", a.ko.Spec.RuleGroup.RulesSource.RulesString, b.ko.Spec.RuleGroup.RulesSource.RulesString)
			} else if a.ko.Spec.RuleGroup.RulesSource.RulesString != nil && b.ko.Spec.RuleGroup.RulesSource.RulesString != nil {
				if *a.ko.Spec.RuleGroup.RulesSource.RulesString != *b.ko.Spec.RuleGroup.RulesSource.RulesString {
					delta.Add("Spec.RuleGroup.RulesSource.RulesString", a.ko.Spec.RuleGroup.RulesSource.RulesString, b.ko.Spec.RuleGroup.RulesSource.RulesString)
				}
			}
			if !reflect.DeepEqual(a.ko.Spec.RuleGroup.RulesSource.StatefulRules, b.ko.Spec.RuleGroup.RulesSource.StatefulRules) {
				delta.Add("Spec.RuleGroup.RulesSource.StatefulRules", a.ko.Spec.RuleGroup.RulesSource.StatefulRules, b.ko.Spec.RuleGroup.RulesSource.StatefulRules)
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions, b.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions) {
				delta.Add("Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions", a.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions, b.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions)
			} else if a.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions != nil && b.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions != nil {
				if !reflect.DeepEqual(a.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.CustomActions, b.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.CustomActions) {
					delta.Add("Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.CustomActions", a.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.CustomActions, b.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.CustomActions)
				}
				if !reflect.DeepEqual(a.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.StatelessRules, b.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.StatelessRules) {
					delta.Add("Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.StatelessRules", a.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.StatelessRules, b.ko.Spec.RuleGroup.RulesSource.StatelessRulesAndCustomActions.StatelessRules)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.StatefulRuleOptions, b.ko.Spec.RuleGroup.StatefulRuleOptions) {
			delta.Add("Spec.RuleGroup.StatefulRuleOptions", a.ko.Spec.RuleGroup.StatefulRuleOptions, b.ko.Spec.RuleGroup.StatefulRuleOptions)
		} else if a.ko.Spec.RuleGroup.StatefulRuleOptions != nil && b.ko.Spec.RuleGroup.StatefulRuleOptions != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder, b.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder) {
				delta.Add("Spec.RuleGroup.StatefulRuleOptions.RuleOrder", a.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder, b.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder)
			} else if a.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder != nil && b.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder != nil {
				if *a.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder != *b.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder {
					delta.Add("Spec.RuleGroup.StatefulRuleOptions.RuleOrder", a.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder, b.ko.Spec.RuleGroup.StatefulRuleOptions.RuleOrder)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RuleGroupName, b.ko.Spec.RuleGroupName) {
		delta.Add("Spec.RuleGroupName", a.ko.Spec.RuleGroupName, b.ko.Spec.RuleGroupName)
	} else if a.ko.Spec.RuleGroupName != nil && b.ko.Spec.RuleGroupName != nil {
		if *a.ko.Spec.RuleGroupName != *b.ko.Spec.RuleGroupName {
			delta.Add("Spec.RuleGroupName", a.ko.Spec.RuleGroupName, b.ko.Spec.RuleGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Rules, b.ko.Spec.Rules) {
		delta.Add("Spec.Rules", a.ko.Spec.Rules, b.ko.Spec.Rules)
	} else if a.ko.Spec.Rules != nil && b.ko.Spec.Rules != nil {
		if *a.ko.Spec.Rules != *b.ko.Spec.Rules {
			delta.Add("Spec.Rules", a.ko.Spec.Rules, b.ko.Spec.Rules)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Type, b.ko.Spec.Type) {
		delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
	} else if a.ko.Spec.Type != nil && b.ko.Spec.Type != nil {
		if *a.ko.Spec.Type != *b.ko.Spec.Type {
			delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
		}
	}

	return delta
}
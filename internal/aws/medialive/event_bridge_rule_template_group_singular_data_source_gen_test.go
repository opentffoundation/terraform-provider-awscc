// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package medialive_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMediaLiveEventBridgeRuleTemplateGroupDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaLive::EventBridgeRuleTemplateGroup", "awscc_medialive_event_bridge_rule_template_group", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSMediaLiveEventBridgeRuleTemplateGroupDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaLive::EventBridgeRuleTemplateGroup", "awscc_medialive_event_bridge_rule_template_group", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}

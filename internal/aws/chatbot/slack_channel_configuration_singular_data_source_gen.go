// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package chatbot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_chatbot_slack_channel_configuration", slackChannelConfigurationDataSource)
}

// slackChannelConfigurationDataSource returns the Terraform awscc_chatbot_slack_channel_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::Chatbot::SlackChannelConfiguration resource.
func slackChannelConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the configuration",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the configuration",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9-_]+$",
		//	  "type": "string"
		//	}
		"configuration_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GuardrailPolicies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"guardrail_policies": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IamRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the IAM role that defines the permissions for AWS Chatbot",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the IAM role that defines the permissions for AWS Chatbot",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoggingLevel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "NONE",
		//	  "description": "Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs",
		//	  "pattern": "^(ERROR|INFO|NONE)$",
		//	  "type": "string"
		//	}
		"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SlackChannelId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the Slack channel",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9]+$",
		//	  "type": "string"
		//	}
		"slack_channel_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the Slack channel",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SlackWorkspaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the Slack workspace",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Z]{1,255}$",
		//	  "type": "string"
		//	}
		"slack_workspace_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the Slack workspace",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"sns_topic_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserRoleRequired
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Enables use of a user role requirement in your chat configuration",
		//	  "type": "boolean"
		//	}
		"user_role_required": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Enables use of a user role requirement in your chat configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Chatbot::SlackChannelConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Chatbot::SlackChannelConfiguration").WithTerraformTypeName("awscc_chatbot_slack_channel_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"configuration_name": "ConfigurationName",
		"guardrail_policies": "GuardrailPolicies",
		"iam_role_arn":       "IamRoleArn",
		"logging_level":      "LoggingLevel",
		"slack_channel_id":   "SlackChannelId",
		"slack_workspace_id": "SlackWorkspaceId",
		"sns_topic_arns":     "SnsTopicArns",
		"user_role_required": "UserRoleRequired",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_wafv2_logging_configuration", loggingConfigurationDataSource)
}

// loggingConfigurationDataSource returns the Terraform awscc_wafv2_logging_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::WAFv2::LoggingConfiguration resource.
func loggingConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: LogDestinationConfigs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"log_destination_configs": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoggingFilter
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.",
		//	  "properties": {
		//	    "DefaultBehavior": {
		//	      "description": "Default handling for logs that don't match any of the specified filtering conditions.",
		//	      "enum": [
		//	        "KEEP",
		//	        "DROP"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Filters": {
		//	      "description": "The filters that you want to apply to the logs.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Behavior": {
		//	            "description": "How to handle logs that satisfy the filter's conditions and requirement. ",
		//	            "enum": [
		//	              "KEEP",
		//	              "DROP"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Conditions": {
		//	            "description": "Match conditions for the filter.",
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "ActionCondition": {
		//	                  "additionalProperties": false,
		//	                  "description": "A single action condition.",
		//	                  "properties": {
		//	                    "Action": {
		//	                      "description": "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
		//	                      "enum": [
		//	                        "ALLOW",
		//	                        "BLOCK",
		//	                        "COUNT",
		//	                        "CAPTCHA",
		//	                        "CHALLENGE",
		//	                        "EXCLUDED_AS_COUNT"
		//	                      ],
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "required": [
		//	                    "Action"
		//	                  ],
		//	                  "type": "object"
		//	                },
		//	                "LabelNameCondition": {
		//	                  "additionalProperties": false,
		//	                  "description": "A single label name condition.",
		//	                  "properties": {
		//	                    "LabelName": {
		//	                      "description": "The label name that a log record must contain in order to meet the condition. This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. ",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "required": [
		//	                    "LabelName"
		//	                  ],
		//	                  "type": "object"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "minItems": 1,
		//	            "type": "array"
		//	          },
		//	          "Requirement": {
		//	            "description": "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
		//	            "enum": [
		//	              "MEETS_ALL",
		//	              "MEETS_ANY"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Behavior",
		//	          "Conditions",
		//	          "Requirement"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "DefaultBehavior",
		//	    "Filters"
		//	  ],
		//	  "type": "object"
		//	}
		"logging_filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DefaultBehavior
				"default_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Default handling for logs that don't match any of the specified filtering conditions.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Filters
				"filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Behavior
							"behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "How to handle logs that satisfy the filter's conditions and requirement. ",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Conditions
							"conditions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ActionCondition
										"action_condition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Action
												"action": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Description: "A single action condition.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: LabelNameCondition
										"label_name_condition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: LabelName
												"label_name": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "The label name that a log record must contain in order to meet the condition. This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. ",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Description: "A single label name condition.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "Match conditions for the filter.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Requirement
							"requirement": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The filters that you want to apply to the logs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ManagedByFirewallManager
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.",
		//	  "type": "boolean"
		//	}
		"managed_by_firewall_manager": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RedactedFields
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Method": {
		//	        "description": "Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform. ",
		//	        "type": "object"
		//	      },
		//	      "QueryString": {
		//	        "description": "Inspect the query string. This is the part of a URL that appears after a ? character, if any. ",
		//	        "type": "object"
		//	      },
		//	      "SingleHeader": {
		//	        "additionalProperties": false,
		//	        "description": "Inspect a single header. Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive.",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "The name of the query header to inspect.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "UriPath": {
		//	        "description": "Inspect the request URI path. This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg. ",
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"redacted_fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Method
					"method": schema.StringAttribute{ /*START ATTRIBUTE*/
						CustomType:  jsontypes.NormalizedType{},
						Description: "Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: QueryString
					"query_string": schema.StringAttribute{ /*START ATTRIBUTE*/
						CustomType:  jsontypes.NormalizedType{},
						Description: "Inspect the query string. This is the part of a URL that appears after a ? character, if any. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SingleHeader
					"single_header": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the query header to inspect.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Inspect a single header. Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: UriPath
					"uri_path": schema.StringAttribute{ /*START ATTRIBUTE*/
						CustomType:  jsontypes.NormalizedType{},
						Description: "Inspect the request URI path. This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::WAFv2::LoggingConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::LoggingConfiguration").WithTerraformTypeName("awscc_wafv2_logging_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                      "Action",
		"action_condition":            "ActionCondition",
		"behavior":                    "Behavior",
		"conditions":                  "Conditions",
		"default_behavior":            "DefaultBehavior",
		"filters":                     "Filters",
		"label_name":                  "LabelName",
		"label_name_condition":        "LabelNameCondition",
		"log_destination_configs":     "LogDestinationConfigs",
		"logging_filter":              "LoggingFilter",
		"managed_by_firewall_manager": "ManagedByFirewallManager",
		"method":                      "Method",
		"name":                        "Name",
		"query_string":                "QueryString",
		"redacted_fields":             "RedactedFields",
		"requirement":                 "Requirement",
		"resource_arn":                "ResourceArn",
		"single_header":               "SingleHeader",
		"uri_path":                    "UriPath",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package evidently

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_evidently_launch", launchDataSource)
}

// launchDataSource returns the Terraform awscc_evidently_launch data source.
// This Terraform data source corresponds to the CloudFormation AWS::Evidently::Launch resource.
func launchDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*/launch/[-a-zA-Z0-9._]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 160,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"execution_status": {
			// Property: ExecutionStatus
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Start or Stop Launch Launch. Default is not started.",
			//   "properties": {
			//     "DesiredState": {
			//       "description": "Provide CANCELLED or COMPLETED as the launch desired state. Defaults to Completed if not provided.",
			//       "type": "string"
			//     },
			//     "Reason": {
			//       "description": "Provide a reason for stopping the launch. Defaults to empty if not provided.",
			//       "type": "string"
			//     },
			//     "Status": {
			//       "description": "Provide START or STOP action to apply on a launch",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Status"
			//   ],
			//   "type": "object"
			// }
			Description: "Start or Stop Launch Launch. Default is not started.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"desired_state": {
						// Property: DesiredState
						Description: "Provide CANCELLED or COMPLETED as the launch desired state. Defaults to Completed if not provided.",
						Type:        types.StringType,
						Computed:    true,
					},
					"reason": {
						// Property: Reason
						Description: "Provide a reason for stopping the launch. Defaults to empty if not provided.",
						Type:        types.StringType,
						Computed:    true,
					},
					"status": {
						// Property: Status
						Description: "Provide START or STOP action to apply on a launch",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"groups": {
			// Property: Groups
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Description": {
			//         "maxLength": 160,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "Feature": {
			//         "type": "string"
			//       },
			//       "GroupName": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "[-a-zA-Z0-9._]*",
			//         "type": "string"
			//       },
			//       "Variation": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "GroupName",
			//       "Feature",
			//       "Variation"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 5,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Type:     types.StringType,
						Computed: true,
					},
					"feature": {
						// Property: Feature
						Type:     types.StringType,
						Computed: true,
					},
					"group_name": {
						// Property: GroupName
						Type:     types.StringType,
						Computed: true,
					},
					"variation": {
						// Property: Variation
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"metric_monitors": {
			// Property: MetricMonitors
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "EntityIdKey": {
			//         "description": "The JSON path to reference the entity id in the event.",
			//         "type": "string"
			//       },
			//       "EventPattern": {
			//         "description": "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
			//         "type": "string"
			//       },
			//       "MetricName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "^[\\S]+$",
			//         "type": "string"
			//       },
			//       "UnitLabel": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": ".*",
			//         "type": "string"
			//       },
			//       "ValueKey": {
			//         "description": "The JSON path to reference the numerical metric value in the event.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MetricName",
			//       "EntityIdKey",
			//       "ValueKey",
			//       "EventPattern"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 3,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"entity_id_key": {
						// Property: EntityIdKey
						Description: "The JSON path to reference the entity id in the event.",
						Type:        types.StringType,
						Computed:    true,
					},
					"event_pattern": {
						// Property: EventPattern
						Description: "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
						Type:        types.StringType,
						Computed:    true,
					},
					"metric_name": {
						// Property: MetricName
						Type:     types.StringType,
						Computed: true,
					},
					"unit_label": {
						// Property: UnitLabel
						Type:     types.StringType,
						Computed: true,
					},
					"value_key": {
						// Property: ValueKey
						Description: "The JSON path to reference the numerical metric value in the event.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "pattern": "[-a-zA-Z0-9._]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"project": {
			// Property: Project
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 0,
			//   "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"randomization_salt": {
			// Property: RandomizationSalt
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 0,
			//   "pattern": ".*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"scheduled_splits_config": {
			// Property: ScheduledSplitsConfig
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "GroupWeights": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "GroupName": {
			//               "maxLength": 127,
			//               "minLength": 1,
			//               "pattern": "[-a-zA-Z0-9._]*",
			//               "type": "string"
			//             },
			//             "SplitWeight": {
			//               "type": "integer"
			//             }
			//           },
			//           "required": [
			//             "GroupName",
			//             "SplitWeight"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "SegmentOverrides": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "EvaluationOrder": {
			//               "type": "integer"
			//             },
			//             "Segment": {
			//               "maxLength": 2048,
			//               "minLength": 1,
			//               "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:segment/[-a-zA-Z0-9._]*)",
			//               "type": "string"
			//             },
			//             "Weights": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "GroupName": {
			//                     "maxLength": 127,
			//                     "minLength": 1,
			//                     "pattern": "[-a-zA-Z0-9._]*",
			//                     "type": "string"
			//                   },
			//                   "SplitWeight": {
			//                     "type": "integer"
			//                   }
			//                 },
			//                 "required": [
			//                   "GroupName",
			//                   "SplitWeight"
			//                 ],
			//                 "type": "object"
			//               },
			//               "type": "array",
			//               "uniqueItems": true
			//             }
			//           },
			//           "required": [
			//             "Segment",
			//             "EvaluationOrder",
			//             "Weights"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "StartTime": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "StartTime",
			//       "GroupWeights"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 6,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"group_weights": {
						// Property: GroupWeights
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"group_name": {
									// Property: GroupName
									Type:     types.StringType,
									Computed: true,
								},
								"split_weight": {
									// Property: SplitWeight
									Type:     types.Int64Type,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"segment_overrides": {
						// Property: SegmentOverrides
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"evaluation_order": {
									// Property: EvaluationOrder
									Type:     types.Int64Type,
									Computed: true,
								},
								"segment": {
									// Property: Segment
									Type:     types.StringType,
									Computed: true,
								},
								"weights": {
									// Property: Weights
									Attributes: tfsdk.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"group_name": {
												// Property: GroupName
												Type:     types.StringType,
												Computed: true,
											},
											"split_weight": {
												// Property: SplitWeight
												Type:     types.Int64Type,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"start_time": {
						// Property: StartTime
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Evidently::Launch",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Evidently::Launch").WithTerraformTypeName("awscc_evidently_launch")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"description":             "Description",
		"desired_state":           "DesiredState",
		"entity_id_key":           "EntityIdKey",
		"evaluation_order":        "EvaluationOrder",
		"event_pattern":           "EventPattern",
		"execution_status":        "ExecutionStatus",
		"feature":                 "Feature",
		"group_name":              "GroupName",
		"group_weights":           "GroupWeights",
		"groups":                  "Groups",
		"key":                     "Key",
		"metric_monitors":         "MetricMonitors",
		"metric_name":             "MetricName",
		"name":                    "Name",
		"project":                 "Project",
		"randomization_salt":      "RandomizationSalt",
		"reason":                  "Reason",
		"scheduled_splits_config": "ScheduledSplitsConfig",
		"segment":                 "Segment",
		"segment_overrides":       "SegmentOverrides",
		"split_weight":            "SplitWeight",
		"start_time":              "StartTime",
		"status":                  "Status",
		"tags":                    "Tags",
		"unit_label":              "UnitLabel",
		"value":                   "Value",
		"value_key":               "ValueKey",
		"variation":               "Variation",
		"weights":                 "Weights",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

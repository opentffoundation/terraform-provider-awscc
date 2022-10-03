// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_ses_configuration_set_event_destination", configurationSetEventDestinationResource)
}

// configurationSetEventDestinationResource returns the Terraform awscc_ses_configuration_set_event_destination resource.
// This Terraform resource corresponds to the CloudFormation AWS::SES::ConfigurationSetEventDestination resource.
func configurationSetEventDestinationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"configuration_set_name": {
			// Property: ConfigurationSetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the configuration set that contains the event destination.",
			//   "type": "string"
			// }
			Description: "The name of the configuration set that contains the event destination.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"event_destination": {
			// Property: EventDestination
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The event destination object.",
			//   "properties": {
			//     "CloudWatchDestination": {
			//       "additionalProperties": false,
			//       "description": "An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.",
			//       "properties": {
			//         "DimensionConfigurations": {
			//           "description": "A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.",
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.",
			//             "properties": {
			//               "DefaultDimensionValue": {
			//                 "description": "The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.",
			//                 "maxLength": 256,
			//                 "minLength": 1,
			//                 "pattern": "^[a-zA-Z0-9_-]{1,256}$",
			//                 "type": "string"
			//               },
			//               "DimensionName": {
			//                 "description": "The name of an Amazon CloudWatch dimension associated with an email sending metric.",
			//                 "maxLength": 256,
			//                 "minLength": 1,
			//                 "pattern": "^[a-zA-Z0-9_:-]{1,256}$",
			//                 "type": "string"
			//               },
			//               "DimensionValueSource": {
			//                 "description": "The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch. To use the message tags that you specify using an X-SES-MESSAGE-TAGS header or a parameter to the SendEmail/SendRawEmail API, specify messageTag. To use your own email headers, specify emailHeader. To put a custom tag on any link included in your email, specify linkTag.",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "DimensionValueSource",
			//               "DefaultDimensionValue",
			//               "DimensionName"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Enabled": {
			//       "description": "Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set. Set to true to enable publishing to this destination; set to false to prevent publishing to this destination. The default value is false.   ",
			//       "type": "boolean"
			//     },
			//     "KinesisFirehoseDestination": {
			//       "additionalProperties": false,
			//       "description": "An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.",
			//       "properties": {
			//         "DeliveryStreamARN": {
			//           "description": "The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.",
			//           "type": "string"
			//         },
			//         "IAMRoleARN": {
			//           "description": "The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "IAMRoleARN",
			//         "DeliveryStreamARN"
			//       ],
			//       "type": "object"
			//     },
			//     "MatchingEventTypes": {
			//       "description": "The type of email sending events, send, reject, bounce, complaint, delivery, open, click, renderingFailure.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "Name": {
			//       "description": "The name of the event destination set.",
			//       "pattern": "^[a-zA-Z0-9_-]{0,64}$",
			//       "type": "string"
			//     },
			//     "SnsDestination": {
			//       "additionalProperties": false,
			//       "description": "An object that contains SNS topic ARN associated event destination.",
			//       "properties": {
			//         "TopicARN": {
			//           "maxLength": 1024,
			//           "minLength": 36,
			//           "pattern": "^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TopicARN"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "MatchingEventTypes"
			//   ],
			//   "type": "object"
			// }
			Description: "The event destination object.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cloudwatch_destination": {
						// Property: CloudWatchDestination
						Description: "An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"dimension_configurations": {
									// Property: DimensionConfigurations
									Description: "A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"default_dimension_value": {
												// Property: DefaultDimensionValue
												Description: "The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 256),
													validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_-]{1,256}$"), ""),
												},
											},
											"dimension_name": {
												// Property: DimensionName
												Description: "The name of an Amazon CloudWatch dimension associated with an email sending metric.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 256),
													validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_:-]{1,256}$"), ""),
												},
											},
											"dimension_value_source": {
												// Property: DimensionValueSource
												Description: "The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch. To use the message tags that you specify using an X-SES-MESSAGE-TAGS header or a parameter to the SendEmail/SendRawEmail API, specify messageTag. To use your own email headers, specify emailHeader. To put a custom tag on any link included in your email, specify linkTag.",
												Type:        types.StringType,
												Required:    true,
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"enabled": {
						// Property: Enabled
						Description: "Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set. Set to true to enable publishing to this destination; set to false to prevent publishing to this destination. The default value is false.   ",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"kinesis_firehose_destination": {
						// Property: KinesisFirehoseDestination
						Description: "An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"delivery_stream_arn": {
									// Property: DeliveryStreamARN
									Description: "The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.",
									Type:        types.StringType,
									Required:    true,
								},
								"iam_role_arn": {
									// Property: IAMRoleARN
									Description: "The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"matching_event_types": {
						// Property: MatchingEventTypes
						Description: "The type of email sending events, send, reject, bounce, complaint, delivery, open, click, renderingFailure.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"name": {
						// Property: Name
						Description: "The name of the event destination set.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_-]{0,64}$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"sns_destination": {
						// Property: SnsDestination
						Description: "An object that contains SNS topic ARN associated event destination.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"topic_arn": {
									// Property: TopicARN
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(36, 1024),
										validate.StringMatch(regexp.MustCompile("^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$"), ""),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::SES::ConfigurationSetEventDestination",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::ConfigurationSetEventDestination").WithTerraformTypeName("awscc_ses_configuration_set_event_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudwatch_destination":       "CloudWatchDestination",
		"configuration_set_name":       "ConfigurationSetName",
		"default_dimension_value":      "DefaultDimensionValue",
		"delivery_stream_arn":          "DeliveryStreamARN",
		"dimension_configurations":     "DimensionConfigurations",
		"dimension_name":               "DimensionName",
		"dimension_value_source":       "DimensionValueSource",
		"enabled":                      "Enabled",
		"event_destination":            "EventDestination",
		"iam_role_arn":                 "IAMRoleARN",
		"id":                           "Id",
		"kinesis_firehose_destination": "KinesisFirehoseDestination",
		"matching_event_types":         "MatchingEventTypes",
		"name":                         "Name",
		"sns_destination":              "SnsDestination",
		"topic_arn":                    "TopicARN",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

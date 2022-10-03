// Code generated by generators/resource/main.go; DO NOT EDIT.

package lookoutmetrics

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
	registry.AddResourceFactory("awscc_lookoutmetrics_alert", alertResource)
}

// alertResource returns the Terraform awscc_lookoutmetrics_alert resource.
// This Terraform resource corresponds to the CloudFormation AWS::LookoutMetrics::Alert resource.
func alertResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"action": {
			// Property: Action
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The action to be taken by the alert when an anomaly is detected.",
			//   "properties": {
			//     "LambdaConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Configuration options for a Lambda alert action.",
			//       "properties": {
			//         "LambdaArn": {
			//           "description": "ARN of a Lambda to send alert notifications to.",
			//           "maxLength": 256,
			//           "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//           "type": "string"
			//         },
			//         "RoleArn": {
			//           "description": "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
			//           "maxLength": 256,
			//           "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "RoleArn",
			//         "LambdaArn"
			//       ],
			//       "type": "object"
			//     },
			//     "SNSConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Configuration options for an SNS alert action.",
			//       "properties": {
			//         "RoleArn": {
			//           "description": "ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.",
			//           "maxLength": 256,
			//           "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//           "type": "string"
			//         },
			//         "SnsTopicArn": {
			//           "description": "ARN of an SNS topic to send alert notifications to.",
			//           "maxLength": 256,
			//           "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "RoleArn",
			//         "SnsTopicArn"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The action to be taken by the alert when an anomaly is detected.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"lambda_configuration": {
						// Property: LambdaConfiguration
						Description: "Configuration options for a Lambda alert action.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"lambda_arn": {
									// Property: LambdaArn
									Description: "ARN of a Lambda to send alert notifications to.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
										validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
									},
								},
								"role_arn": {
									// Property: RoleArn
									Description: "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
										validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
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
					"sns_configuration": {
						// Property: SNSConfiguration
						Description: "Configuration options for an SNS alert action.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"role_arn": {
									// Property: RoleArn
									Description: "ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
										validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
									},
								},
								"sns_topic_arn": {
									// Property: SnsTopicArn
									Description: "ARN of an SNS topic to send alert notifications to.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
										validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"alert_description": {
			// Property: AlertDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the alert.",
			//   "maxLength": 256,
			//   "pattern": ".*\\S.*",
			//   "type": "string"
			// }
			Description: "A description for the alert.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
				validate.StringMatch(regexp.MustCompile(".*\\S.*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"alert_name": {
			// Property: AlertName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the alert. If not provided, a name is generated automatically.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//   "type": "string"
			// }
			Description: "The name of the alert. If not provided, a name is generated automatically.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"alert_sensitivity_threshold": {
			// Property: AlertSensitivityThreshold
			// CloudFormation resource type schema:
			// {
			//   "description": "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
			//   "maximum": 100,
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
			Type:        types.Int64Type,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(0, 100),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"anomaly_detector_arn": {
			// Property: AnomalyDetectorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
				validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN assigned to the alert.",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "ARN assigned to the alert.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::LookoutMetrics::Alert",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutMetrics::Alert").WithTerraformTypeName("awscc_lookoutmetrics_alert")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                      "Action",
		"alert_description":           "AlertDescription",
		"alert_name":                  "AlertName",
		"alert_sensitivity_threshold": "AlertSensitivityThreshold",
		"anomaly_detector_arn":        "AnomalyDetectorArn",
		"arn":                         "Arn",
		"lambda_arn":                  "LambdaArn",
		"lambda_configuration":        "LambdaConfiguration",
		"role_arn":                    "RoleArn",
		"sns_configuration":           "SNSConfiguration",
		"sns_topic_arn":               "SnsTopicArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

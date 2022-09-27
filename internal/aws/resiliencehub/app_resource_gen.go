// Code generated by generators/resource/main.go; DO NOT EDIT.

package resiliencehub

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_resiliencehub_app", appResourceType)
}

// appResourceType returns the Terraform awscc_resiliencehub_app resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ResilienceHub::App resource type.
func appResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_arn": {
			// Property: AppArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the App.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the App.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"app_assessment_schedule": {
			// Property: AppAssessmentSchedule
			// CloudFormation resource type schema:
			// {
			//   "description": "Assessment execution schedule.",
			//   "enum": [
			//     "Disabled",
			//     "Daily"
			//   ],
			//   "type": "string"
			// }
			Description: "Assessment execution schedule.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"Disabled",
					"Daily",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"app_template_body": {
			// Property: AppTemplateBody
			// CloudFormation resource type schema:
			// {
			//   "description": "A string containing full ResilienceHub app template body.",
			//   "maxLength": 5000,
			//   "minLength": 0,
			//   "pattern": "^[\\w\\s:,-\\.'{}\\[\\]:\"]+$",
			//   "type": "string"
			// }
			Description: "A string containing full ResilienceHub app template body.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 5000),
				validate.StringMatch(regexp.MustCompile("^[\\w\\s:,-\\.'{}\\[\\]:\"]+$"), ""),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "App description.",
			//   "maxLength": 500,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "App description.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 500),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the app.",
			//   "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
			//   "type": "string"
			// }
			Description: "Name of the app.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"resiliency_policy_arn": {
			// Property: ResiliencyPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the Resiliency Policy.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"resource_mappings": {
			// Property: ResourceMappings
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of ResourceMapping objects.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Resource mapping is used to map logical resources from template to physical resource",
			//     "properties": {
			//       "LogicalStackName": {
			//         "type": "string"
			//       },
			//       "MappingType": {
			//         "pattern": "CfnStack|Resource|Terraform",
			//         "type": "string"
			//       },
			//       "PhysicalResourceId": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AwsAccountId": {
			//             "pattern": "^[0-9]{12}$",
			//             "type": "string"
			//           },
			//           "AwsRegion": {
			//             "pattern": "^[a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]$",
			//             "type": "string"
			//           },
			//           "Identifier": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Type": {
			//             "pattern": "Arn|Native",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Identifier",
			//           "Type"
			//         ],
			//         "type": "object"
			//       },
			//       "ResourceName": {
			//         "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
			//         "type": "string"
			//       },
			//       "TerraformSourceName": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MappingType",
			//       "PhysicalResourceId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of ResourceMapping objects.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"logical_stack_name": {
						// Property: LogicalStackName
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"mapping_type": {
						// Property: MappingType
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("CfnStack|Resource|Terraform"), ""),
						},
					},
					"physical_resource_id": {
						// Property: PhysicalResourceId
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"aws_account_id": {
									// Property: AwsAccountId
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^[0-9]{12}$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"aws_region": {
									// Property: AwsRegion
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^[a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"identifier": {
									// Property: Identifier
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
								"type": {
									// Property: Type
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("Arn|Native"), ""),
									},
								},
							},
						),
						Required: true,
					},
					"resource_name": {
						// Property: ResourceName
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"terraform_source_name": {
						// Property: TerraformSourceName
						Type:     types.StringType,
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
				Multiset(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
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
		Description: "Resource Type Definition for AWS::ResilienceHub::App.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResilienceHub::App").WithTerraformTypeName("awscc_resiliencehub_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_arn":                 "AppArn",
		"app_assessment_schedule": "AppAssessmentSchedule",
		"app_template_body":       "AppTemplateBody",
		"aws_account_id":          "AwsAccountId",
		"aws_region":              "AwsRegion",
		"description":             "Description",
		"identifier":              "Identifier",
		"logical_stack_name":      "LogicalStackName",
		"mapping_type":            "MappingType",
		"name":                    "Name",
		"physical_resource_id":    "PhysicalResourceId",
		"resiliency_policy_arn":   "ResiliencyPolicyArn",
		"resource_mappings":       "ResourceMappings",
		"resource_name":           "ResourceName",
		"tags":                    "Tags",
		"terraform_source_name":   "TerraformSourceName",
		"type":                    "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

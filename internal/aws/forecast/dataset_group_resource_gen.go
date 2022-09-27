// Code generated by generators/resource/main.go; DO NOT EDIT.

package forecast

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
	registry.AddResourceTypeFactory("awscc_forecast_dataset_group", datasetGroupResourceType)
}

// datasetGroupResourceType returns the Terraform awscc_forecast_dataset_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Forecast::DatasetGroup resource type.
func datasetGroupResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"dataset_arns": {
			// Property: DatasetArns
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.",
			//   "insertionOrder": true,
			//   "items": {
			//     "maxLength": 256,
			//     "pattern": "^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringLenAtMost(256)),
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"dataset_group_arn": {
			// Property: DatasetGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the dataset group to delete.",
			//   "maxLength": 256,
			//   "pattern": "^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the dataset group to delete.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"dataset_group_name": {
			// Property: DatasetGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the dataset group.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z][a-zA-Z0-9_]*",
			//   "type": "string"
			// }
			Description: "A name for the dataset group.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"domain": {
			// Property: Domain
			// CloudFormation resource type schema:
			// {
			//   "description": "The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.",
			//   "enum": [
			//     "RETAIL",
			//     "CUSTOM",
			//     "INVENTORY_PLANNING",
			//     "EC2_CAPACITY",
			//     "WORK_FORCE",
			//     "WEB_TRAFFIC",
			//     "METRICS"
			//   ],
			//   "type": "string"
			// }
			Description: "The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"RETAIL",
					"CUSTOM",
					"INVENTORY_PLANNING",
					"EC2_CAPACITY",
					"WORK_FORCE",
					"WEB_TRAFFIC",
					"METRICS",
				}),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags of Application Insights application.",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
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
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags of Application Insights application.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 200),
			},
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
		Description: "Represents a dataset group that holds a collection of related datasets",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Forecast::DatasetGroup").WithTerraformTypeName("awscc_forecast_dataset_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dataset_arns":       "DatasetArns",
		"dataset_group_arn":  "DatasetGroupArn",
		"dataset_group_name": "DatasetGroupName",
		"domain":             "Domain",
		"key":                "Key",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

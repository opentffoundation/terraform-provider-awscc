// Code generated by generators/resource/main.go; DO NOT EDIT.

package personalize

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
	registry.AddResourceFactory("awscc_personalize_dataset", datasetResource)
}

// datasetResource returns the Terraform awscc_personalize_dataset resource.
// This Terraform resource corresponds to the CloudFormation AWS::Personalize::Dataset resource.
func datasetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"dataset_arn": {
			// Property: DatasetArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the dataset",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The ARN of the dataset",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"dataset_group_arn": {
			// Property: DatasetGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the dataset group to add the dataset to",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the dataset group to add the dataset to",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
				validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"dataset_import_job": {
			// Property: DatasetImportJob
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Initial DatasetImportJob for the created dataset",
			//   "properties": {
			//     "DataSource": {
			//       "additionalProperties": false,
			//       "description": "The Amazon S3 bucket that contains the training data to import.",
			//       "properties": {
			//         "DataLocation": {
			//           "description": "The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.",
			//           "maxLength": 256,
			//           "pattern": "(s3|http|https)://.+",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "DatasetArn": {
			//       "description": "The ARN of the dataset that receives the imported data",
			//       "maxLength": 256,
			//       "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//       "type": "string"
			//     },
			//     "DatasetImportJobArn": {
			//       "description": "The ARN of the dataset import job",
			//       "maxLength": 256,
			//       "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//       "type": "string"
			//     },
			//     "JobName": {
			//       "description": "The name for the dataset import job.",
			//       "maxLength": 63,
			//       "minLength": 1,
			//       "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//       "type": "string"
			//     },
			//     "RoleArn": {
			//       "description": "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.",
			//       "maxLength": 256,
			//       "pattern": "arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Initial DatasetImportJob for the created dataset",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"data_source": {
						// Property: DataSource
						Description: "The Amazon S3 bucket that contains the training data to import.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"data_location": {
									// Property: DataLocation
									Description: "The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
										validate.StringMatch(regexp.MustCompile("(s3|http|https)://.+"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
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
					"dataset_arn": {
						// Property: DatasetArn
						Description: "The ARN of the dataset that receives the imported data",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
							validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"dataset_import_job_arn": {
						// Property: DatasetImportJobArn
						Description: "The ARN of the dataset import job",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
							validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"job_name": {
						// Property: JobName
						Description: "The name for the dataset import job.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 63),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
							validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
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
		"dataset_type": {
			// Property: DatasetType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of dataset",
			//   "enum": [
			//     "Interactions",
			//     "Items",
			//     "Users"
			//   ],
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The type of dataset",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
				validate.StringInSlice([]string{
					"Interactions",
					"Items",
					"Users",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the dataset",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//   "type": "string"
			// }
			Description: "The name for the dataset",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"schema_arn": {
			// Property: SchemaArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the schema to associate with the dataset. The schema defines the dataset fields.",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The ARN of the schema to associate with the dataset. The schema defines the dataset fields.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
				validate.StringMatch(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
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
		Description: "Resource schema for AWS::Personalize::Dataset.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::Dataset").WithTerraformTypeName("awscc_personalize_dataset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_location":          "DataLocation",
		"data_source":            "DataSource",
		"dataset_arn":            "DatasetArn",
		"dataset_group_arn":      "DatasetGroupArn",
		"dataset_import_job":     "DatasetImportJob",
		"dataset_import_job_arn": "DatasetImportJobArn",
		"dataset_type":           "DatasetType",
		"job_name":               "JobName",
		"name":                   "Name",
		"role_arn":               "RoleArn",
		"schema_arn":             "SchemaArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

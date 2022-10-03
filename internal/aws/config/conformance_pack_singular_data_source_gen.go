// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_config_conformance_pack", conformancePackDataSource)
}

// conformancePackDataSource returns the Terraform awscc_config_conformance_pack data source.
// This Terraform data source corresponds to the CloudFormation AWS::Config::ConformancePack resource.
func conformancePackDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"conformance_pack_input_parameters": {
			// Property: ConformancePackInputParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of ConformancePackInputParameter objects.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Input parameters in the form of key-value pairs for the conformance pack.",
			//     "properties": {
			//       "ParameterName": {
			//         "description": "Key part of key-value pair with value being parameter value",
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "ParameterValue": {
			//         "description": "Value part of key-value pair with key being parameter Name",
			//         "maxLength": 4096,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ParameterName",
			//       "ParameterValue"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 60,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of ConformancePackInputParameter objects.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameter_name": {
						// Property: ParameterName
						Description: "Key part of key-value pair with value being parameter value",
						Type:        types.StringType,
						Computed:    true,
					},
					"parameter_value": {
						// Property: ParameterValue
						Description: "Value part of key-value pair with key being parameter Name",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"conformance_pack_name": {
			// Property: ConformancePackName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the conformance pack which will be assigned as the unique identifier.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z][-a-zA-Z0-9]*",
			//   "type": "string"
			// }
			Description: "Name of the conformance pack which will be assigned as the unique identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"delivery_s3_bucket": {
			// Property: DeliveryS3Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "AWS Config stores intermediate files while processing conformance pack template.",
			//   "maxLength": 63,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "AWS Config stores intermediate files while processing conformance pack template.",
			Type:        types.StringType,
			Computed:    true,
		},
		"delivery_s3_key_prefix": {
			// Property: DeliveryS3KeyPrefix
			// CloudFormation resource type schema:
			// {
			//   "description": "The prefix for delivery S3 bucket.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The prefix for delivery S3 bucket.",
			Type:        types.StringType,
			Computed:    true,
		},
		"template_body": {
			// Property: TemplateBody
			// CloudFormation resource type schema:
			// {
			//   "description": "A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.",
			//   "maxLength": 51200,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.",
			Type:        types.StringType,
			Computed:    true,
		},
		"template_s3_uri": {
			// Property: TemplateS3Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "s3://.*",
			//   "type": "string"
			// }
			Description: "Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.",
			Type:        types.StringType,
			Computed:    true,
		},
		"template_ssm_document_details": {
			// Property: TemplateSSMDocumentDetails
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.",
			//   "properties": {
			//     "DocumentName": {
			//       "maxLength": 128,
			//       "minLength": 3,
			//       "type": "string"
			//     },
			//     "DocumentVersion": {
			//       "maxLength": 128,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"document_name": {
						// Property: DocumentName
						Type:     types.StringType,
						Computed: true,
					},
					"document_version": {
						// Property: DocumentVersion
						Type:     types.StringType,
						Computed: true,
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
		Description: "Data Source schema for AWS::Config::ConformancePack",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::ConformancePack").WithTerraformTypeName("awscc_config_conformance_pack")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"conformance_pack_input_parameters": "ConformancePackInputParameters",
		"conformance_pack_name":             "ConformancePackName",
		"delivery_s3_bucket":                "DeliveryS3Bucket",
		"delivery_s3_key_prefix":            "DeliveryS3KeyPrefix",
		"document_name":                     "DocumentName",
		"document_version":                  "DocumentVersion",
		"parameter_name":                    "ParameterName",
		"parameter_value":                   "ParameterValue",
		"template_body":                     "TemplateBody",
		"template_s3_uri":                   "TemplateS3Uri",
		"template_ssm_document_details":     "TemplateSSMDocumentDetails",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

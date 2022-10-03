// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_app_image_config", appImageConfigDataSource)
}

// appImageConfigDataSource returns the Terraform awscc_sagemaker_app_image_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::AppImageConfig resource.
func appImageConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_image_config_arn": {
			// Property: AppImageConfigArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the AppImageConfig.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:app-image-config/.*",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the AppImageConfig.",
			Type:        types.StringType,
			Computed:    true,
		},
		"app_image_config_name": {
			// Property: AppImageConfigName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Name of the AppImageConfig.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}",
			//   "type": "string"
			// }
			Description: "The Name of the AppImageConfig.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kernel_gateway_image_config": {
			// Property: KernelGatewayImageConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The KernelGatewayImageConfig.",
			//   "properties": {
			//     "FileSystemConfig": {
			//       "additionalProperties": false,
			//       "description": "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
			//       "properties": {
			//         "DefaultGid": {
			//           "description": "The default POSIX group ID (GID). If not specified, defaults to 100.",
			//           "maximum": 65535,
			//           "minimum": 0,
			//           "type": "integer"
			//         },
			//         "DefaultUid": {
			//           "description": "The default POSIX user ID (UID). If not specified, defaults to 1000.",
			//           "maximum": 65535,
			//           "minimum": 0,
			//           "type": "integer"
			//         },
			//         "MountPath": {
			//           "description": "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
			//           "maxLength": 1024,
			//           "minLength": 1,
			//           "pattern": "^/.*",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "KernelSpecs": {
			//       "description": "The specification of the Jupyter kernels in the image.",
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "DisplayName": {
			//             "description": "The display name of the kernel.",
			//             "maxLength": 1024,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Name": {
			//             "description": "The name of the kernel.",
			//             "maxLength": 1024,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Name"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 1,
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "KernelSpecs"
			//   ],
			//   "type": "object"
			// }
			Description: "The KernelGatewayImageConfig.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"file_system_config": {
						// Property: FileSystemConfig
						Description: "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"default_gid": {
									// Property: DefaultGid
									Description: "The default POSIX group ID (GID). If not specified, defaults to 100.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"default_uid": {
									// Property: DefaultUid
									Description: "The default POSIX user ID (UID). If not specified, defaults to 1000.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"mount_path": {
									// Property: MountPath
									Description: "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"kernel_specs": {
						// Property: KernelSpecs
						Description: "The specification of the Jupyter kernels in the image.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"display_name": {
									// Property: DisplayName
									Description: "The display name of the kernel.",
									Type:        types.StringType,
									Computed:    true,
								},
								"name": {
									// Property: Name
									Description: "The name of the kernel.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
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
			//   "description": "A list of tags to apply to the AppImageConfig.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of tags to apply to the AppImageConfig.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
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
		Description: "Data Source schema for AWS::SageMaker::AppImageConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::AppImageConfig").WithTerraformTypeName("awscc_sagemaker_app_image_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_image_config_arn":        "AppImageConfigArn",
		"app_image_config_name":       "AppImageConfigName",
		"default_gid":                 "DefaultGid",
		"default_uid":                 "DefaultUid",
		"display_name":                "DisplayName",
		"file_system_config":          "FileSystemConfig",
		"kernel_gateway_image_config": "KernelGatewayImageConfig",
		"kernel_specs":                "KernelSpecs",
		"key":                         "Key",
		"mount_path":                  "MountPath",
		"name":                        "Name",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

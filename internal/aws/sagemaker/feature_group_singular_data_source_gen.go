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
	registry.AddDataSourceFactory("awscc_sagemaker_feature_group", featureGroupDataSource)
}

// featureGroupDataSource returns the Terraform awscc_sagemaker_feature_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::FeatureGroup resource.
func featureGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description about the FeatureGroup.",
			//   "maxLength": 128,
			//   "type": "string"
			// }
			Description: "Description about the FeatureGroup.",
			Type:        types.StringType,
			Computed:    true,
		},
		"event_time_feature_name": {
			// Property: EventTimeFeatureName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Event Time Feature Name.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
			//   "type": "string"
			// }
			Description: "The Event Time Feature Name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"feature_definitions": {
			// Property: FeatureDefinitions
			// CloudFormation resource type schema:
			// {
			//   "description": "An Array of Feature Definition",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "FeatureName": {
			//         "maxLength": 64,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
			//         "type": "string"
			//       },
			//       "FeatureType": {
			//         "enum": [
			//           "Integral",
			//           "Fractional",
			//           "String"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "FeatureName",
			//       "FeatureType"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 2500,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An Array of Feature Definition",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"feature_name": {
						// Property: FeatureName
						Type:     types.StringType,
						Computed: true,
					},
					"feature_type": {
						// Property: FeatureType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"feature_group_name": {
			// Property: FeatureGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Name of the FeatureGroup.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
			//   "type": "string"
			// }
			Description: "The Name of the FeatureGroup.",
			Type:        types.StringType,
			Computed:    true,
		},
		"offline_store_config": {
			// Property: OfflineStoreConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DataCatalogConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Catalog": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "Database": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "TableName": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TableName",
			//         "Catalog",
			//         "Database"
			//       ],
			//       "type": "object"
			//     },
			//     "DisableGlueTableCreation": {
			//       "type": "boolean"
			//     },
			//     "S3StorageConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "KmsKeyId": {
			//           "maxLength": 2048,
			//           "type": "string"
			//         },
			//         "S3Uri": {
			//           "maxLength": 1024,
			//           "pattern": "^(https|s3)://([^/]+)/?(.*)$",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "S3Uri"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "S3StorageConfig"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"data_catalog_config": {
						// Property: DataCatalogConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog": {
									// Property: Catalog
									Type:     types.StringType,
									Computed: true,
								},
								"database": {
									// Property: Database
									Type:     types.StringType,
									Computed: true,
								},
								"table_name": {
									// Property: TableName
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"disable_glue_table_creation": {
						// Property: DisableGlueTableCreation
						Type:     types.BoolType,
						Computed: true,
					},
					"s3_storage_config": {
						// Property: S3StorageConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"kms_key_id": {
									// Property: KmsKeyId
									Type:     types.StringType,
									Computed: true,
								},
								"s3_uri": {
									// Property: S3Uri
									Type:     types.StringType,
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
		"online_store_config": {
			// Property: OnlineStoreConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "EnableOnlineStore": {
			//       "type": "boolean"
			//     },
			//     "SecurityConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "KmsKeyId": {
			//           "maxLength": 2048,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enable_online_store": {
						// Property: EnableOnlineStore
						Type:     types.BoolType,
						Computed: true,
					},
					"security_config": {
						// Property: SecurityConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"kms_key_id": {
									// Property: KmsKeyId
									Type:     types.StringType,
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
		"record_identifier_feature_name": {
			// Property: RecordIdentifierFeatureName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Record Identifier Feature Name.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
			//   "type": "string"
			// }
			Description: "The Record Identifier Feature Name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
			//   "type": "string"
			// }
			Description: "Role Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pair to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of key-value pair to apply to this resource.",
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
		Description: "Data Source schema for AWS::SageMaker::FeatureGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::FeatureGroup").WithTerraformTypeName("awscc_sagemaker_feature_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":                        "Catalog",
		"data_catalog_config":            "DataCatalogConfig",
		"database":                       "Database",
		"description":                    "Description",
		"disable_glue_table_creation":    "DisableGlueTableCreation",
		"enable_online_store":            "EnableOnlineStore",
		"event_time_feature_name":        "EventTimeFeatureName",
		"feature_definitions":            "FeatureDefinitions",
		"feature_group_name":             "FeatureGroupName",
		"feature_name":                   "FeatureName",
		"feature_type":                   "FeatureType",
		"key":                            "Key",
		"kms_key_id":                     "KmsKeyId",
		"offline_store_config":           "OfflineStoreConfig",
		"online_store_config":            "OnlineStoreConfig",
		"record_identifier_feature_name": "RecordIdentifierFeatureName",
		"role_arn":                       "RoleArn",
		"s3_storage_config":              "S3StorageConfig",
		"s3_uri":                         "S3Uri",
		"security_config":                "SecurityConfig",
		"table_name":                     "TableName",
		"tags":                           "Tags",
		"value":                          "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

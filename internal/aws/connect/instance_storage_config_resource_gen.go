// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

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
	registry.AddResourceTypeFactory("awscc_connect_instance_storage_config", instanceStorageConfigResourceType)
}

// instanceStorageConfigResourceType returns the Terraform awscc_connect_instance_storage_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Connect::InstanceStorageConfig resource type.
func instanceStorageConfigResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"association_id": {
			// Property: AssociationId
			// CloudFormation resource type schema:
			// {
			//   "description": "An associationID is automatically generated when a storage config is associated with an instance",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "^[-a-z0-9]*$",
			//   "type": "string"
			// }
			Description: "An associationID is automatically generated when a storage config is associated with an instance",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Connect Instance ID with which the storage config will be associated",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "Connect Instance ID with which the storage config will be associated",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"kinesis_firehose_config": {
			// Property: KinesisFirehoseConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "FirehoseArn": {
			//       "description": "An ARN is a unique AWS resource identifier.",
			//       "pattern": "^arn:aws[-a-z0-9]*:firehose:[-a-z0-9]*:[0-9]{12}:deliverystream/[-a-zA-Z0-9_.]*$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "FirehoseArn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"firehose_arn": {
						// Property: FirehoseArn
						Description: "An ARN is a unique AWS resource identifier.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:firehose:[-a-z0-9]*:[0-9]{12}:deliverystream/[-a-zA-Z0-9_.]*$"), ""),
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
		"kinesis_stream_config": {
			// Property: KinesisStreamConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "StreamArn": {
			//       "description": "An ARN is a unique AWS resource identifier.",
			//       "pattern": "^arn:aws[-a-z0-9]*:kinesis:[-a-z0-9]*:[0-9]{12}:stream/[-a-zA-Z0-9_.]*$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "StreamArn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"stream_arn": {
						// Property: StreamArn
						Description: "An ARN is a unique AWS resource identifier.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:kinesis:[-a-z0-9]*:[0-9]{12}:stream/[-a-zA-Z0-9_.]*$"), ""),
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
		"kinesis_video_stream_config": {
			// Property: KinesisVideoStreamConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "EncryptionConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "EncryptionType": {
			//           "description": "Specifies default encryption using AWS KMS-Managed Keys",
			//           "enum": [
			//             "KMS"
			//           ],
			//           "type": "string"
			//         },
			//         "KeyId": {
			//           "description": "Specifies the encryption key id",
			//           "maxLength": 128,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "EncryptionType",
			//         "KeyId"
			//       ],
			//       "type": "object"
			//     },
			//     "Prefix": {
			//       "description": "Prefixes are used to infer logical hierarchy",
			//       "maxLength": 128,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "RetentionPeriodHours": {
			//       "description": "Number of hours",
			//       "type": "number"
			//     }
			//   },
			//   "required": [
			//     "Prefix",
			//     "RetentionPeriodHours"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_config": {
						// Property: EncryptionConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"encryption_type": {
									// Property: EncryptionType
									Description: "Specifies default encryption using AWS KMS-Managed Keys",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"KMS",
										}),
									},
								},
								"key_id": {
									// Property: KeyId
									Description: "Specifies the encryption key id",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 128),
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
					"prefix": {
						// Property: Prefix
						Description: "Prefixes are used to infer logical hierarchy",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"retention_period_hours": {
						// Property: RetentionPeriodHours
						Description: "Number of hours",
						Type:        types.Float64Type,
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
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the type of storage resource available for the instance",
			//   "enum": [
			//     "CHAT_TRANSCRIPTS",
			//     "CALL_RECORDINGS",
			//     "SCHEDULED_REPORTS",
			//     "MEDIA_STREAMS",
			//     "CONTACT_TRACE_RECORDS",
			//     "AGENT_EVENTS"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the type of storage resource available for the instance",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"CHAT_TRANSCRIPTS",
					"CALL_RECORDINGS",
					"SCHEDULED_REPORTS",
					"MEDIA_STREAMS",
					"CONTACT_TRACE_RECORDS",
					"AGENT_EVENTS",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"s3_config": {
			// Property: S3Config
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BucketName": {
			//       "description": "A name for the S3 Bucket",
			//       "maxLength": 128,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "BucketPrefix": {
			//       "description": "Prefixes are used to infer logical hierarchy",
			//       "maxLength": 128,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "EncryptionConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "EncryptionType": {
			//           "description": "Specifies default encryption using AWS KMS-Managed Keys",
			//           "enum": [
			//             "KMS"
			//           ],
			//           "type": "string"
			//         },
			//         "KeyId": {
			//           "description": "Specifies the encryption key id",
			//           "maxLength": 128,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "EncryptionType",
			//         "KeyId"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "BucketName",
			//     "BucketPrefix"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bucket_name": {
						// Property: BucketName
						Description: "A name for the S3 Bucket",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"bucket_prefix": {
						// Property: BucketPrefix
						Description: "Prefixes are used to infer logical hierarchy",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"encryption_config": {
						// Property: EncryptionConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"encryption_type": {
									// Property: EncryptionType
									Description: "Specifies default encryption using AWS KMS-Managed Keys",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"KMS",
										}),
									},
								},
								"key_id": {
									// Property: KeyId
									Description: "Specifies the encryption key id",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 128),
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
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"storage_type": {
			// Property: StorageType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the storage type to be associated with the instance",
			//   "enum": [
			//     "S3",
			//     "KINESIS_VIDEO_STREAM",
			//     "KINESIS_STREAM",
			//     "KINESIS_FIREHOSE"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the storage type to be associated with the instance",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"S3",
					"KINESIS_VIDEO_STREAM",
					"KINESIS_STREAM",
					"KINESIS_FIREHOSE",
				}),
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
		Description: "Resource Type definition for AWS::Connect::InstanceStorageConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::InstanceStorageConfig").WithTerraformTypeName("awscc_connect_instance_storage_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id":              "AssociationId",
		"bucket_name":                 "BucketName",
		"bucket_prefix":               "BucketPrefix",
		"encryption_config":           "EncryptionConfig",
		"encryption_type":             "EncryptionType",
		"firehose_arn":                "FirehoseArn",
		"instance_arn":                "InstanceArn",
		"key_id":                      "KeyId",
		"kinesis_firehose_config":     "KinesisFirehoseConfig",
		"kinesis_stream_config":       "KinesisStreamConfig",
		"kinesis_video_stream_config": "KinesisVideoStreamConfig",
		"prefix":                      "Prefix",
		"resource_type":               "ResourceType",
		"retention_period_hours":      "RetentionPeriodHours",
		"s3_config":                   "S3Config",
		"storage_type":                "StorageType",
		"stream_arn":                  "StreamArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

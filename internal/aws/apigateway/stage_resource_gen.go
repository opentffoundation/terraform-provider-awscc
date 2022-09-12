// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_apigateway_stage", stageResourceType)
}

// stageResourceType returns the Terraform awscc_apigateway_stage resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ApiGateway::Stage resource type.
func stageResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_log_setting": {
			// Property: AccessLogSetting
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies settings for logging access in this stage.",
			//   "properties": {
			//     "DestinationArn": {
			//       "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. This parameter is required to enable access logging.",
			//       "type": "string"
			//     },
			//     "Format": {
			//       "description": "A single line format of the access logs of data, as specified by selected $context variables (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least $context.requestId. This parameter is required to enable access logging.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specifies settings for logging access in this stage.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination_arn": {
						// Property: DestinationArn
						Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. This parameter is required to enable access logging.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"format": {
						// Property: Format
						Description: "A single line format of the access logs of data, as specified by selected $context variables (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least $context.requestId. This parameter is required to enable access logging.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
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
		"cache_cluster_enabled": {
			// Property: CacheClusterEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether cache clustering is enabled for the stage.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether cache clustering is enabled for the stage.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"cache_cluster_size": {
			// Property: CacheClusterSize
			// CloudFormation resource type schema:
			// {
			//   "description": "The stage's cache cluster size.",
			//   "type": "string"
			// }
			Description: "The stage's cache cluster size.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"canary_setting": {
			// Property: CanarySetting
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies settings for the canary deployment in this stage.",
			//   "properties": {
			//     "DeploymentId": {
			//       "description": "The identifier of the deployment that the stage points to.",
			//       "type": "string"
			//     },
			//     "PercentTraffic": {
			//       "description": "The percentage (0-100) of traffic diverted to a canary deployment.",
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "StageVariableOverrides": {
			//       "additionalProperties": false,
			//       "description": "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UseStageCache": {
			//       "description": "Whether the canary deployment uses the stage cache or not.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specifies settings for the canary deployment in this stage.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"deployment_id": {
						// Property: DeploymentId
						Description: "The identifier of the deployment that the stage points to.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"percent_traffic": {
						// Property: PercentTraffic
						Description: "The percentage (0-100) of traffic diverted to a canary deployment.",
						Type:        types.Float64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatBetween(0.000000, 100.000000),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"stage_variable_overrides": {
						// Property: StageVariableOverrides
						Description: "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"use_stage_cache": {
						// Property: UseStageCache
						Description: "Whether the canary deployment uses the stage cache or not.",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
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
		"client_certificate_id": {
			// Property: ClientCertificateId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the client certificate that API Gateway uses to call your integration endpoints in the stage. ",
			//   "type": "string"
			// }
			Description: "The ID of the client certificate that API Gateway uses to call your integration endpoints in the stage. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"deployment_id": {
			// Property: DeploymentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the deployment that the stage is associated with. This parameter is required to create a stage. ",
			//   "type": "string"
			// }
			Description: "The ID of the deployment that the stage is associated with. This parameter is required to create a stage. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the stage.",
			//   "type": "string"
			// }
			Description: "A description of the stage.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"documentation_version": {
			// Property: DocumentationVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version ID of the API documentation snapshot.",
			//   "type": "string"
			// }
			Description: "The version ID of the API documentation snapshot.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"method_settings": {
			// Property: MethodSettings
			// CloudFormation resource type schema:
			// {
			//   "description": "Settings for all methods in the stage.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Configures settings for all methods in a stage.",
			//     "properties": {
			//       "CacheDataEncrypted": {
			//         "description": "Indicates whether the cached responses are encrypted.",
			//         "type": "boolean"
			//       },
			//       "CacheTtlInSeconds": {
			//         "description": "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.",
			//         "type": "integer"
			//       },
			//       "CachingEnabled": {
			//         "description": "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses.",
			//         "type": "boolean"
			//       },
			//       "DataTraceEnabled": {
			//         "description": "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.",
			//         "type": "boolean"
			//       },
			//       "HttpMethod": {
			//         "description": "The HTTP method. You can use an asterisk (*) as a wildcard to apply method settings to multiple methods.",
			//         "type": "string"
			//       },
			//       "LoggingLevel": {
			//         "description": "The logging level for this method. For valid values, see the loggingLevel property of the Stage (https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/#loggingLevel) resource in the Amazon API Gateway API Reference.",
			//         "type": "string"
			//       },
			//       "MetricsEnabled": {
			//         "description": "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
			//         "type": "boolean"
			//       },
			//       "ResourcePath": {
			//         "description": "The resource path for this method. Forward slashes (/) are encoded as ~1 and the initial slash must include a forward slash. For example, the path value /resource/subresource must be encoded as /~1resource~1subresource. To specify the root path, use only a slash (/). You can use an asterisk (*) as a wildcard to apply method settings to multiple methods.",
			//         "type": "string"
			//       },
			//       "ThrottlingBurstLimit": {
			//         "description": "The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "ThrottlingRateLimit": {
			//         "description": "The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
			//         "minimum": 0,
			//         "type": "number"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Settings for all methods in the stage.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"cache_data_encrypted": {
						// Property: CacheDataEncrypted
						Description: "Indicates whether the cached responses are encrypted.",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"cache_ttl_in_seconds": {
						// Property: CacheTtlInSeconds
						Description: "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"caching_enabled": {
						// Property: CachingEnabled
						Description: "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses.",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"data_trace_enabled": {
						// Property: DataTraceEnabled
						Description: "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"http_method": {
						// Property: HttpMethod
						Description: "The HTTP method. You can use an asterisk (*) as a wildcard to apply method settings to multiple methods.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"logging_level": {
						// Property: LoggingLevel
						Description: "The logging level for this method. For valid values, see the loggingLevel property of the Stage (https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/#loggingLevel) resource in the Amazon API Gateway API Reference.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"metrics_enabled": {
						// Property: MetricsEnabled
						Description: "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"resource_path": {
						// Property: ResourcePath
						Description: "The resource path for this method. Forward slashes (/) are encoded as ~1 and the initial slash must include a forward slash. For example, the path value /resource/subresource must be encoded as /~1resource~1subresource. To specify the root path, use only a slash (/). You can use an asterisk (*) as a wildcard to apply method settings to multiple methods.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"throttling_burst_limit": {
						// Property: ThrottlingBurstLimit
						Description: "The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"throttling_rate_limit": {
						// Property: ThrottlingRateLimit
						Description: "The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
						Type:        types.Float64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatAtLeast(0.000000),
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
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the RestApi resource that you're deploying with this stage.",
			//   "type": "string"
			// }
			Description: "The ID of the RestApi resource that you're deploying with this stage.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"stage_name": {
			// Property: StageName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).",
			//   "type": "string"
			// }
			Description: "The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of arbitrary tags (key-value pairs) to associate with the stage.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Identify and categorize resources.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.",
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
			//   "uniqueItems": false
			// }
			Description: "An array of arbitrary tags (key-value pairs) to associate with the stage.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.",
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"tracing_enabled": {
			// Property: TracingEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether active X-Ray tracing is enabled for this stage.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether active X-Ray tracing is enabled for this stage.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"variables": {
			// Property: Variables
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.",
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
		Description: "Resource Type definition for AWS::ApiGateway::Stage",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Stage").WithTerraformTypeName("awscc_apigateway_stage")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_log_setting":       "AccessLogSetting",
		"cache_cluster_enabled":    "CacheClusterEnabled",
		"cache_cluster_size":       "CacheClusterSize",
		"cache_data_encrypted":     "CacheDataEncrypted",
		"cache_ttl_in_seconds":     "CacheTtlInSeconds",
		"caching_enabled":          "CachingEnabled",
		"canary_setting":           "CanarySetting",
		"client_certificate_id":    "ClientCertificateId",
		"data_trace_enabled":       "DataTraceEnabled",
		"deployment_id":            "DeploymentId",
		"description":              "Description",
		"destination_arn":          "DestinationArn",
		"documentation_version":    "DocumentationVersion",
		"format":                   "Format",
		"http_method":              "HttpMethod",
		"key":                      "Key",
		"logging_level":            "LoggingLevel",
		"method_settings":          "MethodSettings",
		"metrics_enabled":          "MetricsEnabled",
		"percent_traffic":          "PercentTraffic",
		"resource_path":            "ResourcePath",
		"rest_api_id":              "RestApiId",
		"stage_name":               "StageName",
		"stage_variable_overrides": "StageVariableOverrides",
		"tags":                     "Tags",
		"throttling_burst_limit":   "ThrottlingBurstLimit",
		"throttling_rate_limit":    "ThrottlingRateLimit",
		"tracing_enabled":          "TracingEnabled",
		"use_stage_cache":          "UseStageCache",
		"value":                    "Value",
		"variables":                "Variables",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

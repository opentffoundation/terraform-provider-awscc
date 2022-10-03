// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_resolver_rule", resolverRuleDataSource)
}

// resolverRuleDataSource returns the Terraform awscc_route53resolver_resolver_rule data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::ResolverRule resource.
func resolverRuleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the resolver rule.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the resolver rule.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the Resolver rule",
			//   "maxLength": 64,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The name for the Resolver rule",
			Type:        types.StringType,
			Computed:    true,
		},
		"resolver_endpoint_id": {
			// Property: ResolverEndpointId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the endpoint that the rule is associated with.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The ID of the endpoint that the rule is associated with.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resolver_rule_id": {
			// Property: ResolverRuleId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the endpoint that the rule is associated with.",
			//   "type": "string"
			// }
			Description: "The ID of the endpoint that the rule is associated with.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rule_type": {
			// Property: RuleType
			// CloudFormation resource type schema:
			// {
			//   "description": "When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.",
			//   "enum": [
			//     "FORWARD",
			//     "SYSTEM",
			//     "RECURSIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
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
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"target_ips": {
			// Property: TargetIps
			// CloudFormation resource type schema:
			// {
			//   "description": "An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Ip": {
			//         "description": "One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses. ",
			//         "type": "string"
			//       },
			//       "Port": {
			//         "description": "The port at Ip that you want to forward DNS queries to. ",
			//         "maxLength": 65535,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Ip"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"ip": {
						// Property: Ip
						Description: "One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"port": {
						// Property: Port
						Description: "The port at Ip that you want to forward DNS queries to. ",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::Route53Resolver::ResolverRule",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverRule").WithTerraformTypeName("awscc_route53resolver_resolver_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                  "Arn",
		"domain_name":          "DomainName",
		"ip":                   "Ip",
		"key":                  "Key",
		"name":                 "Name",
		"port":                 "Port",
		"resolver_endpoint_id": "ResolverEndpointId",
		"resolver_rule_id":     "ResolverRuleId",
		"rule_type":            "RuleType",
		"tags":                 "Tags",
		"target_ips":           "TargetIps",
		"value":                "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

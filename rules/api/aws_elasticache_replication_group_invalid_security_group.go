// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsElastiCacheReplicationGroupInvalidSecurityGroupRule checks whether attribute value actually exists
type AwsElastiCacheReplicationGroupInvalidSecurityGroupRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsElastiCacheReplicationGroupInvalidSecurityGroupRule returns new rule with default attributes
func NewAwsElastiCacheReplicationGroupInvalidSecurityGroupRule() *AwsElastiCacheReplicationGroupInvalidSecurityGroupRule {
	return &AwsElastiCacheReplicationGroupInvalidSecurityGroupRule{
		resourceType:  "aws_elasticache_replication_group",
		attributeName: "security_group_ids",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsElastiCacheReplicationGroupInvalidSecurityGroupRule) Name() string {
	return "aws_elasticache_replication_group_invalid_security_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheReplicationGroupInvalidSecurityGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheReplicationGroupInvalidSecurityGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheReplicationGroupInvalidSecurityGroupRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeSecurityGroups
func (r *AwsElastiCacheReplicationGroupInvalidSecurityGroupRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeSecurityGroups")
			var err error
			r.data, err = runner.AwsClient.DescribeSecurityGroups()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeSecurityGroups",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		return runner.EachStringSliceExprs(attribute.Expr, func(val string, expr hcl.Expression) {
			if !r.data[val] {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is invalid security group.`, val),
					expr,
				)
			}
		})
	})
}

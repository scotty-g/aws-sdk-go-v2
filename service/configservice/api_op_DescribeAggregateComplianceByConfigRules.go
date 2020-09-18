// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of compliant and noncompliant rules with the number of resources
// for compliant and noncompliant rules. The results can return an empty result
// page, but if you have a nextToken, the results are displayed on the next page.
func (c *Client) DescribeAggregateComplianceByConfigRules(ctx context.Context, params *DescribeAggregateComplianceByConfigRulesInput, optFns ...func(*Options)) (*DescribeAggregateComplianceByConfigRulesOutput, error) {
	stack := middleware.NewStack("DescribeAggregateComplianceByConfigRules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAggregateComplianceByConfigRulesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeAggregateComplianceByConfigRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAggregateComplianceByConfigRules(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeAggregateComplianceByConfigRules",
			Err:           err,
		}
	}
	out := result.(*DescribeAggregateComplianceByConfigRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAggregateComplianceByConfigRulesInput struct {
	// Filters the results by ConfigRuleComplianceFilters object.
	Filters *types.ConfigRuleComplianceFilters
	// The name of the configuration aggregator.
	ConfigurationAggregatorName *string
	// The maximum number of evaluation results returned on each page. The default is
	// maximum. If you specify 0, AWS Config uses the default.
	Limit *int32
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type DescribeAggregateComplianceByConfigRulesOutput struct {
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// Returns a list of AggregateComplianceByConfigRule object.
	AggregateComplianceByConfigRules []*types.AggregateComplianceByConfigRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAggregateComplianceByConfigRulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAggregateComplianceByConfigRules{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAggregateComplianceByConfigRules{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAggregateComplianceByConfigRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeAggregateComplianceByConfigRules",
	}
}
// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the subscribers that are associated with a notification.
func (c *Client) DescribeSubscribersForNotification(ctx context.Context, params *DescribeSubscribersForNotificationInput, optFns ...func(*Options)) (*DescribeSubscribersForNotificationOutput, error) {
	if params == nil {
		params = &DescribeSubscribersForNotificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSubscribersForNotification", params, optFns, addOperationDescribeSubscribersForNotificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSubscribersForNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DescribeSubscribersForNotification
type DescribeSubscribersForNotificationInput struct {

	// The accountId that is associated with the budget whose subscribers you want
	// descriptions of.
	//
	// This member is required.
	AccountId *string

	// The name of the budget whose subscribers you want descriptions of.
	//
	// This member is required.
	BudgetName *string

	// The notification whose subscribers you want to list.
	//
	// This member is required.
	Notification *types.Notification

	// An optional integer that represents how many entries a paginated response
	// contains. The maximum is 100.
	MaxResults *int32

	// The pagination token that you include in your request to indicate the next set
	// of results that you want to retrieve.
	NextToken *string
}

// Response of DescribeSubscribersForNotification
type DescribeSubscribersForNotificationOutput struct {

	// The pagination token in the service response that indicates the next set of
	// results that you can retrieve.
	NextToken *string

	// A list of subscribers that are associated with a notification.
	Subscribers []*types.Subscriber

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSubscribersForNotificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSubscribersForNotification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSubscribersForNotification{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeSubscribersForNotificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSubscribersForNotification(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeSubscribersForNotification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeSubscribersForNotification",
	}
}

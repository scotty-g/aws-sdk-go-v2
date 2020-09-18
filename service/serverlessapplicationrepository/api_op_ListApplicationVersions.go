// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists versions for the specified application.
func (c *Client) ListApplicationVersions(ctx context.Context, params *ListApplicationVersionsInput, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) {
	stack := middleware.NewStack("ListApplicationVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListApplicationVersionsMiddlewares(stack)
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
	addOpListApplicationVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationVersions(options.Region), middleware.Before)

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
			OperationName: "ListApplicationVersions",
			Err:           err,
		}
	}
	out := result.(*ListApplicationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationVersionsInput struct {
	// The total number of items to return.
	MaxItems *int32
	// The Amazon Resource Name (ARN) of the application.
	ApplicationId *string
	// A token to specify where to start paginating.
	NextToken *string
}

type ListApplicationVersionsOutput struct {
	// The token to request the next page of results.
	NextToken *string
	// An array of version summaries for the application.
	Versions []*types.VersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListApplicationVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListApplicationVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListApplicationVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListApplicationVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "ListApplicationVersions",
	}
}
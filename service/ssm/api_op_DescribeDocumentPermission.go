// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the permissions for a Systems Manager document. If you created the
// document, you are the owner. If a document is shared, it can either be shared
// privately (by specifying a user's AWS account ID) or publicly (All).
func (c *Client) DescribeDocumentPermission(ctx context.Context, params *DescribeDocumentPermissionInput, optFns ...func(*Options)) (*DescribeDocumentPermissionOutput, error) {
	stack := middleware.NewStack("DescribeDocumentPermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDocumentPermissionMiddlewares(stack)
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
	addOpDescribeDocumentPermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDocumentPermission(options.Region), middleware.Before)

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
			OperationName: "DescribeDocumentPermission",
			Err:           err,
		}
	}
	out := result.(*DescribeDocumentPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDocumentPermissionInput struct {
	// The name of the document for which you are the owner.
	Name *string
	// The permission type for the document. The permission type can be Share.
	PermissionType types.DocumentPermissionType
}

type DescribeDocumentPermissionOutput struct {
	// The account IDs that have permission to use this document. The ID can be either
	// an AWS account or All.
	AccountIds []*string
	// A list of AWS accounts where the current document is shared and the version
	// shared with each account.
	AccountSharingInfoList []*types.AccountSharingInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDocumentPermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDocumentPermission{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDocumentPermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDocumentPermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeDocumentPermission",
	}
}
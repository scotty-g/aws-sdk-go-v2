// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates phone number product types or calling names. You can update one
// attribute at a time for each UpdatePhoneNumberRequestItem. For example, you can
// update either the product type or the calling name. For product types, choose
// from Amazon Chime Business Calling and Amazon Chime Voice Connector. For
// toll-free numbers, you must use the Amazon Chime Voice Connector product type.
// Updates to outbound calling names can take up to 72 hours to complete. Pending
// updates to outbound calling names must be complete before you can request
// another update.
func (c *Client) BatchUpdatePhoneNumber(ctx context.Context, params *BatchUpdatePhoneNumberInput, optFns ...func(*Options)) (*BatchUpdatePhoneNumberOutput, error) {
	stack := middleware.NewStack("BatchUpdatePhoneNumber", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchUpdatePhoneNumberMiddlewares(stack)
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
	addOpBatchUpdatePhoneNumberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(options.Region), middleware.Before)

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
			OperationName: "BatchUpdatePhoneNumber",
			Err:           err,
		}
	}
	out := result.(*BatchUpdatePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdatePhoneNumberInput struct {
	// The request containing the phone number IDs and product types or calling names
	// to update.
	UpdatePhoneNumberRequestItems []*types.UpdatePhoneNumberRequestItem
}

type BatchUpdatePhoneNumberOutput struct {
	// If the action fails for one or more of the phone numbers in the request, a list
	// of the phone numbers is returned, along with error codes and error messages.
	PhoneNumberErrors []*types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchUpdatePhoneNumberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdatePhoneNumber{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdatePhoneNumber{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUpdatePhoneNumber",
	}
}
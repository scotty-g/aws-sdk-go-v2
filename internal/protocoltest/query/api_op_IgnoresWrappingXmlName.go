// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The xmlName trait on the output structure is ignored in AWS Query. The wrapping
// element is always operation name + "Response", and inside of that wrapper is
// another wrapper named operation name + "Result".
func (c *Client) IgnoresWrappingXmlName(ctx context.Context, params *IgnoresWrappingXmlNameInput, optFns ...func(*Options)) (*IgnoresWrappingXmlNameOutput, error) {
	stack := middleware.NewStack("IgnoresWrappingXmlName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opIgnoresWrappingXmlName(options.Region), middleware.Before)
	addawsAwsquery_serdeOpIgnoresWrappingXmlNameMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "IgnoresWrappingXmlName",
			Err:           err,
		}
	}
	out := result.(*IgnoresWrappingXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IgnoresWrappingXmlNameInput struct {
}

type IgnoresWrappingXmlNameOutput struct {
	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpIgnoresWrappingXmlNameMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpIgnoresWrappingXmlName{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpIgnoresWrappingXmlName{}, middleware.After)
}

func newServiceMetadataMiddleware_opIgnoresWrappingXmlName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "IgnoresWrappingXmlName",
	}
}

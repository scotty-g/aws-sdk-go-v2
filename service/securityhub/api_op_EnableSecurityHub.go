// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables Security Hub for your account in the current Region or the Region you
// specify in the request. When you enable Security Hub, you grant to Security Hub
// the permissions necessary to gather findings from other services that are
// integrated with Security Hub. When you use the EnableSecurityHub operation to
// enable Security Hub, you also automatically enable the following standards.
//
//
// * CIS AWS Foundations
//
//     * AWS Foundational Security Best Practices
//
// You do
// not enable the Payment Card Industry Data Security Standard (PCI DSS) standard.
// To not enable the automatically enabled standards, set EnableDefaultStandards to
// false. After you enable Security Hub, to enable a standard, use the
// BatchEnableStandards () operation. To disable a standard, use the
// BatchDisableStandards () operation. To learn more, see Setting Up AWS Security
// Hub
// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-settingup.html)
// in the AWS Security Hub User Guide.
func (c *Client) EnableSecurityHub(ctx context.Context, params *EnableSecurityHubInput, optFns ...func(*Options)) (*EnableSecurityHubOutput, error) {
	stack := middleware.NewStack("EnableSecurityHub", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpEnableSecurityHubMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableSecurityHub(options.Region), middleware.Before)

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
			OperationName: "EnableSecurityHub",
			Err:           err,
		}
	}
	out := result.(*EnableSecurityHubOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableSecurityHubInput struct {
	// The tags to add to the hub resource when you enable Security Hub.
	Tags map[string]*string
	// Whether to enable the security standards that Security Hub has designated as
	// automatically enabled. If you do not provide a value for EnableDefaultStandards,
	// it is set to true. To not enable the automatically enabled standards, set
	// EnableDefaultStandards to false.
	EnableDefaultStandards *bool
}

type EnableSecurityHubOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpEnableSecurityHubMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpEnableSecurityHub{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableSecurityHub{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableSecurityHub(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "EnableSecurityHub",
	}
}
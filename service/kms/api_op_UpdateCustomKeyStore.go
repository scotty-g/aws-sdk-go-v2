// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the properties of a custom key store. Use the CustomKeyStoreId parameter
// to identify the custom key store you want to edit. Use the remaining parameters
// to change the properties of the custom key store. You can only update a custom
// key store that is disconnected. To disconnect the custom key store, use
// DisconnectCustomKeyStore (). To reconnect the custom key store after the update
// completes, use ConnectCustomKeyStore (). To find the connection state of a
// custom key store, use the DescribeCustomKeyStores () operation. Use the
// parameters of UpdateCustomKeyStore to edit your keystore settings.
//
//     * Use
// the NewCustomKeyStoreName parameter to change the friendly name of the custom
// key store to the value that you specify.
//
//     * Use the KeyStorePassword
// parameter tell AWS KMS the current password of the kmsuser crypto user (CU)
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
// in the associated AWS CloudHSM cluster. You can use this parameter to fix
// connection failures
// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-password)
// that occur when AWS KMS cannot log into the associated cluster because the
// kmsuser password has changed. This value does not change the password in the AWS
// CloudHSM cluster.
//
//     * Use the CloudHsmClusterId parameter to associate the
// custom key store with a different, but related, AWS CloudHSM cluster. You can
// use this parameter to repair a custom key store if its AWS CloudHSM cluster
// becomes corrupted or is deleted, or when you need to create or restore a cluster
// from a backup.
//
// If the operation succeeds, it returns a JSON object with no
// properties. This operation is part of the Custom Key Store feature
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration of
// AWS KMS with the isolation and control of a single-tenant key store.
func (c *Client) UpdateCustomKeyStore(ctx context.Context, params *UpdateCustomKeyStoreInput, optFns ...func(*Options)) (*UpdateCustomKeyStoreOutput, error) {
	stack := middleware.NewStack("UpdateCustomKeyStore", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateCustomKeyStoreMiddlewares(stack)
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
	addOpUpdateCustomKeyStoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCustomKeyStore(options.Region), middleware.Before)

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
			OperationName: "UpdateCustomKeyStore",
			Err:           err,
		}
	}
	out := result.(*UpdateCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCustomKeyStoreInput struct {
	// Associates the custom key store with a related AWS CloudHSM cluster. Enter the
	// cluster ID of the cluster that you used to create the custom key store or a
	// cluster that shares a backup history and has the same cluster certificate as the
	// original cluster. You cannot use this parameter to associate a custom key store
	// with an unrelated cluster. In addition, the replacement cluster must fulfill the
	// requirements
	// (https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore)
	// for a cluster associated with a custom key store. To view the cluster
	// certificate of a cluster, use the DescribeClusters
	// (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	CloudHsmClusterId *string
	// Identifies the custom key store that you want to update. Enter the ID of the
	// custom key store. To find the ID of a custom key store, use the
	// DescribeCustomKeyStores () operation.
	CustomKeyStoreId *string
	// Changes the friendly name of the custom key store to the value that you specify.
	// The custom key store name must be unique in the AWS account.
	NewCustomKeyStoreName *string
	// Enter the current password of the kmsuser crypto user (CU) in the AWS CloudHSM
	// cluster that is associated with the custom key store. This parameter tells AWS
	// KMS the current password of the kmsuser crypto user (CU). It does not set or
	// change the password of any users in the AWS CloudHSM cluster.
	KeyStorePassword *string
}

type UpdateCustomKeyStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateCustomKeyStoreMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCustomKeyStore{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCustomKeyStore{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateCustomKeyStore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "UpdateCustomKeyStore",
	}
}
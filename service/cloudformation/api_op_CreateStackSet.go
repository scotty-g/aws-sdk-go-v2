// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a stack set.
func (c *Client) CreateStackSet(ctx context.Context, params *CreateStackSetInput, optFns ...func(*Options)) (*CreateStackSetOutput, error) {
	stack := middleware.NewStack("CreateStackSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateStackSetMiddlewares(stack)
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
	addIdempotencyToken_opCreateStackSetMiddleware(stack, options)
	addOpCreateStackSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStackSet(options.Region), middleware.Before)

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
			OperationName: "CreateStackSet",
			Err:           err,
		}
	}
	out := result.(*CreateStackSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStackSetInput struct {
	// The structure that contains the template body, with a minimum length of 1 byte
	// and a maximum length of 51,200 bytes. For more information, see Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must specify either the
	// TemplateBody or the TemplateURL parameter, but not both.
	TemplateBody *string
	// The name of the IAM execution role to use to create the stack set. If you do not
	// specify an execution role, AWS CloudFormation uses the
	// AWSCloudFormationStackSetExecutionRole role for the stack set operation. Specify
	// an IAM role only if you are using customized execution roles to control which
	// stack resources users and groups can include in their stack sets.  </p>
	ExecutionRoleName *string
	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack
	// set. Specify an IAM role only if you are using customized administrator roles to
	// control which users or groups can manage specific stack sets within the same
	// administrator account. For more information, see Prerequisites: Granting
	// Permissions for Stack Set Operations
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html)
	// in the AWS CloudFormation User Guide.
	AdministrationRoleARN *string
	// Describes whether StackSets automatically deploys to AWS Organizations accounts
	// that are added to the target organization or organizational unit (OU). Specify
	// only if PermissionModel is SERVICE_MANAGED.
	AutoDeployment *types.AutoDeployment
	// The name to associate with the stack set. The name must be unique in the Region
	// where you create your stack set. A stack name can contain only alphanumeric
	// characters (case-sensitive) and hyphens. It must start with an alphabetic
	// character and can't be longer than 128 characters.
	StackSetName *string
	// The input parameters for the stack set template.
	Parameters []*types.Parameter
	// A description of the stack set. You can use the description to identify the
	// stack set's purpose or other important information.
	Description *string
	// The key-value pairs to associate with this stack set and the stacks created from
	// it. AWS CloudFormation also propagates these tags to supported resources that
	// are created in the stacks. A maximum number of 50 tags can be specified. If you
	// specify tags as part of a CreateStackSet action, AWS CloudFormation checks to
	// see if you have the required IAM permission to tag resources. If you don't, the
	// entire CreateStackSet action fails with an access denied error, and the stack
	// set is not created.
	Tags []*types.Tag
	// The location of the file that contains the template body. The URL must point to
	// a template (maximum size: 460,800 bytes) that's located in an Amazon S3 bucket.
	// For more information, see Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must specify either the
	// TemplateBody or the TemplateURL parameter, but not both.
	TemplateURL *string
	// A unique identifier for this CreateStackSet request. Specify this token if you
	// plan to retry requests so that AWS CloudFormation knows that you're not
	// attempting to create another stack set with the same name. You might retry
	// CreateStackSet requests to ensure that AWS CloudFormation successfully received
	// them. If you don't specify an operation ID, the SDK generates one automatically.
	ClientRequestToken *string
	// In some cases, you must explicitly acknowledge that your stack set template
	// contains certain capabilities in order for AWS CloudFormation to create the
	// stack set and related stack instances.
	//
	//     * CAPABILITY_IAM and
	// CAPABILITY_NAMED_IAM Some stack templates might include resources that can
	// affect permissions in your AWS account; for example, by creating new AWS
	// Identity and Access Management (IAM) users. For those stack sets, you must
	// explicitly acknowledge this by specifying one of these capabilities. The
	// following IAM resources require you to specify either the CAPABILITY_IAM or
	// CAPABILITY_NAMED_IAM capability.
	//
	//         * If you have IAM resources, you can
	// specify either capability.
	//
	//         * If you have IAM resources with custom
	// names, you must specify CAPABILITY_NAMED_IAM.
	//
	//         * If you don't specify
	// either of these capabilities, AWS CloudFormation returns an
	// InsufficientCapabilities error.
	//
	//     If your stack template contains these
	// resources, we recommend that you review all permissions associated with them and
	// edit their permissions if necessary.
	//
	//         * AWS::IAM::AccessKey
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html)
	//
	//
	// * AWS::IAM::Group
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html)
	//
	//
	// * AWS::IAM::InstanceProfile
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)
	//
	//
	// * AWS::IAM::Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html)
	//
	//
	// * AWS::IAM::Role
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html)
	//
	//
	// * AWS::IAM::User
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html)
	//
	//
	// * AWS::IAM::UserToGroupAddition
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html)
	//
	//
	// For more information, see Acknowledging IAM Resources in AWS CloudFormation
	// Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	//
	//
	// * CAPABILITY_AUTO_EXPAND Some templates contain macros. If your stack template
	// contains one or more macros, and you choose to create a stack directly from the
	// processed template, without first reviewing the resulting changes in a change
	// set, you must acknowledge this capability. For more information, see Using AWS
	// CloudFormation Macros to Perform Custom Processing on Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html).
	// Stack sets do not currently support macros in stack templates. (This includes
	// the AWS::Include
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html)
	// and AWS::Serverless
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html)
	// transforms, which are macros hosted by AWS CloudFormation.) Even if you specify
	// this capability, if you include a macro in your template the stack set operation
	// will fail.
	Capabilities []types.Capability
	// Describes how the IAM roles required for stack set operations are created. By
	// default, SELF-MANAGED is specified.
	//
	//     * With self-managed permissions, you
	// must create the administrator and execution roles required to deploy to target
	// accounts. For more information, see Grant Self-Managed Stack Set Permissions
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html).
	//
	//
	// * With service-managed permissions, StackSets automatically creates the IAM
	// roles required to deploy to accounts managed by AWS Organizations. For more
	// information, see Grant Service-Managed Stack Set Permissions
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html).
	PermissionModel types.PermissionModels
}

type CreateStackSetOutput struct {
	// The ID of the stack set that you're creating.
	StackSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateStackSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateStackSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateStackSet{}, middleware.After)
}

type idempotencyToken_initializeOpCreateStackSet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateStackSet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateStackSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateStackSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateStackSetInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateStackSetMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateStackSet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateStackSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "CreateStackSet",
	}
}
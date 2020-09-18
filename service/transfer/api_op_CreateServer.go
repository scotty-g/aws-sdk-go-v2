// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Instantiates an autoscaling virtual server based on the selected file transfer
// protocol in AWS. When you make updates to your file transfer protocol-enabled
// server or when you work with users, use the service-generated ServerId property
// that is assigned to the newly created server.
func (c *Client) CreateServer(ctx context.Context, params *CreateServerInput, optFns ...func(*Options)) (*CreateServerOutput, error) {
	stack := middleware.NewStack("CreateServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateServerMiddlewares(stack)
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
	addOpCreateServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServer(options.Region), middleware.Before)

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
			OperationName: "CreateServer",
			Err:           err,
		}
	}
	out := result.(*CreateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServerInput struct {
	// The virtual private cloud (VPC) endpoint settings that are configured for your
	// file transfer protocol-enabled server. When you host your endpoint within your
	// VPC, you can make it accessible only to resources within your VPC, or you can
	// attach Elastic IPs and make it accessible to clients over the internet. Your
	// VPC's default security groups are automatically assigned to your endpoint.
	EndpointDetails *types.EndpointDetails
	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate.
	// Required when Protocols is set to FTPS.  <p>To request a new public certificate,
	// see <a
	// href="https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html">Request
	// a public certificate</a> in the <i> AWS Certificate Manager User Guide</i>.</p>
	// <p>To import an existing certificate into ACM, see <a
	// href="https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html">Importing
	// certificates into ACM</a> in the <i> AWS Certificate Manager User Guide</i>.</p>
	// <p>To request a private certificate to use FTPS through private IP addresses,
	// see <a
	// href="https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html">Request
	// a private certificate</a> in the <i> AWS Certificate Manager User Guide</i>.</p>
	// <p>Certificates with the following cryptographic algorithms and key sizes are
	// supported:</p> <ul> <li> <p>2048-bit RSA (RSA_2048)</p> </li> <li> <p>4096-bit
	// RSA (RSA_4096)</p> </li> <li> <p>Elliptic Prime Curve 256 bit
	// (EC_prime256v1)</p> </li> <li> <p>Elliptic Prime Curve 384 bit
	// (EC_secp384r1)</p> </li> <li> <p>Elliptic Prime Curve 521 bit (EC_secp521r1)</p>
	// </li> </ul> <note> <p>The certificate must be a valid SSL/TLS X.509 version 3
	// certificate with FQDN or IP address specified and information about the
	// issuer.</p> </note>
	Certificate *string
	// Key-value pairs that can be used to group and search for file transfer
	// protocol-enabled servers.
	Tags []*types.Tag
	// Specifies the mode of authentication for a file transfer protocol-enabled
	// server. The default value is SERVICE_MANAGED, which allows you to store and
	// access user credentials within the AWS Transfer Family service. Use the
	// API_GATEWAY value to integrate with an identity provider of your choosing. The
	// API_GATEWAY setting requires you to provide an API Gateway endpoint URL to call
	// for authentication using the IdentityProviderDetails parameter.
	IdentityProviderType types.IdentityProviderType
	// The RSA private key as generated by the ssh-keygen -N "" -m PEM -f
	// my-new-server-key command.  <important> <p>If you aren't planning to migrate
	// existing users from an existing SFTP-enabled server to a new server, don't
	// update the host key. Accidentally changing a server's host key can be
	// disruptive.</p> </important> <p>For more information, see <a
	// href="https://docs.aws.amazon.com/transfer/latest/userguide/edit-server-config.html#configuring-servers-change-host-key">Change
	// the host key for your SFTP-enabled server</a> in the <i>AWS Transfer Family User
	// Guide</i>.</p>
	HostKey *string
	// Specifies the file transfer protocol or protocols over which your file transfer
	// protocol client can connect to your server's endpoint. The available protocols
	// are:  <ul> <li> <p> <code>SFTP</code> (Secure Shell (SSH) File Transfer
	// Protocol): File transfer over SSH</p> </li> <li> <p> <code>FTPS</code> (File
	// Transfer Protocol Secure): File transfer with TLS encryption</p> </li> <li> <p>
	// <code>FTP</code> (File Transfer Protocol): Unencrypted file transfer</p> </li>
	// </ul> <note> <p>If you select <code>FTPS</code>, you must choose a certificate
	// stored in AWS Certificate Manager (ACM) which will be used to identify your
	// server when clients connect to it over FTPS.</p> <p>If <code>Protocol</code>
	// includes either <code>FTP</code> or <code>FTPS</code>, then the
	// <code>EndpointType</code> must be <code>VPC</code> and the
	// <code>IdentityProviderType</code> must be <code>API_GATEWAY</code>.</p> <p>If
	// <code>Protocol</code> includes <code>FTP</code>, then
	// <code>AddressAllocationIds</code> cannot be associated.</p> <p>If
	// <code>Protocol</code> is set only to <code>SFTP</code>, the
	// <code>EndpointType</code> can be set to <code>PUBLIC</code> and the
	// <code>IdentityProviderType</code> can be set to
	// <code>SERVICE_MANAGED</code>.</p> </note>
	Protocols []types.Protocol
	// The type of VPC endpoint that you want your file transfer protocol-enabled
	// server to connect to. You can choose to connect to the public internet or a VPC
	// endpoint. With a VPC endpoint, you can restrict access to your server and
	// resources only within your VPC.  <note> <p>It is recommended that you use
	// <code>VPC</code> as the <code>EndpointType</code>. With this endpoint type, you
	// have the option to directly associate up to three Elastic IPv4 addresses (BYO IP
	// included) with your server's endpoint and use VPC security groups to restrict
	// traffic by the client's public IP address. This is not possible with
	// <code>EndpointType</code> set to <code>VPC_ENDPOINT</code>.</p> </note>
	EndpointType types.EndpointType
	// Allows the service to write your users' activity to your Amazon CloudWatch logs
	// for monitoring and auditing purposes.
	LoggingRole *string
	// Required when IdentityProviderType is set to API_GATEWAY. Accepts an array
	// containing all of the information required to call a customer-supplied
	// authentication API, including the API Gateway URL. Not required when
	// IdentityProviderType is set to SERVICE_MANAGED.
	IdentityProviderDetails *types.IdentityProviderDetails
}

type CreateServerOutput struct {
	// The service-assigned ID of the file transfer protocol-enabled server that is
	// created.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateServer",
	}
}
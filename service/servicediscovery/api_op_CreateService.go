// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a service. This action defines the configuration for the following
// entities:
//   - For public and private DNS namespaces, one of the following combinations of
//     DNS records in Amazon Route 53:
//   - A
//   - AAAA
//   - A and AAAA
//   - SRV
//   - CNAME
//   - Optionally, a health check
//
// After you create the service, you can submit a RegisterInstance (https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html)
// request, and Cloud Map uses the values in the configuration to create the
// specified entities. For the current quota on the number of instances that you
// can register using the same namespace and using the same service, see Cloud Map
// quotas (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)
// in the Cloud Map Developer Guide.
func (c *Client) CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) {
	if params == nil {
		params = &CreateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateService", params, optFns, c.addOperationCreateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceInput struct {

	// The name that you want to assign to the service. Do not include sensitive
	// information in the name if the namespace is discoverable by public DNS queries.
	// If you want Cloud Map to create an SRV record when you register an instance and
	// you're using a system that requires a specific SRV format, such as HAProxy (http://www.haproxy.org/)
	// , specify the following for Name :
	//   - Start the name with an underscore (_), such as _exampleservice .
	//   - End the name with ._protocol, such as ._tcp .
	// When you register an instance, Cloud Map creates an SRV record and assigns a
	// name to the record by concatenating the service name and the namespace name (for
	// example, _exampleservice._tcp.example.com ). For services that are accessible by
	// DNS queries, you can't create multiple services with names that differ only by
	// case (such as EXAMPLE and example). Otherwise, these services have the same DNS
	// name and can't be distinguished. However, if you use a namespace that's only
	// accessible by API calls, then you can create services that with names that
	// differ only by case.
	//
	// This member is required.
	Name *string

	// A unique string that identifies the request and that allows failed CreateService
	// requests to be retried without the risk of running the operation twice.
	// CreatorRequestId can be any unique string (for example, a date/timestamp).
	CreatorRequestId *string

	// A description for the service.
	Description *string

	// A complex type that contains information about the Amazon Route 53 records that
	// you want Cloud Map to create when you register an instance.
	DnsConfig *types.DnsConfig

	// Public DNS and HTTP namespaces only. A complex type that contains settings for
	// an optional Route 53 health check. If you specify settings for a health check,
	// Cloud Map associates the health check with all the Route 53 DNS records that you
	// specify in DnsConfig . If you specify a health check configuration, you can
	// specify either HealthCheckCustomConfig or HealthCheckConfig but not both. For
	// information about the charges for health checks, see Cloud Map Pricing (http://aws.amazon.com/cloud-map/pricing/)
	// .
	HealthCheckConfig *types.HealthCheckConfig

	// A complex type that contains information about an optional custom health check.
	// If you specify a health check configuration, you can specify either
	// HealthCheckCustomConfig or HealthCheckConfig but not both. You can't add,
	// update, or delete a HealthCheckCustomConfig configuration from an existing
	// service.
	HealthCheckCustomConfig *types.HealthCheckCustomConfig

	// The ID of the namespace that you want to use to create the service. The
	// namespace ID must be specified, but it can be specified either here or in the
	// DnsConfig object.
	NamespaceId *string

	// The tags to add to the service. Each tag consists of a key and an optional
	// value that you define. Tags keys can be up to 128 characters in length, and tag
	// values can be up to 256 characters in length.
	Tags []types.Tag

	// If present, specifies that the service instances are only discoverable using
	// the DiscoverInstances API operation. No DNS records is registered for the
	// service instances. The only valid value is HTTP .
	Type types.ServiceTypeOption

	noSmithyDocumentSerde
}

type CreateServiceOutput struct {

	// A complex type that contains information about the new service.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateService{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateServiceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateService(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

type idempotencyToken_initializeOpCreateService struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateService) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateServiceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateService{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "CreateService",
	}
}

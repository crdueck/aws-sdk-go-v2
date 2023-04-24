// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Submits an updated route configuration for a Multi-Region Access Point. This
// API operation updates the routing status for the specified Regions from active
// to passive, or from passive to active. A value of 0 indicates a passive status,
// which means that traffic won't be routed to the specified Region. A value of 100
// indicates an active status, which means that traffic will be routed to the
// specified Region. At least one Region must be active at all times. When the
// routing configuration is changed, any in-progress operations (uploads, copies,
// deletes, and so on) to formerly active Regions will continue to run to their
// final completion state (success or failure). The routing configurations of any
// Regions that aren’t specified remain unchanged. Updated routing configurations
// might not be immediately applied. It can take up to 2 minutes for your changes
// to take effect. To submit routing control changes and failover requests, use the
// Amazon S3 failover control infrastructure endpoints in these five Amazon Web
// Services Regions:
//   - us-east-1
//   - us-west-2
//   - ap-southeast-2
//   - ap-northeast-1
//   - eu-west-1
//
// Your Amazon S3 bucket does not need to be in these five Regions.
func (c *Client) SubmitMultiRegionAccessPointRoutes(ctx context.Context, params *SubmitMultiRegionAccessPointRoutesInput, optFns ...func(*Options)) (*SubmitMultiRegionAccessPointRoutesOutput, error) {
	if params == nil {
		params = &SubmitMultiRegionAccessPointRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitMultiRegionAccessPointRoutes", params, optFns, c.addOperationSubmitMultiRegionAccessPointRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitMultiRegionAccessPointRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitMultiRegionAccessPointRoutesInput struct {

	// The Amazon Web Services account ID for the owner of the Multi-Region Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// The Multi-Region Access Point ARN.
	//
	// This member is required.
	Mrap *string

	// The different routes that make up the new route configuration. Active routes
	// return a value of 100 , and passive routes return a value of 0 .
	//
	// This member is required.
	RouteUpdates []types.MultiRegionAccessPointRoute

	noSmithyDocumentSerde
}

type SubmitMultiRegionAccessPointRoutesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSubmitMultiRegionAccessPointRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpSubmitMultiRegionAccessPointRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpSubmitMultiRegionAccessPointRoutes{}, middleware.After)
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opSubmitMultiRegionAccessPointRoutesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSubmitMultiRegionAccessPointRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitMultiRegionAccessPointRoutes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addSubmitMultiRegionAccessPointRoutesUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opSubmitMultiRegionAccessPointRoutesMiddleware struct {
}

func (*endpointPrefix_opSubmitMultiRegionAccessPointRoutesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opSubmitMultiRegionAccessPointRoutesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*SubmitMultiRegionAccessPointRoutesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opSubmitMultiRegionAccessPointRoutesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opSubmitMultiRegionAccessPointRoutesMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opSubmitMultiRegionAccessPointRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "SubmitMultiRegionAccessPointRoutes",
	}
}

func copySubmitMultiRegionAccessPointRoutesInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*SubmitMultiRegionAccessPointRoutesInput)
	if !ok {
		return nil, fmt.Errorf("expect *SubmitMultiRegionAccessPointRoutesInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillSubmitMultiRegionAccessPointRoutesAccountID(input interface{}, v string) error {
	in := input.(*SubmitMultiRegionAccessPointRoutesInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addSubmitMultiRegionAccessPointRoutesUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copySubmitMultiRegionAccessPointRoutesInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

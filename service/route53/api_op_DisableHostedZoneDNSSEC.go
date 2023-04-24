// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables DNSSEC signing in a specific hosted zone. This action does not
// deactivate any key-signing keys (KSKs) that are active in the hosted zone.
func (c *Client) DisableHostedZoneDNSSEC(ctx context.Context, params *DisableHostedZoneDNSSECInput, optFns ...func(*Options)) (*DisableHostedZoneDNSSECOutput, error) {
	if params == nil {
		params = &DisableHostedZoneDNSSECInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableHostedZoneDNSSEC", params, optFns, c.addOperationDisableHostedZoneDNSSECMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableHostedZoneDNSSECOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableHostedZoneDNSSECInput struct {

	// A unique string used to identify a hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	noSmithyDocumentSerde
}

type DisableHostedZoneDNSSECOutput struct {

	// A complex type that describes change information about changes made to your
	// hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableHostedZoneDNSSECMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDisableHostedZoneDNSSEC{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDisableHostedZoneDNSSEC{}, middleware.After)
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
	if err = addOpDisableHostedZoneDNSSECValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableHostedZoneDNSSEC(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDisableHostedZoneDNSSEC(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DisableHostedZoneDNSSEC",
	}
}

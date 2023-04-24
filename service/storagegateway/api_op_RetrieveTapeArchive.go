// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves an archived virtual tape from the virtual tape shelf (VTS) to a tape
// gateway. Virtual tapes archived in the VTS are not associated with any gateway.
// However after a tape is retrieved, it is associated with a gateway, even though
// it is also listed in the VTS, that is, archive. This operation is only supported
// in the tape gateway type. Once a tape is successfully retrieved to a gateway, it
// cannot be retrieved again to another gateway. You must archive the tape again
// before you can retrieve it to another gateway. This operation is only supported
// in the tape gateway type.
func (c *Client) RetrieveTapeArchive(ctx context.Context, params *RetrieveTapeArchiveInput, optFns ...func(*Options)) (*RetrieveTapeArchiveOutput, error) {
	if params == nil {
		params = &RetrieveTapeArchiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetrieveTapeArchive", params, optFns, c.addOperationRetrieveTapeArchiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetrieveTapeArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// RetrieveTapeArchiveInput
type RetrieveTapeArchiveInput struct {

	// The Amazon Resource Name (ARN) of the gateway you want to retrieve the virtual
	// tape to. Use the ListGateways operation to return a list of gateways for your
	// account and Amazon Web Services Region. You retrieve archived virtual tapes to
	// only one gateway and the gateway must be a tape gateway.
	//
	// This member is required.
	GatewayARN *string

	// The Amazon Resource Name (ARN) of the virtual tape you want to retrieve from
	// the virtual tape shelf (VTS).
	//
	// This member is required.
	TapeARN *string

	noSmithyDocumentSerde
}

// RetrieveTapeArchiveOutput
type RetrieveTapeArchiveOutput struct {

	// The Amazon Resource Name (ARN) of the retrieved virtual tape.
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetrieveTapeArchiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRetrieveTapeArchive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetrieveTapeArchive{}, middleware.After)
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
	if err = addOpRetrieveTapeArchiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieveTapeArchive(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetrieveTapeArchive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "RetrieveTapeArchive",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a list of asset property values to IoT SiteWise. Each value is a
// timestamp-quality-value (TQV) data point. For more information, see Ingesting
// data using the API (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/ingest-api.html)
// in the IoT SiteWise User Guide. To identify an asset property, you must specify
// one of the following:
//   - The assetId and propertyId of an asset property.
//   - A propertyAlias , which is a data stream alias (for example,
//     /company/windfarm/3/turbine/7/temperature ). To define an asset property's
//     alias, see UpdateAssetProperty (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html)
//     .
//
// With respect to Unix epoch time, IoT SiteWise accepts only TQVs that have a
// timestamp of no more than 7 days in the past and no more than 10 minutes in the
// future. IoT SiteWise rejects timestamps outside of the inclusive range of [-7
// days, +10 minutes] and returns a TimestampOutOfRangeException error. For each
// asset property, IoT SiteWise overwrites TQVs with duplicate timestamps unless
// the newer TQV has a different quality. For example, if you store a TQV {T1,
// GOOD, V1} , then storing {T1, GOOD, V2} replaces the existing TQV. IoT SiteWise
// authorizes access to each BatchPutAssetPropertyValue entry individually. For
// more information, see BatchPutAssetPropertyValue authorization (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-id-based-policies-batchputassetpropertyvalue-action)
// in the IoT SiteWise User Guide.
func (c *Client) BatchPutAssetPropertyValue(ctx context.Context, params *BatchPutAssetPropertyValueInput, optFns ...func(*Options)) (*BatchPutAssetPropertyValueOutput, error) {
	if params == nil {
		params = &BatchPutAssetPropertyValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutAssetPropertyValue", params, optFns, c.addOperationBatchPutAssetPropertyValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutAssetPropertyValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutAssetPropertyValueInput struct {

	// The list of asset property value entries for the batch put request. You can
	// specify up to 10 entries per request.
	//
	// This member is required.
	Entries []types.PutAssetPropertyValueEntry

	noSmithyDocumentSerde
}

type BatchPutAssetPropertyValueOutput struct {

	// A list of the errors (if any) associated with the batch put request. Each error
	// entry contains the entryId of the entry that failed.
	//
	// This member is required.
	ErrorEntries []types.BatchPutAssetPropertyErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchPutAssetPropertyValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchPutAssetPropertyValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchPutAssetPropertyValue{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchPutAssetPropertyValueMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchPutAssetPropertyValueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutAssetPropertyValue(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchPutAssetPropertyValueMiddleware struct {
}

func (*endpointPrefix_opBatchPutAssetPropertyValueMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchPutAssetPropertyValueMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchPutAssetPropertyValueMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchPutAssetPropertyValueMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchPutAssetPropertyValue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "BatchPutAssetPropertyValue",
	}
}

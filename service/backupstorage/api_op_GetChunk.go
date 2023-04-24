// Code generated by smithy-go-codegen DO NOT EDIT.

package backupstorage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backupstorage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Gets the specified object's chunk.
func (c *Client) GetChunk(ctx context.Context, params *GetChunkInput, optFns ...func(*Options)) (*GetChunkOutput, error) {
	if params == nil {
		params = &GetChunkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChunk", params, optFns, c.addOperationGetChunkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChunkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChunkInput struct {

	// Chunk token
	//
	// This member is required.
	ChunkToken *string

	// Storage job id
	//
	// This member is required.
	StorageJobId *string

	noSmithyDocumentSerde
}

type GetChunkOutput struct {

	// Data checksum
	//
	// This member is required.
	Checksum *string

	// Checksum algorithm
	//
	// This member is required.
	ChecksumAlgorithm types.DataChecksumAlgorithm

	// Chunk data
	//
	// This member is required.
	Data io.ReadCloser

	// Data length
	//
	// This member is required.
	Length int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChunkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetChunk{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetChunk{}, middleware.After)
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
	if err = addOpGetChunkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChunk(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetChunk(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-storage",
		OperationName: "GetChunk",
	}
}

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

// Complete upload
func (c *Client) NotifyObjectComplete(ctx context.Context, params *NotifyObjectCompleteInput, optFns ...func(*Options)) (*NotifyObjectCompleteOutput, error) {
	if params == nil {
		params = &NotifyObjectCompleteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyObjectComplete", params, optFns, c.addOperationNotifyObjectCompleteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyObjectCompleteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyObjectCompleteInput struct {

	// Backup job Id for the in-progress backup
	//
	// This member is required.
	BackupJobId *string

	// Object checksum
	//
	// This member is required.
	ObjectChecksum *string

	// Checksum algorithm
	//
	// This member is required.
	ObjectChecksumAlgorithm types.SummaryChecksumAlgorithm

	// Upload Id for the in-progress upload
	//
	// This member is required.
	UploadId *string

	// Optional metadata associated with an Object. Maximum length is 4MB.
	MetadataBlob io.Reader

	// Checksum of MetadataBlob.
	MetadataBlobChecksum *string

	// Checksum algorithm.
	MetadataBlobChecksumAlgorithm types.DataChecksumAlgorithm

	// The size of MetadataBlob.
	MetadataBlobLength int64

	// Optional metadata associated with an Object. Maximum string length is 256 bytes.
	MetadataString *string

	noSmithyDocumentSerde
}

type NotifyObjectCompleteOutput struct {

	// Object checksum
	//
	// This member is required.
	ObjectChecksum *string

	// Checksum algorithm
	//
	// This member is required.
	ObjectChecksumAlgorithm types.SummaryChecksumAlgorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNotifyObjectCompleteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpNotifyObjectComplete{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpNotifyObjectComplete{}, middleware.After)
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
	if err = v4.AddUnsignedPayloadMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = addOpNotifyObjectCompleteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyObjectComplete(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNotifyObjectComplete(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-storage",
		OperationName: "NotifyObjectComplete",
	}
}

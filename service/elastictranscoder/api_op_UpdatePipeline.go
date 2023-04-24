// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use the UpdatePipeline operation to update settings for a pipeline. When you
// change pipeline settings, your changes take effect immediately. Jobs that you
// have already submitted and that Elastic Transcoder has not started to process
// are affected in addition to jobs that you submit after you change settings.
func (c *Client) UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) {
	if params == nil {
		params = &UpdatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePipeline", params, optFns, c.addOperationUpdatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdatePipelineRequest structure.
type UpdatePipelineInput struct {

	// The ID of the pipeline that you want to update.
	//
	// This member is required.
	Id *string

	// The AWS Key Management Service (AWS KMS) key that you want to use with this
	// pipeline. If you use either s3 or s3-aws-kms as your Encryption:Mode , you don't
	// need to provide a key with your job because a default key, known as an AWS-KMS
	// key, is created for you automatically. You need to provide an AWS-KMS key only
	// if you want to use a non-default AWS-KMS key, or if you are using an
	// Encryption:Mode of aes-cbc-pkcs7 , aes-ctr , or aes-gcm .
	AwsKmsKeyArn *string

	// The optional ContentConfig object specifies information about the Amazon S3
	// bucket in which you want Elastic Transcoder to save transcoded files and
	// playlists: which bucket to use, which users you want to have access to the
	// files, the type of access you want users to have, and the storage class that you
	// want to assign to the files. If you specify values for ContentConfig , you must
	// also specify values for ThumbnailConfig . If you specify values for
	// ContentConfig and ThumbnailConfig , omit the OutputBucket object.
	//   - Bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save
	//   transcoded files and playlists.
	//   - Permissions (Optional): The Permissions object specifies which users you
	//   want to have access to transcoded files and the type of access you want them to
	//   have. You can grant permissions to a maximum of 30 users and/or predefined
	//   Amazon S3 groups.
	//   - Grantee Type: Specify the type of value that appears in the Grantee object:
	//   - Canonical: The value in the Grantee object is either the canonical user ID
	//   for an AWS account or an origin access identity for an Amazon CloudFront
	//   distribution. For more information about canonical user IDs, see Access Control
	//   List (ACL) Overview in the Amazon Simple Storage Service Developer Guide. For
	//   more information about using CloudFront origin access identities to require that
	//   users use CloudFront URLs instead of Amazon S3 URLs, see Using an Origin Access
	//   Identity to Restrict Access to Your Amazon S3 Content. A canonical user ID is
	//   not the same as an AWS account number.
	//   - Email: The value in the Grantee object is the registered email address of an
	//   AWS account.
	//   - Group: The value in the Grantee object is one of the following predefined
	//   Amazon S3 groups: AllUsers , AuthenticatedUsers , or LogDelivery .
	//   - Grantee: The AWS user or group that you want to have access to transcoded
	//   files and playlists. To identify the user or group, you can specify the
	//   canonical user ID for an AWS account, an origin access identity for a CloudFront
	//   distribution, the registered email address of an AWS account, or a predefined
	//   Amazon S3 group
	//   - Access: The permission that you want to give to the AWS user that you
	//   specified in Grantee . Permissions are granted on the files that Elastic
	//   Transcoder adds to the bucket, including playlists and video files. Valid values
	//   include:
	//   - READ : The grantee can read the objects and metadata for objects that
	//   Elastic Transcoder adds to the Amazon S3 bucket.
	//   - READ_ACP : The grantee can read the object ACL for objects that Elastic
	//   Transcoder adds to the Amazon S3 bucket.
	//   - WRITE_ACP : The grantee can write the ACL for the objects that Elastic
	//   Transcoder adds to the Amazon S3 bucket.
	//   - FULL_CONTROL : The grantee has READ , READ_ACP , and WRITE_ACP permissions
	//   for the objects that Elastic Transcoder adds to the Amazon S3 bucket.
	//   - StorageClass: The Amazon S3 storage class, Standard or ReducedRedundancy ,
	//   that you want Elastic Transcoder to assign to the video files and playlists that
	//   it stores in your Amazon S3 bucket.
	ContentConfig *types.PipelineOutputConfig

	// The Amazon S3 bucket in which you saved the media files that you want to
	// transcode and the graphics that you want to use as watermarks.
	InputBucket *string

	// The name of the pipeline. We recommend that the name be unique within the AWS
	// account, but uniqueness is not enforced. Constraints: Maximum 40 characters
	Name *string

	// The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic
	// that you want to notify to report job status. To receive notifications, you must
	// also subscribe to the new topic in the Amazon SNS console.
	//   - Progressing: The topic ARN for the Amazon Simple Notification Service
	//   (Amazon SNS) topic that you want to notify when Elastic Transcoder has started
	//   to process jobs that are added to this pipeline. This is the ARN that Amazon SNS
	//   returned when you created the topic.
	//   - Complete: The topic ARN for the Amazon SNS topic that you want to notify
	//   when Elastic Transcoder has finished processing a job. This is the ARN that
	//   Amazon SNS returned when you created the topic.
	//   - Warning: The topic ARN for the Amazon SNS topic that you want to notify
	//   when Elastic Transcoder encounters a warning condition. This is the ARN that
	//   Amazon SNS returned when you created the topic.
	//   - Error: The topic ARN for the Amazon SNS topic that you want to notify when
	//   Elastic Transcoder encounters an error condition. This is the ARN that Amazon
	//   SNS returned when you created the topic.
	Notifications *types.Notifications

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic
	// Transcoder to use to transcode jobs for this pipeline.
	Role *string

	// The ThumbnailConfig object specifies several values, including the Amazon S3
	// bucket in which you want Elastic Transcoder to save thumbnail files, which users
	// you want to have access to the files, the type of access you want users to have,
	// and the storage class that you want to assign to the files. If you specify
	// values for ContentConfig , you must also specify values for ThumbnailConfig
	// even if you don't want to create thumbnails. If you specify values for
	// ContentConfig and ThumbnailConfig , omit the OutputBucket object.
	//   - Bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save
	//   thumbnail files.
	//   - Permissions (Optional): The Permissions object specifies which users and/or
	//   predefined Amazon S3 groups you want to have access to thumbnail files, and the
	//   type of access you want them to have. You can grant permissions to a maximum of
	//   30 users and/or predefined Amazon S3 groups.
	//   - GranteeType: Specify the type of value that appears in the Grantee object:
	//   - Canonical: The value in the Grantee object is either the canonical user ID
	//   for an AWS account or an origin access identity for an Amazon CloudFront
	//   distribution. A canonical user ID is not the same as an AWS account number.
	//   - Email: The value in the Grantee object is the registered email address of an
	//   AWS account.
	//   - Group: The value in the Grantee object is one of the following predefined
	//   Amazon S3 groups: AllUsers , AuthenticatedUsers , or LogDelivery .
	//   - Grantee: The AWS user or group that you want to have access to thumbnail
	//   files. To identify the user or group, you can specify the canonical user ID for
	//   an AWS account, an origin access identity for a CloudFront distribution, the
	//   registered email address of an AWS account, or a predefined Amazon S3 group.
	//   - Access: The permission that you want to give to the AWS user that you
	//   specified in Grantee . Permissions are granted on the thumbnail files that
	//   Elastic Transcoder adds to the bucket. Valid values include:
	//   - READ : The grantee can read the thumbnails and metadata for objects that
	//   Elastic Transcoder adds to the Amazon S3 bucket.
	//   - READ_ACP : The grantee can read the object ACL for thumbnails that Elastic
	//   Transcoder adds to the Amazon S3 bucket.
	//   - WRITE_ACP : The grantee can write the ACL for the thumbnails that Elastic
	//   Transcoder adds to the Amazon S3 bucket.
	//   - FULL_CONTROL : The grantee has READ , READ_ACP , and WRITE_ACP permissions
	//   for the thumbnails that Elastic Transcoder adds to the Amazon S3 bucket.
	//   - StorageClass: The Amazon S3 storage class, Standard or ReducedRedundancy ,
	//   that you want Elastic Transcoder to assign to the thumbnails that it stores in
	//   your Amazon S3 bucket.
	ThumbnailConfig *types.PipelineOutputConfig

	noSmithyDocumentSerde
}

// When you update a pipeline, Elastic Transcoder returns the values that you
// specified in the request.
type UpdatePipelineOutput struct {

	// The pipeline updated by this UpdatePipelineResponse call.
	Pipeline *types.Pipeline

	// Elastic Transcoder returns a warning if the resources used by your pipeline are
	// not in the same region as the pipeline. Using resources in the same region, such
	// as your Amazon S3 buckets, Amazon SNS notification topics, and AWS KMS key,
	// reduces processing time and prevents cross-regional charges.
	Warnings []types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePipeline{}, middleware.After)
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
	if err = addOpUpdatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "UpdatePipeline",
	}
}

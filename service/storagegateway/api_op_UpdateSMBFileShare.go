// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a Server Message Block (SMB) file share. This operation is only
// supported for S3 File Gateways. To leave a file share field unchanged, set the
// corresponding input field to null. File gateways require Security Token Service
// (Amazon Web Services STS) to be activated to enable you to create a file share.
// Make sure that Amazon Web Services STS is activated in the Amazon Web Services
// Region you are creating your file gateway in. If Amazon Web Services STS is not
// activated in this Amazon Web Services Region, activate it. For information about
// how to activate Amazon Web Services STS, see Activating and deactivating Amazon
// Web Services STS in an Amazon Web Services Region (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the Identity and Access Management User Guide. File gateways don't support
// creating hard or symbolic links on a file share.
func (c *Client) UpdateSMBFileShare(ctx context.Context, params *UpdateSMBFileShareInput, optFns ...func(*Options)) (*UpdateSMBFileShareOutput, error) {
	if params == nil {
		params = &UpdateSMBFileShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSMBFileShare", params, optFns, c.addOperationUpdateSMBFileShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSMBFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateSMBFileShareInput
type UpdateSMBFileShareInput struct {

	// The Amazon Resource Name (ARN) of the SMB file share that you want to update.
	//
	// This member is required.
	FileShareARN *string

	// The files and folders on this share will only be visible to users with read
	// access.
	AccessBasedEnumeration *bool

	// A list of users or groups in the Active Directory that have administrator
	// rights to the file share. A group must be prefixed with the @ character.
	// Acceptable formats include: DOMAIN\User1 , user1 , @group1 , and @DOMAIN\group1
	// . Can only be set if Authentication is set to ActiveDirectory .
	AdminUserList []string

	// The Amazon Resource Name (ARN) of the storage used for audit logs.
	AuditDestinationARN *string

	// Specifies refresh cache information for the file share.
	CacheAttributes *types.CacheAttributes

	// The case of an object name in an Amazon S3 bucket. For ClientSpecified , the
	// client determines the case sensitivity. For CaseSensitive , the gateway
	// determines the case sensitivity. The default value is ClientSpecified .
	CaseSensitivity types.CaseSensitivity

	// The default storage class for objects put into an Amazon S3 bucket by the S3
	// File Gateway. The default value is S3_STANDARD . Optional. Valid Values:
	// S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA | S3_ONEZONE_IA
	DefaultStorageClass *string

	// The name of the file share. Optional. FileShareName must be set if an S3 prefix
	// name is set in LocationARN , or if an access point or access point alias is used.
	FileShareName *string

	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false . The default value is true . Valid Values: true | false
	GuessMIMETypeEnabled *bool

	// A list of users or groups in the Active Directory that are not allowed to
	// access the file share. A group must be prefixed with the @ character. Acceptable
	// formats include: DOMAIN\User1 , user1 , @group1 , and @DOMAIN\group1 . Can only
	// be set if Authentication is set to ActiveDirectory .
	InvalidUserList []string

	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or
	// false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used
	// for Amazon S3 server-side encryption. Storage Gateway does not support
	// asymmetric CMKs. This value can only be set when KMSEncrypted is true . Optional.
	KMSKey *string

	// The notification policy of the file share. SettlingTimeInSeconds controls the
	// number of seconds to wait after the last point in time a client wrote to a file
	// before generating an ObjectUploaded notification. Because clients can make many
	// small writes to files, it's best to set this parameter for as long as possible
	// to avoid generating multiple notifications for the same file in a small time
	// period. SettlingTimeInSeconds has no effect on the timing of the object
	// uploading to Amazon S3, only the timing of the notification. The following
	// example sets NotificationPolicy on with SettlingTimeInSeconds set to 60.
	// {"Upload": {"SettlingTimeInSeconds": 60}} The following example sets
	// NotificationPolicy off. {}
	NotificationPolicy *string

	// A value that sets the access control list (ACL) permission for objects in the
	// S3 bucket that a S3 File Gateway puts objects into. The default value is private
	// .
	ObjectACL types.ObjectACL

	// Specifies whether opportunistic locking is enabled for the SMB file share.
	// Enabling opportunistic locking on case-sensitive shares is not recommended for
	// workloads that involve access to files with the same name in different case.
	// Valid Values: true | false
	OplocksEnabled *bool

	// A value that sets the write status of a file share. Set this value to true to
	// set write status to read-only, otherwise set to false . Valid Values: true |
	// false
	ReadOnly *bool

	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true , the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data. RequesterPays is a configuration
	// for the S3 bucket that backs the file share, so make sure that the configuration
	// on the file share is the same as the S3 bucket configuration. Valid Values: true
	// | false
	RequesterPays *bool

	// Set this value to true to enable access control list (ACL) on the SMB file
	// share. Set it to false to map file and directory permissions to the POSIX
	// permissions. For more information, see Using Microsoft Windows ACLs to control
	// access to an SMB file share (https://docs.aws.amazon.com/storagegateway/latest/userguide/smb-acl.html)
	// in the Storage Gateway User Guide. Valid Values: true | false
	SMBACLEnabled *bool

	// A list of users or groups in the Active Directory that are allowed to access
	// the file share. A group must be prefixed with the @ character. Acceptable
	// formats include: DOMAIN\User1 , user1 , @group1 , and @DOMAIN\group1 . Can only
	// be set if Authentication is set to ActiveDirectory .
	ValidUserList []string

	noSmithyDocumentSerde
}

// UpdateSMBFileShareOutput
type UpdateSMBFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the updated SMB file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSMBFileShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSMBFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSMBFileShare{}, middleware.After)
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
	if err = addOpUpdateSMBFileShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSMBFileShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSMBFileShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateSMBFileShare",
	}
}

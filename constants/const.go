package constants

const (
	Account = "964685753046"
	Region  = "eu-north-1"

	FtpUsername = "ftpUsername"
	FtpPassword = "ftpUsername"

	SftpRoleName      = "SftpUserRole" // Role name, not the ARN
	S3BucketName      = "your-sftp-bucket-name"
	SftpHomeDirectory = "/your-sftp-bucket-name/home/${transfer:UserName}"
)

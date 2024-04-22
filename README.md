# My CDK SFTP

### Config
go to file contants/const.go

```bash
    const (
    MikeDevAccount = "<your_account>"
    Stockholm      = "<your_region>"

	FtpUsername = "<your_ftpUsername>"
	FtpPassword = "<your_ftpPassword>"

	SftpRoleName      = "SftpUserRole"
	S3BucketName      = "your-sftp-bucket-name"
	SftpHomeDirectory = "/your-sftp-bucket-name/home/${transfer:UserName}"
)

```

### Deploy

```bash
go mod tidy
make synth
make deploy
```

### Create user for connect sftp

https://www.youtube.com/watch?v=3_HHSnoFsoM
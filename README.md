# My CDK SFTP

### Config
go to file cmd/app/main.go

```bash
    const (
        MikeDevAccount = "<your_account>"
        Stockholm      = "<your_region>"
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
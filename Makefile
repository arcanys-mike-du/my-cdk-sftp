synth:
	npx aws-cdk@2 --app 'go run ./cmd/app/main.go' synth --all

deploy:
	npx aws-cdk@2 --app 'go run ./cmd/app/main.go' deploy --all

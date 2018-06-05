# Install project requirements


# Launch development envrionment

```
GOOS=linux go build -o main
sam local start-api --env-vars configuration.json
```

# Go to http://127.0.0.1:3000/index

# Generate artefact and push it to s3

````
aws cloudformation package \
--template-file template.yaml \
--output-template-file template-out.yaml \
--s3-bucket wallet.lambda
````

# Deploy artefact

````
aws cloudformation deploy \
--template-file template-out.yaml \
--stack-name wallet-ms-go \
--capabilities CAPABILITY_IAM
````

# Invoking function with event file
```
sam local invoke "MigrateFunction" -e event.json --env-vars configuration.json
 ```

# Invoking function with event via stdin
```
echo '{"message": "Hey, are you there?" }' | sam local invoke "MigrateFunction" --env-vars configuration.json
 ```

# For more options
```
sam local invoke --help
 ```
# Setup Lambda

## Create lambda IAM role

`aws iam create-role --role-name go-aws-lambda --assume-role-policy-document file://trust-policy.json`

**Result:**
```json
{
    "Role": {
        "Path": "/",
        "RoleName": "go-aws-lambda",
        "RoleId": "AROA2DT2JTJOE5THTQMKV",
        "Arn": "arn:aws:iam::012345678901:role/go-aws-lambda",
        "CreateDate": "2022-09-24T10:57:58+00:00",
        "AssumeRolePolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "lambda.amazonaws.com"
                    },
                    "Action": "sts:AssumeRole"
                }
            ]
        }
    }
}
```

## Attach Lambda execution policy to role

`aws iam attach-role-policy --role-name go-aws-lambda --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole`

## Build package

`go build main.go`

## Compress package

`zip function.zip main`

## Create Lambda Function

`aws lambda create-function --function-name go-aws-lambda --zip-file fileb://function.zip --handler main --runtime go1.x --role arn:aws:iam::694967442012:role/go-aws-lambda`

**Result:**
```json
{
    "FunctionName": "go-aws-lambda",
    "FunctionArn": "arn:aws:lambda:eu-west-2:012345678901:function:go-aws-lambda",
    "Runtime": "go1.x",
    "Role": "arn:aws:iam::012345678901:role/go-aws-lambda",
    "Handler": "main",
    "CodeSize": 4434162,
    "Description": "",
    "Timeout": 3,
    "MemorySize": 128,
    "LastModified": "2022-09-24T11:14:25.673+0000",
    "CodeSha256": "jyqQrPyfAVuiJzcKyi9+hfu4iE2xV2LPCs7Chr9Lxpc=",
    "Version": "$LATEST",
    "TracingConfig": {
        "Mode": "PassThrough"
    },
    "RevisionId": "512b94ac-b949-4edd-82da-8c8385a0480a",
    "State": "Pending",
    "StateReason": "The function is being created.",
    "StateReasonCode": "Creating",
    "PackageType": "Zip",
    "Architectures": [
        "x86_64"
    ]
}
```

## Invoke Lambda Function

`aws lambda invoke --function-name go-aws-lambda --cli-binary-format raw-in-base64-out --payload '{"What is your name?": "John", "How old are you?": 33}' output.txt`

**Error:**
```json
{
    "StatusCode": 200,
    "FunctionError": "Unhandled",
    "ExecutedVersion": "$LATEST"
}
```

## Output file created

**Error:**
```json
{
   "errorMessage":"fork/exec /var/task/main: exec format error",
   "errorType":"PathError"
}
```

**Fix:** Binary for the lambda has to be built to allow for running on Amazon Linux. Build binary and redeploy with below build command.

`GOARCH=amd64 GOOS=linux go build main.go`

(Delete existing lambda and recreate: `aws lambda delete-function --function-name go-aws-lambda`)

**Result:**
```json
{
  "Answer:": "John is 33 years old!"
}
```
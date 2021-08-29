# aws-lambda-vs-code-debugging-and-local-test

 ```
aws --endpoint-url=http://localhost:4566 s3 ls --profile=localstack - list buckets
 
aws --endpoint-url=http://localhost:4566 s3 ls s3://tutorial-bucket --recursive --profile=localstack - list files in a bucket
 
aws --endpoint-url=http://localhost:4566 s3 cp s3://tutorial-bucket/1.txt 1.txt --profile=localstack - doenload file from S3 bucket
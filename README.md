<p align="center">
  <img width="460" height="300" src="https://miro.medium.com/max/700/1*s3KIaqOXnfT2AAfwruzzIQ.png">
</p>

<h1 align="center"><a href="https://aws.plainenglish.io/aws-lambda-testing-and-debugging-using-visual-studio-code-aws-sam-and-docker-cbd095c2db74">AWS AWS Lambda Testing and Debugging using Visual Studio Code, AWS SAM, and Docker</a></h1>

# Commands used to access S3 bucket locally

```
aws --endpoint-url=http://localhost:4566 s3 ls --profile=localstack - list buckets

aws --endpoint-url=http://localhost:4566 s3 ls s3://tutorial-bucket --recursive --profile=localstack - list files in a bucket

aws --endpoint-url=http://localhost:4566 s3 cp s3://tutorial-bucket/1.txt 1.txt --profile=localstack - doenload file from S3 bucket
```

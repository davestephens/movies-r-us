# Movies-r-us

Movies-r-us is a little setup that takes files from s3, and loads them into a database for querying via a REST API.

## Design Decisions

- Be event driven. As ingress/egress in cloud has a cost attached, only push/pull data when required, rather than poll for changes.
- Be fault tolerant within a region by running across three availability zones.
- Use cloud managed services wherever possible.
- Use one source for data types/structure.

## Contents

`./rest-api` - the API application.
`./helm` - Helm chart to deploy the containerised `rest-api`.
`./lambda` - Lambda function that receives events from s3 object creations, and sends a notification onto the `rest-api` running in k8s.
`./terraform` - Terraform of RDS, K8s, Lambda function, s3 bucket, s3 bucket event notification config.

## Dependencies

- Terraform
- Helm
- Kubectl
- Docker (to build `rest-api` image)
- AWS CLI
- AWS creds configured in `~/.aws/credentials`

## Setup

1. Deploy `terraform/eks`
2. Set up kubectl - `aws eks update-kubeconfig --region eu-west-1 --name movies-r-us-cluster`
2. Deploy terraform/application


## General Info

### Incoming Data Flow

- Files are dropped into an s3 bucket
- The s3 bucket fires an event at a Lambda function when puts occur (saving continual polling and £££).
- The Lambda function unpacks the event, and pushes the filename and bucket onto the `rest-api` application running on k8s.
- The `rest-api` application pulls the new file from s3, and attempts to bind the json to a known data structure, and saves it into a Postgres database if successful.

### s3

The bucket is configured to generate an event notification whenever an object is created. The notifications target a Lambda function. Data is assumed to be in the following format:

```
[
    {
        "title": "National Treasure",
        "year": 2004,
        "cast": [
            "Nicolas Cage",
            "Diane Kruger",
            "Justin Bartha",
            "Jon Voight",
            "Harvey Keitel",
            "Sean Bean",
            "Christopher Plummer"
        ],
        "genres": [
            "Adventure"
        ]
    }
]
```
Receiving single or multiple movies as a JSON array is supported.

### Lambda Function

When events are received from the s3 bucket, the function unpacks the event and sends a JSON payload in the following format to `MRU_API_URL`:

```
{
    "bucket": "movies-r-us",
    "key": "movies.json"
}
```

### Rest API

The `rest-api` receives notification about new files from the Lambda function. The files are collected and attempted to be unmarshalled into the expected data format. If successful, they're written to a Postgres RDS database.

The `rest-api` is configured with the following environment variables:

```
MRU_DB_HOST=<hostname>
MRU_DB_USER=<username>
MRU_DB_PASS=<password>
MRU_DB_NAME=<database_name>
AWS_REGION=<aws_region>
```

## Architecture

[![Movies-r-us Architecture](./movies-r-us-architecture.jpg)](./movies-r-us-architecture.jpg)

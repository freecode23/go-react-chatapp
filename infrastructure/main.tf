provider "aws" {
  region = "us-east-2"  
}


resource "aws_s3_bucket" "gochat-s3" {
  bucket = "gochat-bucket"
}

resource "aws_s3_bucket_ownership_controls" "gochat-s3" {
  bucket = aws_s3_bucket.gochat-s3.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "gochat-s3" {
  bucket = aws_s3_bucket.gochat-s3.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_acl" "gochat-s3" {
  depends_on = [
    aws_s3_bucket_ownership_controls.gochat-s3,
    aws_s3_bucket_public_access_block.gochat-s3,
  ]

  bucket = aws_s3_bucket.gochat-s3.id
  acl    = "public-read"
}

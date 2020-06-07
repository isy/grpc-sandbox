locals {
  project = "grpc-sandbox"
}


provider "aws" {
    regeion = "ap-northeast-1"
}

data "aws_availability_zones" "available" {
  state = "available"
}
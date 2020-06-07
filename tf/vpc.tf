locals {
  vpc_cidr_block = "10.0.0.0/16"
  subnet_count = "2"
}

# ============
# VPC
# ============

resource "aws_vpc" "vpc" {
  cidr_block = "${local.vpc_cidr_block}"
  enable_dns_hostnames = true
  enable_dns_support = true
  instance_tenancy = "default"

  tags = "${local.project}-vpc"
}

# ============
# Subnet
# ============
resource "aws_subnet" "subnet" {
    count = "${local.subnet_count}"
    vpc_id = "${aws_vpc.vpc.id}"
    cidr_block = "${cidrsubnet(local.vpc_cidr_block, 8, count.index)}"
    availability_zone = "${element(data.aws_availability_zones.available.names, count.index)}"
}

# ============
# Internet Gateway
# ============
resource "aws_internet_gateway" "igw" {
    vpc_id = "${aws_vpc.vpc.id}"
}

# ============
# Route Table
# ============
resource "aws_route_table" "route_table" {
    vpc_id = " ${aws_vpc.vpc.id}"
    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = "${aws_internet_gateway.igw.id}"
    }
}

resource "aws_route_table_association" "association" {
  count          = "${local.subnet_count}"
  subnet_id      = "${element(aws_subnet.subnet.*.id, count.index)}"
  route_table_id = "${aws_route_table.route_table.id}"
}

# ============
# Security Group
# ============
# https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html
resource "aws_security_group" "eks-master" {
    name = "eks-master-sec"
    description = "EKS master security group"
    vpc_id = "${aws_vpc.vpc.id}"

    ingress {
        from_port = 443
        to_port = 443
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

resource "aws_security_group" "eks-node" {
  name = "eks-node-sec"
  description = "EKS node security group"

  ingress {
      description = "Allow cluster master to access cluster node"
      from_port = 1025
      to_port = 65535
      protocol = "tcp"
      security_groups = ["${aws_security_group.eks-master.id}"]
  }

  ingress {
      description = "Allow cluster master to access cluster node"
      from_port = 443
      to_port = 443
      protocol = "tcp"
      security_groups = ["${aws_security_group.eks-master.id}"]
      self = false
  }

  ingress {
      description = "Allow inter pods communication"
      from_port = 0
      to_port = 0
      protocol = "-1"
      self = true
  }

  egress {
      from_port = 0
      to_port = 0
      protocol = "-1"
      cidr_blocks = ["0.0.0.0/0"]
  }
}

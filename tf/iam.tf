# ============
# EKS master role
# ============
resource "aws_iam_role" "eks-master" {
    name = "eks-master-role"

    assume_role_policy = <<EOF
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Action": "sts:AssumeRole",
                "Principal": {
                    "Service": "eks.amazonaws.com"
                },
                "Effect": "Allow",
            }
        ]
    }
    EOF
}

resource "aws_iam_role_policy_attachment" "eks_cluster" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role = "${aws_iam_role.eks-master.name}"
}

resource "aws_iam_role_policy_attachment" "eks-service" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
  role = "${aws_iam_role.eks-master.name}"
}

# ============
# EKS node role
# ============
resource "aws_iam_role" "eks-node" {
  name = "eks-node-role"

  assume_role_policy = <<EOF
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Action": "sts:AssumeRole",
                "Principal": {
                    "Service": "ec2.amazonaws.com"
                },
                "Effect": "Allow",
            }
        ]
    }
  EOF
}

resource "aws_iam_role_policy_attachment" "eks-worker-node" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
  role = "${aws_iam_role.eks-node.name}"
}

resource "aws_iam_role_policy_attachment" "eks-node-cni" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
  role       = "${aws_iam_role.eks-node-role.name}"
}

resource "aws_iam_role_policy_attachment" "ecr-read" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
  role       = "${aws_iam_role.eks-node-role.name}"
}

resource "aws_iam_instance_profile" "eks-node-profile" {
  name = "eks-node-profile"
  role = "${aws_iam_role.eks-node-role.name}"
}
locals {
  cluster_name = "sandbox-cluster"
  cluster_version = ""
  userdata = <<USERDATA
    #!/bin/bash
    set -o xtrace
    /etc/eks/bootstrap.sh --apiserver-endpoint "${aws_eks_cluster.cluster.endpoint}" --b64-cluster-ca "${aws_eks_cluster.cluster.certificate_authority.0.data}" "${aws_eks_cluster.cluster.name}"
    USERDATA
}


resource "aws_eks_cluster" "cluster" {
  name = "${local.cluster_name}"
  role_arn = "${aws_iam_role.eks-master.arn}"
  version = "1.16"

  vpc_config {
      security_group_ids = ["${aws_security_group.eks-master.id}"]
      subnet_ids = ["${aws_subnet.subnet.*.id}"]
  }

  depends_on = [
      "aws_iam_role_policy_attachment.eks-cluster",
      "aws_iam_role_policy_attachment.eks-service",
  ]
}

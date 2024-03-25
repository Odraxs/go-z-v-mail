provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "ec2_example" {
  ami                    = "ami-080e1f13689e07408"
  instance_type          = "t2.micro"
  key_name               = "aws_key"
  vpc_security_group_ids = [aws_security_group.main.id]
  user_data              = file("${path.module}/userdata.tpl")

  root_block_device {
    volume_size = 8
  }

  provisioner "remote-exec" {
    inline = [
      "touch hello.txt",
      "echo helloworld remote provisioner >> hello.txt",
    ]
  }

  connection {
    type        = "ssh"
    host        = self.public_ip
    user        = "ubuntu"
    private_key = file("/home/david/.ssh/aws_key")
    timeout     = "4m"
  }
}

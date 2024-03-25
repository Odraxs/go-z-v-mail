## Terraform

This directory contains the terraform code that allows to deploy the project to an aws EC2 instance.

### Instructions

- Get into the terraform directory
    ```bash
        cd terraform
    ```

- Set your aws IAM credentials of a user with programmatic access.
    ```bash
        cp credentials.example credentials.sh
        ##Set your credentials in the created file
        chmod +x credentials.sh
        . credentials.sh
    ```

- Create your ssh key-pair
    ```bash
        ssh-keygen -t rsa -b 2048 
        ...
        Generating public/private RSA key pair.
        Enter file in which to save the key (/home/path_to_ssh/.ssh/id_rsa): /home/path/to/save/keys/aws_key 
        ...
    ```

> [!NOTE]
> Change your own path in the line 27 of the `main.tf` and line 3 of the `security.tf` files.

- Start terraform
    ```bash
        terraform init
    ```

- Look the terraform plan and deploy
    ```bash
        terraform plan
        ...
        terraform apply
        ...
    ```

- Connect to the instance and check if the images are running
    ```bash
        ssh -i "your/private/key/path" ubuntu@your_instance_ip
        sudo docker ps -a
        ...
    ```
    If everything went right all the images should be shown by the last command.

### Additional notes

Everything was done having in mind a free plan deployment, so the images of the project are in a docker repository, IMO the best way should be by creating the images directly in the EC2 instance, because with that way we can modify the `.evn` file of the web application so it sends the requests to the correct endpoint, but with a `t2.micro` is was impossible to build the images on the EC2 instance.

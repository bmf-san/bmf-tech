---
title: Building a VPS with Terraform and Ansible
slug: terraform-ansible-vps-setup
date: 2020-12-20T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
  - VPS
  - OpenStack
  - Terraform
translation_key: terraform-ansible-vps-setup
---



# Overview
This article is the 20th entry in the [Makuake Advent Calendar 2020](https://adventar.org/calendars/5986).

I combined Terraform and Ansible to build a server on a VPS, and I would like to summarize the procedure here.

# Motivation
I wanted to set up the infrastructure environment for an application I'm developing as a hobby using IaC, so I decided to try Terraform as part of my learning process.

# Environment
- Terraform v0.14.0
- Ansible 2.9.10
- ConoHa VPS
- macOS Catalina 10.15.5

I will execute Terraform locally to create and destroy servers on ConoHa VPS.

# OpenStack
OpenStack is an OSS for building IaaS environments.

ConoHa's VPS adopts OpenStack and provides an OpenStack-compliant API.

cf. [www.slideshare.net - "ConoHa" VPS-KVM; OpenStack Grizzly based service](https://www.slideshare.net/chroum/cono-ha-openstackgrizzlyvpsrelease)

By using the OpenStack provider in Terraform, you can build servers on ConoHa VPS.

cf. [conoha.jp - API](https://www.conoha.jp/vps/function/api/)

This time, I will use the OpenStack provider with Terraform, but there is also an OpenStack Ansible module in Ansible, so I think the same can be achieved with Ansible alone. I haven't tried it, though...

cf. [docs.ansible.com - OpenStack Ansible Module](https://docs.ansible.com/ansible/2.9_ja/dev_guide/platforms/openstack_guidelines.html)

# Source Code
The source code created this time is available at [github.com - bmf-san/terraform-ansible-openstack-boilerplate](https://github.com/bmf-san/terraform-ansible-openstack-boilerplate).

# Implementation Policy
I will build the server with Terraform and perform the initial server setup with Ansible.

When using both Terraform and Ansible, you might wonder whether to call Ansible from Terraform or Terraform from Ansible, but the article below suggests that either is fine, and there is no right or wrong answer.

cf. [www.redhat.com - HASHICORP TERRAFORM AND RED HAT ANSIBLE AUTOMATION](https://www.redhat.com/cms/managed-files/pa-terraform-and-ansible-overview-f14774wg-201811-en.pdf)

Terraform is strong in infrastructure resource configuration management, and Ansible is strong in server configuration management, so I tried to structure it so that each can handle their respective areas of expertise, with Ansible being executed within Terraform.

Depending on how you want to manage the code, the opposite pattern might be better in some cases.

# Preparation
- Install Terraform
  - Installed the latest version using brew and tfenv.
- Install Ansible
  - Installed using brew.
- Prepare two sets of public and private key pairs
  - For root user and working user.
- Create an API user in the ConoHa management screen
  - Refer to [support.conoha.jp - Add API User](https://support.conoha.jp/v/addapiuser/) for the procedure.
  - The following information is required, so check it in advance.
    - API Information
      - Tenant ID
      - Tenant Name
    - API User
      - Username
      - Password

# Construction Procedure
The general flow is as follows.

Obtain API token for using ConoHa API

↓

Decide on the image and VM plan to use

↓

Write Terraform code

↓

Write Ansible code

## Obtain API Token for Using ConoHa API
First, obtain an API token to use the ConoHa API.
The API endpoint varies for each user, so refer to the endpoint list in the API information on the ConoHa control panel as needed.

cf. [conoha.jp - Token Issuance](https://www.conoha.jp/docs/identity-post_tokens.php)

```sh
curl -X POST \
-H "Accept: application/json" \
-d '{"auth":{"passwordCredentials":{"username":"USER_NAME","password":"PASSWORD"},"tenantId":"TENANT_ID"}}' \
https://identity.tyo2.conoha.io/v2.0/tokens \
| jq ".access.token.id"
```

## Decide on the Image and VM Plan to Use
Use the obtained API token to get the information and decide on the image and VM plan to use.

### Get List of Available Images
Retrieve a list of available images.

cf. [conoha.jp - Get Image List](https://www.conoha.jp/docs/image-get_images_list.php)

```sh
curl -X GET \
-H 'Content-Type: application/json' \
-H "Accept: application/json" \
-H "X-Auth-Token: API_TOKEN" \
https://compute.tyo2.conoha.io/v2/TENANT_ID/images \
| jq ".images | sort_by(.name) | map(.name)"
```

This time, I used `vmi-ubuntu-20.04-amd64-30gb`.

### Get List of VM Plans
Retrieve a list of available VM plans.

cf. [conoha.jp - Get VM Plan List](https://www.conoha.jp/docs/compute-get_flavors_list.php)

```sh
curl -X GET \
-H 'Content-Type: application/json' \
-H "Accept: application/json" \
-H "X-Auth-Token: API_TOKEN" \
https://compute.tyo2.conoha.io/v2/TENANT_ID/flavors \
| jq ".flavors | sort_by(.name) | map(.name)"
```

This time, I chose `g-1gb`.

It seems that plans below 1gb result in a construction error due to insufficient disk size. (I tried with `g-512mb`, but it didn't work.)

## Terraform
With the necessary information gathered, let's write the code.

This time, I used the following directory structure.

```
.
├── ansible.cfg
├── main.tf
├── playbooks
├── templates
│   └── playbooks
│       ├── hosts.tpl
│       └── setup.tpl
├── terraform.tfvars
└── variable.tf

3 directories, 12 files
```

Since there isn't much to do this time, I didn't split the tf files into smaller parts.

Regarding the management of tfstate files, I think it's better to manage them with external storage using a backend, but since this is executed locally, they are only included in `.gitignore`. (Even locally, it's something that should be done properly...)

As mentioned later, `playbooks` will contain the `hosts` file and `setup` file (yml) generated by Terraform from `templates`.

Since we are building an instance from scratch, we need to pick up the IP address value during the construction process and pass the value from Terraform to Ansible, so I think it's meaningful to template the `hosts` file. However, for the `setup` file (yml), it might be better to separate tasks and variable definitions and template the file for variable definitions. This time, I didn't split them for simplicity.

If you rely too much on Terraform, it might become cumbersome if you later want to extract Ansible, so I would like to know various examples in this regard.

### main.tf
The contents of `main.tf` are as follows.

```
terraform {
  required_version = ">= 0.14"
  required_providers {
    openstack = {
      source = "terraform-provider-openstack/openstack"
      version = "1.33.0"
    }
  }
}

provider "openstack" {
  user_name   = (var.user_name)
  password    = (var.password)
  tenant_name = (var.tenant_name)
  auth_url = (var.auth_url)
}

resource "openstack_compute_keypair_v2" "example_keypair" {
  name       = (var.keypair_name)
  public_key = file(var.path_to_public_key_for_root)
}

resource "openstack_compute_instance_v2" "example_instance" {
  name        = (var.instance_name)
  image_name  = (var.image_name)
  flavor_name = (var.flavor_name)
  key_pair    = (var.keypair_name)

  security_groups = [
    "gncs-ipv4-ssh",
    "gncs-ipv4-web",
  ]

  metadata = {
    instance_name_tag = (var.instance_name_tag)
  }
}

data "template_file" "hosts" {
  template = file("./templates/playbooks/hosts.tpl")

  vars = {
    host = (var.host)
    ip = (openstack_compute_instance_v2.example_instance.access_ip_v4)
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}

resource "local_file" "save_hosts" {
  content  = (data.template_file.hosts.rendered)
  filename = "./playbooks/hosts"

  depends_on = [openstack_compute_instance_v2.example_instance]
}

data "template_file" "setup" {
  template = file("./templates/playbooks/setup.tpl")

  vars = {
    host = (var.host)
    new_user_name = (var.new_user_name)
    new_user_password = (var.new_user_password)
    shell = (var.shell)
    new_user_public_key = file(var.path_to_public_key)
    port = (var.port)
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}

resource "local_file" "save_setup" {
  content  = (data.template_file.setup.rendered)
  filename = "./playbooks/setup.yml"

  depends_on = [openstack_compute_instance_v2.example_instance]
}

resource "null_resource" "example_provisoner" {
  provisioner "local-exec" {
    command = "ansible-playbook ./playbooks/setup.yml -i ./playbooks/hosts --private-key=${var.path_to_private_key_for_root}"
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}
```

#### provider
Since there is an official OpenStack provider, I am using it.

```
provider "openstack" {
  user_name   = (var.user_name)
  password    = (var.password)
  tenant_name = (var.tenant_name)
  auth_url = (var.auth_url)
}
```

`user_name` and `password` are the API user information created in ConoHa.
`tenant_name` is literally the tenant name.
`auth_url` is a bit unclear, but here it is the endpoint of ConoHa's Identity API (e.g., `https://identity.tyo2.conoha.io/v2.0`).

#### openstack_compute_keypair_v2
This is the setup for the public and private key pair used by the root user when building the instance.

cf. [registry.terraform.io - openstack_compute_keypair_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_keypair_v2)

```
resource "openstack_compute_keypair_v2" "example_keypair" {
  name       = (var.keypair_name)
  public_key = file(var.path_to_public_key_for_root)
}
```

If you do not specify a public key, a public and private key pair will be automatically generated.

Since key information is output to the tfstate file, it is necessary to manage the tfstate file appropriately when executing in a real environment.

Although public key authentication is assumed, it seems that there is a way to enable password authentication as well.

cf. [noaboutsnote.hatenablog.com - 【Openstack】Enable Password Login to Instance OS](http://noaboutsnote.hatenablog.com/entry/openstack_instance_password)

#### openstack_compute_instance_v2
This is the setup for building an instance, including the instance image, VM plan, and network configuration.

cf. [registry.terraform.io - openstack_compute_instance_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_instance_v2)

```
resource "openstack_compute_instance_v2" "example_instance" {
  name        = (var.instance_name)
  image_name  = (var.image_name)
  flavor_name = (var.flavor_name)
  key_pair    = (var.keypair_name)

  security_groups = [
    "gncs-ipv4-ssh",
    "gncs-ipv4-web",
  ]

  metadata = {
    instance_name_tag = (var.instance_name_tag)
  }
}
```

`instance_name` is an arbitrary name, and `image_name` is literally the image name.
`flavor_name` might not be intuitive at first glance, but here it refers to the VM plan name.

The `instance_name_tag` part will be the name tag displayed on the ConoHa control panel.

Although not used this time, you can also use cloud-init by specifying user_data.

ex.
```
user_data = data.template_file.user_data.rendered

data "template_file" "user_data" {
  template = file("user_data.sh")
}
```

#### null_resource
null_resource is a resource that performs provisioning triggered by other resources. The trigger is specified with depends_on.

I want to provision the constructed instance with Ansible, so I set the completion of instance construction (or more accurately, the completion of Terraform execution, as Terraform execution completion does not necessarily mean instance construction completion, so I prepared a process to wait for instance construction with Ansible, as mentioned later) as the trigger.

```
resource "null_resource" "example_provisoner" {
  provisioner "local-exec" {
    command = "ansible-playbook ./playbooks/setup.yml -i ./playbooks/hosts --private-key=${var.path_to_private_key_for_root}"
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}
```

Since this is executed locally, I am using `local-exec`.

While searching for a suitable resource, I found [github.com - jonmorehouse/terraform-provisioner-ansible](https://github.com/jonmorehouse/terraform-provisioner-ansible), but it seems to be unmaintained now.

## Ansible
This time, I prepared templates under `templates/playbooks`, generated actual files based on the templates during Terraform execution, and executed Ansible using the generated files.

The provisioning content includes creating an execution user and adjusting ssh settings.

I got stuck trying to provision without waiting for instance connectivity... (I used `wait_for_connection` to address this.)

# Execution
It's all done with terraform commands.

```
terraform init
terraform plan
terraform apply
terraform show
ssh username@ipaddress -i path_to_private_key
terraform destroy
```

# Impressions
It was my first time using Terraform, so it was a good learning experience.

I found OpenStack to be an interesting technology, so I want to create more opportunities to explore it further.

# References
- [blog.mosuke.tech](https://blog.mosuke.tech/entry/2019/03/26/terraform-integrated-with-ansible/)
- [www.redhat.com - HASHICORP TERRAFORM AND RED HAT ANSIBLE AUTOMATION](https://www.redhat.com/cms/managed-files/pa-terraform-and-ansible-overview-f14774wg-201811-en.pdf)
- [note.com - Ansible and Terraform](https://note.com/santak/n/n22dbb35f6c01)
- [qiita.com - Trying to Set Up a ConoHa Instance with Terraform](https://qiita.com/kaminchu/items/d0776c381213d54a3a69)
- [qiita.com - Executing Ansible Playbook from Terraform](https://qiita.com/hayaosato/items/ee0d6eabb7b3d0a22136)
- [conoha.jp - ConoHa API Documentation](https://www.conoha.jp/docs/)
- [blog.adachin.me - [OpenStack] Setting Up a ConoHa VPS Instance with Terraform!](https://blog.adachin.me/archives/10378)
- [gist.github.com - san-tak/tf-bookmark.md](https://gist.github.com/san-tak/9ef4d15eafb4c8a42af33ffa04464739)
- [github.com - maki0922/terraform_in_conoha](https://github.com/maki0922/terraform_in_conoha)
- [github.com - hobbby-kube](https://github.com/hobby-kube/provisioning)
- [github.com - r0b2g1t/terraform-openstack-ansible](https://github.com/r0b2g1t/terraform-openstack-ansible)
- [github.com - dan4ex/Terraform](https://github.com/dan4ex/Terraform)
- [registry.terraform.io - openstack_compute_keypair_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_keypair_v2)
- [registry.terraform.io - openstack_compute_instance_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_instance_v2)

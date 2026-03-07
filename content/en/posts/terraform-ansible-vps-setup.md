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
This article is the 20th entry of the [Makuake Advent Calendar 2020](https://adventar.org/calendars/5986).

I combined Terraform and Ansible to build a server on a VPS, and I would like to summarize the steps.

# Motivation
I wanted to set up the infrastructure for an application I am developing as a hobby using IaC, so I decided to try Terraform as a learning experience.

# Environment
- Terraform v0.14.0
- Ansible 2.9.10
- ConoHa VPS
- macOS Catalina 10.15.5

I will run Terraform locally to create and destroy servers on ConoHa VPS.

# OpenStack
OpenStack is an OSS for building IaaS environments.

ConoHa's VPS uses OpenStack and provides an OpenStack-compliant API.

cf. [www.slideshare.net - "ConoHa" VPS-KVM; OpenStack Grizzly based service](https://www.slideshare.net/chroum/cono-ha-openstackgrizzlyvpsrelease)

By using the OpenStack provider in Terraform, you can build servers on ConoHa VPS.

cf. [conoha.jp - API](https://www.conoha.jp/vps/function/api/)

This time, I will use the OpenStack provider with Terraform, but Ansible also has an OpenStack Ansible module, so I believe similar functionality can be achieved with just Ansible. I haven't tried it though...

cf. [docs.ansible.com - OpenStack Ansible Module](https://docs.ansible.com/ansible/2.9_ja/dev_guide/platforms/openstack_guidelines.html)

# Source Code
The source code created this time is available at [github.com - bmf-san/terraform-ansible-openstack-boilerplate](https://github.com/bmf-san/terraform-ansible-openstack-boilerplate).

# Implementation Policy
I will build the server with Terraform and perform the initial setup with Ansible.

When using both Terraform and Ansible, it can be confusing whether to call Ansible from Terraform or vice versa, but the article below suggests that either approach is fine, and there is no right or wrong answer.

cf. [www.redhat.com - HASHICORP TERRAFORM AND RED HAT ANSIBLE AUTOMATION](https://www.redhat.com/cms/managed-files/pa-terraform-and-ansible-overview-f14774wg-201811-en.pdf)

Terraform is strong in managing infrastructure resource configurations, while Ansible excels in configuration management within the server. I aimed to execute Ansible within Terraform while being aware of their respective roles. Depending on how you want to manage the code, the opposite pattern might also be preferable.

# Preparation
- Install Terraform
  - I installed tfenv via brew to prepare the latest version.
- Install Ansible
  - I installed it via brew.
- Prepare two sets of public and private key pairs
  - One for the root user and one for the working user.
- Create an API user in the ConoHa management panel
  - Refer to [support.conoha.jp - Adding an API User](https://support.conoha.jp/v/addapiuser/) for the procedure.
  - The following information is required, so please check it in advance:
    - API Information
      - Tenant ID
      - Tenant Name
    - API User
      - Username
      - Password

# Construction Steps
The general flow is as follows:

Obtain API token for ConoHa API usage

↓

Decide on the desired image and VM plan

↓

Write Terraform code

↓

Write Ansible code

## Obtain API Token for ConoHa API Usage
First, obtain the API token to use the ConoHa API. The API endpoint differs for each user, so please refer to the list of endpoints in the API information of the ConoHa control panel as needed.

cf. [conoha.jp - Token Issuance](https://www.conoha.jp/docs/identity-post_tokens.php)

```sh
curl -X POST \
-H "Accept: application/json" \
-d '{"auth":{"passwordCredentials":{"username":"USER_NAME","password":"PASSWORD"},"tenantId":"TENANT_ID"}}}' \
https://identity.tyo2.conoha.io/v2.0/tokens \
| jq ".access.token.id"
```

## Decide on the Desired Image and VM Plan
Using the obtained API token, retrieve the necessary information and decide on the desired image and VM plan.

### Retrieve List of Available Images
Retrieve the list of available images.

cf. [conoha.jp - Retrieve Image List](https://www.conoha.jp/docs/image-get_images_list.php)

```sh
curl -X GET \
-H 'Content-Type: application/json' \
-H "Accept: application/json" \
-H "X-Auth-Token: API_TOKEN" \
https://compute.tyo2.conoha.io/v2/TENANT_ID/images \
| jq ".images | sort_by(.name) | map(.name)"
```

This time, I used `vmi-ubuntu-20.04-amd64-30gb`.

### Retrieve List of VM Plans
Retrieve the list of available VM plans.

cf. [conoha.jp - Retrieve VM Plan List](https://www.conoha.jp/docs/compute-get_flavors_list.php)

```sh
curl -X GET \
-H 'Content-Type: application/json' \
-H "Accept: application/json" \
-H "X-Auth-Token: API_TOKEN" \
https://compute.tyo2.conoha.io/v2/TENANT_ID/flavors \
| jq ".flavors | sort_by(.name) | map(.name)"
```

This time, I selected `g-1gb`.

It seems that plans below 1GB result in a disk size error during construction (I tried `g-512mb`, but it didn't work).

## Terraform
Now that I have the necessary information, I will write the code.

This time, I organized the directory structure as follows:

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

Since there are not many tasks this time, I haven't divided the tf files in detail.

Regarding the management of tfstate files, it would be better to manage them with external storage using a backend, but since I am executing locally, they are only included in the `.gitignore` target. (Even though it's local, it's something I want to handle properly...)

As mentioned later, the `playbooks` directory will contain the `hosts` file and `setup` file (yml) generated by Terraform from `templates`.

Since I will be building an instance from scratch, I need to capture the IP address value during the construction process and pass it from Terraform to Ansible, so I think it makes sense to template the `hosts` file. However, for the `setup` file (yml), I think it would be better to separate the task and variable definitions and template the file that defines the variables. This time, I skipped that and did not divide it.

If I lean too much towards Terraform, I feel it might become cumbersome later if I want to extract Ansible, so I would like to know various examples in this regard.

### main.tf
The contents of `main.tf` look like this.

```
tf
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
Since there is an official OpenStack provider, I am using that.

```
provider "openstack" {
  user_name   = (var.user_name)
  password    = (var.password)
  tenant_name = (var.tenant_name)
  auth_url = (var.auth_url)
}
```

`user_name` and `password` are the information of the API user created in ConoHa.
`tenant_name` is literally the tenant name.
`auth_url` is somewhat confusing, but here it refers to the endpoint of ConoHa's Identity API (e.g., `https://identity.tyo2.conoha.io/v2.0`).

#### openstack_compute_keypair_v2
This sets up the public and private key pair used by the root user during instance construction.

cf. [registry.terraform.io - openstack_compute_keypair_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_keypair_v2)

```
resource "openstack_compute_keypair_v2" "example_keypair" {
  name       = (var.keypair_name)
  public_key = file(var.path_to_public_key_for_root)
}
```

If a public key is not specified, a public and private key pair will be automatically generated.

The key information will be output to the tfstate file, so if you run it in a real environment, you need to manage the tfstate file appropriately.

Public key authentication is assumed, but there seems to be a way to enable password authentication as well.

cf. [noaboutsnote.hatenablog.com - [Openstack] Enabling Password Login to Instance OS](http://noaboutsnote.hatenablog.com/entry/openstack_instance_password)

#### openstack_compute_instance_v2
This sets up the instance's image, VM plan, network configuration, etc., to build the instance.

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

`instance_name` is an arbitrary name, `image_name` is literally the image name.
`flavor_name` might be hard to guess at first glance, but here it refers to the VM plan name.

The `instance_name_tag` part is the name tag displayed in the ConoHa control panel.

Although I am not using it this time, you can also use cloud-init by specifying `user_data`.

ex.
```
user_data = data.template_file.user_data.rendered

data "template_file" "user_data" {
  template = file("user_data.sh")
}
```

#### null_resource
The null_resource is a resource that performs provisioning triggered by other resources. The trigger is specified with depends_on.

Since I want to provision the instance built with Ansible, I am using the completion of the instance construction (which might be more accurately described as the completion of the Terraform execution) as the trigger. (Even though the Terraform execution is complete, the instance construction may not be finished, so I will prepare a process to wait for the instance construction with Ansible later.)

```
resource "null_resource" "example_provisoner" {
  provisioner "local-exec" {
    command = "ansible-playbook ./playbooks/setup.yml -i ./playbooks/hosts --private-key=${var.path_to_private_key_for_root}"
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}
```

This time, I am using `local-exec` since I will execute it locally.

I searched for a suitable resource and found [github.com - jonmorehouse/terraform-provisioner-ansible](https://github.com/jonmorehouse/terraform-provisioner-ansible), but it seems to be no longer maintained.

## Ansible
This time, I prepared templates under `templates/playbooks`, and during the execution of Terraform, I generated actual files based on the templates and used the generated files to execute Ansible.

The provisioning content includes creating the execution user and adjusting SSH settings.

I got stuck trying to provision without waiting for the instance to be reachable... (I used `wait_for_connection` to handle this.)

# Execution
It can be completed with just Terraform commands.

```
terraform init
terraform plan
terraform apply
terraform show
ssh username@ipaddress -i path_to_private_key
tf
terraform destroy
```

# Impressions
Since this was my first experience with Terraform, it was a great learning opportunity.

I found OpenStack to be an interesting technology, so I would like to create more opportunities to explore it further.

# References
- [blog.mosuke.tech](https://blog.mosuke.tech/entry/2019/03/26/terraform-integrated-with-ansible/)
- [www.redhat.com - HASHICORP TERRAFORM AND RED HAT ANSIBLE AUTOMATION](https://www.redhat.com/cms/managed-files/pa-terraform-and-ansible-overview-f14774wg-201811-en.pdf)
- [note.com - Ansible and Terraform](https://note.com/santak/n/n22dbb35f6c01)
- [qiita.com - Creating an Instance on ConoHa with Terraform](https://qiita.com/kaminchu/items/d0776c381213d54a3a69)
- [qiita.com - Executing Ansible Playbook from Terraform](https://qiita.com/hayaosato/items/ee0d6eabb7b3d0a22136)
- [conoha.jp - ConoHa API Documentation](https://www.conoha.jp/docs/)
- [blog.adachin.me - [OpenStack] Trying to Launch a ConoHa VPS Instance with Terraform!](https://blog.adachin.me/archives/10378)
- [gist.github.com - san-tak/tf-bookmark.md](https://gist.github.com/san-tak/9ef4d15eafb4c8a42af33ffa04464739)
- [github.com - maki0922/terraform_in_conoha](https://github.com/maki0922/terraform_in_conoha)
- [github.com - hobbby-kube](https://github.com/hobby-kube/provisioning)
- [github.com - r0b2g1t/terraform-openstack-ansible](https://github.com/r0b2g1t/terraform-openstack-ansible)
- [github.com - dan4ex/Terraform](https://github.com/dan4ex/Terraform)
- [registry.terraform.io - openstack_compute_keypair_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_keypair_v2)
- [registry.terraform.io - openstack_compute_instance_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_instance_v2)
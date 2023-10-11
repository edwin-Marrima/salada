## Problem
Cloud adoption is growing rapidly year after year, the value proposition has been increasingly approved by companies, the benefits behind it have attracted more and more users. While cloud adoption offers numerous benefits, it also comes with inherent risks that should not be underestimated. The high level of automation in cloud environments can be a double-edged sword, as a single misstep can lead to severe consequences. A simple mistake, such as a few unintended clicks, may result in irreversible damage, like unintentional deletion of critical databases in production, skyrocketing costs, or even data leakage. It only takes 1 second for everything to go wrong.

For this reason, companies have guidelines that help parameterize the use of resources provided by cloud providers. The problem is that these guides are found in long documentation, so whenever a Software Engineer needs to provision resources in the cloud (in "you build it, you run it" mentality companies) he needs to check the documentation to verify whether or not it is compliant. For security reasons, the verification done by the Software Engineer is not considered enough, usually many companies have a production gatekeeping process, in which Production, Platform or DevOps Engineer do a doublecheck to verify if the settings are correct or not. This is where Salada comes to play.

## Salada

We all know that the most revolutionary things in Software Engineering are automation tools, and that's what Salada is. Salada is an automation tool that automatically enforces guidelines related to provisioning resources in the cloud through infrastructure as code, specifically Terraform. You only need to provide a file with the policies, and Salada will take care of the rest for you. The platform, production, DevOps, or any other engineer who controls the environments in any cloud provider only needs to create the policy file once. During the CI/CD process, Salada will check the policies on their behalf, providing greater safety and speed.


### Working Principle

In some organizations, data layer resources must not be in specific subnets; in others, including the environment name in resource names is mandatory. Some organizations require all non-global resources to be provisioned in more than one region, while others mandate enabling encryption at rest for data-storing resources using specific keys. Moreover, specific organizations only allow updates to database management systems on Sundays from 4 am to 11 am. The examples presented show how critical organizations are in relation to the cloud, most organization policies go far beyond what was listed, and Salada allows you to enforce all scenarios.

Many organizations implement GitOps using Terraform to automate infrastructure management. During the CI/CD process, there is a segmentation into Terraform plan and apply stages for obvious reasons, to verify what will be changed. Salada resides between the plan and apply stages, taking the Terraform plan file and Terraform module as input to verify whether the changes are compliant with the policies, as shown in the code snippet below:

```
salada policy=/policies/production.yaml plan-file=/tf-plan.json
```

#### Policy File

It is a yaml file that contains settings that enforce organizations policies regarding cloud resource provisioning. The file is segmented into tree sections, namely: `provider`, `variables` and `resources`.

* Providers: contains the list of allowed and disallowed terraform providers, versions, and rules regarding provider configuration parameters.
* Variables: contains variables configurations.
* Resources: contains the list of allowed and/or disallowed resources, conditionals for deleting, updating and creating, and conditionals that enforce the configuration of resource attributes. 

```
providers:
  list:
    - name: "aws"
      source: "hashicorp/aws"
      version: greaterThan("5.9")
      allowed: true
      configuration_parameters:
        - name: region
          value: oneOf(["af-south-1","us-west-1"])
        - name: access_key:
          allowed: false
        - name: secret_key:
          allowed: false  
    - source: "hashicorp/azurerm"
      version: "3.5.5"
        
variables:
  description:
    value: `match_regex("^*")` 
  name:
    value: `match_regex("^*")`
  

resources:
  - type: aws_s3_bucket
    allowed:
      when: always  
    change_actions:
      update:
        time: "* * * * *"
    attributes:
      - name: bucket_prefix
        value: this.object_lock_enabled == true ? endsWith("-${ENVIRONMENT}-"):endsWith("abc")
      - name: force_destroy:
        value: false
  - type: aws_s3_bucket_public_access_block
    allowed:
      when: 
        - resource: aws_s3_bucket
          condition: exists
    attributes:
      - name: block_public_acls:
        value: false  
  - type: aws_cloudwatch_log_group
    allowed:
      when: always
    attributes:
      - name: retention_in_days:
        value: greatherThan(29)
      - name: tags_all.Owners.0.Gender:
        value: M
      - name: tags_all.Owner:
        value: oneOf(["team-a","team-b"])

```
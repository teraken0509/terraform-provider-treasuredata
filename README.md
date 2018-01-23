Terraform provider for Treasure Data
==================================

A [Terraform](https://www.terraform.io/) plugin that provides resources for [Treasure Data](https://www.treasuredata.com/).

[![Build Status](https://travis-ci.org/kterada0509/terraform-provider-treasuredata.svg?branch=master)](https://travis-ci.org/kterada0509/terraform-provider-treasuredata)


## Install
---
* Download the latest release for your platform.
* Rename the executable to `terraform-provider-treasuredata`


## Provider Configuration
---
### Example
```
provider "treasuredata" {
  api_key = "xxx"
}
```

or

```
provider "treasuredata" {}
```

### Reference

* `api_key` - (Optional)API Key of the Treasure Data. If empty, provider try use `TD_API_KEY` envirenment variable.


## Resources
---

### `treasuredata_database`
Configure a Database.

#### Example

```
resource "treasuredata_database" "foobar" {
    name = "terraform_for_treasuredata_test_foobar"
}
```

### `treasuredata_schedule`
Configure a Schedule(Query).

#### Example

```
```


## Build
---
```
$ make build
```


## Test
---
```
$ TD_API_KEY=xxxx make test
```

## Licence
---
Mozilla Public License, version 2.0


## Author
---
[Kentaro Terada](https://github.com/kterada0509)
provider "bizfly" {
	host = "https://manage.bizflycloud.vn"
	username = "/"
	password = "/"
}

resource "bizfly_instance" "test" {
  name = "test"
  flavorname = "2c_2g"
  sshkey=  "test"
  password = "test"
  rootdisk  {
    size = 40
    type = "HDD"
    }
  datadisk  {
    size = 40
    type = "HDD"
  }
  type = "premium"
  availabilityzone = "HN1"
  os {
    id = "cbf5f34b-751b-42a5-830f-6b2324f61d5a"
    type = "image"
  }
}

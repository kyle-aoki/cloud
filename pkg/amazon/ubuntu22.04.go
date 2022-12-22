package amazon

// ubuntu 22.04
var Ubuntu2204AmiRegionMap = map[string]string{
	"us-gov-west-1":  "ami-04e77113f128920b2",
	"us-gov-east-1":  "ami-08fcbec75b43330b7",
	"us-west-2":      "ami-0ee8244746ec5d6d4",
	"us-west-1":      "ami-0dc5e9ff792ec08e3",
	"us-east-2":      "ami-0aeb7c931a5a61206",
	"us-east-1":      "ami-09d56f8956ab235b3",
	"sa-east-1":      "ami-0deebba34ef22f5a9",
	"me-south-1":     "ami-0ca63b0f7ddb12615",
	"eu-west-3":      "ami-0042da0ea9ad6dd83",
	"eu-west-2":      "ami-0a244485e2e4ffd03",
	"eu-west-1":      "ami-00c90dbdc12232b58",
	"eu-south-1":     "ami-069aaf99166131204",
	"eu-north-1":     "ami-01ded35841bc93d7f",
	"eu-central-1":   "ami-015c25ad8763b2f11",
	"ca-central-1":   "ami-0fb99f22ad0184043",
	"ap-southeast-3": "ami-0c1460efd8855de7c",
	"ap-southeast-2": "ami-0b21dcff37a8cd8a4",
	"ap-southeast-1": "ami-04d9e855d716f9c99",
	"ap-south-1":     "ami-0756a1c858554433e",
	"ap-northeast-3": "ami-0d6dbe860474011f3",
	"ap-northeast-2": "ami-063454de5fe8eba79",
	"ap-northeast-1": "ami-081ce1b631be2b337",
	"ap-east-1":      "ami-0f2ea204cd818ce8e",
	"af-south-1":     "ami-0df5b771d6e8bfdf9",
}

// gets correct ubuntu 22.04 ami based on user's aws config region
func Ubuntu2204Ami() string {
	if val, ok := Ubuntu2204AmiRegionMap[userRegion]; ok {
		return val
	}
	panic(`did not find an Ubuntu 22.04 ami in your region.
	your aws config is most likely missing or corrupted`)
}

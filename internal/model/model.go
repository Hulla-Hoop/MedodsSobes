package model

type Session struct {
	BcryptTocken      string `bson:"bcryptTocken"`
	TimeCreatedTocken string `bson:"timeCreatedTocken"`
	Guid              string `bson:"guid"`
	ExpireTime        int64  `bson:"expiretime"`
}

var Users = map[string]string{
	"3825c945-8843-4b7d-995e-30b16c173c65": "user1",
	"019ed7ca-8286-40b8-ac80-1950c92dccfd": "user2",
}

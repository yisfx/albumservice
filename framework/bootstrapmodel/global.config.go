package bootstrapmodel

type GlobalConf struct {
	AdminPwd  map[string]string `json:"AdminPwd"`
	AlbumPath string            `json:"AlbumPath"`
	SHAKEYOrg string            `json:"SHAKEYOrg"`
	SHAIVOrg  string            `json:"SHAIVOrg"`
	SHAKey    string            `json:"SHAKey"`
	SHAIV     string            `json:"SHAIV"`
	Redis     RedisConf         `json:"RedisConf"`
	WordFile  string            `json:"WordFile"`
}

type RedisConf struct {
	Port int    `json:"Port"`
	Pwd  string `json:"Pwd"`
}

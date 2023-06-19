package model


type Word struct{
	Chinese  string `json:"Chinese"`
	Japanese string `json:"Japanese"`
	Romanic string 	`json:"Romanic`
}

type Chapter struct{
	Title string `json:"Title"`
	Section []*Word `json:"Section"`
}


type WordRecord struct{
	WordRecord []*Chapter `json:"WordRecord"`
}
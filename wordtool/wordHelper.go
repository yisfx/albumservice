package wordtool

import (
	"albumservice/framework/fileTool"
	"albumservice/framework/utils"
	"albumservice/model"
	"encoding/json"
	"sort"
)



func GetWord(path string) []*model.Chapter{
	if !fileTool.FileExists(path) {
		return []*model.Chapter{}
	}
	content := fileTool.GetFileContentByName(path)

	result := utils.StringTool.DeSerializerFromJson(content,&model.WordRecord{}).(*model.WordRecord)
	res:= []*model.Chapter{}

	for _,v := range result.WordRecord{
		res=append(res,&model.Chapter{Title:v.Title,Section:v.Section})
		
	}
	return res
}
func GetSection(path string) []string{
	Chapter:= GetWord(path)
	res:=[]string{}

	for _,v:=range Chapter {
			res=append(res,v.Title)
	}
	return res
}

func AddWord(chapter *model.Chapter,path string){

	result := &model.WordRecord{}

	if fileTool.FileExists(path) {
		content := fileTool.GetFileContentByName(path)
		result = utils.StringTool.DeSerializerFromJson(content,&model.WordRecord{}).(*model.WordRecord)
	}

	section:=GetSection(path)

	index:= sort.SearchStrings(section,chapter.Title)

	if index < len(section) && section[index]==chapter.Title {
		return
	}

	result.WordRecord=append(result.WordRecord,chapter)
	content, _ :=json.Marshal(result)
	fileTool.WriteFile(string(content), path)
}
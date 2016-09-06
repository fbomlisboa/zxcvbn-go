package frequency

import (
	"encoding/json"
	"github.com/fbomlisboa/zxcvbn-go/data"
	"log"
)

type FrequencyList struct {
	Name string
	List []string
}

var FrequencyLists = make(map[string]FrequencyList)

func init() {
	// maleFilePath := getAsset("data/MaleNames.json")
	// femaleFilePath := getAsset("data/FemaleNames.json")
	// surnameFilePath := getAsset("data/Surnames.json")
	// englishFilePath := getAsset("data/English.json")
	passwordsFilePath := getAsset("data/Passwords.json")
	biblicWordsFilePath := getAsset("data/Biblic-words-pt-br.json")
	brazilianSoccerPath := getAsset("data/Brazilian-soccer-teams.json")
	dicPtbr2 := getAsset("data/Dic-ptbr-utf8-2.json")
	// dicPtbr := getAsset("data/Dic-ptbr-utf8.json")
	umbandaCandomble := getAsset("data/Umbanda_candomble_terms_small.json")
	firstName := getAsset("data/First-name-pt-br.json")

	// FrequencyLists["MaleNames"] = GetStringListFromAsset(maleFilePath, "MaleNames")
	// FrequencyLists["FemaleNames"] = GetStringListFromAsset(femaleFilePath, "FemaleNames")
	// FrequencyLists["Surname"] = GetStringListFromAsset(surnameFilePath, "Surname")
	// FrequencyLists["English"] = GetStringListFromAsset(englishFilePath, "English")
	FrequencyLists["Passwords"] = GetStringListFromAsset(passwordsFilePath, "Passwords")
	FrequencyLists["Biblic"] = GetStringListFromAsset(biblicWordsFilePath, "Biblic")
	FrequencyLists["Brazilian"] = GetStringListFromAsset(brazilianSoccerPath, "Brazilian")
	FrequencyLists["DicPtbr2"] = GetStringListFromAsset(dicPtbr2,"DicPtbr2")
	// FrequencyLists["DicPtbr"] = GetStringListFromAsset(dicPtbr,"DicPtbr")
	FrequencyLists["UmbandaCandomble"] = GetStringListFromAsset(umbandaCandomble,"UmbandaCandomble")
	FrequencyLists["FirstName"] = GetStringListFromAsset(firstName,"FirstName")

}
func getAsset(name string) []byte {
	data, err := zxcvbn_data.Asset(name)
	if err != nil {
		panic("Error getting asset " + name)
	}

	return data
}
func GetStringListFromAsset(data []byte, name string) FrequencyList {

	var tempList FrequencyList
	err := json.Unmarshal(data, &tempList)
	if err != nil {
		log.Fatal(err)
	}
	tempList.Name = name
	return tempList
}

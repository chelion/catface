package unipal
import(
	"log"
	"time"
	"io/ioutil"
	"net/http"
	"strings"
	"errors"
	"encoding/json"
)
const(
	TOKEN_URL = "https://aip.baidubce.com/oauth/2.0/token"
	THRESHOLD = 0.9
)

type CatFace struct{
	tokenExpireUnixStamp int64
	apiUrl string
	apiKey string
	secretKey string
	accessToken string
}

type TokenApiResponse struct{
	ExpiresIn int64 `json:"expires_in"`
	Sessionkey string `json:"session_key"`
	AccessToken string `json:"access_token"`
}

func (catface *CatFace)Config(apiUrl string,apiKey string,secretKey string){
	catface.apiUrl = apiUrl
	catface.apiKey = apiKey
	catface.secretKey = secretKey
}

func (catface *CatFace)getToken()bool{
	var tar TokenApiResponse
	resp, err := http.Post(TOKEN_URL,
		"application/x-www-form-urlencoded",
        strings.NewReader("grant_type=client_credentials&client_id="+catface.apiKey+"&client_secret="+catface.secretKey))
    if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Println(err)
        return false
    }
    err = json.Unmarshal(body, &tar)
    if err != nil {
		log.Println(err)
		return false
	}
	catface.tokenExpireUnixStamp = tar.ExpiresIn+time.Now().Unix()
	catface.accessToken = tar.AccessToken
	return true
}

func (catface *CatFace)doPost(data string)(string,error){
	var param map[string]interface{}
	if time.Now().Unix() > catface.tokenExpireUnixStamp{
		if false == catface.getToken(){
			return "",errors.New("get token fail")
		}
	}
	param = make(map[string]interface{})
	param["image"] = data
	param["threshold"] = THRESHOLD
	dataType , _ := json.Marshal(param)
	dataString := string(dataType)
	resp, err := http.Post(catface.apiUrl+"?access_token="+catface.accessToken,
        "application/json",
        strings.NewReader(dataString))
    if err != nil {
		log.Println(err)
		return "",err
	}
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		log.Println(err)
        return "",err
	}
	return string(body),nil
}

func (catface *CatFace)DetectCatFace(imageBase64Data string)(string,error){
	return catface.doPost(imageBase64Data)
}

package main
import(
	"log"
	"time"
	"encoding/base64"
	"io/ioutil"
	"github.com/chelion/catface/unipal"
)
const(
	UNIPAL_CATFACE_SERVERURL	= "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/detection/unipal_cat"
	APIKEY = "E3Gz952ySviRchHMUzUsoy1L"
	SECRETKEY = "WxNH1oBNsDdbrN3Mvl7fIFMgh0A5PwrW"
)
func main(){
	var catface unipal.CatFace
	catface.Config(UNIPAL_CATFACE_SERVERURL,APIKEY,SECRETKEY)
	imageData0, err := ioutil.ReadFile("0.jpg")
	if err != nil {
		log.Fatal(err)
		return
	}
	imageBase64Data0 := base64.StdEncoding.EncodeToString(imageData0)
	if "" != imageBase64Data0{
		t1 := time.Now()
		info,err := catface.DetectCatFace(imageBase64Data0)
		if nil == err{
			elapsed := time.Since(t1)
			log.Println("Detect elapsed: ", elapsed)
			log.Println(info)
		}
	}

	imageData1, err := ioutil.ReadFile("1.jpg")
	if err != nil {
		log.Fatal(err)
		return
	}
	imageBase64Data1 := base64.StdEncoding.EncodeToString(imageData1)
	if "" != imageBase64Data1{
		t1 := time.Now()
		info,err := catface.DetectCatFace(imageBase64Data1)
		if nil == err{
			elapsed := time.Since(t1)
			log.Println("Detect elapsed: ", elapsed)
			log.Println(info)
		}
	}
}
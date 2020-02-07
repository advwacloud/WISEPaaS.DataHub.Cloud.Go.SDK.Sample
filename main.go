package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	sdk "WISEPaaS.SCADA.Cloud.Go.SDK"
	"github.com/gorilla/mux"
)

var cloudAgent sdk.Agent

func getRealDataHandler(w http.ResponseWriter, r *http.Request) {

	s, _ := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	// bearerToken := r.Header.Get("Authorization")
	// splitToken := strings.Split(bearerToken, "Bearer")
	// // fmt.Println(string(s))
	// // fmt.Println(splitToken[1])
	// bearerToken = splitToken[1]
	// options := &sdk.DataStoreOptions{
	// 	URL:         "https://portal-scada-develop.wise-paas.com",
	// 	AccessToken: bearerToken,
	// }
	req := realDataReqDataFormat(s)
	response := cloudAgent.GetRealData(req)
	jsondata, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}
func getHisDataHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	var hisRawDataTags sdk.HistRawDataRequest
	json.Unmarshal(s, &hisRawDataTags)
	response := cloudAgent.GetHistoryData(hisRawDataTags)
	jsondata, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}
func realDataReqDataFormat(req []byte) sdk.RealDataReq {
	var realTags sdk.RealDataReq
	json.Unmarshal(req, &realTags)
	return realTags
}
func main() {
	r := mux.NewRouter()
	options := &sdk.DataStoreOptions{
		URL:         "https://portal-scada-develop.wise-paas.com",
		AccessToken: "YRI27pfMH1HItc3Yk4MqdcExhQlSjg",
	}
	cloudAgent = sdk.DataStore(options)
	r.HandleFunc("/api/RealData/raw", getRealDataHandler).Methods("POST")
	r.HandleFunc("/api/HistData/raw", getHisDataHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

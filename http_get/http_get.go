package main
import (
   "fmt"
   "io/ioutil"
   "net/http"
   "crypto/tls"
)

//    tr := &http.Transport{
//    	TLSClientConfig:    &tls.Config{RootCAs: pool},
//    	DisableCompression: true,
//    }
//    client := &http.Client{Transport: tr}
//    resp, err := client.Get("https://example.com")

func main() {

// RootCAs: pool,
   tr := &http.Transport{ TLSClientConfig: &tls.Config{ InsecureSkipVerify: true}, DisableCompression: true, }
   client := &http.Client{Transport: tr}
   resp, _ := client.Get("http://vm17/")

   // resp , _ := http.Get("http://vm6/tmtrack/tmtrack.dll?")
   
   body , _ := ioutil.ReadAll(resp.Body)
   fmt.Println(string(body))
   resp.Body.Close()
}



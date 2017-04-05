package main

import (
   "fmt"
   "io/ioutil"
	"log"
   "golang.org/x/net/html"
)


//    tr := &http.Transport{
//    	TLSClientConfig:    &tls.Config{RootCAs: pool},
//    	DisableCompression: true,
//    }
//    client := &http.Client{Transport: tr}
//    resp, err := client.Get("https://example.com")

// Directory of C:\Users\pthompson\Documents\Clients\CACI\Projects\OPM II\Daily Load\01-24-2017
// Directory of C:/Users/pthompson/Documents/Clients/CACI/Projects/OPM II/Daily Load/01-24-2017
//
// 01/24/2017  07:11 AM        48,885,481 PIP97639-01-24-2017.txt
// 01/24/2017  07:11 AM            14,943 01-24-2017-No FW-SCReport.xlsx
//
// 01/24/2017  07:36 AM            30,163 01-24-2017-WLReport.xls
// 01/24/2017  07:25 AM           157,383 01-24-2017-No FW-CLReport.xls
// 01/24/2017  07:35 AM         1,085,978 01-24-2017-Testimony.xls
// 01/24/2017  07:46 AM         1,818,768 01-24-2017-XDB.xls
// 01/24/2017  07:37 AM         2,666,022 01-24-2017-FFReport.xls
// 01/24/2017  07:09 AM         3,639,780 01-24-2017-CLReport.xls
// 01/24/2017  07:18 AM        91,302,470 01-24-2017-SCReport2.xls
// 01/24/2017  07:24 AM       396,646,843 01-24-2017-SCReport1.xls
// 01/24/2017  07:42 AM       427,762,700 01-24-2017-APIR.xls



//    01/24/2017  06:36 AM            30,163 01-24-2017-WLReport.xls
//
//       1 File(s)         30,163 bytes
//       0 Dir(s)  1,454,528,417,792 bytes free

func main() {

   // func ReadFile(filename string) ([]byte, error)
   dload_file , err := ioutil.ReadFile("C:/Users/pthompson/Documents/Clients/CACI/Projects/OPM II/Daily Load/01-24-2017/01-24-2017-WLReport.xls")
   if err != nil {
		log.Fatal(err)
	}
   // <nil>
   // 30163
   fmt.Println("Error:",err)
   fmt.Println("Length of file: " , len(dload_file))

   parse_xls()
}



func parse_xls() {
   println("set filename")
   excelFileName := "C:/Users/pthompson/Documents/Clients/CACI/Projects/OPM II/Daily Load/01-24-2017/01-24-2017-WLReport.xls"
   println("Open Excel file")
   // xlFile, err := xlsx.OpenFile(excelFileName)
   xlFile, err := xls.OpenFile(excelFileName)
   if err != nil {
		log.Fatal(err)
	}

   println("Iterate it")
   for _, sheet := range xlFile.Sheets {
      for _, row := range sheet.Rows {
         for _, cell := range row.Cells {
            text, _ := cell.String()
            fmt.Printf("%s\n", text)
         }
      }
   }
}
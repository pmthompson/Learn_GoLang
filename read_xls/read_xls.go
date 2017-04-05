package main

// To export a name it must begin with a Capital Letter.
// When importing a package, you can refer only to its exported names which must be Capitalize.
// the "if" short statement can include an "init" statement before the comparison.


// Directory of C:\Users\pthompson\Documents\Clients\CACI\Projects\OPM II\Daily Load\01-24-2017
// Directory of C:/Users/pthompson/Documents/Clients/CACI/Projects/OPM II/Daily Load/01-24-2017
//
//    01/24/2017  07:11 AM        48,885,481 PIP97639-01-24-2017.txt
//    01/24/2017  07:11 AM            14,943 01-24-2017-No FW-SCReport.xlsx
//
//    03/20/2017  11:18 PM            31,744 01-24-2017-WLReport.excel97.xls
//    03/20/2017  11:20 PM            43,612 01-24-2017-WLReport.XMLSpreadsheet2003.xml
//
//    01/24/2017  07:36 AM            30,163 01-24-2017-WLReport.xls
//    01/24/2017  07:25 AM           157,383 01-24-2017-No FW-CLReport.xls
//    01/24/2017  07:35 AM         1,085,978 01-24-2017-Testimony.xls
//    01/24/2017  07:46 AM         1,818,768 01-24-2017-XDB.xls
//    01/24/2017  07:37 AM         2,666,022 01-24-2017-FFReport.xls
//    01/24/2017  07:09 AM         3,639,780 01-24-2017-CLReport.xls
//    01/24/2017  07:18 AM        91,302,470 01-24-2017-SCReport2.xls
//    01/24/2017  07:24 AM       396,646,843 01-24-2017-SCReport1.xls
//    01/24/2017  07:42 AM       427,762,700 01-24-2017-APIR.xls



import (
    "fmt"
    "github.com/extrame/xls"
)

/*
func (w *Work Book) ReadAllCells() (res [][]string) {
    for _, sheet := range w.Sheets {
        w.PrepareSheet(sheet)
        if sheet.MaxRow != 0 {
            temp := make([][]string, sheet.MaxRow+1)
            for k, row := range sheet.Rows {
                data := make([]string, 0)
                if len(row.Cols) > 0 {
                    for _, col := range row.Cols {
                        if uint16(len(data)) <= col.LastCol() {
                            data = append(data, make([]string, col.LastCol()-uint16(len(data))+1)...)
                        }
                        str := col.String(w)
                        for i := uint16(0); i < col.LastCol()-col.FirstCol()+1; i++ {
                            data[col.FirstCol()+i] = str[i]
                        }
                    }
                    temp[k] = data
                }
            }
            res = append(res, temp...)
        }
    }
    return
}
*/

func main() {

   // 01-24-2017-WLReport.xls                            Really an HTML file
   // 01-24-2017-WLReport.excel97.xls                    That file saves as Excel97
   // 01-24-2017-WLReport.XMLSpreadsheet2003.xml         Same file saved as XML SpreadSheet 2003

   // This fails "Not an Excel file"
   xlsFname := "C:/Users/pthompson/Documents/Clients/CACI/Projects/OPM II/Daily Load/01-24-2017/01-24-2017-WLReport.xls"
   // This works but the extra manual step of converting it to real Excel 97 is not acceptable.
   xlsFname = "C:/Users/pthompson/Documents/Clients/CACI/Projects/OPM II/Daily Load/01-24-2017/01-24-2017-WLReport.excel97.xls"


   var xlFile *xls.WorkBook
   var err error

   xlFile, err = xls.Open(xlsFname, "utf-8")
   if err != nil {
      fmt.Println(err)
   }
   fmt.Println("xlFile is a %T\n",xlFile)
}


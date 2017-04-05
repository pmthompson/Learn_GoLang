package main
import (
   // Import go-mssqldb strictly for side-effects
   _ "github.com/denisenkom/go-mssqldb"
   "database/sql"
   "log"
)

func main() {

   var n_tables int

   println (sql.Drivers())

   // select count(*) from [servername\instance].database1.schema.tablename with (nolock)
   // Works in query :  select COUNT(*) from {{SERVER}}_SBM_AE_OE_REPO_CL.dbo.ts_tables

   // sqlserver://sa:mypass@localhost?database=master&connection+timeout=30 // username=sa, password=mypass.
   // sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30" // port 1234 on localhost.
   const server = "{{SERVER}}"
   const server_port = "1433"
   const instance = "{{SERVER}}_SQLSERVER"

   const DB = "{{SERVER}}_SBM_AE_OE_REPO_CL"
   const schema = "dbo"

   // Using SA doesn't make any difference
   // const userid = "{{SERVER}}_SBM"
   const userid = "SA"
   // const pwd = "{{PASSWORD}}"
   const pwd = "{{PASSWORD}}"
   
   const pwd_encoded = "P%4055word"

   // Connection Properties
   // applicationIntent       "ReadOnly" or "ReadWrite"
   // authentication          Use SqlPassword to connect using userName/user and password.
   //                            Note: When authentication property is set to any value
   //                            other than NotSpecified, the driver by default uses
   //                             Secure Sockets Layer (SSL) encryption.
   // databaseName, database
   // instanceName
   // integratedSecurity      Set to "true" to indicate that Windows credentials will be
   //                            used by SQL Server on Windows operating systems.
   // password
   // portNumber, port
   // serverName, server
   // trustServerCertificate
   // userName, user



   // ADO?
   // server=localhost\\SQLExpress;user id=sa;database=master;connection timeout=30
   // server=localhost;user id=sa;database=master;connection timeout=30
   // Server=myServerAddress;Database=myDataBase;User Id=myUsername;Password=myPassword;

   str_ADO_ConnectString := "Server=" + server + ";Database=" + DB + ";User Id=" + userid + ";Password=" + pwd + ";"
   _ = str_ADO_ConnectString

   // ODBC ???
   // odbc:server=localhost\\SQLExpress;user id=sa;database=master;connection timeout=30
   // odbc:server=localhost;user id=sa;database=master;connection timeout=30
   // odbc:server=localhost;user id=sa;password={foo;bar} // Value marked with {}, password is "foo;bar"
   str_ODBC_ConnectString := "odbc:server=" + server + ";user id=" + userid + ";password=" + pwd + ";database=" + DB + "." + schema + ";connection timeout=30"
   _ = str_ODBC_ConnectString

   // URL ??????
   // MS jdbc format does NOT work
   // URL: with sqlserver scheme. username and password appears before the host.
   // Any instance appears as the first segment in the path. All other options are query parameters. Examples:

   // Login error: read tcp 192.168.91.1:6132->192.168.91.135:1433: wsarecv: An existing connection was forcibly closed by the remote host.
   // Unable to get instances from Sql Server Browser on host {{SERVER}}: read udp 192.168.91.1:49332->192.168.91.135:1434: i/o timeout

   // sqlserver://username:password@host/instance?param1=value&param2=value
   // sqlserver://username:password@host:port?param1=value&param2=value
   // sqlserver://sa@localhost/SQLExpress?database=master&connection+timeout=30     // SQLExpress instance.
   // sqlserver://sa:mypass@localhost?database=master&connection+timeout=30         // username=sa, password=mypass.
   // sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30"   // port 1234 on localhost.
   // sqlserver://sa:my%7Bpass@somehost?connection+timeout=30                       // password is "my{pass"
   // str_URL_ConnectString :="sqlserver://" + userid + ":" + pwd_encoded + "@" + server + "?database=" + DB + "&connection+timeout=30"
   // str_URL_ConnectString :="sqlserver://" + userid + ":" + pwd_encoded + "@" + server + "/" + instance + "?database=" + DB + "&connection+timeout=30"
   str_URL_ConnectString :="sqlserver://" + userid + ":" + pwd + "@" + server +":" + server_port + "?database=" + DB + "&connection+timeout=30&encrypt=disable"
   // str_URL_ConnectString :="sqlserver://" + userid + ":" + pwd_encoded + "@" + server +":" + server_port + "/" + instance + "?database=" + DB + "&connection+timeout=30"
   _ = str_URL_ConnectString



   connectString := str_URL_ConnectString
   println("Connection string=" , connectString )

   println("open connection")
   db, err := sql.Open("mssql", connectString)
   defer db.Close()
   println ("Open Error:" , err)

   if err != nil {
      log.Fatal(err)
   }


   println("count records in TS_TABLES & scan")
   err = db.QueryRow("Select count(*) from ts_tables").Scan(&n_tables)
   if err != nil {
      log.Fatal(err)
   }
   println ("count of tables" , n_tables)

   println("closing connection")
   db.Close()

}

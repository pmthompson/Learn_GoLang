package main

import (
   "fmt"
   "strings"
   "syscall"
   "os"
   "golang.org/x/sys/windows"
   "golang.org/x/sys/windows/svc/eventlog"
)

// ----------------------------------------------------------------------------
// - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- -
// ----------------------------------------------------------------------------

// https://play.golang.org/p/PDKP6ZfV0X

func toShort(path string) (string, error) {
	p, err := syscall.UTF16FromString(path)
	if err != nil {
		return "", err
	}
	b := p // GetShortPathName says we can reuse buffer
	n := uint32(len(b))
	for {
		n, err = syscall.GetShortPathName(&p[0], &b[0], uint32(len(b)))
		if err != nil {
			return "", err
		}
		if n <= uint32(len(b)) {
			return syscall.UTF16ToString(b[:n]), nil
		}
		b = make([]uint16, n)
	}
}

// ----------------------------------------------------------------------------

//    Type and function names:
//       If the name of the type or function starts with an uppercase letter,
//          it’s visible outside the package.
//       If it starts with a lowercase letter, it isn’t.
//
//    Same rule applies to structure fields.
//       If a structure field name starts with a lowercase letter, only code
//          within the same package will be
//          able to access them.

func main() {

   println("\n")

   var str , str2 = "This is the string `str`." , "the one after str"

   // Strings have len but not cap
   fmt.Printf("str ::= (%T) == '%s'\n",str,str)
   fmt.Printf("str2 ::= (%T) == '%s'\n",str2,str2)
   println("len(str)=",len(str))

   // str[0] returns the character code; ASCII or UTF-8
   fmt.Printf("str[0]=%s\n",str[0])
   fmt.Printf("str[0]='%s'\n",str[0])
   fmt.Printf("str[0:5]='%s'\n",str[0:5])

   print("strptr = &str\n")
   strptr := &str
   fmt.Printf("strptr ::= (%T) == %d ; *strptr ::= (%T) == %s\n",strptr,strptr,*strptr,*strptr)
   str2ptr := &str2
   fmt.Printf("str2ptr ::= (%T) == %d ; *strptr ::= (%T) == %s\n",str2ptr,str2ptr,*strptr,*strptr)

   str = "THIS IS A NEW VERSION OF 'str'"
   fmt.Printf("str ::= (%T) == '%s'\n",str,str)
   fmt.Printf("strptr ::= (%T) == %d ; *strptr ::= (%T) == %s\n",strptr,strptr,*strptr,*strptr)
   fmt.Printf("str2ptr ::= (%T) == %d ; *strptr ::= (%T) == %s\n",str2ptr,str2ptr,*strptr,*strptr)

   /* ************************************** */   println(strings.Repeat("-",100) + "\n\n")  /* ************************************** */

   str_EnvStrs := windows.Environ()      // []string
   fmt.Printf("str_EnvStrs ::= %T\n",str_EnvStrs)

   /*
   for idx, str := range str_EnvStrs {
      fmt.Printf("str_EnvStrs :: idx == %v == %#v ; str == %#v\n",idx,idx,str)
   }
   */

   //    str_EnvStrs[0] == =::=::\
   //    str_EnvStrs[1] == =C:=C:\Users\pthompson\Documents\Source\go\src\github.com\pmthompson\maps_arrays_slices

   fmt.Printf("str_EnvStrs[0] == %v == %#v\n",str_EnvStrs[0],str_EnvStrs[0])
   fmt.Printf("str_EnvStrs[1] == %v == %#v\n",str_EnvStrs[1],str_EnvStrs[1])

   // argv[0] == "C:\\Users\\PTHOMP~1\\AppData\\Local\\Temp\\go-build882645766\\command-line-arguments\\_obj\\exe\\Maps_Arrays_Slices.exe"
   // var Args []string
   str_Args := os.Args
   fmt.Printf("str_Args ::= %T\n",str_Args)
   for idx, str := range str_Args {
      fmt.Printf("str_Args :: idx == %v == %#v ; str == %#v\n",idx,idx,str)
   }


   str_Exe , err := os.Executable()
   if err != nil {
      panic(err)
   }
   fmt.Printf("str_Exe ::= %T == %#v\n",str_Exe,str_Exe)


   // (name string, err error)
   str_CompName , err := windows.ComputerName()
   if err != nil {
      panic(err)
   }
   println("str_CompName=",str_CompName)


   // Convert a path name to it's short name
   // Compile error:    non-hex character in escape sequence: s
   //                   unknown escape sequence
   // str , err = toShort("C:\Users\pthompson\Documents\Source\go\src\github.com\pmthompson\maps_arrays_slices\Maps_Arrays_Slices.go")

   // result: str= C:\Users\PTHOMP~1\DOCUME~1\Source\go\src\github.com\PMTHOM~1\MAPS_A~1\MAPS_A~1.GO
   // str , err = toShort("C:\\Users\\pthompson\\Documents\\Source\\go\\src\\github.com\\pmthompson\\maps_arrays_slices\\Maps_Arrays_Slices.go")

   // result: str= C:/Users/PTHOMP~1/DOCUME~1/Source/go/src/github.com/PMTHOM~1/MAPS_A~1/MAPS_A~1.GO
   str , err = toShort(str_EnvStrs[1])

   if err != nil {
      panic(err)
   }
   println("str=",str)



   // func (*Log) Info
   // func (l *Log) Info(eid uint32, msg string) error
   //    Info writes an information event msg with event id eid to the end of event log l. When EventCreate.exe is used, eid must be between 1 and 1000.
   //  "golang.org/x/sys/windows/svc/eventlog"
   ptr_infoLog , err := eventlog.Open("Info")
   if err != nil {
      panic(err)
   }

   defer ptr_infoLog.Close()
   fmt.Printf("ptr_infoLog ::= (%T) == %d\n",ptr_infoLog,ptr_infoLog)

   err = ptr_infoLog.Info(999,"This is a test INFO log entry from my go program")
   if err != nil {
      panic(err)
   }


   ptr_infoLog.Close()

   /* ************************************** */   println(strings.Repeat("-",100) + "\n\n")  /* ************************************** */

   print ("\nvar scores [5] int\n")
   var scores [5] int
   scores[0] = 1
   fmt.Printf("scores array is a (%T) : len(scores)=%d ; iterate using for idx:=0 ; idx<len(scores) ; idx++  \n",scores,len(scores))
   for idx:=0 ; idx<len(scores) ; idx++ {
      fmt.Printf("idx ::= (%T) =%d ; scores[idx] = %d\n",idx,idx,scores[idx])
   }

   print ("\nscores = [5]int{0,11,222}\n")
   scores = [5]int{0,11,222}

   // idx is the index ; val is the value of the array at that index
   print("\niterate array using : for idx := range slice.  This automatically discards 'val', the second value returned by 'range'.\n")
   for idx := range scores {
      fmt.Printf("idx ::= (%T) =%d ; scores[idx] == %d\n",idx,idx,scores[idx])
   }

   print("\niterate array using : for idx,_ := range slice\n")
   for idx,_ := range scores {
      fmt.Printf("idx ::= (%T) =%d ; scores[idx]=%d\n",idx,idx,scores[idx])
   }

   print("\niterate array using : for idx, val := range scores\n")
   // index and value
   for idx, val := range scores {
      fmt.Printf("idx ::= (%T) =%d ; val ::= (%T) =%d \n",idx,idx,val,val)
   }

   print("\niterate array using : for _, val := range scores\n")
   // value only
   for _, val := range scores {
      fmt.Printf("val ::= (%T)  = %d\n",val,val)
   }

   /* ************************************** */   println(strings.Repeat("-",100) + "\n\n")  /* ************************************** */

}

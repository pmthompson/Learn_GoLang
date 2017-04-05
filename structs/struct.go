package main

import (
   "fmt"
   "strings"
)

// ----------------------------------------------------------------------------
// - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- -
// ----------------------------------------------------------------------------

// Create a struct to be included in another struct. Called a "Composition" ??
type g_score struct {
   Current int
   High int
}


// Our main struct
type g_struct struct {
   Name  string
   Val   string
   // We created the g_score element as anonymous.
   // GO secretly gave the element a name which is the struct's type name: "g_score"
   // It's members can be addressed using the struct type name and the element name or just the element name:
   //    " xxxx.g_score.Current"  OR    " xxxx.Current"
   *g_score
   // Adding another instance of the anonymous element returns compiler error:
   //       duplicate field g_score
   // *g_score

}

// - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- - -- -

// Function that expects a g_struct param
// Default is pass-by-value
func f1(gs g_struct) {
   gs.Val = gs.Val + " -- f1 was here"
   gs.g_score.Current--
   gs.g_score.High++
}

// Use a pointer
func f2(gs *g_struct) {
   gs.Val = gs.Val + " -- f2 was here"
}
// ----------------------------------------------------------------------------

// a g_struct method.
// Note the different function definition syntax compared to a function with a g_struct param.
// "this" is the "receiver" of the "ProperCase" method.  Because it's a pointer we can modify the object on which it's called.
// Notice we don't need to dereference the pointer (i.e "this->Name = strings.Title(this->Name)") .  Go does it for us.
// If the "g_struct" had many methods and 1 needed a pointer, then they should all get pointers for consistency's sake.
func (this *g_struct) ProperCase() {
   this.Name = strings.Title(this.Name)
   this.Val = strings.Title(this.Val)
}

// We can refere
func (this *g_struct) Prt() string {
   // the g_struct's "g_score" members can be addressed using the struct type name and element name or just the element name:
   return fmt.Sprintf("[" + this.Name + "]::[" + this.Val + "]::[%d , %d]" , this.g_score.Current , this.High )
}

// ----------------------------------------------------------------------------

func main() {

   // Not sure how to initialize this without the "g_score" type name
   my_var  := g_struct { Name:"All on one line" , Val:"Trailing COMMA not required" , g_score: &g_score{1,2} }

   my_var2  := g_struct { Name:"When last element isn't on same line as closing curly brace" ,
      Val:"closing COMMA required" ,
      &g_score{1,2} ,
   }

   // Basically the same as : tptr := new(g_struct)
   tptr := &g_struct { "tptr is a pointer to me" , "and I'm happy" , &g_score{1,2} }


   //  What is my_var and tptr?
   fmt.Printf("my_var is a %T\n" , my_var)
   fmt.Printf("tptr is a %T\n" , tptr)

   println("myvar= " , my_var.Prt())

   println("my_var2= " , my_var2.Prt())

   // printing elements using our pointer seems to work the same as using the variable itself
   println("tptr= " , tptr.Prt())

   // Don't modify the struct
   f1(my_var)
   println("After calling f1 :: my_var=" , my_var.Prt())

   // modify the struct
   f2(&my_var)
   println("After calling f2 :: my_var=" , my_var.Prt())

   // do it with our pointer variable.  Don't modify the struct.
   // To call this we need to de-reference the pointer and pass the actual objecctt.
   f1(*tptr)
   println("After calling f1 :: tptr=" , tptr.Prt())

   // modify the struct
   f2(tptr)
   println("After calling f2 :: tptr=" , tptr.Prt())

   // Call our method
   tptr.ProperCase()

   // println("After calling tptr.ProperCase :: tptr=[" , tptr.Name , "]::[" , tptr.Val , "]::[" , tptr.Score.Current , "]" )

   println("After calling tptr.ProperCase :: tptr=" , my_var.Prt())
}

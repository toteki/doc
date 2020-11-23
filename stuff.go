package doc

import "fmt"

//FunctionOne prints 'one' to console
func FunctionOne() {
  fmt.Println("one")
}

//FunctionTwo prints the string 'two ' concatenated with the argument to the screen.
func FunctionTwo(arg string) {
  fmt.Println("two " + arg)
}

//TypeOne is a structure with exported and unexported fields.
type TypeOne struct {
  Exported string
  unexported string
}

//FuncThree prints both the exported and unexported fields of a TypeOne struct.
func (to *TypeOne) FuncThree() {
  fmt.Println("Exported: " + to.Exported)
  fmt.Println("unexported: " + to.unexported)
}

//FuncFour returns a new TypeOne with fields based off of the argument.
func FuncFour(arg string) TypeOne {
  to := TypeOne{
    Exported: "ex" + arg,
    unexported: "unex" + arg,
  }
  return to
}

package main
import (
    "bufio"
    "io"
    "os"
    "log"
    "fmt"
)
func main(){
    //os.Create("hello.go")
   f,err:=os.Open("read_count")
    if err != nil {
	log.Fatal(err)
     }
     r:=bufio.NewReader(f)
    stat:=make(map[rune]int)
     for {
         ch,_,err:=r.ReadRune()
         log.Println(ch)
         if err==io.EOF{
             break
         }else if err!=nil{
             log.Fatal(err)
         }
         stat[ch]=stat[ch]+1
     }
     for r,i:=range stat{
         fmt.Println(string(r),i)
     }
      
     

     
     
     

}

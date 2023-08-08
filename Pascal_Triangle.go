//Create a Pascal Triangle
package main

import "fmt"

func main(){
    var size int
    fmt.Scanln(&size)
    a:=make([][]int,size)
    for i:=range a{
        a[i]=make([]int,i+1)
    }
    a[0][0]=1
    a[1][0]=1
    a[1][1]=1
    for i:=2;i<size;i++{
        a[i][0]=1
        for j:=0;j<i-1;j++{
           a[i][j+1]=a[i-1][j]+a[i-1][j+1]
        }
        a[i][i]=1
    }
   
    for i := 0; i < size; i++ {
        for k:=0;k<size-i-1;k++{
            fmt.Print(" ")
        }
        for j := 0; j <= i; j++ {
            fmt.Print(a[i][j], "  ")
        }
        fmt.Println()
    }

}

package main

import (
	"fmt"
    "time"
    "sync"
    "runtime"
)

var wg sync.WaitGroup

func computeFAPC(helperSource_x int, helperSource_y int, subRSD_width int, subRSD_height int, fapc [][]int, rows int, cols int){
    
    fapc[helperSource_x][helperSource_y]=1

    //setting bottommost row to 1
    for i:=1; i<subRSD_width; i++ {
        if(fapc[helperSource_x][helperSource_y+i]!=0){
            fapc[helperSource_x][helperSource_y+i]=fapc[helperSource_x][helperSource_y+i-1]
        }    
    }

    //leftmost column set to 1
    for i:=1; i<subRSD_height; i++ {
        if(fapc[helperSource_x+i][helperSource_y]!=0){
            fapc[helperSource_x+i][helperSource_y]=fapc[helperSource_x+i-1][helperSource_y]
        }    
    }

    //rest of the nodes' fapcs are calculated
    for i:=1; i<subRSD_height; i++{
        for j:=1; j<subRSD_width; j++{
            if(fapc[helperSource_x+i][helperSource_y+j]!=0){
                fapc[helperSource_x+i][helperSource_y+j] = fapc[helperSource_x+i-1][helperSource_y+j] + fapc[helperSource_x+i][helperSource_y+j-1]       
            }        
        }
    }
    
}


func main() {

    runtime.GOMAXPROCS(256)
    start := time.Now()    
 
	rows := 6
	cols := 6

	faultyNodes := [][]int{{0,1}, {0,2}, {0,3}, {0,4}, {0,5}, {1,2}, {1,3}, {1,4}, {1,5}, {2,3}, {2,4}, {2,5}, {3,4}, {3,5}, {4,5} }

	// Declare a slice of empty slices
	fapc := make([][]int, rows)
	// Declare those empty slices
	for i := 0; i < rows; i++ {
		fapc[i] = make([]int, cols)
	}

    // Declare a slice of empty slices
	rapc := make([][]int, rows)
	// Declare those empty slices
	for i := 0; i < rows; i++ {
		rapc[i] = make([]int, cols)
	}
    

    //Initialize all values as -1
    for i:=0; i<rows;i++{
        for j:=0; j<cols; j++{
            fapc[i][j]=-1        
            rapc[i][j]=-1
        }
    
    }

    //Set faulty nodes to 0
    for i:=0; i<len(faultyNodes); i++{    
        fapc[faultyNodes[i][0]][faultyNodes[i][1]]=0
        rapc[faultyNodes[i][0]][faultyNodes[i][1]]=0
    } 
    
    wg.Add(1)

    func() {
        computeFAPC(0,0,rows,cols, fapc, rows, cols)
        wg.Done()
    }()
    //go func() {
    //    computeFAPC(0,3,3,9, fapc, rows, cols)
    //    wg.Done()
    //}()
    //go func() {
    //    computeFAPC(0,6,3,9, fapc, rows, cols)
    //    wg.D*one()
    //}()

    wg.Wait()

    for i:=rows-1;i>=0;i--{
        for j:=0;j<cols;j++{
            fmt.Printf("%d\t",fapc[i][j])
        }
       fmt.Printf("\n")
    }
       
    elapsed := time.Since(start)    

    fmt.Println("Program took time =", elapsed)

}


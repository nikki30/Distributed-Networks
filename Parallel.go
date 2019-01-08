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
            
    wg.Done()
    
}


func computeAddMatrix(rightSubRSD_y int, rightSubRSD_width int, rightSubRSD_height int, fapc [][]int, addMatrix [][]int){
    j:= rightSubRSD_y-1

    if(fapc[0][j]==0){
        addMatrix[0][rightSubRSD_y]=-1
    }
        
      
    for i:=1; i<rightSubRSD_height; i++{
        if(addMatrix[i][rightSubRSD_y]==-1){
            addMatrix[i][rightSubRSD_y]=0
            continue
        }
        addMatrix[i][rightSubRSD_y]=addMatrix[i-1][rightSubRSD_y]+fapc[i][j]
    }

    for j=rightSubRSD_y+1; j<rightSubRSD_y+rightSubRSD_width; j++{
        for i:=0; i<rightSubRSD_height; i++{
            if(addMatrix[i][j]==-1){
                addMatrix[i][j]=0
                continue
            }
            if(i==0){
                addMatrix[i][j]=addMatrix[i][j-1]
                continue
            }            
            addMatrix[i][j]=addMatrix[i-1][j]+addMatrix[i][j-1]
        }    
    }

    wg.Done()
  
}


func main() {   

    runtime.GOMAXPROCS(128) 
    start := time.Now()
   
	rows :=32
	cols :=32

	faultyNodes := [][]int{{0,3}, {1,1}, {2,5}, {3,3}, {5,6}, {6,0}, {7,4}}

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

    size:= 2
    i:=0

    wg.Add(cols/size)
    for i<cols/size{
        go computeFAPC(0, i*size, size, rows, fapc, rows, cols)
        i++    
    }
    wg.Wait()
    
//    for it:=0; it<cols/size; it++{              
//        fmt.Printf("FAPC %d:\n",it+1)
//        for i:=rows-1;i>=0;i--{
//            for j:=0;j<size;j++{
//               fmt.Printf("%d\t",fapc[i][j+it*size])
//            }
//            fmt.Printf("\n")
//       }
//        fmt.Printf("\n")
//    }    
    
    //declaring addMatrix
    addMatrix := make([][]int, rows)
	for i:= 0; i < rows; i++ {
		addMatrix[i] = make([]int, cols)
	}


    for size<=cols{

        for i:= 0; i<rows; i++ {
            for j:=0; j<cols; j++{
		        addMatrix[i][j] = 0
            }
	    }
        
        for i:= 0; i <len(faultyNodes); i++ {
	        addMatrix[faultyNodes[i][0]][faultyNodes[i][1]] = -1
        }

        size=size*2

        wg.Add(cols/size)
        
        for i:=1; i<=cols/size; i++{
            go computeAddMatrix(i*size - size/2, size/2, rows, fapc, addMatrix)
        }

        wg.Wait()
    
//        for it:=size/2; it<=cols-size/2; it=it+size{
//            fmt.Printf("Add Matrix at %d:\n",it)
//            for i:=rows-1;i>=0;i--{
//                for j:=it;j<it+size/2;j++{
//                    fmt.Printf("%d\t",addMatrix[i][j])
//                }
//                fmt.Printf("\n")
//            }
//            fmt.Printf("\n")
//       } 
        
        for i:=0;i<rows;i++{
            for j:=0;j<cols;j++{
                if(addMatrix[i][j]>0){
                    fapc[i][j]+=addMatrix[i][j]
                }
            }
        }

//        for it:=0; it<cols/size; it++{            
//            fmt.Printf("FAPC %d:\n",it+1)
//            for i:=rows-1;i>=0;i--{
//               for j:=0;j<size;j++{
//                    fmt.Printf("%d\t",fapc[i][j+it*size])
//                }
//                fmt.Printf("\n")
//            }
//            fmt.Printf("\n")
//        }  
      
    }

 
    elapsed := time.Since(start)
    
    for i:=rows-1;i>=0;i--{
        for j:=0;j<cols;j++{
            fmt.Printf("%d\t",fapc[i][j])
        }
       fmt.Printf("\n")
    }
    

    fmt.Println("Program took time =", elapsed)

}


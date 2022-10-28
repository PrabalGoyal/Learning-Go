package main

import "fmt"

func main(){
	
	 arr := [7]string{"This", "is", "the", "tutorial",
	 "of", "Go", "language"}  // Creating an array


fmt.Println("Array:", arr) // Display array


myslice := arr[1:6]  // Creating a slice


fmt.Println("Slice:", myslice)  // Display slice


fmt.Printf("Length of the slice: %d", len(myslice))  // Display length of the slice

 
fmt.Printf("\nCapacity of the slice: %d", cap(myslice))  // Display the capacity of the slice

oRignAl_slice := []int{90, 60, 40, 50,
	34, 49, 30}

// Creating slices from the given slice
var my_slice_1 = oRignAl_slice[1:5]
my_slice_2 := oRignAl_slice[0:]
my_slice_3 := oRignAl_slice[:6]
my_slice_4 := oRignAl_slice[:]
my_slice_5 := my_slice_3[2:4]

// Display the result
fmt.Println("Original Slice:", oRignAl_slice)
fmt.Println("New Slice 1:", my_slice_1)
fmt.Println("New Slice 2:", my_slice_2)
fmt.Println("New Slice 3:", my_slice_3)
fmt.Println("New Slice 4:", my_slice_4)
fmt.Println("New Slice 5:", my_slice_5)

// Creating a zero value slice
arr := [6]int{55, 66, 77, 88, 99, 22}
slc := arr[0:4]

// Before modifying

fmt.Println("Original_Array: ", arr)
fmt.Println("Original_Slice: ", slc)

// After modification
slc[0] = 100
slc[1] = 1000
slc[2] = 1000

fmt.Println("\nNew_Array: ", arr)
fmt.Println("New_Slice: ", slc)
}
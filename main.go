package main

import (
    "fmt"
    "syscall"
    "unsafe"
)

var (
    user32           = syscall.NewLazyDLL("user32.dll")
    procMessageBoxW  = user32.NewProc("MessageBoxW")
)

const (
    MB_OK = 0x00000000
)

func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
    ret, _, _ := procMessageBoxW.Call(
        hwnd,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
        uintptr(flags))
    return int(ret)
}

func main() {
    var num1, num2 int
    fmt.Print("Enter the first number: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter the second number: ")
    fmt.Scanln(&num2)
    sum := num1 + num2
    sumStr := fmt.Sprintf("The sum of %d and %d is: %d", num1, num2, sum)
    MessageBox(0, sumStr, "Sum Result", MB_OK)
}

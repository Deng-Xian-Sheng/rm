//go build:windows
package WindowsAPI

import (
	"syscall"
	"unsafe"
)

// 定义 SHFILEOPSTRUCT 结构体
type SHFILEOPSTRUCT struct {
   hwnd   uintptr
   wFunc  uint32
   pFrom  *uint16
   pTo    *uint16
   fFlags uint32
}

// 定义 SHFileOperation 函数
var shell32 = syscall.NewLazyDLL("shell32.dll")
var procSHFileOperation = shell32.NewProc("SHFileOperationW")

func SHFileOperation(op *SHFILEOPSTRUCT) int {
   rc, _, _ := procSHFileOperation.Call(uintptr(unsafe.Pointer(op)))
   return int(rc)
}

// 定义一些常量
const (
   FO_DELETE          = 3
   FOF_ALLOWUNDO      = 0x0040
   FOF_NOCONFIRMATION = 0x0010
)
// 删除文件到回收站，适用于windows系统的函数
func DeleteToBin(s string) bool {

   // 构造 SHFILEOPSTRUCT 结构体
   op := &SHFILEOPSTRUCT{
      hwnd:   0,
      wFunc:  FO_DELETE,
      pFrom:  syscall.StringToUTF16Ptr(s), // 注意要以两个 \0 结尾
      pTo:    nil,
      fFlags: FOF_ALLOWUNDO | FOF_NOCONFIRMATION, // 使用 FOF_ALLOWUNDO 可以将文件移动到回收站
   }

   // 调用 SHFileOperation 函数
   ret := SHFileOperation(op)
   if ret != 0 {
      // 删除失败，输出错误信息
      err := syscall.Errno(ret)
      println("Error deleting file:", err)
   } else {
      // 删除成功
      println("File deleted successfully.")
      return true
   }
   return false
}



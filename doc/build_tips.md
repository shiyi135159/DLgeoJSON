## 1. 最小编译go程序
   把Go程序变小的办法是：
   go build -ldflags "-s -w"

   相关解释：
   -s去掉符号表,panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果，
   -w去掉DWARF调试信息，得到的程序就不能用gdb调试了。 -s和-w也可以分开使用.

## 2. windows程序（UAC）以管理员身份运行
   1> go get github.com/akavel/rsrc
   2> rsrc -manifest uac.manifest -o uac.syso
   3> 将生成的.syso文件拷贝到Go项目根目录
   4> go build

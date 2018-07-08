package main
import(
	"crypto/md5"
	"fmt"
	"crypto/sha1"
	"os"
	"io"
)
func main() {
	Content := "hello!"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(Content))
	Result :=Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(Content))
	Result =Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	File := "test.txt"
	infile, inerr := os.Open(File)
	if inerr == nil{
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%x %s\n", md5h.Sum([]byte("")), File)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%x %s\n", sha1h.Sum([]byte("")), File)

	}else{
		fmt.Println(inerr)
		os.Exit(1)
	}
}

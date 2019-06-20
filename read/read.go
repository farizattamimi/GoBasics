package main
import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)


//This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	//read whole text
	dat, err := ioutil.ReadFile("text.txt")
    check(err)
    fmt.Print(string(dat))


    //open text
    f, err := os.Open("text.txt")
    check(err)


    //read first five bytes
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))


    //read 6 bytes starting at 17
    o2, err := f.Seek(17, 0)
    check(err)
    b2 := make([]byte, 6)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    //implements io
    o3, err := f.Seek(17, 0)
    check (err)
    b3 := make([]byte, 6)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))


     _, err = f.Seek(0, 0)
    check(err)

	//The bufio package implements a buffered reader.
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

	//Close the file
    f.Close()

}
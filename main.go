package main

import (
	"flag"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var total *int = flag.Int("n", 0, "")
	var file *string = flag.String("f", "", "")
	flag.Parse()

	ReadWrite(file, total)
}

func ReadWrite(file *string, total *int) {
	var wg = sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int, 10)

	go func(file string, total int, wg *sync.WaitGroup, ch chan <- int) {
		f, err := os.OpenFile(file, os.O_CREATE, 0777)
		if err != nil {
			log.Fatal("Cannot create file " + file)
		}
		f.Close()

		for i:=0; i<total; {
			n := min(100, total - i)
			write(file, n)
			i+= n
			ch <- n
		}

		close(ch)
		wg.Done()
	}(*file, *total, &wg, ch)


	writer := func(s string) {
		os.Stdout.Write([]byte(s))
		os.Stdout.Write([]byte("\n"))
	}

	go func(file string, writer func(s string), ch <- chan int) {
		var start int = 0
		for n := range ch {
			total, err := read(file, writer, start, start + n)
			if err != nil {
				log.Fatalf("Unexpect error %v", err)
			}
			start += total
		}
		wg.Done()
	}(*file, writer, ch)

	wg.Wait()
}

func write(file string, total int) (int, error) {
	f, err := os.OpenFile(file, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		log.Fatal("Cannot open file " + file)
		return 0, err
	}

	rand.Seed(time.Now().Unix())
	retry := 3
	for i:=0; i<total; {
		num := strconv.Itoa(rand.Intn(math.MaxInt64))
		length := len(num)
		var n int = func(length int, remain int) int {
			if length <= remain {
				return length
			} else {
				return remain
			}
		}(length, total - i)

		_, err := f.WriteString(num[:n])
		if err != nil {
			if retry <= 0 {
				log.Fatalf("Some problems occur while writing to the file %v, retry %v",file, retry)
				return i, err
			}
			retry--
		} else {
			i += length
		}
	}
	return total, nil
}

func read(file string, writer func(s string), startIndex int, endIndex int) (int, error) {
	f, err1 := os.OpenFile(file, os.O_RDONLY | os.O_CREATE, 0644);
	defer f.Close()
	if err1 != nil {
		log.Fatalf("Cannot read file %v because of %v", file, err1)
		return 0, err1
	}

	buffer := make([]byte, 1)

	var err error = nil
	var i int = startIndex
	for ; err != io.EOF && i<endIndex; i++ {
		_, err := f.ReadAt(buffer, int64(i))
		if err == nil {
			s := string(buffer[0])
			if s[0] % 2 == 0 {
				writer(s + " is even number")
			} else {
				writer(s + " is odd number")
			}
		} else if err == io.EOF {
			return i, nil
		} else {
			return i, err
		}
	}
	return i, nil
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
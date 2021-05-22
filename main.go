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

	go write(*file, *total, &wg)

	writer := func(s string) {
		os.Stdout.Write([]byte(s))
		os.Stdout.Write([]byte("\n"))
	}

	go read(*file, &wg, writer)
	wg.Wait()
}

func write(file string, total int, wg *sync.WaitGroup) (int, error) {
	f, err := os.OpenFile(file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0777)
	defer f.Close()
	defer wg.Done()
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

func read(file string, wg *sync.WaitGroup, writer func(s string)) (int64, error) {
	f, err1 := os.OpenFile(file, os.O_RDONLY | os.O_CREATE, 0644);
	defer f.Close()
	defer wg.Done()
	if err1 != nil {
		log.Fatalf("Cannot read file %v because of %v", file, err1)
		return 0, err1
	}

	buffer := make([]byte, 1)

	var err error = nil
	var i int64 = 0
	for ; err != io.EOF; i++ {
		_, err := f.ReadAt(buffer, i)
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
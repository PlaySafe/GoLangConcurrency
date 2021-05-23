Please write a programme to complete the following function and finish the unit-test using go test.

1. Open a go routine to create a file named "input.file" and write a random number in string
   (the string length is 2048, number in [0-9]),when finished, close the file.
2. Open another go routine to open this file, read the string by digit. If the digit is odd number,
   print odd number otherwise print even number, e.g. if the digit read from file is 0, it's an even number,
   print "0 is even number". If the digit read from file is3, it's an odd number, print "3 is odd number"

The go routine in requirement 1 and requirement 2 should be able to run at the same time, in another word,
they should run concurrently.

------------------------------------------------------------------------------------------------------------

To run the application, use command `go run main.go -n <total> -f <file>`

**total** specifies how many numbers

**file** specifies either relative, or absolute path

To run test, use command `go test`
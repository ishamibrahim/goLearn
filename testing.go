package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const s string = "I'm Isham"

func typeChange() {
	const s string = "Im Isham inside change"
	fmt.Println(s)
}

//////////////////////////////////////
func typeCheck() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(3000)
}

//////////////////////////////////////
func arrayCheck() {
	var a [5]int
	fmt.Println("emp:", a)
}

//////////////////////////////////////
func sliceCheck() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "Zero"
	s[1] = "One"
	s[2] = "Two"
	s = append(s, "three")
	fmt.Println("s:", s)

	c := make([]string, len(s))
	copy(c, s)
	c = append(c, "Five")
	fmt.Println("s:", s)
	fmt.Println("c:", c)

}

//////////////////////////////////////
func mapCheck() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println("map:", m)
	delete(m, "two")
	apr, prs := m["two"]
	fmt.Println("key:", apr)
	fmt.Println("prs:", prs)
}

//////////////////////////////////////
func rangeCheck() {
	nums := []string{"first", "second", "third", "fourth", "fifth"}
	for i, num := range nums {
		fmt.Println("number : ", num)
		fmt.Println("index : ", i)
	}

	for i, c := range "ABCDEFGH" {
		fmt.Println(i, c)
	}
}

//////////////////////////////////////
func addThree(a, b, c int) int {
	return a + b + c
}
func funcCheck() {
	var a, b, c int = 3, 5, 7
	fmt.Println(addThree(a, b, c))
}

//////////////////////////////////////
func printPars(nums ...int) {
	for _, r := range nums {
		fmt.Println(r)
	}
}
func variardicCheck() {
	pars := []int{3, 5, 7, 4, 9}
	printPars(pars...)
}

//////////////////////////////////////
func getSeq() func() int {
	n := 5
	return func() int {
		n++
		return n
	}

}
func closureCheck() {
	nextNum := getSeq()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())
}

//////////////////////////////////////
func facto(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("Recursive call No. ", n)
	return n * facto(n-1)
}
func recursionCheck() {
	fmt.Println(facto(8))
}

//////////////////////////////////////
func changeVal(i int) {
	i = 0
}
func changePtr(i *int) {
	*i = 0
}
func pointerCheck() {
	var b = 345
	fmt.Println("Value of ib: ", b)
	changeVal(b)
	fmt.Println("Value after changing b : ", b)
	changePtr(&b)
	fmt.Println("Value after changing addresss of b : ", b)
	fmt.Println("Value of the address of b  ", &b)

}

//////////////////////////////////////
type rectangle struct {
	length  int
	breadth int
}

func structCheck() {
	f := rectangle{length: 30, breadth: 20}
	fmt.Println("rectangle details : ", f)
	fb := &f
	fb.breadth = 25
	fmt.Println("rectangle details : ", f)
}

//////////////////////////////////////
func (r rectangle) area() float64 {
	return float64(r.length * r.breadth)
}
func (r rectangle) perimeter() float64 {
	return float64(2*r.length + 2*r.breadth)
}

func (r *rectangle) perimeterptr() float64 {
	return float64(2*r.length + 2*r.breadth)
}
func structMethodCheck() {
	didi := rectangle{length: 20, breadth: 30}
	fmt.Println("Rectangle details : ", didi)

	fmt.Println("Rectangle area : ", didi.area())
	didir := &didi
	fmt.Println("Rectangle perimeter : ", didir.perimeterptr())
}

//////////////////////////////////////
func (r rectangle) diagonal() float64 {
	diag := math.Sqrt(float64((r.length * r.length) + (r.breadth * r.breadth)))
	return diag
}

type geometry interface {
	area() float64
	perimeter() float64
}

func measurements(g geometry) {
	fmt.Println("Geometry : ", g)
	fmt.Println("Area : ", g.area())
	fmt.Println("Perimeter : ", g.perimeter())
}
func interfaceCheck() {
	red := rectangle{breadth: 33, length: 24}
	measurements(red)
	fmt.Println("Diagonal : ", red.diagonal())
}

//////////////////////////////////////

func errorCheck(b int) (float64, error) {
	if b == 0 {
		return -1, errors.New("Anything other than 0")
	}
	return float64(b / 10), nil
}

//////////////////////////////////////
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func oddSquare(arg int) (int, error) {

	if arg%2 == 0 {
		return -1, &argError{arg, "Cannot continue with an even number"}
	}
	return arg * arg, nil

}
func customErrorCheck() {
	for _, i := range []int{2, 3, 5, 6, 8, 9} {
		if j, errMsg := oddSquare(i); errMsg != nil {
			fmt.Println("FAILED: ", errMsg)
		} else {
			fmt.Println("SUCCESS: ", j)
		}
	}
}

//////////////////////////////////////
func printMany(st int, rtype string) {
	for i := 0; i < st; i++ {
		fmt.Println("Calling ", rtype, " Goroutine no : ", i)
	}
}
func goroutineCheck() {
	printMany(4, "DIRECT")
	go printMany(12, "A")
	go printMany(10, "B")
	fmt.Scanln()
	fmt.Println("done")
}

//////////////////////////////////////
func changeChan(achan chan int, somval int) {
	fmt.Println("Channel : ", somval)
	achan <- somval
}
func channelCheck() {
	//wg := new(sync.WaitGroup)
	somchan := make(chan int, 2)

	changeChan(somchan, 5)
	changeChan(somchan, 20)

	firstval := <-somchan
	secondval := <-somchan
	fmt.Println(firstval)
	fmt.Println(secondval)
	fmt.Scanln()
	fmt.Println("done")

}

//////////////////////////////////////
func bufferChancheck() {
	somchan := make(chan string, 3)

	somchan <- "First"
	somchan <- "Second"
	somchan <- "third"

	fmt.Println(<-somchan)
	fmt.Println(<-somchan)
	fmt.Println(<-somchan)
}

//////////////////////////////////////
func workerWaits(b chan int) {
	fmt.Println("Starting ..")
	for i := 0; i < 5; i++ {
		fmt.Println("Doing...", i)
	}
	fmt.Println("Done.")
	b <- 3
}
func syncChannelCheck() {
	done := make(chan int, 1)
	go workerWaits(done)

	// If the below line is removed. The go routine runs asyncrounously.
	// The below line makes sure that the program waits till it gets the value for 'done'
	<-done
}

//////////////////////////////////////
func setChannel(v chan<- int, val int) {
	v <- val
	//fmt.Println("The value of the channel set is : ", <-v)
}
func getChannel(v <-chan int) {
	val := <-v

	fmt.Println("The value of the channel got is : ", val)
}

func channelDirectionCheck() {
	v := make(chan int, 1)
	setChannel(v, 31)
	getChannel(v)

}

//////////////////////////////////////
func runSystemCommand() {
	cmd := exec.Command("git", "describe", "--all")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed: %s", err)
	}
	fmt.Printf("The output is %s\n", out)
}

/////////////////////////////////////////

func selectChannels() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func withoutSelectChannels() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	msg1 := <-c1
	fmt.Println("received", msg1)
	msg2 := <-c2
	fmt.Println("received", msg2)

}

/////////////////////////////////////////

func channelTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

/////////////////////////////////////////

func nonBlockingChannels() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//fmt.Println("Message is  ", <-messages)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

/////////////////////////////////////////

func closingChannels() {
	messages := make(chan int, 2)
	done := make(chan bool)

	go func() {
		for {
			m, more := <-messages

			if more {
				fmt.Println("Recieved message", m)
			} else {
				fmt.Println("recieved all messages")
				done <- true
				break
			}
		}
	}()

	for i := 0; i < 4; i++ {
		messages <- i
		fmt.Printf("Message %d sent\n", i)
	}
	close(messages)
	fmt.Println("Sent all messages")
	<-done
}

/////////////////////////////////////////

func rangeChannel() {
	numbers := make(chan int, 5)

	for i := 0; i < 5; i++ {
		numbers <- i
	}
	close(numbers)

	for num := range numbers {
		fmt.Println("Recieving  ", num)
	}
}

/////////////////////////////////////////

func timerCheck() {
	timer1 := time.NewTimer(3 * time.Second)
	done := make(chan bool)

	<-timer1.C

	fmt.Println("Timer1 expired now")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired now")
		done <- true
	}()

	//This is used to stop a timer midway
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 Stopped")
	}

}

/////////////////////////////////////////

func tickerCheck() {
	ticker := time.NewTicker(time.Second)
	r := "Random wod"
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticking at ", t)
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped.", r)
}

/////////////////////////////////////////
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Job", j, "Has started by worker ", id)
		time.Sleep(time.Second)
		fmt.Println("Job", j, "Has ended by worker", id)
		results <- j * 10
	}
}

func workerPoolCheck() {
	jobs := make(chan int)
	results := make(chan int)

	for wId := 0; wId < 5; wId++ {
		go worker(wId, jobs, results)
	}

	for jNum := 1; jNum <= 5; jNum++ {
		jobs <- jNum
	}
	for rC := 1; rC <= 5; rC++ {
		fmt.Println("Rpinting result :", <-results)
	}

	close(jobs)
}

/////////////////////////////////////////

func rateLimiting() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Request number ", req, time.Now())
	}

	burstyLim := make(chan time.Time, 3)

	burstyLim <- time.Now()
	burstyLim <- time.Now()
	burstyLim <- time.Now()

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLim <- t
		}
	}()

	burstyReq := make(chan int, 6)
	for i := 0; i < 6; i++ {
		burstyReq <- i
	}
	for burst := range burstyReq {
		<-burstyLim
		fmt.Println("Bursty Request number ", burst, time.Now())
	}
	close(burstyLim)
}

///////////////////////////////////////////////////////////
func atomicCounter() {
	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {

				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

///////////////////////////////////////////////////////////

func mutexCheck() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state in : I ", state)
	time.Sleep(time.Millisecond)
	fmt.Println("state in : II ", state)
	time.Sleep(time.Millisecond)
	fmt.Println("state in: III ", state)
	mutex.Unlock()
	time.Sleep(time.Millisecond)
	fmt.Println("state out : I ", state)
	time.Sleep(time.Millisecond)
	fmt.Println("state out : II ", state)
	time.Sleep(time.Millisecond)
	fmt.Println("state out: III ", state)
}

///////////////////////////////////////////////////////////

func statefulGoRoutineCheck() {
	type readOp struct {
		key  int
		resp chan int
	}
	type writeOp struct {
		key  int
		val  int
		resp chan bool
	}

	var readOps uint64
	var writeOps uint64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

///////////////////////////////////////////////////////////

func sortCheck() {
	names := []string{"Bhim", "Arjun", "Nakul", "Yudhistir", "Sahadev"}
	fmt.Println("Unsorted names", names)
	sort.Strings(names)
	fmt.Println("Sorted names", names)

	fibonacci := []int{1, 0, 8, 2, 1, 5, 3}
	fmt.Println("unsorted fibonacci", fibonacci)
	sort.Ints(fibonacci)
	fmt.Println("Sorted fibonacci ", fibonacci)

}

///////////////////////////////////////////////////////////

type planet struct {
}

func sortByFuncCheck() {

}

//getInputFilePath get input template file path
func getInputFilePath(workDir, fileName string) string {
	ext := filepath.Ext(fileName)
	fileName = strings.TrimRight(fileName, ext)
	epoch := time.Now().Unix()
	fileName = fmt.Sprintf("%s_%d%s", fileName, epoch, ext)
	return path.Join(workDir, fileName)
}

//getOutputFilePath get output file path
func getOutputFilePath(inputFilePath string) string {
	ext := filepath.Ext(inputFilePath)
	filePath := strings.TrimRight(inputFilePath, ext)
	ouptputFilePath := filePath + "_out" + ext
	return ouptputFilePath
}

func reverselist(fl []string) {
	lenList := len(fl)
	for i := 0; i < lenList/2; i++ {
		fl[i], fl[lenList-1-i] = fl[lenList-1-i], fl[i]
	}
}

func main() {
	f := []string{"one", "two", "three", "four", "five", "six", "seven"}

	fmt.Println("Before reverse     ", f)
	reverselist(f)
	fmt.Println("After reverse     ", f)
}

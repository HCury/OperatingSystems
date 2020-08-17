//“I Henrique Cury(he200230) affirm that this program is entirely my own
// work and that I have neither developed my code together with
//any another person, nor copied any code from any other person,
//nor permitted my code to be copied or otherwise used by any
//other person, nor have I copied, modified, or otherwise used
//programs created by others. I acknowledge that any violation of
//the above terms will be treated as academic dishonesty.”

package main

import "os"
import "fmt"
import "strings"
import "bufio"
import "strconv"
import "sort"
import "io"

var f *os.File
var w io.Writer

type algo struct {
  processcount int
  runfor int
  use string
  quantum int
}

type process struct{
  processname string
  arrival int
  burst int
  wait int
  turnaround int
  finish int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func updateFinish(a algo, s []process, time int, name string ) {
  for i := 0; i < a.processcount; i++ {
    if(s[i].processname == name) {
      s[i].finish = time
    }
  }
}

func calcFWaitTime(a algo, slice []process) {
  for i := 0; i < a.processcount; i++ {
    slice[i].wait = slice[i].turnaround - slice[i].burst
  }
}

func calcTurnaround(a algo, slice []process) {
  for i := 0; i < a.processcount; i++ {
    slice[i].turnaround = slice[i].finish - slice[i].arrival
  }
}

func fcfs(a algo, slice []process) {
  w := bufio.NewWriter(f)
  queue := make([]process, 0)
  for i := 0; i < a.runfor; i++ {

    if len(queue) != 0 {
      queue[0].burst = queue[0].burst - 1
    }

    for j:= 0; j < a.processcount; j++ {

      if slice[j].arrival == i {
        if len(queue) == 0 {
          queue = append(queue, slice[j])
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s arrived\n", i, slice[j].processname)
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s selected (burst %3d)\n", i, queue[0].processname, queue[0].burst)
          w.Flush()
        }else {
          queue = append(queue, slice[j])
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s arrived\n", i, slice[j].processname)
          w.Flush()
        }
    }
    }

    if len(queue) > 0 {
      if queue[0].burst == 0 {
        w.Flush()
        fmt.Fprintf(w,"Time %3d : %s finished\n", i, queue[0].processname)
        updateFinish(a, slice, i, queue[0].processname)
        w.Flush()
        queue = queue[1:]
        if len(queue) > 0 {
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s selected (burst %3d)\n", i, queue[0].processname, queue[0].burst)
          w.Flush()
        }
      }
    }

    if len(queue) == 0 {
      w.Flush()
      fmt.Fprintf(w,"Time %3d : Idle\n", i)
      w.Flush()
    }
  }
    fmt.Fprintf(w,"Finished at time  %2d\n", a.runfor)
    fmt.Fprintf(w, "\n")
    w.Flush()
}

func sjf(a algo, s []process) {
  w := bufio.NewWriter(f)
  queue := make([]process, 0)

  for i := 0; i < a.runfor; i++ {
    head := ""
    if len(queue) != 0 {
      queue[0].burst = queue[0].burst - 1
      head = queue[0].processname
    }

    for j:= 0; j < a.processcount; j++ {

      if s[j].arrival == i {
        if len(queue) == 0 {
          queue = append(queue, s[j])
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s arrived\n", i, s[j].processname)
          w.Flush()
        }else {
          queue = append(queue, s[j])
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s arrived\n", i, s[j].processname)
          w.Flush()
          }
        }
      }

      if len(queue) > 0 {
        sort.Slice(queue, func(i, j int) bool {
          return queue[i].burst < queue[j].burst
        })
        if head != queue[0].processname {
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s selected (burst %3d)\n", i, queue[0].processname, queue[0].burst)
          w.Flush()
        }
      }

      if len(queue) > 0 {
        if queue[0].burst == 0 {
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s finished\n", i, queue[0].processname)
          updateFinish(a, s, i, queue[0].processname)
          w.Flush()
          queue = queue[1:]
          if len(queue) > 0 {
            w.Flush()
            fmt.Fprintf(w,"Time %3d : %s selected (burst %3d)\n", i, queue[0].processname, queue[0].burst)
            w.Flush()
          }
        }
      }
      if len(queue) == 0 {
        w.Flush()
        fmt.Fprintf(w,"Time %3d : Idle\n", i)
        w.Flush()
      }
    }
    fmt.Fprintf(w,"Finished at time %3d\n", a.runfor)
    fmt.Fprintf(w, "\n")
    w.Flush()
  }

func rr(a algo, s []process) {
  w := bufio.NewWriter(f)
  queue := make([]process, 0)

  fmt.Fprintf(w,"Quantum %3d\n\n", a.quantum)
  w.Flush()
  quantum := 0

  for i := 0; i < a.runfor; i++ {

    if len(queue) != 0 {
      queue[0].burst = queue[0].burst - 1
      quantum = quantum - 1
    }

    for j:= 0; j < a.processcount; j++ {

      if s[j].arrival == i {
        if len(queue) == 0 {
          queue = append(queue, s[j])
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s arrived\n", i, s[j].processname)
          w.Flush()
        }else {
          queue = append(queue, s[j])
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s arrived\n", i, s[j].processname)
          w.Flush()
          }
        }
      }

      if len(queue) > 0 {
        if queue[0].burst == 0 {
          w.Flush()
          fmt.Fprintf(w,"Time %3d : %s finished\n", i, queue[0].processname)
          updateFinish(a, s, i, queue[0].processname)
          w.Flush()
          queue = queue[1:]
          quantum = a.quantum
          if len(queue) > 0 {
            w.Flush()
            fmt.Fprintf(w,"Time %3d : %s selected (burst %3d)\n", i, queue[0].processname, queue[0].burst)
            w.Flush()
          }
        }
      }

      if quantum == 0 {
        quantum = a.quantum
        if len(queue) > 0 {
              if queue[0].burst == 0 {
                w.Flush()
                fmt.Fprintf(w,"Time %3d : %s finished\n", i, queue[0].processname)
                queue[0].finish = i
                w.Flush()
                quantum = a.quantum
                queue = queue[1:]
              } else {
                temp := queue[0]
                queue = queue[1:]
                queue = append(queue, temp)
                quantum = a.quantum
              }
              w.Flush()
              fmt.Fprintf(w,"Time %3d : %s selected (burst %3d)\n", i, queue[0].processname, queue[0].burst)
              w.Flush()
            }
        }

      if len(queue) == 0 {
        w.Flush()
        fmt.Fprintf(w,"Time %3d : Idle\n", i)
        w.Flush()
      }
    }
    fmt.Fprintf(w,"Finished at time %3d\n", a.runfor)
    fmt.Fprintf(w,"\n")
    w.Flush()
}

func main() {
 var a algo
 var b process
 var err error
 f, err = os.Create(os.Args[2])
 check(err)
 defer f.Close()

 w := bufio.NewWriter(f)

 arg := os.Args[1]; //declaring a variable with argument 1

 file, err := os.Open(arg) //opening that argument file
 check(err) //checking if its ok

 var lines []string //declaring new string with all the lines
 scanner := bufio.NewScanner(file) //creating scanner
 for(scanner.Scan()) {
   lines = append(lines, scanner.Text())
 } //adding every line to my string

 length := len(lines)

 //processcount
 splitter := strings.Split(lines[0], " ") //splitline by spaces
 splitterTest:= strings.Split(splitter[1], "\t") //split entire line
 count, err := strconv.Atoi(splitterTest[0]) //conv 5string to 5int
 a.processcount = count

 //runfor
 splitter = strings.Split(lines[1], " ") //splitline by spaces
 splitterTest = strings.Split(splitter[1], "\t") //split entire line
 count, err = strconv.Atoi(splitterTest[0]) //conv 5string to 5int
 a.runfor = count

//use
 splitter = strings.Split(lines[2], " ") //splitline by spaces
 splitterTest = strings.Split(splitter[1], "\t") //split entire line
 a.use = splitterTest[0]

//round robin
 if a.use == "rr" {
   splitter = strings.Split(lines[3], " ") //splitline by spaces
   splitterTest = strings.Split(splitter[1], "\t") //split entire line
   count, err = strconv.Atoi(splitterTest[0]) //conv 5string to 5int
   a.quantum = count
}

 struc := make([]process, a.processcount) //making a slice of type process
//proccesses
if a.use == "rr"{
  for i := 4; i <= length - 2; i++ {
    splitter = strings.Split(lines[i], " ") //splitline by spaces
    b.processname = splitter[2] //assigning processname
    count, err = strconv.Atoi(splitter[4])//convert
    b.arrival = count //arrival assign
    count, err = strconv.Atoi(splitter[6]) // convert
    b.burst = count //burst assign
    struc[i - 4] = b // assign struct to slice
  }
}else {
   for i := 3; i <= length - 2; i++ {
     splitter = strings.Split(lines[i], " ") //splitline by spaces
     b.processname = splitter[2] //assigning processname
     count, err = strconv.Atoi(splitter[4])//convert
     b.arrival = count //arrival assign
     count, err = strconv.Atoi(splitter[6]) // convert
     b.burst = count //burst assign
     struc[i - 3] = b // assign struct to slice
   }
}
 sort.Slice(struc, func(i, j int) bool {
   return struc[i].arrival < struc[j].arrival}) //sort slices by arrival time

 fmt.Fprintf(w, "%3d processes\n", a.processcount)
 w.Flush()


 switch a.use {
  case "fcfs":
   fmt.Fprintf(w ,"Using First-Come First-Served\n")
   w.Flush()
   fcfs(a, struc)
   sort.Slice(struc, func(i, j int) bool {
     return struc[i].processname < struc[j].processname})
   calcTurnaround(a, struc)
   calcFWaitTime(a, struc)
   for i := 0; i < a.processcount; i++ {
     fmt.Fprintf(w, "%s wait %3d turnaround %3d\n", struc[i].processname, struc[i].wait, struc[i].turnaround)
   }

   w.Flush()
  case "sjf":
   fmt.Fprintf(w, "Using preemptive Shortest Job First\n")
   w.Flush()
   sjf(a, struc)
   sort.Slice(struc, func(i, j int) bool {
     return struc[i].processname < struc[j].processname})
   calcTurnaround(a, struc)
   calcFWaitTime(a, struc)
   for i := 0; i < a.processcount; i++ {
     fmt.Fprintf(w, "%s wait %3d turnaround %3d\n", struc[i].processname, struc[i].wait, struc[i].turnaround)
   }

   w.Flush()
  case "rr":
    fmt.Fprintf(w, "Using Round-Robin\n")
    w.Flush()
    rr(a, struc)
    sort.Slice(struc, func(i, j int) bool {
      return struc[i].processname < struc[j].processname})
    calcTurnaround(a, struc)
    calcFWaitTime(a, struc)
    for i := 0; i < a.processcount; i++ {
      fmt.Fprintf(w, "%s wait %3d turnaround %3d\n", struc[i].processname, struc[i].wait, struc[i].turnaround)
    }

    w.Flush()
 }


}
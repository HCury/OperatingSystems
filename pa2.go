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





type algo struct {
  lowerCYL int
  upperCYL int
  initialCYL int
  use string
  currentCYL int
}

type cylinder struct{
  cyl int
  closeVal int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func calcCloseVal(a int, s []cylinder, length int) {
  for i := 0; i < len(s); i++ {
    s[i].closeVal = s[i].cyl - a
    if s[i].closeVal < 0 {
      s[i].closeVal = s[i].closeVal * -1
    }
}
  sort.Slice(s, func(i, j int) bool {
  return s[i].closeVal < s[j].closeVal}) //sort slices by closest value
}

func calcLimits(a algo) int{
  high := a.initialCYL - a.upperCYL
  low := a.initialCYL - a.lowerCYL

  if high < 0 {
    high = high * -1
  }
  if low < 0 {
    low = low * -1
  }

  if low < high {
    return a.lowerCYL
  } else {
    return a.upperCYL
  }
}

func calcCloseCyl(a int,s []cylinder, length int) {
  for i := 0; i < len(s); i++ {
    s[i].closeVal = s[i].cyl - a
  }
}

func fcfs(a algo, slice []cylinder, length int) {
  head := a.initialCYL
  temp := 0
  value := 0
  for i := 0; i < len(slice); i++ {
    fmt.Printf("Servicing%6d\n", slice[i].cyl)
    temp = head - slice[i].cyl
    if temp < 0 {
      temp = temp * -1
    }
    head = slice[i].cyl
    value = value + temp
  }
  
  fmt.Printf("FCFS traversal count =%6d\n", value)
} //done

func sstf(a algo, s []cylinder, length int) {
  temp := a.initialCYL
  queue := s
  traversal := 0
  head := 0
  
  
  for i := 0; i < length - 5; i++ {
    calcCloseVal(temp, s, length)
    fmt.Printf("Servicing%6d\n",s[0].cyl)
    temp = s[0].cyl
    if i == 0 {
      head = s[0].cyl
    }
    s = s[1:]
  }
  
  sort.Slice(queue, func(i, j int) bool {
    return queue[i].cyl < queue[j].cyl})
  
  traversal = head - queue[len(queue) - 1].cyl
  if traversal < 0 {
    traversal = traversal * -1
  }
  
  t := queue[0].cyl - queue[len(queue) - 1].cyl
  if t < 0 {
    t = t * -1
  }
  traversal = traversal + t
  
  t = head - a.initialCYL
  if t < 0 {
    t = t * -1
  }
  traversal = traversal + t
  
  fmt.Printf("SSTF traversal count =%6d\n", traversal)
} //sorta done

func scan(a algo, s []cylinder, length int) {
  queue := make([]cylinder, 0)
  tempU := 0
  tempL := 0
  head := a.initialCYL
  value := 0
  t := 0
  // lengthL := 0
  lengthU := 0
   
  calcCloseCyl(a.initialCYL, s, length)
  sort.Slice(s, func(i, j int) bool { return s[i].cyl < s[j].cyl})

  for i := 0; i <len(s); i++ {
    if s[i].closeVal < 0 {
      t++
    }
  }

  queue = s[:t]
  s = s[t:]
  // lengthL = len(queue)
  lengthU = len(s)
  
    for i := 0; i < len(s); i++ {
      fmt.Printf("Servicing%6d\n",s[i].cyl)
      tempU = head - s[i].cyl
      if tempU < 0 {
        tempU = tempU * -1
      }
      head = s[i].cyl
      value = value + tempU
    }
    
    if len(queue) != 0 {
      tempU = s[lengthU - 1].cyl - a.upperCYL
      if tempU < 0 {
        tempU = tempU * -1
      }
      
      value = value + tempU 
        head = a.upperCYL
        
        for i := len(queue) - 1; i >= 0; i-- {
          fmt.Printf("Servicing%6d\n",queue[i].cyl)
          tempL = head - queue[i].cyl
          if tempL < 0 {
            tempL = tempL * -1
          }
          head = queue[i].cyl
          value = value + tempL
      }
    }
  fmt.Printf("SCAN traversal count =%6d\n", value)
} //done

func cscan(a algo, s []cylinder, length int) {
  queue := make([]cylinder, 0)
  t := 0
  traversal := 0
  l := 0
  he := len(s)

  calcCloseCyl(a.initialCYL, s, length)
  sort.Slice(s, func(i, j int) bool {
    return s[i].cyl < s[j].cyl})

  for i := 0; i <len(s); i++ {
    if s[i].closeVal < 0 {
      t++
    }
  }

  queue = s[:t]
  s = s[t:]

    for i := 0; i < len(s); i++ {
      fmt.Printf("Servicing%6d\n",s[i].cyl)
    }
    for i := 0; i < len(queue); i++ {
      fmt.Printf("Servicing%6d\n", queue[i].cyl)
      l = queue[i].cyl
  }

if(len(queue) != 0) {
    traversal = a.initialCYL - a.upperCYL
    if traversal < 0 {
      traversal = traversal * -1
    }
    traversal = traversal + (a.upperCYL - a.lowerCYL)
    traversal = traversal + (l - a.lowerCYL)
  } else {
    traversal = s[he - 1].cyl - a.initialCYL
    if traversal < 0 {
      traversal = traversal * -1
    }
  }
  fmt.Printf("C-SCAN traversal count =%6d\n", traversal)
} // done

func look(a algo, s []cylinder, length int) {
  queue := make([]cylinder, 0)
  tempU := 0
  tempL := 0
  head := a.initialCYL
  value := 0
  t := 0
  lengthU := 0

  calcCloseCyl(a.initialCYL, s, length)
  sort.Slice(s, func(i, j int) bool {
    return s[i].cyl < s[j].cyl})

  for i := 0; i <len(s); i++ {
    if s[i].closeVal < 0 {
      t++
    }
  }

  queue = s[:t]
  s = s[t:]
  lengthU = len(s) - 1

    for i := 0; i < len(s); i++ {
      fmt.Printf("Servicing%6d\n",s[i].cyl)
      tempU = head - s[i].cyl
      if tempU < 0 {
        tempU = tempU * -1
      }
      head = s[i].cyl
      value = value + tempU
    }
    
    head = s[lengthU].cyl
    
    for i := len(queue) - 1; i >= 0; i-- {
      fmt.Printf("Servicing%6d\n",queue[i].cyl)
      tempL = head - queue[i].cyl
      if tempL < 0 {
        tempL = tempL * -1
      }
      head = queue[i].cyl
      value = value + tempL
  }
  
  fmt.Printf("LOOK traversal count =%6d\n", value)
}//done

func clook(a algo, s []cylinder, length int) {
  queue := make([]cylinder, 0)
  t := 0
  le := 0
  l := 0
  traversal := 0


  calcCloseCyl(a.initialCYL, s, length)
  sort.Slice(s, func(i, j int) bool {
    return s[i].cyl < s[j].cyl})

  for i := 0; i <len(s); i++ {
    if s[i].closeVal < 0 {
      t++
    }
  }

  queue = s[:t]
  s = s[t:]
  le = len(s) - 1


    for i := 0; i < len(s); i++ {
      fmt.Printf("Servicing%6d\n",s[i].cyl)
    }
    for i := 0; i < len(queue); i++ {
      fmt.Printf("Servicing%6d\n",queue[i].cyl)
      l = queue[i].cyl
  }
  
  if(len(queue) != 0) {
      traversal = a.initialCYL - s[le].cyl
      if traversal < 0 {
        traversal = traversal * -1
      }
      traversal = traversal + (s[le].cyl - queue[0].cyl)
      traversal = traversal + (l - queue[0].cyl)
    } else {
      traversal = s[le].cyl - a.initialCYL
      if traversal < 0 {
        traversal = traversal * -1
      }
    }
    fmt.Printf("C-LOOK traversal count =%6d\n", traversal)
}//done

func main() {
 var a algo
 var b cylinder
 var err error
 
 arg := os.Args[1]; //declaring a variable with argument 1
 file, err := os.Open(arg) //opening that argument file
 check(err) //checking if its ok

 var lines []string //declaring new string with all the lines
 scanner := bufio.NewScanner(file) //creating scanner
 for(scanner.Scan()) {
   lines = append(lines, scanner.Text())
 } //adding every line to my string

 length := len(lines)
 
 //use
 splitter := strings.Split(lines[0], " ") //splitline by spaces
 splitterTest := strings.Split(splitter[1], "\t") //split entire line
 a.use = splitterTest[0]

 //lowerCYL
 splitter = strings.Split(lines[1], " ") //splitline by spaces
 splitterTest= strings.Split(splitter[1], "\t") //split entire line
 count, err := strconv.Atoi(splitterTest[0]) //conv 5string to 5int
 a.lowerCYL = count

 //upperCYL
 splitter = strings.Split(lines[2], " ") //splitline by spaces
 splitterTest = strings.Split(splitter[1], "\t") //split entire line
 count, err = strconv.Atoi(splitterTest[0]) //conv 5string to 5int
 a.upperCYL = count
 

 //initialCYL
 splitter = strings.Split(lines[3], " ") //splitline by spaces
 splitterTest = strings.Split(splitter[1], "\t") //split entire line
 count, err = strconv.Atoi(splitterTest[0]) //conv 5string to 5int
 a.initialCYL = count
 
 if a.upperCYL < a.lowerCYL {
   fmt.Printf("ABORT(13):upper (%d) < lower (%d)\n", a.upperCYL, a.lowerCYL)
   return
  } else if a.initialCYL > a.upperCYL {
     fmt.Printf("ABORT(11):initial (%d) > upper (%d)\n", a.initialCYL, a.upperCYL)
     return
  } else if a.initialCYL < a.lowerCYL {
     fmt.Printf("ABORT(12):initial (%d) < lower (%d)\n", a.initialCYL, a.lowerCYL)
     return
  }
 
 t := length
 
 for i := 4; i <= length - 2; i++ {
   splitter = strings.Split(lines[i], " ") //splitline by spaces
   count, err = strconv.Atoi(splitter[1]) //conv 5string to 5int
   if count < a.lowerCYL || count > a.upperCYL {
   t = t - 1
   }
 }
 
struc := make([]cylinder, length - 5) //making a slice of type process

 for i := 4; i <= length - 2; i++ {
   splitter = strings.Split(lines[i], " ") //splitline by spaces
   count, err = strconv.Atoi(splitter[1]) //conv 5string to 5int
   if count < a.lowerCYL || count > a.upperCYL {
     fmt.Printf("ERROR(15):Request out of bounds: req (%d) > upper (%d) or  < lower (%d)\n", count, a.upperCYL, a.lowerCYL)
     
     continue
   }else {
     b.cyl = count
     struc[i- 4] = b // assign struct to slice
   }
 }


 switch a.use {
  case "fcfs":
   fmt.Printf("Seek algorithm: FCFS\n")
   fmt.Printf("\tLower cylinder:%6d\n", a.lowerCYL)
   fmt.Printf("\tUpper cylinder:%6d\n", a.upperCYL)
   fmt.Printf("\tInit cylinder:%7d\n", a.initialCYL)
   fmt.Printf("\tCylinder requests:\n")
   for i := 0; i < len(struc); i++ {
     fmt.Printf("\t\tCylinder%6d\n", struc[i].cyl)
   }
   fcfs(a, struc, length)

 case "sstf":
   fmt.Printf("Seek algorithm: SSTF\n")
   fmt.Printf("\tLower cylinder:%6d\n", a.lowerCYL)
   fmt.Printf("\tUpper cylinder:%6d\n", a.upperCYL)
   fmt.Printf("\tInit cylinder:%7d\n", a.initialCYL)
   fmt.Printf("\tCylinder requests:\n")
   a.currentCYL = a.initialCYL

   for i := 0; i < len(struc); i++ {
     fmt.Printf("\t\tCylinder%6d\n", struc[i].cyl)
   }
   sstf(a, struc, length)

 case "scan":
   fmt.Printf("Seek algorithm: SCAN\n")
   fmt.Printf("\tLower cylinder:%6d\n", a.lowerCYL)
   fmt.Printf("\tUpper cylinder:%6d\n", a.upperCYL)
   fmt.Printf("\tInit cylinder:%7d\n", a.initialCYL)
   fmt.Printf("\tCylinder requests:\n")
   a.currentCYL = a.initialCYL

   for i := 0; i < len(struc); i++ {
     fmt.Printf("\t\tCylinder%6d\n", struc[i].cyl)
   }
   scan(a, struc, length)

 case "c-scan":
   fmt.Printf("Seek algorithm: C-SCAN\n")
   fmt.Printf("\tLower cylinder:%6d\n", a.lowerCYL)
   fmt.Printf("\tUpper cylinder:%6d\n", a.upperCYL)
   fmt.Printf("\tInit cylinder:%7d\n", a.initialCYL)
   fmt.Printf("\tCylinder requests:\n")
   a.currentCYL = a.initialCYL

   for i := 0; i < len(struc); i++ {
     fmt.Printf("\t\tCylinder%6d\n", struc[i].cyl)
   }
   cscan(a, struc, length)

 case "look":
   fmt.Printf("Seek algorithm: LOOK\n")
   fmt.Printf("\tLower cylinder:%6d\n", a.lowerCYL)
   fmt.Printf("\tUpper cylinder:%6d\n", a.upperCYL)
   fmt.Printf("\tInit cylinder:%7d\n", a.initialCYL)
   fmt.Printf("\tCylinder requests:\n")
   a.currentCYL = a.initialCYL

   for i := 0; i < len(struc); i++ {
     fmt.Printf("\t\tCylinder%6d\n", struc[i].cyl)
   }
   look(a, struc, length)

 case "c-look":
   fmt.Printf("Seek algorithm: C-LOOK\n")
   fmt.Printf("\tLower cylinder:%6d\n", a.lowerCYL)
   fmt.Printf("\tUpper cylinder:%6d\n", a.upperCYL)
   fmt.Printf("\tInit cylinder:%7d\n", a.initialCYL)
   fmt.Printf("\tCylinder requests:\n")
   a.currentCYL = a.initialCYL

   for i := 0; i < len(struc); i++ {
     fmt.Printf("\t\tCylinder%6d\n", struc[i].cyl)
   }
   clook(a, struc, length)
 }


}
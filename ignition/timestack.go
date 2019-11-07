package ignition

import(
"fmt"
"time"
)

var timeset  int = 5 
var	enableTimer bool 
var	enableEngine bool 

func SetTimer(set int){
		timeset = set
		enableTimer = true
}

func StartTimer(){
	if enableTimer == true{
	t1:=time.NewTicker(1*time.Second)
	defer t1.Stop()
	for{
		timeset--
		select {
		case <- t1.C:
			fmt.Println(timeset)
		}
		if (timeset < 0 || enableTimer == false){
			enableTimer = false
			enableEngine = true
			break
		}
	}
	}else{
		fmt.Println("Please reset timer")
	}
}

func StopTimer(){
	enableTimer = false
}

func SetEngine(callback func()bool) {
	if  callback() != true{
		StartTimer()
	}else{
		time.Sleep(time.Second)
	}
}

func SetPower()bool{
	if enableEngine != false{
		return true
	}else{
		return false
	}
}
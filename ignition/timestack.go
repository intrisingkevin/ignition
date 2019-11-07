package ignition

import(
"fmt"
"time"
)

var timeset  int = 5
var	enableTimer bool = false
var	enableEngine bool = false

func SetTimer(set int){
		timeset = set
		enableTimer = true
}

func StartTimer(){
	t1:=time.NewTicker(1*time.Second)
	defer t1.Stop()
	for{
		select {
		case <- t1.C:
			fmt.Println(timeset)
			timeset--
			}
		if timeset == 0|| enableTimer == false{
			enableTimer = false
			enableEngine = true
			break
		}
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

func SetPower(callback func()bool){
	if enableEngine != false{
	}else{
	}
}
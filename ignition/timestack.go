package ignition

import(
"fmt"
"time"
)

var Timeset  int
var	EnableTimer bool
var	EnableEngine bool
var	Pin bool

func SetTimer(timeset int){
		ptr := &Timeset 
		*ptr = timeset
		pt := &EnableTimer
		*pt = true
}

func StartTimer(){
	ptr := &EnableTimer
	pt := &EnableEngine
		for i:=Timeset; i>0; i--{
			if Timeset == 0|| *ptr == false{
				*ptr = false
				*pt = true
				break
			}
		fmt.Println(i)
		time.Sleep(time.Second)
		}
}

func StopTimer(){
	ptr := &EnableTimer
	*ptr = false
}

func SetEngine(pin *bool){
	if *pin != true{
		StartTimer()
	}else{
		time.Sleep(time.Second)
	}
}

func SetPower()bool{

	if EnableEngine != false{
		return true
	}else{
		return false
	}
}
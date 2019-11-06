package ignition

import(
"fmt"
"time"
)
var Timeset = 5 
var Enable = false 

func SetTime(timeset int){
		Timeset = timeset
}
func StartTime(){
	for i:=0; i<Timeset; i++{
		if Enable == true{
			break
		}
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
func StopTime(){
	Enable = true
}

func SetEngine(*pin){

}

func SetPower(){

}




func Getignition(timeset int, signalIn bool)bool{
	t1:=time.Tick(1*time.Second)
	for{
		select {
		case <-t1:
			enable := detection(signalIn)
			}
			if enable == true{
				break
			}
		}
	t2:=time.Tick(time.Second)
	for {
			select {
			case <-t2:
				}
			timeset-=1
			fmt.Println(timeset)
				if timeset < 1{
					break
				}
		}
	return true
}

func detection(engine bool) bool{
	signalIn := engine
	if signalIn == true{
		fmt.Println("timer is unenable")
		return false
	}else{
		fmt.Println("timer is enable")
		return true
	}
}
package ignition

import(
"fmt"
"time"
)

func Getignition(timeset int, signalIn bool)bool{
	t1:=time.Tick(1*time.Second)
	for{
		enable := detection(signalIn)
		if enable == true{
			break
		}
		select {
		case <-t1:
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
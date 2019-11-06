package ignition

import(
"fmt"
"time"
)

func ignition(timeset int, signalIn bool)bool{
	t1:=time.Tick(1*time.Second)
	for{
		enable := detection(signalIn)
		if !enable{
			break
		}
		select {
			case <-t1:
			}
		}
	t2:=time.Tick(time.Duration(timeset) * time.Second)
	for {
			select {
			case <-t2:
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
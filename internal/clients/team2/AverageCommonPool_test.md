
///////////////////////////  TESTBENCH ////////////////////////////////
package main
import "fmt"
func main(){
 fmt.Println(AverageCommonPoolDilemma())
 return
}

func AverageCommonPoolDilemma() float64 {
	ResourceHistory := make(map[uint]float64)
	ResourceHistory[0] = 0     
	ResourceHistory[1] = 100
	ResourceHistory[2] = 200
	ResourceHistory[3] = 300
	
	var turn uint=3
	var default_strat float64 = 20 
	var fair_sharer float64 
	var altruist float64
 
	var no_freeride float64 = 1
	var freeride float64 = 5  

	if turn==0 { 
		decreasing_pool = 0
		return default_strat
	}

	altruist = determine_altruist(turn,ResourceHistory)
	fair_sharer = determine_fair(turn,ResourceHistory)  

	prevTurn := turn - 1
	prevTurn2 := turn -2
	if ResourceHistory[prevTurn] > ResourceHistory[turn] {
		if ResourceHistory[prevTurn2] > ResourceHistory[prevTurn] {
			return altruist
		}
	}

	if float64(turn) > no_freeride { 
		if ResourceHistory[prevTurn] < (ResourceHistory[turn] * freeride) {
			if ResourceHistory[prevTurn2] < (ResourceHistory[prevTurn] * freeride) {
				return 0
			}
		}
	}
	return fair_sharer
}

func determine_altruist(turn uint, ResourceHistory map[uint]float64 ) float64 { 
	var tune_alt float64 = 2    
	for j := turn; j > 0; j-- { 
		prevTurn := j - 1
		if ResourceHistory[j]-ResourceHistory[prevTurn] > 0 {
			return ((ResourceHistory[j] - ResourceHistory[prevTurn]) / 6) * tune_alt
		}
	}
	return 0
}

func determine_fair(turn uint,ResourceHistory map[uint]float64) float64 {
	var tune_average float64 = 1
	for j := turn; j > 0; j-- {  
		prevTurn := j - 1
		if ResourceHistory[j]-ResourceHistory[prevTurn]> 0 {
			return ((ResourceHistory[j]- ResourceHistory[prevTurn]) / 6) * tune_average 
		}
	}
	return 0
}

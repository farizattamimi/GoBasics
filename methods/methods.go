m
import "math/rand"

type score struct {
	currentscore int
}

func (s *score) threepoints() int{
	s.currentscore = s.currentscore+3
	return s.currentscore		
}

func (s *score) twopoints() int{
	s.currentscore = s.currentscore+2
	return s.currentscore		
}

func (s *score) onepoint() int{
	s.currentscore = s.currentscore+1
	return s.currentscore		
}

func main(){
	s := score{currentscore: 0}
	s2 := score{currentscore: 0}

	for i := 0; i < 50; i++{
		//team 1
		random := rand.Intn(100)
		if random%3 == 0 {
			fmt.Println("team 1 threepoints: ", s.threepoints())
		}else if random%2 == 0{
			fmt.Println("team 1 twopoints: ", s.twopoints())
		}else{
			fmt.Println("team 1 onepoint: ", s.onepoint())
		}

		//team 2
		randomtwo := rand.Intn(100)
		if randomtwo%3 == 0 {
			fmt.Println("team 2 threepoints: ", s2.threepoints())
		}else if randomtwo%2 == 0{
			fmt.Println("team 2 twopoints: ", s2.twopoints())
		}else{
			fmt.Println("team 2 onepoint: ", s2.onepoint())
		}
	}

fmt.Println("final score")
fmt.Println("team 1: ", s.currentscore)
fmt.Println("team 2: ", s2.currentscore)

}
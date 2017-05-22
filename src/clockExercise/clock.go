package clockExercise

type Clock struct {
	Hour   int
	Minute int
}

func (c *Clock) AddMinutesToTime(minutes int) {
	minutesSum := c.Minute + minutes

	if minutesSum%60 == 0 {
		c.Hour += minutesSum / 60
		c.Minute = 0
	} else {

		c.Hour += minutesSum / 60
		c.Minute = minutesSum % 60
	}

}

func (c *Clock) SubMinutesToTime(minutes int) {

	minutesSub := c.Minute - minutes
	if minutesSub >= 0 {
		c.Minute = minutesSub
	} else {
		if minutesSub%60 == 0 {
			c.Hour += minutesSub / 60
			c.Minute = 0
		} else {
			c.Hour += minutesSub / 60
			c.Minute =60+ minutesSub % 60
		}

	}

}

func CompareToTimes(c1, c2 Clock) (equal bool) {
	if c1.Hour == c2.Hour && c1.Minute == c2.Minute  {
		equal = true
		return
	}
	equal = false
	return

}

//func main() {
//
//	var time1 *clockExercise.Clock = &clockExercise.Clock{5, 10}
//
//	var time2 *clockExercise.Clock = &clockExercise.Clock{5, 10}
//	
//	var time3 *clockExercise.Clock = &clockExercise.Clock{1, 43}
//	
//	var time4 *clockExercise.Clock = &clockExercise.Clock{23, 15}
//
//	//execute the compare
//
//	if clockExercise.CompareToTimes(*time1, *time2) {
//		fmt.Println(" The two times are equal")
//	} else {
//		fmt.Println(" The two times are not equal")
//	}
//
//	//execute the add
//
//	time3.AddMinutesToTime(123)
//	fmt.Printf("The new time is %v : %v  \n", time3.Hour, time3.Minute)
//
//	//execute the Sub
//
//	time4.SubMinutesToTime(10)
//	fmt.Printf("The new time is %v : %v \n ", time4.Hour, time4.Minute)
//
//}


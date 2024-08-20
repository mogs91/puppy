package puppy

import(
	"github.com/mogs91/wolf"
)

func Bark() string{
	return "Woof!!"
}

func Barks()string{
	return "Woof! Woof! Woof!"
}


func BigBark()string{
	return wolf.WhenGrownUps(Bark())
}

func Barking(){
	return "The Dog is barking"
}
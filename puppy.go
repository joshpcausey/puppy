package puppy

import "github.com/joshpcausey/dog"

func Bark() string {
	return "Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}

func AddVersion() string {
	return "version updated"
}

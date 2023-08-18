package annalyn


// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
    if !knightIsAwake{
        return true
    }else{
    return false
    }
  
	panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {

    if knightIsAwake || archerIsAwake || prisonerIsAwake{
        return true
    }else{
    return false
    }
	panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {

    if (!archerIsAwake) && prisonerIsAwake{
        return true
    }else{
    return false
    }
	panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {

  switch {

      	case knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == true:
  return true
        case knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true:
    		return true
        case knightIsAwake == false && archerIsAwake == false && petDogIsPresent == true:
    		return true
        case archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == true:
    		return true
        default:
			return false
    }
  
	panic("Please implement the CanFreePrisoner() function")
}

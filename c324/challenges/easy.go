package challenges

// babylonian method
func recursiveSqRoot(num float64, guess float64) float64 {
    if !guess {
        // Take an initial guess at the square root
        guess = num / 2.0
    }
    var d = num / guess               // Divide our guess into the number
    var newGuess = (d + guess) / 2.0  // Use average of g and d as our new guess
    if newGuess == guess { // The guess hasn't changed, return our guess
        return guess
    }
    // Recursively solve for closer and closer approximations of the square root
    return recursiveSqRoot(num, n)
}

func SqRoot(num, precision) float64 {
  recursiveSqRoot(num)
}

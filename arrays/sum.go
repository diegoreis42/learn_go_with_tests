package arrays

func Sum(list []int) int {
  var result int

  for _, number := range list {
    result += number
  }

  return result
}


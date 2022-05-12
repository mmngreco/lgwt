package sum

func Sum(numbers []int) (total int) {

    for _, number := range numbers {
        total += number
    }
    return total
}

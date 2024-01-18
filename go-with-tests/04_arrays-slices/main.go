package main

func Sum(numbers []int) (sum int) {
    sum = 0
    for _, v := range numbers {
        sum += v
    }
    return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
    for _, v := range numbersToSum {
        if len(v) == 0 {
            sums = append(sums, 0)
        } else {
            sums = append(sums, Sum(v))
        }
    }
    return
}

func SumAllTails(tailsToSum ...[]int) (tailSums []int) {
    for _, v := range tailsToSum {
        if len(v) == 0 {
            tailSums = append(tailSums, 0)
        } else {
            tailSums = append(tailSums, Sum(v[1:]))
        }
    }
    return
}

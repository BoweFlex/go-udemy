package iteration

const repetitions int = 5

func Repeat(rep string) string {
    var repeated string
    for i := 0; i < repetitions; i++ {
        repeated += rep
    }
    return repeated
}

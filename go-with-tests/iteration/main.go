package iteration


func Repeat(rep string, repetitions int) string {
    var repeated string
    for i := 0; i < repetitions; i++ {
        repeated += rep
    }
    return repeated
}

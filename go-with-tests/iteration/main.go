package iteration

func Repeat(rep string) string {
    var repeated string
    for i := 0; i < 5; i++ {
        repeated = repeated + rep
    }
    return repeated
}

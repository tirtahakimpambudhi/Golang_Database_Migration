package helper


func PanicIFError(err error){
	if err != nil {
		panic(err)
	}
}
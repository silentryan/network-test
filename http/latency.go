package http

// please also consider upload and download performance

// situation: user report login a website slow through a proxy
func Get(url string) {}

// situation: user report login a website slow without forward proxy
func GetFromProxy(url string, proxy string) {}

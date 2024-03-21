package main

func main() {
	svc := NewLoggingService(&pricefetcher{})

	jsonApi := NewJsonApiServer(svc)
	jsonApi.Run()
}

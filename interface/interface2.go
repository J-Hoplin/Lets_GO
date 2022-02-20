package main

import (
	"awesomeProject1/fedex"
	"awesomeProject1/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func interface2() {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("Prince", koreaPostSender)
	SendBook("Princess", koreaPostSender)

	fedexPostSender := &fedex.FedexSender{}
	SendBook("parcel1", fedexPostSender)
	SendBook("parcel2", fedexPostSender)
}

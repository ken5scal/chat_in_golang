package main

type room struct {
	fwd_chan chan []byte //channel that olds message pending to transfer
}

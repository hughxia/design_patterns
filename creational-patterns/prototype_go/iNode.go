package main

type iNode interface {
	print(string)
	clone() iNode
}

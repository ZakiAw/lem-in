package main

type Room struct{
	Name string
	Jeran []*Room
}

type Start struct{
	StartRoom Room
}

type End struct{
	EndRoom Room
}

type Ant struct{
	AntNum int
}

package framework 

type HandlerFunc func (*Context)

type HandlerChain []HandlerFunc
package public

var Redis = redis{
	ExpireSeconds: 300,
}

type redis struct {
	ExpireSeconds int
}

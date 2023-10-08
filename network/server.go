package network

type serveropts struct {
	transport []Transport
}

type server struct {
	serveropts
}

func Newserver(opts serveropts) server {
	return *&server{
		serveropts: opts,
	}
}

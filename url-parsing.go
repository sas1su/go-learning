package main
import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	fmt.Println(u)

	if err != nil {
		panic(err)
	}
	//find sheme
	fmt.Println(u.Scheme)
	//find user string, usernaeme
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	// get password, if exist return true
	p, perr := u.User.Password()
	if !perr {
		panic(perr)
	}
	fmt.Println(p)
	//Split hostport
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	//path
	fmt.Println(u.Path)
	//fragment after #
	fmt.Println(u.Fragment)
	//raw query
	fmt.Println(u.RawQuery)
	//parse rawquery
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])


}
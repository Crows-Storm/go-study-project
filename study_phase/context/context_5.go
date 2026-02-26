// ----------------
// Request/Response
// ----------------

// Sample program that implements a web request with a context that is
// used to timeout the request if it takes too long.

package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type Response struct {
	Body []byte
	Err  error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 11150*time.Millisecond)
	//ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.ardanlabs.com/blog/post/index.xml", nil)
	if err != nil {
		log.Println(err)
		return
	}

	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := http.Client{
		Transport: &tr,
	}

	ch := make(chan Response, 1)
	go func() {
		log.Println("Starting Request")

		resp, err := client.Do(req)

		if err != nil {
			ch <- Response{Err: err}
			return
		} else {
			//io.Copy(os.Stdout, resp.Body)	// Write directly to standard output (stdout)
			body, _ := io.ReadAll(resp.Body)
			ch <- Response{Body: body, Err: nil}
		}

		defer resp.Body.Close()
	}()

	select {
	case <-ctx.Done():
		log.Println("timeout, cancel work...")
		log.Println(<-ch)
	case r := <-ch:
		if r.Err != nil {
			log.Println(err)
		} else {
			log.Println(string(r.Body))
		}
	}
}

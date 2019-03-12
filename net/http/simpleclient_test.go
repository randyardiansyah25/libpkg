package http

import (
	"fmt"
	"testing"
)

func TestNewSimpleClient(t *testing.T) {
	client := NewSimpleClient("POST", "https://devpds.pegadaian.co.id:9090/param/hargaemaschannel", 30)
	client.SetAuthorization("Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTIzNTkzMDQsInVzZXJfbmFtZSI6IjUyMDEiLCJzY29wZSI6WyJSRUFEIiwiV1JJVEUiXSwiYXV0aG9yaXRpZXMiOlsiVVNTRVIiXSwianRpIjoiYmNiYWIyM2EtNmNlNC00ZTQ3LWFjOWYtZDMyMWM3OWI1ZTU3IiwiY2xpZW50X2lkIjoiYXBsaWthc2lpYnMifQ.dZfvBjmuwsC-E1s7n1bPQWhuw7To3uh2OFirzJ563NgOtUX0qHCp3oLCRPyjzwHRZnOuEomPApeb3FNxUNa6ItpeZYfIodiSUljb9sPc-RIAcOGtp8xWfVZeu-_FpNM3ut4wY-CGTjdEJSvKApV8JG7LNcyKahrlU6a4ypRjpmCCox84QjdhJ_04GVlHyqYFpMCeuWZYlyZdlXV4OgTpoH9RIwMaLnFUApjQ36E4wvq1Q7qfDDnMljYTQwBoFn8bUSCbl8aYGPkbbGVdiY1XFMI59V6zMWjiiYtfsli4RPBDtLR-PIi7Aqhm8bvnQ3hmel_XwxJkgBS4OlGe5pjOtA")
	res := client.DoRawRequest(`{"channelId": "6017", "clientId": "5201", "flag": "K"}`)
	fmt.Println(res)
}

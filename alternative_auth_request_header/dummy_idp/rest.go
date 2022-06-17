package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Backend request to the OAuth Server hosted on OpenShift
func ChallengeRequest(s Settings, clientId string, codeChallenge string, user string) string {
	// Read the key pair to create certificate
	cert, err := tls.LoadX509KeyPair(s.Cert, s.Key)
	if err != nil {
		log.Fatal("Unable to load the keypair", err)
	}

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile(s.CA)
	if err != nil {
		log.Fatal("Unable to load the CA", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Construct the Code Challenge URL
	// Refer to https://docs.openshift.com/container-platform/4.10/authentication/configuring-internal-oauth.html
	url := s.Backend + "?client_id=" + clientId + "&code_challenge=" + codeChallenge +
		"&code_challenge_method=S256&response_type=code&scope=user:full"

	// Request URL via the created HTTPS client over port 8443 via GET
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Remote-User", user)

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("no such host", err)
	}

	// Read the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Response ", err)
	}

	// Print the response body to stdout
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Printf("%s\n", body)

	u, _ := response.Location()
	fmt.Println("Location: ", u)

	for name, values := range response.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	// This is the important bit...
	loc, _ := response.Location()
	fmt.Println("Location: ", loc)
	return loc.String()
}

// Backend request to the OAuth Server hosted on OpenShift
// This is a web-login request (as redirected from the OpenShift Console)
func WebLoginRequest(s Settings, clientId string, idp string, state string, user string) string {
	// Read the key pair to create certificate
	cert, err := tls.LoadX509KeyPair(s.Cert, s.Key)
	if err != nil {
		log.Fatal("Unable to load the keypair", err)
	}

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile(s.CA)
	if err != nil {
		log.Fatal("Unable to load the CA", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a mTLS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	url := s.Backend + "?client_id=" + clientId + "&idp=" + idp +
		"&scope=user:full" + "&state=" + state +
		"&response_type=code"

	// Request URL via the created HTTPS client over port 8443 via GET
	req, _ := http.NewRequest("GET", url, nil)

	// This could be read from settings
	req.Header.Add("X-Remote-User", user)

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("no such host", err)
	}

	// Read the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Response ", err)
	}

	// Print the response body to stdout
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Printf("%s\n", body)

	// This is the important bit...
	loc, _ := response.Location()
	fmt.Println("Location: ", loc)

	// Very helpful during development confirmation
	// It should be removed if in production.
	for name, values := range response.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	return loc.String()
}

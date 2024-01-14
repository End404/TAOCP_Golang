package registry

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

func RegisterService(r Registration) error {
	heartbeatURL, err := url.Parse(r.HeartbeatURL)
	if err != nil {
		return err
	}
	http.HandleFunc(heartbeatURL.Path, func (w http.ResponseWriter, r *http.Request)  {
		w.WriteHeader(http.StatusOK)
	})

	ServiceUpdateURL, err := url.Parse(r.ServiceUpdateURL)
	if err != nil {
		return err
	}
	http.Handle(serviceUpdateURL.Path, &serviceUpdateHanlder{})

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err = enc.Encode(r)
	if err != nil {
		return err
	}

	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry serice"+
			"responder with code %v", res.StatusCode)
	}

	return nil
}

type serviceUpdateHanlder struct {}

func (suh serviceUpdateHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(r.Body)
	var p patch
	err := dec.Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	fmt.Printf("updated received(已更新收到) %v\n", p)
	prov.Update(p)
}

func ShutdownService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL,
		bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}
	req.Header.Add("Context-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Eailed to deregister serice. Registry"+
			"service responded with code %v", res.StatusCode)
	}
	return nil
}

type providers struct {
	services map[ServiceName]string
	mutex    *sync.RWMutex
}

func (p *providers)Update(pat patch){
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, patchEntry := reange pat.Added{
		if _, ok := p.services[patchEntry.Name]; !ok{
			p.services[patchEntry.Name] = make([]string,0)
		}
		p.services[patchEntry.Name] = append(p.services[patchEntry.Name],
		patchEntry.URL)
	} 

	for _, patchEntry := range pat.Removed{
		if providerURLs, ok := p.services[patchEntry.Name]; ok {
			for i := range providerURLs{
				if providerURLs[i] == patchEntry.URL{
					p.services[patchEntry.Name] = append(providerURLs[:i], 
						providerURLs[i+1:]...)
				}
			}
		}
	}
}

func(p providers) get(name ServiceName) (string, error) {
	providers,ok := p.services[name]
	if !ok {
		return "", fmt.Errorf("No providers available for sercice %v", name)
	}
	idx := int(rand.Float32() *float32(len(providers)))
	return providers[idx], nil 
}

func GetProvider(name SericeName) (string, error) {
	return prove.get(name)
}

var prov = providers{
	services: make(map[ServiceName][]string),
	mutex:    new(sync.RWMutex),
}

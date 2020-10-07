package main

// Sample multi IP call and return for ip-api.com batch api
// https://ip-api.com/docs/api:batch

//Try https://github.com/BenB196/ip-api-go-pkg for full functionality

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
)

func main() {
        apiURL := "http://ip-api.com/batch"

        var ips []string
        // sample input. you can construct an array of IPs in other ways
        ips = append(ips, "1.1.1.1", "8.8.8.8")
        payload, err := json.Marshal(ips)
        if err != nil {
                log.Fatal("Error Marshaling IP list", err)
        }

        client := &http.Client{}
        req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(string(payload)))
        if err != nil {
                log.Fatal("Failed creating HTTP request", err)
        }

        resp, err := client.Do(req)
        if err != nil {
                fmt.Println("The HTTP request failed with error.", err)
        }

        // You can implement throttling based on these values.
        // Currently the batch api allows 15 calls per minute with 100 IPs max per call
        fmt.Printf("\nRequest per minute remaining: %v", resp.Header["X-Rl"][0])
        fmt.Printf("\nWait for %v seconds", resp.Header["X-Ttl"][0])

        f, err := ioutil.ReadAll(resp.Body)
        var loc ipLoc
        if err := json.Unmarshal(f, &loc); err != nil {
        log.Println("Failed unmarshaling JSON", err)
        }

        for i := range loc {
                fmt.Printf("\n%d: %s \t%s \t%s  - %f/%f ", i+1, loc[i].Query, loc[i].City, loc[i].Isp, loc[i].Lat, loc[i].Lon)

        }
        fmt.Println()
        //      fmt.Printf("\nRaw Json %+v", loc)

        resp.Body.Close()

}

type ipLoc []struct {
        As          string  `json:"as"`
        City        string  `json:"city"`
        Country     string  `json:"country"`
        CountryCode string  `json:"countryCode"`
        Isp         string  `json:"isp"`
        Lat         float64 `json:"lat"`
        Lon         float64 `json:"lon"`
        Org         string  `json:"org"`
        Query       string  `json:"query"`
        Region      string  `json:"region"`
        RegionName  string  `json:"regionName"`
        Status      string  `json:"status"`
        Timezone    string  `json:"timezone"`
        Zip         string  `json:"zip"`
}
                    

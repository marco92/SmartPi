/*
    Copyright (C) Jens Ramhorst
  	This file is part of SmartPi.
    SmartPi is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    SmartPi is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with SmartPi.  If not, see <http://www.gnu.org/licenses/>.
    Diese Datei ist Teil von SmartPi.
    SmartPi ist Freie Software: Sie können es unter den Bedingungen
    der GNU General Public License, wie von der Free Software Foundation,
    Version 3 der Lizenz oder (nach Ihrer Wahl) jeder späteren
    veröffentlichten Version, weiterverbreiten und/oder modifizieren.
    SmartPi wird in der Hoffnung, dass es nützlich sein wird, aber
    OHNE JEDE GEWÄHRLEISTUNG, bereitgestellt; sogar ohne die implizite
    Gewährleistung der MARKTFÄHIGKEIT oder EIGNUNG FÜR EINEN BESTIMMTEN ZWECK.
    Siehe die GNU General Public License für weitere Details.
    Sie sollten eine Kopie der GNU General Public License zusammen mit diesem
    Programm erhalten haben. Wenn nicht, siehe <http://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type JSONMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var epoch = time.Unix(0, 0).Format(time.RFC1123)

var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

func NoCache(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Delete any ETag headers that may have been set
		for _, v := range etagHeaders {
			if r.Header.Get(v) != "" {
				r.Header.Del(v)
			}
		}

		// Set our NoCache headers
		for k, v := range noCacheHeaders {
			w.Header().Set(k, v)
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func stringInSlice(list1 []string, list2 []string) bool {
	for _, a := range list1 {
		for _, b := range list2 {
			if b == a {
				return true
			}
		}
	}
	return false
}

var appVersion = "No Version Provided"

type Softwareinformations struct {
	Softwareversion string
}

func getSoftwareInformations(w http.ResponseWriter, r *http.Request) {
	data := Softwareinformations{Softwareversion: appVersion}

	// JSON output of request
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// func main() {

// 	config := smartpi.NewConfig()
// 	user := smartpi.NewUser()

// 	version := flag.Bool("v", false, "prints current version information")
// 	flag.Parse()
// 	if *version {
// 		fmt.Println(appVersion)
// 		os.Exit(0)
// 	}

// 	fmt.Println("SmartPi server started")

// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/{phaseId}/{valueId}/now", smartpi.ServeMomentaryValues)
// 	r.HandleFunc("/api/{phaseId}/{valueId}/now/{format}", smartpi.ServeMomentaryValues)

// 	r.HandleFunc("/api/version", getSoftwareInformations)

// 	fmt.Println("SmartPi server listening: " + strconv.Itoa(config.WebserverPort))
// 	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.WebserverPort), r))
// }

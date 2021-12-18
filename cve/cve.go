package cve

func dlcvelist() {
	checkUrl, err := http.Get(URL)
	if err != nil {
		return err
	}
}

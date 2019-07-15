package crons

import (
	"fmt"
	"github.com/robfig/cron"
	"net/http"
	"sync"
	"time"
)


func InititateCrons()  {
	c := cron.New()

	c.AddFunc(" */1 * * * *", startHealthCheckup)

	c.Start()
}




var wgCron sync.WaitGroup

func startHealthCheckup()  {
	// TODO add lock file here
	fmt.Println("checking health status")

	var terminals []Terminal
	db.Find(&terminals)

	for _,terminal := range terminals {
		if terminal.Url !="" {
			wgCron.Add(1)
			go checkHealth(terminal)
		} else {
			continue
		}
	}
	wgCron.Wait()
	fmt.Println("Health Check Completed")
}

func checkHealth(term Terminal)  {


	tr := &http.Transport{
		IdleConnTimeout:    time.Duration(term.Timeout) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(term.Url)

	termHealth := TerminalHealth{}
	termHealth.TerminalId = term.ID

	db.Create(&termHealth)

	termHealthHit := TerminalHealthHit{}
	termHealthHit.TerminalHealthId =termHealth.ID

	if(err != nil){
		fmt.Println("got unexpected error")
		fmt.Println(err.Error())
		termHealth.Result = 2  // wrong url
		termHealthHit.Result = -1  //wrong hit
		db.Save(&termHealthHit)
		db.Save(&termHealth)
	}else {
		if(resp.StatusCode == 200){
			fmt.Println("Terminal",term.Url , "is working ")
			termHealth.Result = 1  // working
			termHealthHit.Result = 200  //successful hit
			db.Save(&termHealthHit)
			db.Save(&termHealth)
		}else {
			fmt.Println("Terminal",term.Url , "is not working ")
			termHealthHit.Result = resp.StatusCode  //wrong hit
			db.Save(&termHealthHit)
			time.Sleep(time.Duration(term.Frequency) * time.Millisecond)
			retryHealthHit(term,termHealth,2)
		}
	}

	wgCron.Done()
}

func retryHealthHit(term Terminal, termHealth TerminalHealth, retryCount int)  {
	if(retryCount > term.Threshold){
		termHealth.Result = 0 // not working
		db.Save(&termHealth)
		return
	}

	termHealthHit := TerminalHealthHit{}
	termHealthHit.TerminalHealthId =termHealth.ID

	tr := &http.Transport{
		IdleConnTimeout:    time.Duration(term.Timeout) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(term.Url)

	if(err != nil){
		fmt.Println("got unexpected error")
		fmt.Println(err.Error())
		termHealth.Result = 2  // wrong url
		termHealthHit.Result = -1  //wrong hit
		db.Save(&termHealthHit)
		db.Save(&termHealth)
		return
	}else {
		if(resp.StatusCode == 200){
			fmt.Println("Terminal",term.Url , "is working ")
			termHealth.Result = 1  //  working
			termHealthHit.Result = 200  //successful hit
			db.Save(&termHealthHit)
			db.Save(&termHealth)
			return
		}else {
			fmt.Println("Terminal",term.Url , "is not working ")
			termHealthHit.Result = resp.StatusCode  //wrong hit
			db.Save(&termHealthHit)
			time.Sleep(time.Duration(term.Frequency) * time.Millisecond)
			retryHealthHit(term,termHealth,retryCount+1)
			return
		}
	}
}



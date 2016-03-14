GoPush
---

First you need to see what -> [Pushbullet](https://pushbullet.com) is.

A Go CLI wrapper around the Pushbullet API that helps you retrieve your saved
links from your Pushbullet account and also do searches and also push links or
notes to you account and have them synced across all your devices.  


#### How-To
NOTE: Needs go 1.5 +   
Enabled the vendor experiment with `export GO15VENDOREXPERIMENT=1`  
If you have go 1.6 it's on by default.  
* Git clone this repo.  
* In the project folder `go build -o gopush cmd/gopush/main.go`  

NOTE: at this point you might wanna add it to your `GOPATH/bin` by copying it there so that you can execute. That is if you have your `GOPATH/bin` added to `PATH`.

* Create a new Pushbullet API Token by logging into your account and going to `Settings`-> `Create Access Token`  
* Login with your generated token with `gopush login --token <my_generated_access_token>`  

NOTE: Your user details and token are stored on disk under the `~/.gopush` folder in gopush.json

#### Not Functional yet
* List your pushes with `gopush list-puses`  
* Filter your retrieved pushes with `gopush list-pushes --filter golang`
* Filter your retrieved pushes by getting only the ones that are not deleted. `gopush list-pushes --active=true`  


Licensed under: [Apache V2](http://www.apache.org/licenses/)  

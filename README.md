Gunner
---
[![CircleCI](https://circleci.com/gh/PI-Victor/gunner/tree/master.svg?style=svg)](https://circleci.com/gh/PI-Victor/gunner/tree/master)
[![Apache](https://img.shields.io/badge/license-Apache%20License%202.0-E91E63.svg?style=flat-square)](http://www.apache.org/licenses/)

#### Not functional yet!!!
---

What is [Pushbullet](https://pushbullet.com)?

`gunner` wraps the Pushbullet API.  

`gunner` - helps you retrieve your saved links, do searches and also push links, ephemerals (SMS, mirrored notifications) to your account and have them sync across all of your devices.

#### How-To
NOTE: Needs go 1.5 +  
* make install  

* Create a new Pushbullet API Token by logging into your account and going to `Settings`-> `Create Access Token` (this will be replaced at some point by OAuth loggin with your favorite provider)  
* Log in with your generated token with `gunner login --token <my_generated_access_token>`  

NOTE: Your user details and token are stored on disk under the folder
`~/.gunner/user/`

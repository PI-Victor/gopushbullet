Gunner & Sidekick
---
[![CircleCI](https://circleci.com/gh/PI-Victor/gunner/tree/master.svg?style=svg)](https://circleci.com/gh/PI-Victor/gunner/tree/master)
[![Apache](https://img.shields.io/badge/license-Apache%20License%202.0-E91E63.svg?style=flat-square)](http://www.apache.org/licenses/)

#### Not functional yet!!!
---

What is [Pushbullet](https://pushbullet.com)?

`gunner` & `sidekick` wrap the Pushbullet API.  

`gunner` - helps you retrieve your saved links, do searches and also push links, ephemerals (SMS, mirrored notifications) to your account and have them sync across all of your devices.

`sidekick` - stores the data locally and can provide functionality that the pushbullet mobile and browser app are missing. Like sorting your saved links, grouping them and a bunch of other stuff.

#### How-To
NOTE: Needs go 1.5 +  
* make install  

* Create a new Pushbullet API Token by logging into your account and going to `Settings`-> `Create Access Token` (this will be replaced at some point by OAuth loggin with your favorite provider)  
* Log in with your generated token with `gunner login --token <my_generated_access_token>`  

NOTE: Your user details and token are stored on disk under the folder
`~/.gunner/user/`

Note: If you want to store your data locally there's also `sidekick` to help
you manage your data.  
`sidekick sync` - this will retrieve and store locally all of your links, device list, created chats and subscriptions.

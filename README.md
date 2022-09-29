# README

## About

This is a tool to check if your password is in the compromized database. Use the api in here https://haveibeenpwned.com/Passwords, they do have a check but I do not trust them so I wrote this - be sure my password is not sent over the internet. Only 5 first char of your SHA1 hash of the password is sent 



## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

More info see https://wails.io/ . You need to install the wails cli to build it in addition to golang and some other dependencies (not much, just webkit development)

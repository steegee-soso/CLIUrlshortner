# CLIUrlshortener
CLIURl shortener Golang Project

##
This is a simple CLI tool to create URL shortener. Application uses  TinyUrl External API to generate the Shortener at CLI environment

##
How to use  the application
###
1.On your CLI environment, Enter this command : go run CLI $: go run URLShortnerCLI.go https://www.google.com

###
2.Wait for application to generate the URL shortner for the long url sumitted,while waiting the application initially validate 
the submitted long URL before submitting it to third party service.

###
3. A sample output :   go run URLShortnerCLI.go https://www.google.com  => https://tinyurl.com/bjnwp7u%.
###
4.The next upgrame will be a custom uild URL shortner without the use of External service.

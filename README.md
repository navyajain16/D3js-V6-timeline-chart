# D3js-V6-timeline-chart

Follow these initial steps first:

1. Create a project in Goland
2. Create three new directories: css, html and js.
3. Download latest D3 library and save it into this js directory. Create new file homepage.js in this js directory.
4. Create new file homepage.html in html directory
5. Create new file homepage.css in css directory
6. Create three new files in main project directory: main.go, homepage.go, timeline.go and log.go

## Backend

### main.go
In main.go complete go code for creating and running a web server service is given.

### log.go
log.go is responsible for logging to console only.

### homepage.go
homepage.go is a function for serving the homepage.

### timeline.go
timeline.go code is the basis, on which we will build right on. The function getTimelineData() will generate some pseudorandom data in the future and send it back to the browser.

### updated_timeline.go
This contains the updated code after generation of random data in 24 hours.

## Initial test 
Synchronise all Go modules either using built-in Goland functionality or using command go get -u all. Run the main.go and don’t forget to set it to package mode.

## Frontend

### homepage.html
This file loads two javascript files, creating heading and one<div> element, that will hold our timeline.

### homepage.css
This is used in a tooltip.

### homepage.js
this consists of code that sends a JSON request to /get_timeline_data and prints whatever it gets back. This also contains code to access code, set dimensions, draw canvas, set scales, prepare data, prepare tooltip and to draw data.

This project was taken from "https://itnext.io/how-to-create-d3js-v6-timeline-1a32e586fe56". For any queries refer this link.

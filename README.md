# Golang scraper

## About

A command line application with a set of tools originally intended for being run regularly to acquire whatever data neccessary from the web and possibly sending e-mail notifications when neccessary.

## Available tools/functionality

 - Fetch data about the current interest rates of the government issued savings bonds of the Republic of Latvia and send e-mail notification in case the 12 months bonds interest rate is equal or higher than the desired configured value.

## Initial setup

Before running, you must create and fill an `.env` file; use this [example](.env) to fill out the values.

To run locally, you can press `F5` if using VS Code to run via a launch profile or just use the CMD command `go run main.go` in the root of the project.
 
## Use

Intended to be run regularly with the help of automatization tools or something like a Windows Scheduler.
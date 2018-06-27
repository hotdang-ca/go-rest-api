# Go Rest Api

## What is this?
Just a simple rest api example with Go.

## Why is this?
Because I want to learn go. Geeze it's tricky.

## How to Build
Clone this repo. Perform a `go get`. Then, `go run main`. Base Http Port is 8000, you can change it if you want.

## How to Use?
`GET /`: Version number
`GET /people`: List in-memory array of people
`GET /people/{id}`: List specific person
`DELETE /people/{id}`: Delete specific person
`POST /people`: Create a new person with JSON in the body:

```
{
	"firstname": "",
	"lastname": "",
	"address": {
		"street": "",
		"city": "",
		"state": ""
	}
}
```

## LICENSE
Copyright (c) 2018 James Robert Perih

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

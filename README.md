## Answers to Go Problem Sheet 2
## Author: Matthew Shiel 
## Student ID: G00338622

These are my set of completed programs to the problems from https://github.com/data-representation/data-representation.github.io/blob/master/problems/go-web-applications.md
## How to Run The Programs

*Assumes that Git and Go are installed along with the prerequisites.*
**If not, they can be found from https://golang.org/dl/ and https://git-scm.com/downloads 

**1. Clone the Repository**
```bash
> git clone https://github.com/mattshiel/02-go-problem-sheet.git
```
**2. Change Directory to the Folder**

```bash
Open the terminal/command line and navigate into the folder 
eg. > cd 02-go-problem-sheet
```

**2. Compile the Game**

```bash
Before you run the program you can compile it using go build 
eg. > go build webApp.go
```

**3. Run the Game**

```bash
To run the game enter './' followed by the executable produced
For Mac:
> ./webApp
For Windows:
> ./webApp.go.exe
```

**4. Examine Respone with Curl**

```curl
curl localhost:8080 -v
curl localhost:8080/guess -v
```
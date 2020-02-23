# Hacker News Scraper

### Library Choice
Initially the colly library was going to be used for this project, however the only way to traverse the web page was
using a for loop. This iteration was not useful when truing to traverse from post to post as it would get all instances
of a certain tag. This resulted in 6 different arrays containing the different information for each post which then had 
to be matched and a post would be built. This quickly lead to an increase in complexity, so in order to reduce the 
complexity custom traversal functions were used.

## How to run:
#####This step does not require any libraries/software to be installed
In the main directory there is an executable file called __hackernews__.
Open a terminal and run the executable file with the following command:

`./hackernews --posts n`

Where 'n' is the number of posts to return. 

## How to install:

#### Using docker:
##### 1. Docker Build:
Please install docker using the following guide: https://docs.docker.com/. Once installed run the following command to 
build the docker container:

`sudo docker build -t hacker-news .`

Once built th output will contain the result of the test as well as an execution of the system with the required post 
number at 35.

##### 2. Docker Run:
To run the container that has just been built, run the following command:

`sudo docker run hacker-news "./hackernews" "--posts" "n"`

Where 'n' is the number of posts to return.

#### Using Golang
###### 1. Install golang
To install golang please follow the install steps in the following guide: https://golang.org/doc/install

###### 2. Run without installing the project (optional)
(Optional Step) To run tests run the following command from the project directory: `test -v ./...`

To run the application without an executable run the following commands: 

`go build ./...`

`go run main.go --posts n` 
Where 'n' is the number of posts to return.

###### 3. Install Project
In the project directory run the following command to install the dependencies: `get -d -v ./...` 

To generate an executable file run the following command: `install -v ./...`

The executable file will be saved in the go bin directory. The location of this directory can be found running the 
following command  `go env` and looking at the GOPATH variable.

###### 4. Run Application Executable

Navigate to the bin directory in the GOPAth directory and run the following command:
`./HackerNewsScraper --posts n`
Where 'n' is the number of posts to return.

# fizzbuzz
The goal is to implement a web server that will expose a REST API endpoint that: 
* Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
* Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server is:
* Ready for production
* Easy to maintain by other developers

* We add a statistics endpoint allowing users to know what the most frequent request has been. 
This endpoint:
* Accept no parameter
* Return the parameters corresponding to the most used request, as well as the number of hits for this request

# Technologies used
* GoLang: Programming langage. Official url: `https://golang.org/doc/`

* Kubernetes: is the container-orchestration system we use for automating the application deployment, scaling, and management. Official url: `https://kubernetes.io/docs/home/`

* Helm: helps us manage our Kubernetes application. With the Helm Charts in `fizzbuzz/FizzBuzzChart/`, we define how install our Kubernetes application (we can also use it to upgrade the application)

* Promotheus: systems monitoring and alerting toolkit used to produce our endoinpts statistics. Official url: `https://prometheus.io/docs/introduction/overview/`

# How to run the application
* Clone project from github: `https://github.com/dfkossi/fizzbuzz.git`

* Install go and setup all the environnement variables

* Install Kubernetes local cluster 

* Go to the home location of your project, to deploy the application by running in command lines: 
- `export TAG=fizzbuzz`
- `helm install $TAG FizzBuzzChart/ -n $TAG`

* Use command:  `kubectl get pods`to see if all the pods are running 

* Then use Postman to run by hunting: `http://localhost:3000` 
* use `http://localhost:3000/metrics`to see the metrics(statistics)
* To run the tests of our application: 
- `go test -run TestCreateFizzBuzz` or 
- `go test -run TestMain`
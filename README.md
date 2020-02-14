# fizzbuzz
The goal is to implement a web server that will expose a REST API endpoint that: 
* Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
* Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.


# Technologies used
* GoLang: Programming langage. `https://golang.org/doc/`

* Kubernetes: is the container-orchestration system we use for automating the application deployment, scaling, and management. `https://kubernetes.io/docs/home/`

* Helm: helps us manage our Kubernetes application. With the Helm Charts in `fizzbuzz/FizzBuzzChart/`, we define how install our Kubernetes application (we can also use it to upgrade the application)

* Promotheus: systems monitoring and alerting toolkit used to produce our endoinpts statistics. `https://prometheus.io/docs/introduction/overview/`


# How to run the application
* Clone project from github: `https://github.com/dfkossi/fizzbuzz.git`

* Install go and setup all the environnement variables: `https://golang.org/doc/install`

* Install Kubernetes local cluster: `https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/`

* Install Helm `https://helm.sh/docs/using_helm/#installing-helm`


# Deploy to Production
* A build image of the application is pushed on `hub.docker.com` (`dfkossi/fizzbuzz`)

* Go to the home location of your project, to deploy the application by running in command lines: 
- `export TAG=master`
- `helm install $TAG FizzBuzzChart/ -n $TAG`

* Use command:  `kubectl get pods`to see if all the pods are running 

* Then use Postman to run by hunting: `http://localhost:{TCP Port}/` example`http://localhost:3000` 


# Get endpoint statistics (metric documentation is also there: `https://prometheus.io/docs/introduction/overview/`):
* Use `http://localhost:{TCP Port}/metrics`to see the metrics(statistics); example `http://localhost:3000/metrics`

# How to unit test the application 
* To run the tests of our application in your local repository: 
- `go test -run TestCreateFizzBuzz` or 
- `go test -run TestMain` or
- `go test -v` to run all the tests


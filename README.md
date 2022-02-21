#Ecosia challenge.
***
Considerations

There are many ways to configure minikube, and it varies depending on which distribution you are using
and the way you chose to setup your minikube network you might need or not to use the minikube IP/localhost,
you will know better than me how to access an exposed service in your test environment after running minikube tunnel.

Another point is the use or not of the environment variable $KUBECONFIG, I don't know how you manage or export
this in your system, in my case in my laptop running exherbo I need to provide $KUBECONFIG if the deploy through
the script fails, check the messages carefully it will check if this environment variable is empty or not.
If this is empty, and you need it to deploy please export (the script will warn you).
Otherwise it will warn you but the deploy will just work if this is not needed.

Last thing is I have created a /status endpoint to provide information about the app and to use for the liveness probe
but I wasn't sure if I should use it or not because it wasn't requested in the challenge so maybe can be used
in the future, it can be ignored.
***
1. To run this code you can use the build_deploy.sh script that was sent separately by email.
    * the script has a help function, but it has basically 3 functions, clone, build and deploy.
    * everytime you need to pass a directory path you should provide the trailing slash.
2. How to test?
    * I don't know if you have local.ecosia.org in your hosts file or if you have dnsmasq in your system, I am
not considering any of that, please use curl with the respective header and everything should work just fine.
    * curl --header "HOST: local.ecosia.org" http://localhost/tree or http://localhost/tree/<NAME>
    * curl --header "HOST: local.ecosia.org" http://$(minikube ip)/tree or http://$(minikube ip)/tree/<NAME> 

Thank you!
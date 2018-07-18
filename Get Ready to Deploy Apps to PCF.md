


## Get Ready to Deploy Apps to PCF

### Setup PCF Workspace
### Using On-Prem PCF:
When using an **on-prem** version of PCF, login using your organization's uri.
`cf login [PCF_URI]`
Target the org and space where you want to deploy your apps:
`cf target -o [PCF_ORG]`

### Using Pivotal Web Services Account
If you've setup a Pivotal Web Services Account, use the URI given when setting up. If you can't are looking for that login URI, click on `Tools` in the left navigation. 
`cf login -a api.run.pivotal.io`

Use the email/password combination you used when signing up for your Pivotal Web Services account. 

There's no need to change the targeted org or space in case you've just created this, but if you do want to change it, use the `target` command. 
 `cf target -o [PCF_ORG]`

### Build the Akka Application

You can deploy an Akka application using PCF’s [java-buildpack](https://github.com/cloudfoundry/java-buildpack.git). Our sample application is inspired by [akka-sample-cluster](https://github.com/akka/akka-samples/tree/2.5/akka-sample-cluster-scala). 

It has backend nodes that calculate a factorial upon receiving messages from frontend nodes. Frontend nodes also expose `HTTP interface GET <frontend-hostname>/info` that shows number of jobs completed.

If you haven't already, clone or download the project folder from GitHub. 
* Example ssh clone `git clone git@github.com:cloudfoundry/cf-networking-examples.git`

Go to your local repo's folder:
`cd akka-sample-cluster-on-cloudfoundry/akka-sample-cluster/`

If you are deploying on the Pivotal Web Services, you will need to make one change to your application name to ensure it's uniqueness in the shared domain. (More on `Shared Domain` later.)

Open the file `src/main/resources/application.conf` and edit the `seed-host` entry. 

It currently says: `seed-host = "akka-backend.apps.internal"` but you'll need to make it unique, for example: `seed-host = "akka-backend-[special_name].apps.internal"`

Compile and package the Akka backend component:
`sbt backend:assembly # backend`  

Compile and package both Akka  frontend component:
`sbt frontend:assembly # frontend`

### Deploying Akka Backend Application

Next, are going to deploy sample Akka backend without starting it. We will use the `--no-start` flag to make sure it's not started after deployment. We are also using options `--no-route ` and `--health-check-type none` since the backend doesn't expose any HTTP ports.

_**Note 1**: you must use the correct java build pack based on your deployment option, on-prem or with Pivotal Web Services. When deploying on-premise, you should use java_buildpack_offline. When deploying into Pivotal Web Services, use instead java_buildpack_

_**Note 2**: substitute [YOUR_BACKEND_APP_NAME] with the name you've chosen for the akka sample backend_

**On Prem Deployment Command**
`cf push --no-start --health-check-type none [YOUR_BACKEND_APP_NAME] -p target/scala-2.11/akka-sample backend.jar -b java_buildpack_offline -d apps.internal`

**Pivotal Web Services Deployment Command**
`cf push --no-start --health-check-type none [YOUR_BACKEND_APP_NAME] -p target/scala-2.11/akka-sample-backend.jar -b java_buildpack -d apps.internal`

Take a look at the option `-d apps.internal`. This option configures the application to use a new shared domain for internal routes called `apps.internal`. Applications that use internal routes will be directly communicating over the container network, instead of having traffic routed out and back in through the load balancer and GoRouter.

Next, you'll add this network policy to allow the backend nodes to communicate:

`cf add-network-policy [YOUR_BACKEND_APP_NAME] --destination-app [YOUR_BACKEND_APP_NAME] --port 2551 --protocol tcp`

### Start and Scale Akka Backend Application

Start the backend app:
`cf start [YOUR_BACKEND_APP_NAME]`

Check the log to see that first node joined itself:
`cf logs [YOUR_BACKEND_APP_NAME]`

**IMPORTANT: To prevent cluster split, verify that the first node is running before scaling it.**

Scale backend to 2 instances:
`cf scale [YOUR_BACKEND_APP_NAME] -i 2`

### Deploying Akka Frontend Application

Similarly to the backend, we will deploy the Frontend without starting it. 

**On Prem Deployment Command**
`cf push --no-start --health-check-type none [YOUR_FRONTEND_APP_NAME] -p target/scala-2.11/akka-sample frontend.jar -b java_buildpack_offline -d apps.internal`

**Pivotal Web Services Deployment Command**
`cf push --no-start --health-check-type none [YOUR_FRONTEND_APP_NAME] -p target/scala-2.11/akka-sample-frontend.jar -b java_buildpack -d apps.internal`

Add this network policy to allow the frontend app to communicate with the backend app via backend's TCP:2551 port:

`cf add-network-policy [YOUR_FRONTEND_APP_NAME] --destination-app [YOUR_BACKEND_APP_NAME] --port 2551 --protocol tcp`

### Start and View Akka Frontend Application

Start the frontend app:
`cf start [YOUR_FRONTEND_APP_NAME]`

In separate windows or terminal sessions, check logs from both frontend and backend to ensure all client/server and server-to-server communications are working fine:

View backend logs:
`cf logs [YOUR_BACKEND_APP_NAME]`  

View frontend logs:
`cf logs [YOUR_FRONTEND_APP_NAME]`

Verify that it all works together:
`curl [YOUR_FRONTEND_APP_NAME].<YOUR_PCF_DOMAIN>/info`

If all is working as expected, it should show the number of completed jobs! 

## Summary

We’ve shown you how to get Akka Cluster up and running quickly.  You’ve seen how easy it is to scale your cluster using the cli. Now, you are ready to run your own Akka Cluster on PCF, and immediately enjoy many production-grade features out-of-the-box!

Want to learn more about Pivotal Cloud Foundry? Join us at [SpringOne Platform](https://springoneplatform.io/) in Washington, D.C., September 24 to 27. [Register now](https://2018.event.springoneplatform.io/register).


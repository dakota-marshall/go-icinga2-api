* Create of existing object fails with '500 Object could not be created' 
This should only be an error is passed parameters do not match 100% else its not an error
* Tests involving the package API require waiting on icinga to reload its config, currently we are just sleeping. Need to more intellegently call the endpoints to try until we no longer get "icinga is reload"

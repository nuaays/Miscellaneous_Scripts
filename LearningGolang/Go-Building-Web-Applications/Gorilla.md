


#


```
Gorilla is less of a framework and more of a set of very useful tools that are generally bundled in frameworks. Specifically, Gorilla contains:

* gorilla/context: This is a package for creating a globally-accessible variable from the request. It is useful for sharing a value from the URL without repeating the code to access it across your application.
* gorilla/rpc: This implements RPC-JSON, which is a system for remote code services and communication without implementing specific protocols. This relies on the JSON format to define the intentions of any request.
* gorilla/schema: This is a package that allows simple packing of form variables into a struct, which is an otherwise cumbersome process.
* gorilla/securecookie: This, unsurprisingly, implements authenticated and encrypted cookies for your application.
* gorilla/sessions: Similar to cookies, this provides unique, long-term, and repeatable data stores by utilizing a file-based and/or cookie-based session system.
* gorilla/mux: This is intended to create flexible routes that allow regular expressions to dictate available variables for routers.
The last package is the one we are most interested in here, and it comes with a related package called gorilla/reverse, which essentially allows you to reverse the process of creating regular expression-based muxes. We will cover that topic in detail in the later section.

```

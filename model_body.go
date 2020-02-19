/*
 * wallet-api
 *
 * <img src=\"logo.png\" style=\"display: block; margin-left: auto; margin-right: auto; width: 70%;\">  # Intro  This website documents the API methods available for the `wallet-api` program.  The default API url is [127.0.0.1:8070](http://127.0.0.1:8070).  Note that this program is distinct from `turtle-service`, which uses a different API. The api docs for `turtle-service` can be found [here](https://api-docs.turtlecoin.lol/).  When using the docs, be aware that you can click the 'Model' tab to get a description of all available parameters, their types, and an example value.  # Initialization  * Start by launching the `wallet-api` program. * From a terminal (or via a programming language):  ``` ./wallet-api --rpc-password somepassword ```  * Or if you're on windows:  ``` wallet-api.exe --rpc-password somepassword ```  * An RPC password is always required, and you should then provide this password with every request to the API, in the `X-API-KEY` header.  * If you want to see the available configuration options for the `wallet-api`, launch the program with: ``` ./wallet-api --help ```  * Note that you cannot generate/open a wallet via the command line, only via the API.   This is by design to prevent application developers having to parse command line output.  # Try it Out  If you want to test wallet-api out without doing any programming, you can use the 'Try it out' section in swagger. * Start by launching your wallet with CORS enabled:  ``` ./wallet-api --rpc-password \"mypassword\" --enable-cors \"*\" ```  * Click the 'Authorize' button in Swagger, and type in the RPC password you just chose - in this case, \"mypassword\", and click 'Authorize'. * Click a method to expand it, then click the 'Try it Out' button. * You can modify the parameters you want to send here, and then hit 'Execute' to send the request. * Note that this may not work in Firefox, I had to use Chrome to get it to work. F12 (developer console) may help diagnose CORS issues.  # API Wrappers  Finally, whilst you may send raw HTTP requests to the API url, you may be interested in an API/RPC wrapper.  These can be generated for many languages, by using the `Generate Client` option in the top menu. Be patient, this can take some time to complete.  # Support  If you are having issues, please stop by our [discord](http://chat.turtlecoin.lol) and visit the `#dev_learning` channel for assistance.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Body struct {
	// The filepath to save the wallet JSON to. Note that this is relative to where wallet-api was launched from - it is recommended you use an absolute path.
	Filename string `json:"filename,omitempty"`
}

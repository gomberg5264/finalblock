#### Blockchain Technology - Exam Assignment (Bitcoin Block Explorer)
##### Feature list
* Search either by block height or block hash
* See detailed information about a block
* Responsive design
* Show best block on page load

##### Dependencies
*Back-end* is written in Go (Golang):
* BTCd
* Gin

*Front-end* requires JavaScript 6 minimum and depends on these stylesheets:
* Material Icons
* Quicksand font

##### How it works
JavaScript uses the Fetch API to send GET requests to the Gin server in the back-end.

JSONs are returned containing requested information and are applied to
elements in the DOM for viewing.

Endpoints list:
* "/" - index.html, styling and JS 
* "/getBlockCount" - string containing block count in the blockchain
* "/getLatestBlock" - JSON containing block Hash and Height
* "/getBlockHeaderFromHash/:hash" - JSON containing block Header information, found from Hash
* "/getBlockHeaderFromHeight/:height" - Same as above, but found from Height instead
* "/getBlockTransactions/:hash" - Return a JSON containing count of all transactions in the block,
as well as an array of all the transaction Hashes
* "/getBlockHashFromHeight/:height" - Return JSON containing block Hash found by Height
* "/getBlockHeightFromHash/:hash" - Return JSON containing block height found by Hash


#### Version history
##### [v0.1](https://github.com/Electrostasy/BCT-Explorer/releases/tag/v0.1) / 2020-01-16:
* Finalized front-end and back-end
* Added README.md

# algodexidx

This is an early-stage service exposing a REST API for watching accounts as well as providing a debugging endpoint for inspecting a msgpack transaction.

## Services

### Account watcher

* **Watching** one or more Algorand Accounts (cumulative)
  * `POST /account {"address":["address1", "address2", ...]}`
    * This will cause the server to watch the chain for any transaction impacting those accounts and queuing parallel updates of those accounts against current node state.
* **Fetching** information for a **specific** account:
  * `GET /account/address`
    * Note: ID "1" represents the holdings ALGO balance. 
    * eg return:
    ```json
    {
      "address": "6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY",
      "holdings": {
        "1": {
          "asset": 1,
          "amount": 9991000,
          "decimals": 6,
          "metadataHash": "",
          "name": "ALGO",
          "unitName": "ALGO",
          "url": ""
        },
        "17574184": {
          "asset": 17574184,
          "amount": 3,
          "decimals": 0,
          "metadataHash": "",
          "name": "Fluffy Catcher",
          "unitName": "NET",
          "url": "https://fluffybunnycoin.store/15"
        }
      }
    }
    ```
* **Fetching** information for **all** accounts:
  * `GET /account`
    ```json
      [{"address": "6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY"},
       {"address": "xxxxxxx"}]
    ```
  * `GET /account?view=full`
    ```json
      [
        {
          "address": "6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY",
          "holdings": {
            "1": {
              "asset": 1,
              "amount": 9991000,
              "decimals": 6,
              "metadataHash": "",
              "name": "ALGO",
              "unitName": "ALGO",
              "url": ""
            }
          }
        },
        {"address": "xxxx", "holdings": {"1": {}}}
      ]
      ```

### Debug helper

* Returns output from `goal clerk inspect` for base64 encoded msgpack transaction data. 
  * `POST /inspect/unpack {"msgpack": "base64 encoded data"}`
    * Returns text/plain response from 'goal clerk inspect' of body data.  
  
## Building (for testing)

From project root:

```
docker build -t algodexidxsvr:latest .
```

## Running (for testing)

```
docker run --rm -p 8000:8000 algodexidxsvr:latest [args]
```
### Optional Arguments

```
 -debug (adds logging ot show all incoming/outgoing requests/responses)
 -network (testnet|mainnet)  (defaults to testnet)  
````
#### Swagger

Contents of gen/openapi3.yaml can be pasted into https://editor.swagger.io/ for API view/testing.  

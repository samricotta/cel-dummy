# Celestia Dummy Application


The application's purpose is to send and receive messages to/from the celestia network.

To submit a message, an HTTP POST request is made with the a namespace id, message and gas limit included in the body of the request and a block height is returned upon success.

To retrieve a message, a GET request is made to the /namespaced_data/{nID}/height/{height} endpoint that includes the namespaceID under which the user submitted their message, and the block height at which the message was included.

## Installation

```sh
make build
```

### Run `Submit`

Command:

```sh
./cel-dummy submit <namespace_id> <message> <gas_limit>
```

Response:

```
Message successfully submitted at block height: <block_height> 
```


### Run `Retrieve`

Command:

```sh
./cel-dummy retrieve <namespace_id> <block_height>
```

Response:

```sh
Message retrieved: <message>
```

## Resources
https://docs.celestia.org/developers/celestia-node

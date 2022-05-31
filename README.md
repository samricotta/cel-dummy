# Celestia Dummy Application

The application's purpose is to send and receive messages to both the celestia app and celestia node. 
A POST request is sent with the a namespace id, message and gas limit. A response is sent back which is inclusive of the
block height. To receive a message, a GET request is sent along with the block height and retrieves the assigned message.

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


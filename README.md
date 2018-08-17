# NSQ Producer flogo activity
This activity allows your flogo application to publish a topic on NSQ


## Installation

```bash
flogo install github.com/abasse/folgonsqproducer
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "NsqdTcpAddress",
      "type": "string",
      "required": true
    },
    {
      "name": "NsqdHttpAddress",
      "type": "string",
      "required": true
    },
    {
      "name": "Topic",
      "type": "string",
      "required": true
    },
    {
      "name": "Message",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```


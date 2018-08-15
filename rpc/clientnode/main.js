const net = require('net')

const client = new net.Socket()

client.connect(1234, '127.0.0.1', function () {
  client.write(JSON.stringify({
    'method': 'DemoService.Div',
    'params': [{
      'A': 3,
      'B': 4
    }],
    'id': 1
  }))
})

client.on('data', function (data) {
  const obj = JSON.parse(data.toString())
  console.log(obj)
  client.destroy()
})

client.on('close', function () {
  console.log('closed')
})

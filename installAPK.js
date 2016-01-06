var Promise = require('bluebird')
var adb = require('adbkit')
var client = adb.createClient()
var apk = 'app.apk'

client.listDevices()
  .then(function(devices) {
    return Promise.map(devices, function(device) {
      return client.install(device.id, apk)
    })
  })
  .then(function() {
    console.log('Installed %s on all connected devices', apk)
  })
  .catch(function(err) {
    console.error('Something went wrong:', err.stack)
  })
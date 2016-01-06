var Promise = require('bluebird')
var adb = require('adbkit')
var client = adb.createClient()

client.listDevices()
  .then(function(devices) {
    return Promise.filter(devices, function(device) {
      return client.getFeatures(device.id)
        .then(function(features) {
          return features['android.hardware.nfc']
        })
    })
  })
  .then(function(supportedDevices) {
    console.log('The following devices support NFC:', supportedDevices)
  })
  .catch(function(err) {
    console.error('Something went wrong:', err.stack)
  })
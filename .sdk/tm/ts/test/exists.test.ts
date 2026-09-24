
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { FreeMealSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = FreeMealSDK.test()
    equal(testsdk instanceof FreeMealSDK, true,
      'FreeMealSDK.test() must return a client synchronously')
  })

})

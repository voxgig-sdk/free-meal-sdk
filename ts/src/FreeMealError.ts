
import { Context } from './Context'


class FreeMealError extends Error {

  isFreeMealError = true

  sdk = 'FreeMeal'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  FreeMealError
}


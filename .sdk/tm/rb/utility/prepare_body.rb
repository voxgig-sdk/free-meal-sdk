# FreeMeal SDK utility: prepare_body
require_relative 'media'
module FreeMealUtilities
  PrepareBody = ->(ctx) {
    return nil unless ctx.op.input == "data"
    return FreeMealUtilities.raw_body(ctx.reqdata) if FreeMealUtilities.raw_request?(ctx.point)
    ctx.utility.transform_request.call(ctx)
  }
end

# FreeMeal SDK utility: make_context
require_relative '../core/context'
module FreeMealUtilities
  MakeContext = ->(ctxmap, basectx) {
    FreeMealContext.new(ctxmap, basectx)
  }
end

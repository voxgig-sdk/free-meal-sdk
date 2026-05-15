-- FreeMeal SDK error

local FreeMealError = {}
FreeMealError.__index = FreeMealError


function FreeMealError.new(code, msg, ctx)
  local self = setmetatable({}, FreeMealError)
  self.is_sdk_error = true
  self.sdk = "FreeMeal"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function FreeMealError:error()
  return self.msg
end


function FreeMealError:__tostring()
  return self.msg
end


return FreeMealError

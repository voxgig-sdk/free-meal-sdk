# FreeMeal SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module FreeMealFeatures
  def self.make_feature(name)
    case name
    when "base"
      FreeMealBaseFeature.new
    when "ratelimit"
      FreeMealRatelimitFeature.new
    when "retry"
      FreeMealRetryFeature.new
    when "test"
      FreeMealTestFeature.new
    when "timeout"
      FreeMealTimeoutFeature.new
    else
      FreeMealBaseFeature.new
    end
  end
end

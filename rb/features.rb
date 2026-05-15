# FreeMeal SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module FreeMealFeatures
  def self.make_feature(name)
    case name
    when "base"
      FreeMealBaseFeature.new
    when "test"
      FreeMealTestFeature.new
    else
      FreeMealBaseFeature.new
    end
  end
end

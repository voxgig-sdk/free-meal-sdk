# FreeMeal SDK exists test

require "minitest/autorun"
require_relative "../FreeMeal_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = FreeMealSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end

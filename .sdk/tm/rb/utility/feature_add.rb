# FreeMeal SDK utility: feature_add
module FreeMealUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end

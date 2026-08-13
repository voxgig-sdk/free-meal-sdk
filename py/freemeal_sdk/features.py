# FreeMeal SDK feature factory

from freemeal_sdk.feature.base_feature import FreeMealBaseFeature
from freemeal_sdk.feature.test_feature import FreeMealTestFeature


def _make_feature(name):
    features = {
        "base": lambda: FreeMealBaseFeature(),
        "test": lambda: FreeMealTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()

# FreeMeal SDK utility: make_context

from freemeal_sdk.core.context import FreeMealContext


def make_context_util(ctxmap, basectx):
    return FreeMealContext(ctxmap, basectx)

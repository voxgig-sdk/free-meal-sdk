# ProjectName SDK exists test

import pytest
from freemeal_sdk import FreeMealSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = FreeMealSDK.test(None, None)
        assert testsdk is not None

# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'Sha1Result',
    'AwaitableSha1Result',
    'sha1',
    'sha1_output',
]

@pulumi.output_type
class Sha1Result:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableSha1Result(Sha1Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Sha1Result(
            result=self.result)


def sha1(input: Optional[str] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSha1Result:
    """
    Returns a hexadecimal representation of the SHA-1 hash of the given string.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:sha1', __args__, opts=opts, typ=Sha1Result).value

    return AwaitableSha1Result(
        result=pulumi.get(__ret__, 'result'))
def sha1_output(input: Optional[pulumi.Input[str]] = None,
                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[Sha1Result]:
    """
    Returns a hexadecimal representation of the SHA-1 hash of the given string.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:sha1', __args__, opts=opts, typ=Sha1Result)
    return __ret__.apply(lambda __response__: Sha1Result(
        result=pulumi.get(__response__, 'result')))

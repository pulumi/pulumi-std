# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'Md5Result',
    'AwaitableMd5Result',
    'md5',
    'md5_output',
]

@pulumi.output_type
class Md5Result:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableMd5Result(Md5Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Md5Result(
            result=self.result)


def md5(input: Optional[str] = None,
        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableMd5Result:
    """
    Returns a (conventional) hexadecimal representation of the MD5 hash of the given string.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:md5', __args__, opts=opts, typ=Md5Result).value

    return AwaitableMd5Result(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(md5)
def md5_output(input: Optional[pulumi.Input[str]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Md5Result]:
    """
    Returns a (conventional) hexadecimal representation of the MD5 hash of the given string.
    """
    ...

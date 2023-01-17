# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'Filemd5Result',
    'AwaitableFilemd5Result',
    'filemd5',
    'filemd5_output',
]

@pulumi.output_type
class Filemd5Result:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableFilemd5Result(Filemd5Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Filemd5Result(
            result=self.result)


def filemd5(input: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableFilemd5Result:
    """
    Reads the contents of a file into a string and returns the MD5 hash of it.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:filemd5', __args__, opts=opts, typ=Filemd5Result).value

    return AwaitableFilemd5Result(
        result=__ret__.result)


@_utilities.lift_output_func(filemd5)
def filemd5_output(input: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Filemd5Result]:
    """
    Reads the contents of a file into a string and returns the MD5 hash of it.
    """
    ...

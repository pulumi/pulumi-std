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
    'Filesha1Result',
    'AwaitableFilesha1Result',
    'filesha1',
    'filesha1_output',
]

@pulumi.output_type
class Filesha1Result:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableFilesha1Result(Filesha1Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Filesha1Result(
            result=self.result)


def filesha1(input: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableFilesha1Result:
    """
    Reads the contents of a file into a string and returns the SHA1 hash of it.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:filesha1', __args__, opts=opts, typ=Filesha1Result).value

    return AwaitableFilesha1Result(
        result=__ret__.result)


@_utilities.lift_output_func(filesha1)
def filesha1_output(input: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Filesha1Result]:
    """
    Reads the contents of a file into a string and returns the SHA1 hash of it.
    """
    ...
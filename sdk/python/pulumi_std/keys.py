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
    'KeysResult',
    'AwaitableKeysResult',
    'keys',
    'keys_output',
]

@pulumi.output_type
class KeysResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[str]:
        return pulumi.get(self, "result")


class AwaitableKeysResult(KeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return KeysResult(
            result=self.result)


def keys(input: Optional[Mapping[str, Any]] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableKeysResult:
    """
    Returns a lexically sorted list of the map keys.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:keys', __args__, opts=opts, typ=KeysResult).value

    return AwaitableKeysResult(
        result=pulumi.get(__ret__, 'result'))
def keys_output(input: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[KeysResult]:
    """
    Returns a lexically sorted list of the map keys.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:keys', __args__, opts=opts, typ=KeysResult)
    return __ret__.apply(lambda __response__: KeysResult(
        result=pulumi.get(__response__, 'result')))

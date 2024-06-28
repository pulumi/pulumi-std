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
    'CidrsubnetResult',
    'AwaitableCidrsubnetResult',
    'cidrsubnet',
    'cidrsubnet_output',
]

@pulumi.output_type
class CidrsubnetResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableCidrsubnetResult(CidrsubnetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CidrsubnetResult(
            result=self.result)


def cidrsubnet(input: Optional[str] = None,
               netnum: Optional[int] = None,
               newbits: Optional[int] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCidrsubnetResult:
    """
    Takes an IP address range in CIDR notation (like 10.0.0.0/8) and extends its prefix
    to include an additional subnet number. For example, cidrsubnet("10.0.0.0/8", netnum: 2, newbits: 8)
    returns 10.2.0.0/16; cidrsubnet("2607:f298:6051:516c::/64", netnum: 2, newbits: 8) returns
    2607:f298:6051:516c:200::/72.
    """
    __args__ = dict()
    __args__['input'] = input
    __args__['netnum'] = netnum
    __args__['newbits'] = newbits
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:cidrsubnet', __args__, opts=opts, typ=CidrsubnetResult).value

    return AwaitableCidrsubnetResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(cidrsubnet)
def cidrsubnet_output(input: Optional[pulumi.Input[str]] = None,
                      netnum: Optional[pulumi.Input[int]] = None,
                      newbits: Optional[pulumi.Input[int]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[CidrsubnetResult]:
    """
    Takes an IP address range in CIDR notation (like 10.0.0.0/8) and extends its prefix
    to include an additional subnet number. For example, cidrsubnet("10.0.0.0/8", netnum: 2, newbits: 8)
    returns 10.2.0.0/16; cidrsubnet("2607:f298:6051:516c::/64", netnum: 2, newbits: 8) returns
    2607:f298:6051:516c:200::/72.
    """
    ...

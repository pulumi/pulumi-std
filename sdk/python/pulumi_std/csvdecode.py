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
    'CsvdecodeResult',
    'AwaitableCsvdecodeResult',
    'csvdecode',
    'csvdecode_output',
]

@pulumi.output_type
class CsvdecodeResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[Mapping[str, str]]:
        return pulumi.get(self, "result")


class AwaitableCsvdecodeResult(CsvdecodeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CsvdecodeResult(
            result=self.result)


def csvdecode(input: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCsvdecodeResult:
    """
    Decodes a string containing CSV-formatted data and produces a list of maps representing that data.
    	The first line of the CSV data is interpreted as a "header" row: the values given
    	are used as the keys in the resulting maps.
    	Each subsequent line becomes a single map in the resulting list,
    	matching the keys from the header row with the given values by index.
    	All lines in the file must contain the same number of fields,
    	or this function will produce an error.
    	Follows the format defined in RFC 4180.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:csvdecode', __args__, opts=opts, typ=CsvdecodeResult).value

    return AwaitableCsvdecodeResult(
        result=__ret__.result)


@_utilities.lift_output_func(csvdecode)
def csvdecode_output(input: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[CsvdecodeResult]:
    """
    Decodes a string containing CSV-formatted data and produces a list of maps representing that data.
    	The first line of the CSV data is interpreted as a "header" row: the values given
    	are used as the keys in the resulting maps.
    	Each subsequent line becomes a single map in the resulting list,
    	matching the keys from the header row with the given values by index.
    	All lines in the file must contain the same number of fields,
    	or this function will produce an error.
    	Follows the format defined in RFC 4180.
    """
    ...

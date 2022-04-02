# 用途

- 汉字转拼音
- 中文文本检索预处理

# 用法

- ``Easy("中文")`` => ``"zhongwen"``
- ``Get("中文").StringStyle(FullUpper)`` => ``"ZHONGWEN"``
- ``Get("中文").StringStyle(Initial)`` => ``"zhw""``
- ``Get("中文").StringStyle(InitialUpper)`` => ``"ZHW"``
- ``Get("中文").StringStyle(First)`` => ``"zw"``
- ``Get("中文").StringStyle(FirstUpper)`` => ``"ZW"``

